package File

import (
	"github.com/tealeg/xlsx"
	"reflect"
)

//文件操作

// ReadExcel 读取 Excel 文件并映射为特定的结构体数组
func ReadExcel(filename string, dataType reflect.Type) (interface{}, error) {
	// 创建结构体切片
	slicePtr := reflect.New(reflect.SliceOf(dataType)).Elem()

	// 打开 Excel 文件
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	// 获取结构体字段数量
	numFields := dataType.NumField()

	// 获取结构体字段名到字段索引的映射
	fieldMap := make(map[string]int)
	for i := 0; i < numFields; i++ {
		fieldMap[dataType.Field(i).Name] = i
	}

	// 遍历 Excel 文件中的每个工作表
	for _, sheet := range xlFile.Sheets {
		// 遍历工作表中的每一行
		for rowIndex, row := range sheet.Rows {
			// 忽略表头行
			if rowIndex == 0 {
				continue
			}

			// 创建结构体实例
			structPtr := reflect.New(dataType).Elem()

			// 获取每行中的单元格
			for i, cell := range row.Cells {
				// 获取单元格的值
				value := cell.String()

				// 查找字段名对应的结构体字段索引
				fieldIndex, found := fieldMap[sheet.Cell(0, i).String()]
				if !found {
					continue // 如果字段名不在结构体中，则忽略该列
				}

				// 将值设置到结构体字段中
				fieldPtr := structPtr.Field(fieldIndex)
				switch fieldPtr.Kind() {
				case reflect.String:
					fieldPtr.SetString(value)
				case reflect.Int:
					number, err := cell.Int()
					if err != nil {
						return nil, err
					}
					fieldPtr.SetInt(int64(number))
				}
			}

			// 将映射好的结构体添加到数组中
			slicePtr = reflect.Append(slicePtr, structPtr)
		}
	}

	return slicePtr.Interface(), nil
}
