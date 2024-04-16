package Dynamic

import "BindingCalculation/DataStructure"

//动态羁绊处理：例如s11的尊者

// AddMiddle 添加新的英雄羁绊对应到切片中并返回
func AddMiddle(middleSlice []DataStructure.Middle, bindingName string, legendNames ...string) []DataStructure.Middle {
	for _, legendName := range legendNames {
		middleSlice = append(middleSlice, DataStructure.Middle{LegendName: legendName, BindingName: bindingName})
	}
	return middleSlice
}
