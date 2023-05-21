package handler

import (
	"github.com/Ki4EH/lib-service/catalog/entities"
	"math"
	"sort"
)

func cosR(s1 string, s2 string) float64 {

	var sum1, sum2, sum3 int
	n := int(math.Min(float64(len(s1)), float64(len(s2))))
	for i := 0; i < n; i++ {
		sum1 += (int(rune(s1[i])) + 200) * (200 + int(rune(s2[i])))
		sum2 += (int(rune(s1[i])) + 200) * (200 + int(rune(s1[i])))
		sum3 += (200 + int(rune(s2[i]))) * (200 + int(rune(s2[i])))
	}
	cos := float64(sum1) / (math.Sqrt(float64(sum2)) * math.Sqrt(float64(sum3)))
	return cos
}

func maxEl(m map[float64][]entities.Book) float64 {
	var i float64
	keys := make([]float64, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	i = keys[len(m)-1]
	return i
}
