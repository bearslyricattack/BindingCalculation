package Out

import (
	"BindingCalculation/Calulation"
	"BindingCalculation/DataStructure"
	"fmt"
)

type HeroCount struct {
	Hero  []DataStructure.Legend
	Count int
}

// 将 map 转换为字符串
func mapToString(m map[string]int) string {
	result := "{"
	for key, value := range m {
		result += fmt.Sprintf("%s: %d, ", key, value)
	}
	// 去掉末尾的逗号和空格
	if len(result) > 1 {
		result = result[:len(result)-2]
	}
	result += "}"
	return result
}

// IncreaseBindingCount 根据输入的字符串和数字增加 map 中对应值
func IncreaseBindingCount(bindings map[string]int, name string, count int) map[string]int {
	bindings[name] += count
	return bindings
}

func mergeMaps(map1, map2 map[string]int) map[string]int {
	mergedMap := make(map[string]int)

	// 遍历第一个map，将键和值复制到合并后的map中
	for key, value := range map1 {
		mergedMap[key] = value
	}

	// 遍历第二个map，如果键已存在，则将值相加；如果键不存在，则直接复制到合并后的map中
	for key, value := range map2 {
		if _, ok := mergedMap[key]; ok {
			mergedMap[key] += value
		} else {
			mergedMap[key] = value
		}
	}

	return mergedMap
}

// PrintHeroCounts1 输出 HeroCount 切片的内
func PrintHeroCounts1(heroCounts []HeroCount, middleTable []DataStructure.Middle) {
	legendNames := make(map[string]int)
	for i, heroCount := range heroCounts {
		hero := heroCount.Hero
		legendNames = Calulation.CalculateBindingsForCombo(legendNames, hero, middleTable)
		fmt.Printf("第%d名:\n", i+1)
		fmt.Printf("评估函数计算羁绊得分为: %d\n", heroCount.Count)
		fmt.Println("包含英雄:")
		for _, hero := range heroCount.Hero {
			fmt.Printf("名称: %s,费用: %d\n", hero.Name, hero.Cost)
		}
		fmt.Println("最终羁绊:")
		fmt.Println(mapToString(legendNames))
		fmt.Println("")
		legendNames = make(map[string]int)
	}
}

// OutputData 表示API要返回的数据
type OutputData struct {
	//排名
	Index int `json:"index"`
	//羁绊得分
	Number int `json:"number"`
	//包含英雄
	Results []DataStructure.Legend `json:"results"`
	//最终羁绊
	Binding string `json:"binding"`
}

// PrintHeroCounts 将输出转换为一个结构体，以便被Gin框架的API返回
func PrintHeroCounts(heroCounts []HeroCount, middleTable []DataStructure.Middle, mmap map[string]int) []OutputData {
	legendNames := make(map[string]int)
	var outputs []OutputData
	for i, heroCount := range heroCounts {
		hero := heroCount.Hero
		legendNames = Calulation.CalculateBindingsForCombo(legendNames, hero, middleTable)
		outputs = append(outputs, OutputData{
			Index:   i,
			Number:  heroCount.Count,
			Results: heroCount.Hero,
			Binding: mapToString(mergeMaps(legendNames, mmap)),
		})
		legendNames = make(map[string]int)
	}
	return outputs
}
