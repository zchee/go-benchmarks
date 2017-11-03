package serialize_bench

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import "github.com/glycerine/zebrapack/msgp"

// fieldsNotEmpty supports omitempty tags
func (z *A) fieldsNotEmpty(isempty []bool) uint32 {
	return 6
}

// ZMarshalMsg implements msgp.Marshaler
func (z *A) ZMarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.ZMsgsize())
	// map header, size 6
	// string "Name"
	o = append(o, 0x86, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "BirthDay"
	o = append(o, 0xa8, 0x42, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79)
	o = msgp.AppendTime(o, z.BirthDay)
	// string "Phone"
	o = append(o, 0xa5, 0x50, 0x68, 0x6f, 0x6e, 0x65)
	o = msgp.AppendString(o, z.Phone)
	// string "Siblings"
	o = append(o, 0xa8, 0x53, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73)
	o = msgp.AppendInt(o, z.Siblings)
	// string "Spouse"
	o = append(o, 0xa6, 0x53, 0x70, 0x6f, 0x75, 0x73, 0x65)
	o = msgp.AppendBool(o, z.Spouse)
	// string "Money"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x6e, 0x65, 0x79)
	o = msgp.AppendFloat64(o, z.Money)
	return
}

// ZUnmarshalMsg implements msgp.Unmarshaler
func (z *A) ZUnmarshalMsg(bts []byte) (o []byte, err error) {
	cfg := &msgp.RuntimeConfig{UnsafeZeroCopy: true}
	return z.ZUnmarshalMsgWithCfg(bts, cfg)
}
func (z *A) ZUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields0zbrb = 6

	// -- templateZUnmarshalMsg starts here--
	var totalEncodedFields0zbrb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zbrb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft0zbrb := totalEncodedFields0zbrb
	missingFieldsLeft0zbrb := maxFields0zbrb - totalEncodedFields0zbrb

	var nextMiss0zbrb int32 = -1
	var found0zbrb [maxFields0zbrb]bool
	var curField0zbrb string

doneWithStruct0zbrb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zbrb > 0 || missingFieldsLeft0zbrb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zbrb, missingFieldsLeft0zbrb, msgp.ShowFound(found0zbrb[:]), unmarshalMsgFieldOrder0zbrb)
		if encodedFieldsLeft0zbrb > 0 {
			encodedFieldsLeft0zbrb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField0zbrb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zbrb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zbrb = 0
			}
			for nextMiss0zbrb < maxFields0zbrb && (found0zbrb[nextMiss0zbrb] || unmarshalMsgFieldSkip0zbrb[nextMiss0zbrb]) {
				nextMiss0zbrb++
			}
			if nextMiss0zbrb == maxFields0zbrb {
				// filled all the empty fields!
				break doneWithStruct0zbrb
			}
			missingFieldsLeft0zbrb--
			curField0zbrb = unmarshalMsgFieldOrder0zbrb[nextMiss0zbrb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zbrb)
		switch curField0zbrb {
		// -- templateZUnmarshalMsg ends here --

		case "Name":
			found0zbrb[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "BirthDay":
			found0zbrb[1] = true
			z.BirthDay, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Phone":
			found0zbrb[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Siblings":
			found0zbrb[3] = true
			z.Siblings, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Spouse":
			found0zbrb[4] = true
			z.Spouse, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Money":
			found0zbrb[5] = true
			z.Money, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				panic(err)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zbrb != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of A
var unmarshalMsgFieldOrder0zbrb = []string{"Name", "BirthDay", "Phone", "Siblings", "Spouse", "Money"}

var unmarshalMsgFieldSkip0zbrb = []bool{false, false, false, false, false, false}

// ZMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) ZMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.TimeSize + 6 + msgp.StringPrefixSize + len(z.Phone) + 9 + msgp.IntSize + 7 + msgp.BoolSize + 6 + msgp.Float64Size
	return
}
