package main

import (
	"errors"
	"fmt"
)

// Node 创建节点结构/类型
type Node struct {
	Data     interface{} //数据
	PreNote  *Node       //父级节点 等于null则为第一个节点
	NextNode *Node       //子级节点 等于null则为最后一个节点
}

// Link 创建链表结构
type Link struct {
	HeadNode *Node //头部节点（不存在数据，只保存头部节点信息关联后续节点）
	TailNode *Node //尾部节点（不存在数据，只保存尾部节点信息关联后续节点）
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
	node = &Node{Data: v, PreNote: nil, NextNode: nil}
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

// Length 返回链表节点长度
func (l *Link) Length() (count int) {
	//定义临时变量存储链表头结点
	tempNode := l.HeadNode

	//循环位移到指定索引的节点位置
	for count = 0; tempNode != nil; count++ {
		//如果子级节点不为nil则一直递归
		tempNode = tempNode.NextNode
	}
	return
}

// Contain 查询指定值的节点位置
func (l *Link) Contain(v interface{}) (index int) {
	//定义变量用于存储递归子级节点
	tempNode := l.HeadNode

	//默认索引位为0
	index = 0

	//循环位移到指定索引的节点位置
	for tempNode != nil {
		//如果该节点的值等于指定的值则返回true，否则继续位移子级节点
		if tempNode.Data == v {
			return
		}

		//取出该节点的子级节点
		tempNode = tempNode.NextNode

		//索引前进一位
		index++
	}

	//所有节点扫描完还是不能匹配上则返回-1
	return -1
}

// Add 从链表的头部添加节点
func (l *Link) Add(v interface{}) {
	//创建新节点
	node := CreateNode(v)
	//如果链表为空则直接将头部节点和尾部节点都指向新节点
	if l.IsEmpty() {
		l.HeadNode, l.TailNode = node, node
	} else {
		//1.将原头部节点的父级节点指向新节点
		//2.将原头部节点指向新节点的子级节点
		//3.将新节点指向链表的头部节点
		l.HeadNode.PreNote, node.NextNode, l.HeadNode = node, l.HeadNode, node
	}
}

// Append 从链表尾部添加节点
func (l *Link) Append(v interface{}) {
	//创建新节点
	node := CreateNode(v)

	//如果链表为空则直接将头部节点和尾部节点都指向新节点
	if l.IsEmpty() {
		l.HeadNode, l.TailNode = node, node
	} else {
		//定义变量用于存储递归子级节点
		tempNode := l.HeadNode

		//循环子级节点
		for tempNode.NextNode != nil {
			//如果子级节点不为nil则一直位移
			tempNode = tempNode.NextNode
		}

		//1.链表最后一个节点的子级节点指向新节点
		//2.链表尾部节点指向新节点
		//3.新节点的父级节点指向链表最后一个节点
		tempNode.NextNode, l.TailNode, node.PreNote = node, node, tempNode
	}
}

// Insert 在链表指定位置添加节点
func (l *Link) Insert(index int, v interface{}) {
	//如果索引小于0则从头部添加
	if index < 0 {
		l.Add(v)
	} else if index > l.Length() {
		//如果索引大于链表长度则从尾部添加
		l.Append(v)
	} else {
		//定义变量用于存储递归子级节点
		tempNode := l.HeadNode

		//循环位移到指定索引的节点位置
		for i := 0; i < index; i++ {
			//取出该节点的子级节点
			tempNode = tempNode.NextNode
		}

		//创建新节点
		node := CreateNode(v)
		//TODO 不能使用多重赋值的写法，否则报未对指针初始化的错误
		//1.将新节点的父级节点指向指定节点的父级节点
		node.PreNote = tempNode.PreNote
		//2.将新节点的父级节点的子级节点指向新节点
		node.PreNote.NextNode = node
		//3.将指定节点的父级节点指向新节点
		tempNode.PreNote = node
		//4.将新节点的子级节点指向指定节点
		node.NextNode = tempNode
	}
}

// RemoveAtData 删除指定值的节点
func (l *Link) RemoveAtData(v interface{}) {
	//定义变量用于存储递归子级节点
	tempNode := l.HeadNode

	//如果该节点的值等于指定值则直接接将链表的头部节点和尾部节点都指向该节点的子级节点
	if tempNode.Data == v {
		l.HeadNode, l.TailNode = tempNode.NextNode, tempNode.NextNode
	} else {
		//循环子级节点
		for tempNode.NextNode != nil {
			//如果该节点指向的子级节点的值等于指定值
			if tempNode.NextNode.Data == v {
				//1.将该节点的子级节点指向该节点的子级节点的子级节点
				tempNode.NextNode = tempNode.NextNode.NextNode
				//2.将该节点的子级节点的父级节点指向该节点
				tempNode.NextNode.PreNote = tempNode
			} else {
				//取出该节点的子级节点
				tempNode = tempNode.NextNode
			}
		}
	}
}

// RemoveAtIndex 删除指定位置的节点
func (l *Link) RemoveAtIndex(index int) (err error) {
	//定义变量用于存储递归子级节点
	tempNode := l.HeadNode

	//如果index小于0则直接用头部节点的子级节点指向链表的头部节点 （删除原头部节点）
	if index < 0 {
		l.HeadNode = tempNode.NextNode
	} else if index > l.Length() {
		err = errors.New("超出链表长度")
		return
	} else {
		//循环位移至指定位置节点前一位，并且该节点的子节点不可能为nil (改节点的子节点指向要删除的节点)
		for i := 0; i != index && tempNode.NextNode != nil; i++ {
			//取出该节点的子级节点
			tempNode = tempNode.NextNode
		}

		//1.将该节点的子级节点的父级节点指向该节点的父级节点
		tempNode.NextNode.PreNote = tempNode.PreNote
		//2.将该节点的父级节点的子级节点指向该节点的子级节点
		tempNode.PreNote.NextNode = tempNode.NextNode
	}

	return
}

// ShowList 遍历所有节点
func (l *Link) ShowList() {
	//如果链表不为空则打印子级节点信息
	if !l.IsEmpty() {
		//定义变量用于存储递归子级节点
		tempNode := l.HeadNode

		//循环子级节点
		fmt.Println("原数据-------")
		for index := 0; tempNode != nil; index++ {
			fmt.Printf("该节点位置：%d 值：%v\n", index, tempNode)

			//取出该节点的子级节点
			tempNode = tempNode.NextNode
		}

		//定义变量用于存储递归子级节点
		tempNode = l.HeadNode

		//循环子级节点
		fmt.Println("子级节点-------")
		for index := 0; tempNode != nil; index++ {
			fmt.Printf("该节点位置：%d 值：%v\n", index, tempNode.Data)

			//取出该节点的子级节点
			tempNode = tempNode.NextNode
		}

		//定义变量用于存储递归父级节点
		tempNode = l.TailNode

		//循环父级节点
		fmt.Println("父级节点-------")
		for index := 0; tempNode != nil; index++ {
			fmt.Printf("该节点位置：%d 值：%v\n", index, tempNode.Data)

			//取出该节点的父级节点
			tempNode = tempNode.PreNote
		}
	} else {
		fmt.Printf("link empty")
	}
	fmt.Println()
}

func main() {
	//创建链表
	link := CreateLink()

	//实现接口
	var m Method = link

	//头部加入节点
	m.Add("从头部添加的")
	//显示链表
	m.ShowList()

	/*//从头部加入节点二
	m.Add("从头部加入节点二")
	m.ShowList()

	//从头部加入节点三
	m.Add("从头部加入节点三")
	m.ShowList()*/

	//尾部添加节点
	m.Append("从尾部添加的")
	m.ShowList()

	m.Append("从尾部再添加一个")
	m.ShowList()

	//链表长度
	fmt.Printf("链表长度为：%d\n", m.Length())

	//指定索引2的位置添加
	m.Insert(1, "指定索引1的位置添加")
	m.ShowList()

	//获取指定值位置
	fmt.Println("该值位置：", m.Contain("从尾部添加的"))

	//删除指定值的节点
	m.RemoveAtData("从尾部添加的")
	m.ShowList()

	//删除指定位置的节点
	err := m.RemoveAtIndex(1)
	if err != nil {
		fmt.Println(err)
	}
	m.ShowList()
}
