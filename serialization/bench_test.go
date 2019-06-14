package serialization_bench

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
	"unsafe"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/neovim/go-client/msgpack"
)

var validate = os.Getenv("VALIDATE")
var commonData []*A

func init() {
	commonData = generate()
}

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func generate() []*A {
	rand.Seed(42)
	a := make([]*A, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &A{
			Name:     randString(16),
			BirthDay: time.Now(),
			Phone:    randString(10),
			Siblings: rand.Intn(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

// UnsafeString converts the []byte to string without a heap allocation.
func UnsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(&b[0])),
		Len:  len(b),
	}))
}

// UnsafeBytes converts the string to []byte without a heap allocation.
func UnsafeBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  len(s),
		Cap:  len(s),
		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data,
	}))
}

type Serializer interface {
	Marshal(o interface{}) []byte
	Unmarshal(d []byte, o interface{}) error
	String() string
}

func benchMarshal(b *testing.B, s Serializer) {
	b.StopTimer()
	data := commonData
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Marshal(data[rand.Intn(len(data))])
	}
}

func benchUnmarshal(b *testing.B, s Serializer) {
	b.StopTimer()
	data := commonData
	ser := make([][]byte, len(data))
	for i, d := range data {
		o := s.Marshal(d)
		t := make([]byte, len(o))
		copy(t, o)
		ser[i] = t
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &A{}
		err := s.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("%s failed to unmarshal: %s (%s)", s, err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay.String() == i.BirthDay.String() //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// encoding/gob

type GobSerializer struct {
	b   bytes.Buffer
	enc *gob.Encoder
	dec *gob.Decoder
}

func (g *GobSerializer) Marshal(o interface{}) []byte {
	g.b.Reset()
	err := g.enc.Encode(o)
	if err != nil {
		panic(err)
	}
	return g.b.Bytes()
}

func (g *GobSerializer) Unmarshal(d []byte, o interface{}) error {
	g.b.Reset()
	g.b.Write(d)
	err := g.dec.Decode(o)
	return err
}

func (g *GobSerializer) String() string {
	return "gob"
}

func NewGobSerializer() *GobSerializer {
	s := &GobSerializer{}
	s.enc = gob.NewEncoder(&s.b)
	s.dec = gob.NewDecoder(&s.b)
	err := s.enc.Encode(A{})
	if err != nil {
		panic(err)
	}
	var a A
	err = s.dec.Decode(&a)
	if err != nil {
		panic(err)
	}
	return s
}

func BenchmarkGobMarshal(b *testing.B) {
	s := NewGobSerializer()
	benchMarshal(b, s)
}

func BenchmarkGobUnmarshal(b *testing.B) {
	s := NewGobSerializer()
	benchUnmarshal(b, s)
}

// github.com/neovim/go-client/msgpack

type GoClientMsgpackSerializer struct {
	b   bytes.Buffer
	enc *msgpack.Encoder
	dec *msgpack.Decoder
}

func (g *GoClientMsgpackSerializer) Marshal(o interface{}) []byte {
	g.b.Reset()
	err := g.enc.Encode(o)
	if err != nil {
		panic(err)
	}
	return g.b.Bytes()
}

func (g *GoClientMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	g.b.Reset()
	g.b.Write(d)
	err := g.dec.Decode(o)
	return err
}

func (g *GoClientMsgpackSerializer) String() string {
	return "neovim/go-client/msgpack"
}

func NewGoClientMsgpackSerializer() *GoClientMsgpackSerializer {
	s := &GoClientMsgpackSerializer{}
	s.enc = msgpack.NewEncoder(&s.b)
	s.dec = msgpack.NewDecoder(&s.b)
	err := s.enc.Encode(A{})
	if err != nil {
		panic(err)
	}
	var a A
	err = s.dec.Decode(&a)
	if err != nil {
		panic(err)
	}
	return s
}

func BenchmarkGoClientMsgpackMarshal(b *testing.B) {
	s := NewGoClientMsgpackSerializer()
	benchMarshal(b, s)
}

func BenchmarkGoClientMsgpackUnmarshal(b *testing.B) {
	s := NewGoClientMsgpackSerializer()
	benchUnmarshal(b, s)
}

// github.com/google/flatbuffers/go

type FlatBuffersSerializer struct {
	builder *flatbuffers.Builder
}

func (s *FlatBuffersSerializer) Marshal(o interface{}) []byte {
	a := o.(*A)
	builder := s.builder

	builder.Reset()

	name := builder.CreateString(a.Name)
	phone := builder.CreateString(a.Phone)

	FlatBufferAStart(builder)
	FlatBufferAAddName(builder, name)
	FlatBufferAAddPhone(builder, phone)
	FlatBufferAAddBirthDay(builder, a.BirthDay.UnixNano())
	FlatBufferAAddSiblings(builder, int32(a.Siblings))
	var spouse byte
	if a.Spouse {
		spouse = byte(1)
	}
	FlatBufferAAddSpouse(builder, spouse)
	FlatBufferAAddMoney(builder, a.Money)
	builder.Finish(FlatBufferAEnd(builder))
	return builder.Bytes[builder.Head():]
}

func (s *FlatBuffersSerializer) Unmarshal(d []byte, i interface{}) error {
	a := i.(*A)
	o := FlatBufferA{}
	o.Init(d, flatbuffers.GetUOffsetT(d))
	a.Name = string(o.Name())
	a.BirthDay = time.Unix(0, o.BirthDay())
	a.Phone = string(o.Phone())
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse() == byte(1)
	a.Money = o.Money()
	return nil
}

func (s *FlatBuffersSerializer) String() string {
	return "FlatBuffer"
}

func benchUnmarshalFlatBuffers(b *testing.B, s *FlatBuffersSerializer) {
	b.StopTimer()
	data := commonData

	ser := make([][]byte, len(data))
	for i, d := range data {
		o := s.Marshal(d)
		t := make([]byte, len(o))
		copy(t, o)
		ser[i] = t
	}

	o := &A{}
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		err := s.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("%s failed to unmarshal: %s (%s)", s, err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay.String() == i.BirthDay.String() //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

func BenchmarkFlatBuffersMarshal(b *testing.B) {
	benchMarshal(b, &FlatBuffersSerializer{flatbuffers.NewBuilder(0)})
}

func BenchmarkFlatBuffersUnmarshal(b *testing.B) {
	// benchUnmarshal(b, &FlatBuffersSerializer{flatbuffers.NewBuilder(0)})
	benchUnmarshalFlatBuffers(b, &FlatBuffersSerializer{flatbuffers.NewBuilder(0)})
}

// github.com/tinylib/msgp

func BenchmarkMsgpMarshal(b *testing.B) {
	b.StopTimer()
	data := commonData

	b.ReportAllocs()
	_, err := data[rand.Intn(len(data))].MarshalMsg(nil)
	if err != nil {
		panic(err)
	}
	/*
		// compute bytes written
		writ := 0
		for i := range data {
			o, _ := data[i].MarshalMsg(nil)
			writ += len(o)
		}
		b.SetBytes(int64(writ / len(data)))
	*/
	buf := make([]byte, 0, 100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		data[rand.Intn(len(data))].MarshalMsg(buf)
	}
}

func BenchmarkMsgpUnmarshal(b *testing.B) {
	b.StopTimer()
	data := commonData

	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i], _ = d.MarshalMsg(nil)
	}
	o := &A{}
	z := A{}
	var n int
	var err error
	b.ResetTimer()
	b.ReportAllocs()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rand.Intn(len(data))
		*o = z // clear
		_, err = o.UnmarshalMsg(ser[n])
		if err != nil {
			b.Fatalf("zebrapack failed to unmarshal: %s (%s)", err, ser[n])
		}

		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}

	}
}

// github.com/glycerine/zebrapack

func BenchmarkZebraPackMarshal(b *testing.B) {
	b.StopTimer()
	data := commonData

	b.ReportAllocs()
	_, err := data[rand.Intn(len(data))].ZMarshalMsg(nil)
	if err != nil {
		panic(err)
	}
	// compute bytes written
	writ := 0
	for i := range data {
		o, _ := data[i].ZMarshalMsg(nil)
		writ += len(o)
	}
	b.SetBytes(int64(writ / len(data)))
	buf := make([]byte, 0, 100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		data[rand.Intn(len(data))].ZMarshalMsg(buf)
	}
}

func BenchmarkZebraPackUnmarshal(b *testing.B) {
	b.StopTimer()
	//	data := generateZebraPack()
	data := commonData

	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i], _ = d.ZMarshalMsg(nil)
		//ser[i], _ = proto.Marshal(d)
	}
	o := &A{}
	z := A{}
	var n int
	var err error
	b.ResetTimer()
	b.ReportAllocs()

	// compute bytes read
	red := 0
	for i := range ser {
		red += len(ser[i])
	}
	b.SetBytes(int64(red / len(ser)))

	b.SetBytes(int64(len(ser[0])))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rand.Intn(len(ser))
		*o = z // clear
		_, err = o.ZUnmarshalMsg(ser[n])
		if err != nil {
			b.Fatalf("zebrapack failed to unmarshal: %s (%s)", err, ser[n])
		}

		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}

	}
}
