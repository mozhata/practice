package performancetest

import (
	"strconv"
	"testing"
)

/*go test -bench=. -benchmem*/
var (
	slic5, slic50, slic250, slic1250, slic7500, slic750000 []string
	map5, map50, map250, map1250, map7500, map750000       map[string]int
)

func TestMain(m *testing.M) {
	slic5 = buildSlic(5)
	map5 = buildMap(5)

	m.Run()
}

func buildSlic(n int) []string {
	result := make([]string, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result
}

func buildMap(n int) map[string]int {
	result := make(map[string]int, n)
	for i := 0; i < n; i++ {
		result[strconv.Itoa(i)] = i
	}
	return result
}

func rangeSlic(slic []string) {
	for i, v := range slic {
		_, _ = i, v
	}
}

func rangeMap(dic map[string]int) {
	for key, val := range dic {
		_, _ = key, val
	}
}

func BenchmarkBuildSlice5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildSlic(5)
	}
}
func BenchmarkBuildMap5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildMap(5)
	}
}
func BenchmarkRangeSlic5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeSlic(slic5)
	}
}
func BenchmarkRangeMap5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeMap(map5)
	}
}

func BenchmarkBuildSlice50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildSlic(50)
	}
}
func BenchmarkBuildMap50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildMap(50)
	}
}
func BenchmarkRangeSlic50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeSlic(slic50)
	}
}
func BenchmarkRangeMap50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeMap(map50)
	}
}
func BenchmarkBuildSlice250(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildSlic(250)
	}
}
func BenchmarkBuildMap250(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildMap(250)
	}
}
func BenchmarkRangeSlic250(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeSlic(slic250)
	}
}
func BenchmarkRangeMap250(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeMap(map250)
	}
}
func BenchmarkBuildSlice1250(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildSlic(1250)
	}
}
func BenchmarkBuildMap1250(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildMap(1250)
	}
}
func BenchmarkRangeSlic1250(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeSlic(slic1250)
	}
}
func BenchmarkRangeMap1250(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeMap(map1250)
	}
}
func BenchmarkBuildSlice7500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildSlic(7500)
	}
}
func BenchmarkBuildMap7500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildMap(7500)
	}
}
func BenchmarkRangeSlic7500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeSlic(slic7500)
	}
}
func BenchmarkRangeMap7500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeMap(map7500)
	}
}
func BenchmarkBuildSlice750000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildSlic(750000)
	}
}
func BenchmarkBuildMap750000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildMap(750000)
	}
}
func BenchmarkRangeSlic750000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeSlic(slic750000)
	}
}
func BenchmarkRangeMap750000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeMap(map750000)
	}
}
