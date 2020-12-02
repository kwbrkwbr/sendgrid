package internal

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

func Benchmark_echoHealthz(b *testing.B) {
	hcEcho()
}

func Benchmark_fiberHealthz(b *testing.B) {
	hcFiber()
}

func Benchmark_echoHealthzMulti(b *testing.B) {
	wg := &sync.WaitGroup{}   // WaitGroupの値を作る
	for i := 0; i < 10; i++ { // （例として）10回繰り返す
		wg.Add(1) // wgをインクリメント
		go func() {
			hcEcho()
			wg.Done() // 完了したのでwgをデクリメント
		}()
	}
	wg.Wait() // メインのgoroutineはサブgoroutine 10個が完了するのを待つ
}

func Benchmark_fiberHealthzMulti(b *testing.B) {
	wg := &sync.WaitGroup{}    // WaitGroupの値を作る
	for i := 0; i < 100; i++ { // （例として）10回繰り返す
		wg.Add(1) // wgをインクリメント
		go func() {
			hcFiber()
			wg.Done() // 完了したのでwgをデクリメント
		}()
	}
	wg.Wait() // メインのgoroutineはサブgoroutine 10個が完了するのを待つ
}

func hcEcho() {
	_, err := http.Get("http://127.0.0.1:1323/healthz")
	if err != nil {
		fmt.Println(err)
	}
}

func hcFiber() {
	_, err := http.Get("http://127.0.0.1:3000/healthz")
	if err != nil {
		fmt.Println(err)
	}
}

const tests int = 400

func Test_PlainTextFiber(t *testing.T) {
	b := time.Now()
	for i := 0; i < tests; i++ {
		hcFiber()
	}
	a := time.Now()

	t.Log(a.Sub(b))
}

func Test_PlainTextEcho(t *testing.T) {
	b := time.Now()
	for i := 0; i < tests; i++ {
		hcEcho()
	}
	a := time.Now()

	t.Log(a.Sub(b))
}
