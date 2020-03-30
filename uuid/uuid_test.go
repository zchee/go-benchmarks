package uuid

import (
	"testing"

	googleuuid "github.com/google/uuid"
	vcuuid "github.com/kata-containers/runtime/virtcontainers/pkg/uuid"
	shortuuid "github.com/lithammer/shortuuid/v3"
	m4rw3ruuid "github.com/m4rw3r/uuid"
	"github.com/rogpeppe/fastuuid"
	satoriuuid "github.com/satori/go.uuid"
)

func BenchmarkFastUUID(b *testing.B) {
	g := fastuuid.MustNewGenerator()
	for i := 0; i < b.N; i++ {
		_ = fastuuid.Hex128(g.Next())
	}
}

func BenchmarkVirtContainersUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = vcuuid.Generate().String()
	}
}

func Benchmark_m4rw3rUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uid, _ := m4rw3ruuid.V4()
		_ = uid.String()
	}
}

func BenchmarkGoogleUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = googleuuid.New().String()
	}
}

func BenchmarkSatoriUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = satoriuuid.NewV4().String()
	}
}

func BenchmarkShortUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = shortuuid.New()
	}
}

func BenchmarkShortUUID2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New()
	}
}
