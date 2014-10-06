package main

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

type Vector []float64

func (v Vector) String() string {
	sb := new(bytes.Buffer)

	n, trunc := len(v), ""
	if n > 8 {
		n, trunc = 8, "..."
	}
	for i := 0; i < n; i++ {
		fmt.Fprintf(sb, "%.4g ", v[i])
	}
	fmt.Fprint(sb, trunc)

	return sb.String()
}

func Convolve(u, v Vector) (w Vector) {
	n := len(u) + len(v) - 1
	w = make(Vector, n)
	for k := 0; k < n; k++ {
		w[k] = mul(u, v, k)
	}
	return
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func mul(u, v Vector, k int) (res float64) {
	n := min(k+1, len(u))
	j := min(k, len(v)-1)

	for i := k - j; i < n; i, j = i+1, j-1 {
		res += u[i] * v[j]
	}
	return
}

func ConcurrentConvolve(u, v Vector) (w Vector) {
	n := len(u) + len(v) - 1
	w = make(Vector, n)
	size := max(1, 1<<20/n)

	wg := new(sync.WaitGroup)
	wg.Add(1 + (n-1)/size)

	for i := 0; i < n && i >= 0; i += size {
		j := i + size
		if j > n || j < 0 {
			j = n
		}
		// These goroutines share memory, but only for reading.
		go func(i, j int) {
			for k := i; k < j; k++ {
				w[k] = mul(u, v, k)
			}
			wg.Done()
		}(i, j)
	}
	wg.Wait()
	return
}

func init() {
	log.SetFlags(0) // no extra info in log messages
	//log.SetOutput(ioutil.Discard) // turns off logging

	numcpu := runtime.NumCPU()
	log.Println("CPU count:", numcpu)
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}

func main() {
	v := make(Vector, 32*1000)
	for i := range v {
		v[i] = 1
	}

	before := time.Now()
	w := Convolve(v, v)
	fmt.Println("time:", time.Now().Sub(before))

	before = time.Now()
	w = ConcurrentConvolve(v, v)
	fmt.Println("time:", time.Now().Sub(before))

	fmt.Println(w)
}
