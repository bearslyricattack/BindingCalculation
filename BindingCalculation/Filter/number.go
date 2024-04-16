package Filter

import "BindingCalculation/DataStructure"

// FilterByCost 根据指定费用过滤英雄
func FilterByCost(legends []DataStructure.Legend, costThreshold int) []DataStructure.Legend {
	var filteredLegends []DataStructure.Legend

	for _, legend := range legends {
		if legend.Cost < costThreshold {
			filteredLegends = append(filteredLegends, legend)
		}
	}

	return filteredLegends
}
