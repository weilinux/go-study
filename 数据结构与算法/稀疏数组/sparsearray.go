package main

import "fmt"

// ValModel 稀疏数组结构体
type ValModel struct {
	Row int //行
	Col int //列
	Val int //值
}

func main() {
	//原始数组  (五子棋 1：黑子 2：白子)
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[1][4] = 1
	chessMap[2][4] = 1
	chessMap[5][6] = 1
	chessMap[3][2] = 1
	chessMap[8][3] = 2
	chessMap[4][5] = 2
	chessMap[7][6] = 2
	chessMap[9][9] = 2
	chessMap[4][6] = 2

	//输出还原棋盘
	for _, rowVal := range chessMap {
		for _, colVal := range rowVal {
			switch colVal {
			case 1:
				fmt.Printf("白\t")
			case 2:
				fmt.Printf("黑\t")
			default:
				fmt.Printf("+\t")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	//转成稀缺数组    （只存储有意义的值  用行、列来定位并保存改值）
	var sparseArr []ValModel
	//存储该数据的最大行和列
	nodeArr := ValModel{
		Row: 11,
		Col: 11,
	}
	sparseArr = append(sparseArr, nodeArr)

	//整个数组循环
	for i, rowVal := range chessMap {
		for j, colVal := range rowVal {
			if colVal != 0 {
				//保存有意义的值
				val := ValModel{
					Row: i,
					Col: j,
					Val: colVal,
				}
				sparseArr = append(sparseArr, val)
			}
		}
	}

	//格式化稀疏数组
	for i, val := range sparseArr {
		fmt.Printf("索引\t row(行)\t col(列)\t val(值)\n")
		fmt.Printf("%d\t %d\t\t %d\t\t %d\n", i, val.Row, val.Col, val.Val)
	}
	fmt.Println()

	//取出最大行列
	nodeArr2 := sparseArr[0]

	//根据最大行列来开辟切片内存
	var chessMap2 [][]int
	chessMap2 = make([][]int, nodeArr2.Row)
	//将稀疏数组转成普通数组格式输出
	for i, v := range sparseArr {
		//跳过结构体切片的第一行（只用来存储最大行列）
		if i != 0 {
			//根据索引下标开辟最大内存
			for j := 0; j < nodeArr2.Row; j++ {
				//如果子切片内存不存在则开辟最大内存（在确定最大范围的情况下直接make完，不确定的情况下用append追加元素）
				if cap(chessMap2[j]) == 0 {
					chessMap2[j] = make([]int, nodeArr2.Col)
				}
			}
			//将结构体赋值给已经开辟完内存的切片
			chessMap2[v.Row][v.Col] = v.Val
		}
	}
	fmt.Println(chessMap2)
	fmt.Println()

	//将切片格式化输出 (还原棋盘)
	for _, rowVal := range chessMap2 {
		for _, colVal := range rowVal {
			switch colVal {
			case 1:
				fmt.Printf("白\t")
			case 2:
				fmt.Printf("黑\t")
			default:
				fmt.Printf("+\t")
			}
		}
		fmt.Println()
	}
}
