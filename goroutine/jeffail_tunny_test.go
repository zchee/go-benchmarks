package workerpool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/Jeffail/tunny"
)

func benchmarkTunny(b *testing.B) {
	for _, test := range tests {
		runtime.GC()

		for _, concur := range concurrency {
			b.Run(fmt.Sprintf("%s_Concurrency=%d", test.name, concur), func(b *testing.B) {
				mu := &sync.Mutex{}
				response := make([]any, 0)

				job := func(_ any) any {
					defer wg.Done()
					s, _ := test.executor(b.Context(), "1")

					return s
				}

				wp := tunny.NewFunc(runtime.GOMAXPROCS(0), job)
				wp.SetSize(50)

				b.ReportAllocs()
				b.SetParallelism(concur)
				b.ResetTimer()

				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						wg.Add(1)
						s := wp.Process("")

						mu.Lock()
						response = append(response, s)
						mu.Unlock()
					}
				})

				wp.Close()
				wg.Wait()

				b.StopTimer()
			})
		}
	}
}
