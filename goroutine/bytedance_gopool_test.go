package workerpool

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/bytedance/gopkg/util/gopool"
)

func benchmarkBytedance(b *testing.B) {
	for _, test := range tests {
		runtime.GC()

		for _, concur := range concurrency {
			b.Run(fmt.Sprintf("%s_Concurrency=%d", test.name, concur), func(b *testing.B) {
				job := func() {
					defer wg.Done()
					test.executor(b.Context(), "1")
				}

				wp := gopool.NewPool("", 50, gopool.NewConfig())

				b.ReportAllocs()
				b.SetParallelism(concur)
				b.ResetTimer()

				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						wg.Add(1)
						wp.CtxGo(b.Context(), job)
					}
				})

				wg.Wait()

				b.StopTimer()
			})
		}
	}
}
