package Filter

import "BindingCalculation/DataStructure"

// RemoveHeroes 根据指定英雄名称去掉英雄，并返回新切片
func RemoveHeroes(allHeroes []DataStructure.Legend, namesToRemove ...string) []DataStructure.Legend {
	var result []DataStructure.Legend

	// 创建一个 map 来存储需要去掉的英雄名称，以便更快地检查
	heroesToRemove := make(map[string]bool)
	for _, name := range namesToRemove {
		heroesToRemove[name] = true
	}

	// 遍历所有英雄，将不在需要去掉的英雄名称中的英雄添加到结果切片中
	for _, hero := range allHeroes {
		if !heroesToRemove[hero.Name] {
			result = append(result, hero)
		}
	}

	return result
}
