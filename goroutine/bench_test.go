package workerpool

import (
	"flag"
	"os"
	"testing"
)

var (
	alittoPond         = flag.Bool("alitto_pond", false, "run alitto_pond")
	devchatAiGopool    = flag.Bool("devchat_ai_gopool", false, "run devchat_ai_gopool")
	jeffailTunny       = flag.Bool("jeffail_tunny", false, "run jeffail_tunny")
	maurice2kUltrapool = flag.Bool("maurice2k_ultrapool", false, "run maurice2k_ultrapool")
	panjf2000Ants      = flag.Bool("panjf2000_ants", false, "run panjf2000_ants")
	bytedanceGopool    = flag.Bool("bytedance_gopool", false, "run bytedance_gopool")
)

func TestMain(m *testing.M) {
	flag.Parse()

	code := m.Run()

	os.Exit(code)
}

func Benchmark(b *testing.B) {
	switch {
	case *alittoPond:
		benchmarkPond(b)
	case *devchatAiGopool:
		benchmarkPond(b)
	case *jeffailTunny:
		benchmarkTunny(b)
	case *maurice2kUltrapool:
		benchmarkUltraPool(b)
	case *panjf2000Ants:
		benchmarkAnts(b)
	case *bytedanceGopool:
		benchmarkBytedance(b)
	}
}
