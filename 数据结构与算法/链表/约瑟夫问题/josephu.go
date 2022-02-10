package main

import (
	"fmt"
)

// Node 创建节点结构/类型
type Node struct {
	No       int   //编号
	PreNote  *Node //父级节点 等于null则为第一个节点
	NextNode *Node //子级节点 等于null则为最后一个节点
}

// Link 创建链表结构
type Link struct {
	HeadNode *Node //头部节点（不存在数据，只保存头部节点信息关联后续节点）
	TailNode *Node //尾部节点（不存在数据，只保存尾部节点信息关联后续节点）
}

// Method 设计接口：
type Method interface {
	IsEmpty() bool          // 链表是否为空
	Add(n int)              // 从链表的头部添加节点
	Append(n []int)         // 从链表尾部添加节点
	Calculation(outNum int) // 模拟约瑟夫问题出队
	ShowList()              // 遍历所有节点
	Out(tempNode *Node)     // 删除指定的节点
}

// CreateNode 创建节点
func CreateNode(n int) (node *Node) {
	node = &Node{No: n, PreNote: nil, NextNode: nil}
	return
}

// CreateLink 创建空链表
func CreateLink() *Link {
	return &Link{}
}

// IsEmpty 链表是否为空
func (l *Link) IsEmpty() bool {
	if l.HeadNode == nil && l.TailNode == nil {
		return true
	}
	return false
}

// Add 从链表的头部添加节点
func (l *Link) Add(n int) {
	//创建新节点
	node := CreateNode(n)
	//如果链表为空则直接将头部节点和尾部节点都指向新节点,且该节点的父级节点和子级节点都指向自己（形成单节点闭环）
	if l.IsEmpty() {
		l.HeadNode, l.TailNode, node.PreNote, node.NextNode = node, node, node, node
	} else {
		//1.将原头部节点的父级节点指向新节点
		//2.将原头部节点指向新节点的子级节点
		//3.将新节点指向链表的头部节点
		//4.将链表尾部节点的子级节点指向新节点
		//5.将新节点的父级节点指向链表尾部节点
		l.HeadNode.PreNote, node.NextNode, l.HeadNode, l.TailNode.NextNode, node.PreNote = node, l.HeadNode, node, node, l.TailNode
	}
}

// Append 从链表尾部添加节点
func (l *Link) Append(n []int) {
	//如果链表为空则直接从头部添加
	if l.IsEmpty() {
		l.Add(n[0])
	}

	if len(n) > 1 {
		//剔除第一个元素
		n = append(n[:0], n[1:]...)

		for _, v := range n {
			//过滤默认值0
			if v > 0 {
				//创建新节点
				node := CreateNode(v)

				//获取链表尾部节点
				tempNode := l.TailNode

				//1.链表最后一个节点的子级节点指向新节点
				//2.链表尾部节点指向新节点
				//3.新节点的父级节点指向链表最后一个节点
				//4.链表头部节点的父级节点指向新节点
				//5.新节点的子级节点指向链表头部节点
				tempNode.NextNode, l.TailNode, node.PreNote, l.HeadNode.PreNote, node.NextNode = node, node, tempNode, node, l.HeadNode
			}
		}
	}
}

// ShowList 遍历所有节点
func (l *Link) ShowList() {
	//链表头部节点
	fmt.Println("链表头部节点：", l.HeadNode)

	//链表尾部节点
	fmt.Println("链表尾部节点：", l.TailNode)

	//定义变量用于存储递归子级节点
	tempNode := l.HeadNode

	//默认索引位为0
	index := 0

	//循环子级节点
	fmt.Println("原数据-------")
	for {
		fmt.Printf("该节点位置：%d 值：%v\n", index, tempNode)

		//取出该节点的子级节点
		tempNode = tempNode.NextNode

		//索引前进一位
		index++

		if tempNode == l.HeadNode {
			break
		}
	}
	fmt.Println()
}

// Out 删除指定的节点
func (l *Link) Out(tempNode *Node) {
	//将辅助节点指向临时节点(为了不影响临时节点进行循环出队)
	assistNode := tempNode

	//如果链表尾部节点等于指定节点
	if l.TailNode == assistNode {
		//将链表尾部节点指向该节点的父级节点
		l.TailNode = assistNode.PreNote
	}

	//如果链表头部节点等于指定节点
	if l.HeadNode == assistNode {
		//将链表头部节点指向该节点的子级节点
		l.HeadNode = assistNode.NextNode
	}

	//1.将该节点的父级节点的子级节点指向该节点的子级节点
	assistNode.PreNote.NextNode = assistNode.NextNode
	//2.将该节点的子级节点的父级节点指向该节点的父级节点
	assistNode.NextNode.PreNote = assistNode.PreNote
}

// Calculation 模拟约瑟夫问题出队
func (l *Link) Calculation(outNum int) {
	//定义变量用于存储递归子级节点
	tempNode := l.HeadNode

	//存储出队顺序 (单维，不判断圈数)
	singleQueue := make([]int, 0)

	//存储出队顺序 (多维，存储出队圈数)
	manyQueue := make([][]int, 0)

	//圈数默认为1
	circleNum := 1

	for {
		//循环报数
		for j := 1; j < outNum; j++ {
			//取出该节点的子级节点
			tempNode = tempNode.NextNode

			//如果链表头部节点等于当前节点则代表已循环完一整圈
			if l.HeadNode == tempNode {
				circleNum++
			}
		}

		//如果出队顺序(多维)容量为0或者容量已经大于当前切片长度时则追加长度
		if cap(manyQueue) == 0 || circleNum > len(manyQueue) {
			manyQueue = append(manyQueue, []int{})
		}

		//保存出队顺序(单维)
		singleQueue = append(singleQueue, tempNode.No)

		//保存出队顺序，根据循环圈数和出队序号(多维)
		manyQueue[len(manyQueue)-1] = append(manyQueue[len(manyQueue)-1], tempNode.No)

		//出队
		l.Out(tempNode)

		//该节点需向前进一位(此时该节点已出队)
		tempNode = tempNode.NextNode

		//如果头部等于尾部则队列只剩最后一个节点
		if l.HeadNode == l.TailNode {
			//保存出队顺序(单维)
			singleQueue = append(singleQueue, tempNode.No)

			//保存出队顺序，根据循环圈数和出队序号(多维)
			manyQueue[len(manyQueue)-1] = append(manyQueue[len(manyQueue)-1], tempNode.No)

			goto complete
		}
	}

	//出队完成
complete:
	fmt.Println("出队顺序：", singleQueue)

	for i := 0; i <= len(manyQueue)-1; i++ {
		fmt.Printf("出队圈数：%v 顺序：%v\n", i+1, manyQueue[i])
	}
}

func main() {
	//创建链表
	link := CreateLink()

	//实现接口
	var m Method = link

	fmt.Println("请输入入队个数：")
	var enterNum int
	_, err := fmt.Scanln(&enterNum)
	if err != nil {
		fmt.Println(err)
	}

	s := make([]int, enterNum)
	for i := 0; i < enterNum; i++ {
		s[i] = i + 1
	}
	fmt.Println(s)

	//添加节点
	m.Append(s)

	fmt.Println("请输入出队数字：")
	var outNum int
	_, err = fmt.Scanln(&outNum)
	if err != nil {
		fmt.Println(err)
	}

	//队列初始
	link.ShowList()

	//出队
	m.Calculation(outNum)
}
