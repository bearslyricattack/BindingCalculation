package main

import (
	cache "BindingCalculation/Cache"
	"BindingCalculation/Calulation"
	"BindingCalculation/DataStructure"
	"BindingCalculation/Dynamic"
	"BindingCalculation/File"
	"BindingCalculation/Filter"
	"BindingCalculation/Out"
	"BindingCalculation/Sort"
	"context"
	_ "context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

var binding []DataStructure.Binding
var legends []DataStructure.Legend
var middles []DataStructure.Middle

func init() {
	//读取 Excel 文件到结构体数组
	//读取羁绊列表
	data, err := File.ReadExcel("羁绊.xlsx", reflect.TypeOf(DataStructure.Binding{}))
	if err != nil {
		fmt.Printf("Error reading Excel File: %v\n", err)
		os.Exit(1)
	}
	binding = data.([]DataStructure.Binding)
	//读取英雄列表
	data1, err := File.ReadExcel("英雄.xlsx", reflect.TypeOf(DataStructure.Legend{}))
	if err != nil {
		fmt.Printf("Error reading Excel File: %v\n", err)
		os.Exit(1)
	}
	legends = data1.([]DataStructure.Legend)
	//读取英雄-羁绊列表
	data2, err := File.ReadExcel("英雄-羁绊.xlsx", reflect.TypeOf(DataStructure.Middle{}))
	if err != nil {
		fmt.Printf("Error reading Excel File: %v\n", err)
		os.Exit(1)
	}
	middles = data2.([]DataStructure.Middle)
}

func main() {
	//创建内存数据库
	baseClient := cache.New()

	//创建路由
	r := gin.Default()

	//尊者处理
	r.GET("/dynamic", func(c *gin.Context) {
		name1 := c.Query("name1")
		name2 := c.Query("name1")
		name3 := c.Query("name1")
		name4 := c.Query("name1")
		name5 := c.Query("name1")
		//尊者处理
		middles = Dynamic.AddMiddle(middles, "尊者", name1, name2, name3, name4, name5)
		c.JSON(http.StatusOK, gin.H{})
		return
	})

	//获取闭环羁绊
	r.GET("/binding", func(c *gin.Context) {
		//获取参数

		//人口约束
		number := c.Query("number")
		//费用约束
		cost := c.Query("cost")
		//转职约束，暂定3个转职
		item1 := c.Query("item1")
		num1 := c.Query("num1")
		item2 := c.Query("item2")
		num2 := c.Query("num2")
		item3 := c.Query("item3")
		num3 := c.Query("num3")

		//使用内存数据库缓存查询结果，如果存在直接返回，不存在才走下面的计算
		key := fmt.Sprintf("%s_%s_%s_%s_%s_%s_%s_%s", number, cost, item1, num1, item2, num2, item3, num3)

		fmt.Println("搜索的key为：%s", key)
		res, ok := baseClient.Get(context.Background(), key)
		if ok {
			c.IndentedJSON(200, res)
			return
		}
		//转换参数
		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("转换出错:", err)
			return
		}
		cos, err := strconv.Atoi(cost)
		if err != nil {
			fmt.Println("转换出错:", err)
			return
		}
		mmap := make(map[string]int)
		if len(item1) != 0 {
			number1, err := strconv.Atoi(num1)
			if err != nil {
				fmt.Println("转换出错:", err)
				return
			}
			mmap[item1] = number1
		}
		if len(item2) != 0 {
			number2, err := strconv.Atoi(num2)
			if err != nil {
				fmt.Println("转换出错:", err)
				return
			}
			mmap[item2] = number2
		}
		if len(item3) != 0 {
			number3, err := strconv.Atoi(num3)
			if err != nil {
				fmt.Println("转换出错:", err)
				return
			}
			mmap[item3] = number3
		}

		//过滤人口
		legends = Filter.FilterByCost(legends, cos)

		//计算所有的英雄组合
		combinations := Calulation.GenerateCombinations(num, legends)

		var heroCounts []Out.HeroCount
		// 遍历每种组合
		for _, combo := range combinations {
			bindingCounts := make(map[string]int)
			//增加前置条件
			// 将初始map的内容复制到当前map中
			for key, value := range mmap {
				bindingCounts[key] = value
			}
			res := Calulation.CalculateClosedBindings(Calulation.CalculateBindingsForCombo(bindingCounts, combo, middles), binding, mmap)
			heroCounts = append(heroCounts, Out.HeroCount{
				Hero:  combo,
				Count: res,
			})
		}
		//根据result排序
		result := Sort.TopFiveHeroes(heroCounts)

		//转换格式
		out := Out.PrintHeroCounts(result, middles, mmap)

		//存入内存数据库
		fmt.Println("存入的key为：%s", key)
		baseClient.Set(context.Background(), key, out, 0)

		c.IndentedJSON(200, out)
		return
	})
	r.Run(":8082")
}
