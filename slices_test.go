package slices

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/zhangyunhao116/pdqsort"
)

func TestSlices(t *testing.T) {
	fuzzTestSort(t, func(data []int) {
		Sort(data, func(a, b int) bool {
			return a < b
		})
	})
}

func fuzzTestSort(t *testing.T, f func(data []int)) {
	const times = 2048
	randomTestTimes := rand.Intn(times)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := rand.Intn(times)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		v2 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			randomValue := rand.Intn(randomLenth)
			v1[j] = randomValue
			v2[j] = randomValue
		}
		sort.Ints(v1)
		f(v2)
		for idx := range v1 {
			if v1[idx] != v2[idx] {
				t.Fatal("invalid sort:", idx, v1[idx], v2[idx])
			}
		}
	}
}

var sizes = []int{1 << 6, 1 << 8, 1 << 10, 1 << 12, 1 << 16}

type benchTask struct {
	name string
	f    func([]int)
}

var benchTasks = []benchTask{
	{
		name: "pdqsort",
		f: func(i []int) {
			Sort(i, func(a, b int) bool {
				return a < b
			})
		},
	},
	{
		name: "pdqsortDirectly",
		f: func(i []int) {
			pdqsort.Slice(i)
		},
	},
	{
		name: "stdsort",
		f: func(i []int) {
			sort.Slice(i, func(a, b int) bool {
				return i[a] < i[b]
			})
		},
	},
}

func benchmarkBase(b *testing.B, dataset func(x []int)) {
	for _, size := range sizes {
		for _, task := range benchTasks {
			b.Run(fmt.Sprintf(task.name+"_%d", size), func(b *testing.B) {
				b.StopTimer()
				for i := 0; i < b.N; i++ {
					data := make([]int, size)
					dataset(data)
					b.StartTimer()
					task.f(data)
					b.StopTimer()
				}
			})
		}
	}
}

func BenchmarkRandom(b *testing.B) {
	benchmarkBase(b, func(x []int) {
		for i := range x {
			x[i] = rand.Int()
		}
	})
}

func BenchmarkSorted(b *testing.B) {
	benchmarkBase(b, func(x []int) {
		for i := range x {
			x[i] = i
		}
	})
}
