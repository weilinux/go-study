package main

import (
	"errors"
	"fmt"
)

// Node 创建节点结构/类型
type Node struct {
	Data     interface{} // 数据
	NextNode *Node       // 子级节点
}

// Link 创建链表结构
type Link struct {
	HeadNode *Node // 头部节点（不存在数据，只保存头部节点信息关联后续节点）
}

// Method 设计接口：
type Method interface {
	IsEmpty() bool                       // 链表是否为空
	Length() (count int)                 // 获取链表节点长度
	Contain(v interface{}) (index int)   // 查询指定值的节点位置
	Add(v interface{})                   // 从链表的头部添加节点
	Append(v interface{})                // 从链表尾部添加节点
	Insert(index int, v interface{})     // 在链表指定位置添加节点
	RemoveAtData(v interface{})          // 删除指定值的节点
	RemoveAtIndex(index int) (err error) // 删除指定位置的节点
	ShowList()                           // 遍历所有节点
}

// CreateNode 创建节点
func CreateNode(v interface{}) (node *Node) {
	node = &Node{Data: v, NextNode: nil}
	return
}

// CreateLink 创建空链表
func CreateLink() *Link {
	return &Link{}
}

// IsEmpty 链表是否为空
func (l *Link) IsEmpty() bool {
	if l.HeadNode == nil {
		return true
	}
	return false
}

// Length 返回链表节点长度
func (l *Link) Length() (count int) {
	// 定义临时变量存储链表头结点
	tempNode := l.HeadNode

	// 循环位移到指定索引的节点位置
	for count = 0; tempNode != nil; count++ {
		// 如果子级节点不为nil则一直递归
		tempNode = tempNode.NextNode
	}
	return
}

// Contain 查询指定值的节点位置
func (l *Link) Contain(v interface{}) (index int) {
	// 定义变量用于存储递归子级节点
	tempNode := l.HeadNode

	// 默认索引位为0
	index = 0

	// 循环位移到指定索引的节点位置
	for tempNode != nil {
		// 如果该节点的值等于指定的值则返回true，否则继续位移子级节点
		if tempNode.Data == v {
			return
		}

		// 取出该节点的子级节点
		tempNode = tempNode.NextNode

		// 索引前进一位
		index++
	}

	// 所有节点扫描完还是不能匹配上则返回-1
	return -1
}

// Add 从链表的头部添加节点
func (l *Link) Add(v interface{}) {
	// 创建新节点
	node := CreateNode(v)

	// 如果链表为空则直接将头部节点指向新节点
	if l.IsEmpty() {
		l.HeadNode = node
	} else {
		// 将原头部节点指向新节点的子级节点上，再将新节点指向链表的头部节点上
		node.NextNode, l.HeadNode = l.HeadNode, node
	}
}

// Append 从链表尾部添加节点
func (l *Link) Append(v interface{}) {
	// 如果链表为空则直接从头部添加
	if l.IsEmpty() {
		l.Add(v)
	} else {
		// 创建新节点
		node := CreateNode(v)

		// 定义变量用于存储递归子级节点
		tempNode := l.HeadNode

		// 循环子级节点
		for tempNode.NextNode != nil {
			// 如果子级节点不为nil则一直位移
			tempNode = tempNode.NextNode
		}

		// 链表最后一个节点的子级节点指向新节点
		tempNode.NextNode = node
	}
}

// Insert 在链表指定位置添加节点
func (l *Link) Insert(index int, v interface{}) {
	// 如果索引小于0则从头部添加
	if index < 0 {
		l.Add(v)
	} else if index > l.Length() {
		// 如果索引大于链表长度则从尾部添加
		l.Append(v)
	} else {
		// 定义变量用于存储递归子级节点
		tempNode := l.HeadNode

		// 循环位移到指定索引的节点位置
		for i := 0; i < index-1; i++ {
			// 取出该节点的子级节点
			tempNode = tempNode.NextNode
		}

		// 创建新节点
		node := CreateNode(v)

		// 将指定节点的子级节点指向新节点的子级节点上，再将新节点指向指定节点的子级节点上
		node.NextNode, tempNode.NextNode = tempNode.NextNode, node
	}
}

// RemoveAtData 删除指定值的节点
func (l *Link) RemoveAtData(v interface{}) {
	// 定义变量用于存储递归子级节点
	tempNode := l.HeadNode

	// 如果该节点的值等于指定值则直接接将链表的头部节点指向该节点的子级节点
	if tempNode.Data == v {
		l.HeadNode = tempNode.NextNode
	} else {
		// 循环子级节点
		for tempNode.NextNode != nil {
			// 如果该节点指向的子级节点的值等于指定值则直接接将链表的头部节点指向该节点的子级节点
			if tempNode.NextNode.Data == v {
				tempNode.NextNode = tempNode.NextNode.NextNode
			} else {
				// 取出该节点的子级节点
				tempNode = tempNode.NextNode
			}
		}
	}
}

// RemoveAtIndex 删除指定位置的节点
func (l *Link) RemoveAtIndex(index int) (err error) {
	// 定义变量用于存储递归子级节点
	tempNode := l.HeadNode

	// 如果index小于0则直接用头部节点的子级节点指向链表的头部节点 （删除原头部节点）
	if index < 0 {
		l.HeadNode = tempNode.NextNode
	} else if index > l.Length() {
		err = errors.New("超出链表长度")
		return
	} else {
		// 循环位移至指定位置节点前一位，并且该节点的子节点不可能为nil (改节点的子节点指向要删除的节点)
		for i := 0; i != index-1 && tempNode.NextNode != nil; i++ {
			// 取出该节点的子级节点
			tempNode = tempNode.NextNode
		}

		// 指定位置节点前一位节点的子级节点指向要删除的节点的子级节点
		tempNode.NextNode = tempNode.NextNode.NextNode
	}

	return
}

// ShowList 遍历所有节点
func (l *Link) ShowList() {
	// 如果链表不为空则打印子级节点信息
	if !l.IsEmpty() {
		// 定义变量用于存储递归子级节点
		tempNode := l.HeadNode

		// 默认索引位为0
		index := 0

		// 循环子级节点
		for tempNode != nil {
			fmt.Printf("该节点位置：%d 值：%v\n", index, tempNode.Data)

			// 取出该节点的子级节点
			tempNode = tempNode.NextNode

			// 索引前进一位
			index++
		}
	} else {
		fmt.Printf("link empty")
	}
	fmt.Println()
}

func main() {
	// 创建链表
	link := CreateLink()

	// 实现接口
	var m Method = link

	// 头部加入节点
	m.Add("从头部添加的")

	// 尾部添加节点
	m.Append("从尾部添加的")
	m.Append("从尾部再添加一个")

	// 链表长度
	fmt.Printf("链表长度为：%d\n", m.Length())

	// 显示链表
	m.ShowList()

	// 指定索引2的位置添加
	m.Insert(2, "指定索引2的位置添加")
	m.ShowList()

	// 获取指定值位置
	fmt.Println("该值位置：", m.Contain("从尾部添加的"))

	// 删除指定值的节点
	m.RemoveAtData("从尾部添加的")
	m.ShowList()

	// 删除指定位置的节点
	err := m.RemoveAtIndex(1)
	if err != nil {
		fmt.Println(err)
	}
	m.ShowList()
}
