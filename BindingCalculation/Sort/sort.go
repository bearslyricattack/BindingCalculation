package Sort

import (
	"BindingCalculation/Out"
	"sort"
)

// ByCount 实现排序接口
type ByCount []Out.HeroCount

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

// TopFiveHeroes 返回按照 count 字段从高到低排序的前五个元素
func TopFiveHeroes(heroCounts []Out.HeroCount) []Out.HeroCount {
	// 对切片进行排序
	sort.Sort(ByCount(heroCounts))

	// 如果切片长度大于五，返回前五个元素，否则返回整个切片
	if len(heroCounts) > 5 {
		return heroCounts[:5]
	}
	return heroCounts
}
