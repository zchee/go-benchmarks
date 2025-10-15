package workerpool

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	ants "github.com/panjf2000/ants/v2"
)

func benchmarkAnts(b *testing.B) {
	for _, test := range tests {
		runtime.GC()

		for _, concur := range concurrency {
			b.Run(fmt.Sprintf("%s_Concurrency=%d", test.name, concur), func(b *testing.B) {
				job := func(_ any) {
					defer wg.Done()

					test.executor(b.Context(), "1")
				}

				wp, _ := ants.NewPoolWithFunc(50, job, ants.WithPreAlloc(false), ants.WithExpiryDuration(time.Second*5))

				defer wp.Release()

				b.ReportAllocs()
				b.SetParallelism(concur)
				b.ResetTimer()

				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						wg.Add(1)
						wp.Invoke("")
					}
				})

				wg.Wait()

				b.StopTimer()
			})
		}
	}
}
