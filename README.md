# 云顶之弈 金铲铲之战羁绊计算器：

# 1.设计意图：

s11 赛季加入动态的羁绊。并且会出现很多给纹章的奇遇和海克斯。（比如漫游训练师，混沌召唤等）

很容易让人猪脑过载，无法短时间内拼凑出好的闭环羁绊。

另外，市面上的羁绊计算器大多无法限定条件，无法限定人口，最高弈子的费用等数据，计算速度过慢。而且不能换赛季用。

# 2.工作流程：

## 2.1 羁绊数据导入：

每个赛季一次。

为了让此工具能在任意赛季使用，考虑不写死数据，而是从文件中读取。比如excel文件。

文件需要与项目在同一个住文件夹下。

## 2.2 输入约束：

### 2.2.1 尊者：

通过接口输入尊者羁绊

### 2.2.2 获取闭环羁绊：

暂时支持三个纹章及以下。

接口文档：

![image-20240416181505992](https://bearsblog.oss-cn-beijing.aliyuncs.com/img/image-20240416181505992.png)

获取的结果json已经自动排序，举例如下：

六人口，三费卡及以下，一个永恒之森转。

```
curl -X GET "http://localhost:8082/binding?number=6&cost=4&item1=永恒之森&num1=1"
[
    {
        "index": 0,
        "number": 10,
        "results": [
            {
                "Name": "卡兹克",
                "Cost": 1
            },
            {
                "Name": "雷克塞",
                "Cost": 1
            },
            {
                "Name": "锐雯",
                "Cost": 2
            },
            {
                "Name": "千珏",
                "Cost": 2
            },
            {
                "Name": "纳尔",
                "Cost": 2
            },
            {
                "Name": "索拉卡",
                "Cost": 3
            }
        ],
        "binding": "{斗士: 2, 永恒之森: 4, 剪纸仙灵: 1, 武仙子: 2, 灵魂莲华: 1, 护卫: 1, 天将: 2, 死神: 2}"
    },
    {
        "index": 1,
        "number": 10,
        "results": [
            {
                "Name": "锐雯",
                "Cost": 2
            },
            {
                "Name": "妮蔻",
                "Cost": 2
            },
            {
                "Name": "亚托克斯",
                "Cost": 2
            },
            {
                "Name": "纳尔",
                "Cost": 2
            },
            {
                "Name": "俄洛伊",
                "Cost": 3
            },
            {
                "Name": "索拉卡",
                "Cost": 3
            }
        ],
        "binding": "{剪纸仙灵: 1, 斗士: 2, 墨之影: 1, 护卫: 2, 天将: 2, 法师: 2, 幽魂: 2, 武仙子: 2, 永恒之森: 2}"
    },
    {
        "index": 2,
        "number": 10,
        "results": [
            {
                "Name": "雷克塞",
                "Cost": 1
            },
            {
                "Name": "千珏",
                "Cost": 2
            },
            {
                "Name": "亚托克斯",
                "Cost": 2
            },
            {
                "Name": "纳尔",
                "Cost": 2
            },
            {
                "Name": "俄洛伊",
                "Cost": 3
            },
            {
                "Name": "永恩",
                "Cost": 3
            }
        ],
        "binding": "{法师: 1, 夜幽: 1, 斗士: 2, 永恒之森: 4, 灵魂莲华: 1, 死神: 2, 墨之影: 1, 护卫: 2, 幽魂: 2}"
    },
    {
        "index": 3,
        "number": 10,
        "results": [
            {
                "Name": "雷克塞",
                "Cost": 1
            },
            {
                "Name": "千珏",
                "Cost": 2
            },
            {
                "Name": "亚托克斯",
                "Cost": 2
            },
            {
                "Name": "纳尔",
                "Cost": 2
            },
            {
                "Name": "俄洛伊",
                "Cost": 3
            },
            {
                "Name": "佐伊",
                "Cost": 3
            }
        ],
        "binding": "{永恒之森: 4, 法师: 2, 吉星: 1, 斗士: 2, 灵魂莲华: 1, 死神: 1, 墨之影: 1, 护卫: 2, 幽魂: 2, 剪纸仙灵: 1}"
    },
    {
        "index": 4,
        "number": 10,
        "results": [
            {
                "Name": "雷克塞",
                "Cost": 1
            },
            {
                "Name": "千珏",
                "Cost": 2
            },
            {
                "Name": "妮蔻",
                "Cost": 2
            },
            {
                "Name": "亚托克斯",
                "Cost": 2
            },
            {
                "Name": "纳尔",
                "Cost": 2
            },
            {
                "Name": "俄洛伊",
                "Cost": 3
            }
        ],
        "binding": "{灵魂莲华: 1, 死神: 1, 永恒之森: 4, 护卫: 2, 法师: 2, 天将: 1, 墨之影: 1, 幽魂: 2, 斗士: 2}"
    }
]%                                               
```

# 3.设计：

## 3.1 数据结构：

英雄与羁绊之间是多对多关系，采用中间表形式维护

### 3.1.1 英雄：

名称

费用

### 3.1.2 羁绊：

名称

层级数 数组

### 3.1.3 英雄-羁绊中间表：

英雄名

羁绊名

## 3.2 系统架构：

### 3.2.1 整体架构：

整体架构图如下：

![image-20240416182021799](https://bearsblog.oss-cn-beijing.aliyuncs.com/img/image-20240416182021799.png)

下面分模块介绍。

### 3.2.2 cache：

缓存模块。

为了防止重复计算导致用户体验过差的问题，使用ornncache内存数据库缓存数据。

### 3.2.3 caculation：

羁绊计算模块。

使用评估函数评估每一个阵容的闭合程度。

评估函数待优化。

```
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
```

### 3.2.4 datastructure：

数据结构。

### 3.2.5 dynamic：

动态羁绊处理。例如s11的尊者。

### 3.2.6 file：

文件处理。

### 3.2.7 filter:

约束条件过滤，例如过滤卡的价值和特定的英雄。

### 3.2.8 out：

输出。使用gin的IndentedJSON方法在输出的时候整理json格式的文件。

### 3.2.9 sort：

排序模块。使用go自带的sort函数进行排序。

# 4.总结：

一次全新的尝试，希望大家都能把自己的热爱，生活与工作结合在一起，创造出更多更好的作品。

> 青鸟一去 不觉此间 百年悄然
>
> 自有白鹿 踏歌如梦来
>
> -- 《万梦星》黄诗扶

 
