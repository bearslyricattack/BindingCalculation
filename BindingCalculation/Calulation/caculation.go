package Calulation

import "BindingCalculation/DataStructure"

// GenerateCombinations 生成所有可能的英雄组合
func GenerateCombinations(n int, legends []DataStructure.Legend) [][]DataStructure.Legend {
	var combinations [][]DataStructure.Legend

	var backtrack func(start int, path []DataStructure.Legend)
	backtrack = func(start int, path []DataStructure.Legend) {
		if len(path) == n {
			combinations = append(combinations, append([]DataStructure.Legend{}, path...))
			return
		}
		for i := start; i < len(legends); i++ {
			backtrack(i+1, append(path, legends[i]))
		}
	}

	backtrack(0, []DataStructure.Legend{})
	return combinations
}

// CalculateBindingsForCombo 计算给定英雄组合的羁绊数量，并返回羁绊名称和数量的映射
func CalculateBindingsForCombo(bindingCounts map[string]int, combo []DataStructure.Legend, middleTable []DataStructure.Middle) map[string]int {
	legendNames := make(map[string]bool)

	// 统计每个英雄的出现次数
	for _, hero := range combo {
		legendNames[hero.Name] = true
	}

	// 遍历每个英雄
	for _, hero := range combo {
		// 查找与当前英雄相关联的羁绊
		for _, middle := range middleTable {
			if middle.LegendName == hero.Name {
				bindingCounts[middle.BindingName]++
			}
		}
	}

	return bindingCounts
}

// CalculateClosedBindings 计算羁绊闭合个数
func CalculateClosedBindings(bindingCounts map[string]int, binding []DataStructure.Binding, tag map[string]int) int {
	closedCount := 0
	for _, value := range binding {
		bind := bindingCounts[value.Name]
		//天将特别判断
		if value.Name == "天将" {
			if bind > 4 {
				closedCount += 2
				continue
			}
			if bind > 2 {
				closedCount += 1
				continue
			}
		}
		//天龙之子特别判断
		if value.Name == "天龙之子" {
			if bind > 3 {
				closedCount += 2
				continue
			}
			if bind > 2 {
				closedCount += 1
				continue
			}
		}
		// 一般情况
		if bind != 0 {
			if bind >= value.NumberOne && bind < value.NumberTwo {
				closedCount += 1
			}
			if bind >= value.NumberTwo && bind < value.NumberThree {
				closedCount += 2
			}
			if bind >= value.NumberThree && bind < value.NumberFour {
				closedCount += 3
			}
			if bind >= value.NumberFour && bind < value.NumberFive {
				closedCount += 4
			}
			if bind >= value.NumberFive && bind < value.NumberSix {
				closedCount += 5
			}
			if bind >= value.NumberSix {
				closedCount += 6
			}
		}
		//不等于零 说明是带纹章的羁绊 优先级会提高
		if tag[value.Name] != 0 {
			if bind != 0 {
				if bind >= value.NumberOne && bind < value.NumberTwo {
					closedCount += 3
				}
				if bind >= value.NumberTwo && bind < value.NumberThree {
					closedCount += 4
				}
				if bind >= value.NumberThree && bind < value.NumberFour {
					closedCount += 5
				}
				if bind >= value.NumberFour && bind < value.NumberFive {
					closedCount += 6
				}
				if bind >= value.NumberFive && bind < value.NumberSix {
					closedCount += 7
				}
				if bind >= value.NumberSix {
					closedCount += 8
				}
			}
		}
	}
	return closedCount
}
