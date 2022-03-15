package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	//TODO container包详解

	// 双向链表
	// 创建一个 List 对象的指针 并自动调用Init方法初始化链表 返回该双向链表
	doubleLink := list.New()
	// 初始化链表或清除链表成初始状态
	//doubleLink.Init()
	// 新增
	// 插入链表
	// 从链表头部插入单个值，并返回该节点（*Element）
	node1 := doubleLink.PushFront("front node")
	// 从链表头部插入一个链表
	//element.PushFrontList(list *List)
	// 从链表尾部插入单个值，并返回该节点（*Element）
	node2 := doubleLink.PushBack("back node")
	// 从链表尾部插入一个链表
	//element.PushBackList(list *List)
	// 在指定的节点（*Element）之前插入一个值，并返回该节点（*Element） 如果该节点不是该链表的元素则链表不会被修改
	node3 := doubleLink.InsertBefore("insert front node", node1)
	// 在指定的节点（*Element）之后插入一个值，并返回该节点（*Element） 如果该节点不是该链表的元素则链表不会被修改
	node4 := doubleLink.InsertAfter("insert back node", node2)
	// 查询
	// 获取链表的第一个节点
	frontNode := doubleLink.Front()
	fmt.Println("链表第一个节点：", frontNode, " 值：", frontNode.Value, " 上一个节点：", frontNode.Prev(), " 下一个节点：", frontNode.Next())
	// 获取链表的最后一个节点
	backNode := doubleLink.Back()
	fmt.Println("链表最后一个节点：", backNode)
	// 获取指定节点的上一个节点
	fmt.Println("该节点的上一个节点", node3.Prev())
	// 获取指定节点的下一个节点
	fmt.Println("该节点的下一个节点", node4.Next())
	// 删除
	// 删除该链表指定的节点，并返回该节点的值（interface{}）
	value := doubleLink.Remove(node3)
	fmt.Println("删除元素 值：", value)
	// 修改
	// 将指定的节点移动到该链表的头部 如果该节点不是该链表的元素则链表不会被修改
	doubleLink.MoveToFront(node1)
	// 将指定的节点移动到该链表的尾部 如果该节点不是该链表的元素则链表不会被修改
	doubleLink.MoveToBack(node2)
	// 将指定的节点移动到节点mark的前面 如果e或者mark不是该链表的元素，或者e==mark则链表不会被修改
	//doubleLink.MoveBefore(e, mark *Element)
	// 将指定的节点移动到节点mark的后面 如果e或者mark不是该链表的元素，或者e==mark则链表不会被修改
	//doubleLink.MoveAfter(e, mark *Element)
	fmt.Println("该双向链表的长度为：", doubleLink.Len())

	// 环形链表
	// 创建一个 Ring 对象的指针 需要指定链表的长度 返回该链表的头节点
	node := ring.New(10)
	// 给环形链表设置值
	for i := 0; i < 10; i++ {
		// 设置节点的值
		node.Value = i
		// 递归子级节点
		node = node.Next()
	}
	fmt.Println("该环形链表的长度为：", node.Len())
	fmt.Println("该节点的上一个节点：", node.Prev().Value)
	fmt.Println("该节点的下一个节点：", node.Next().Value)
	// 移动指定数量n个位置 （n>=0向前移动，n<0向后移动） 并返回该元素节点
	node5 := node.Move(5)
	// 将指定的元素节点做为该节点的下一个级节点 并返回该节点原本的下一个节点
	node6 := node.Link(node5)
	// 删除该节点指定长度n的子级节点 如果指定长度<=0或者超出链表长度则返回nil，否则返回被移除的节点（如果被移除多个节点则返回最后一个节点）
	node7 := node6.Unlink(2)
	// 统计链表值的集合
	sum := 0
	// 对链表的每一个元素节点都执行该自定义函数（正向顺序）
	node7.Do(func(i interface{}) {
		sum += i.(int)
	})
	fmt.Println("链表值得集合结果：", sum)

	// 堆
	// 我们要使用go标准库给我们提供的heap，那么必须自己实现这些接口定义的方法
	h := &IntHeap{1, 2, 3, 4, 5}
	// Init()函数建立此包中其他例程所需的堆不变量。对于堆不变量是幂等的(可被多次调用)，并且可以在堆不变量失效时调用
	heap.Init(h)
	// 往堆里面插入一个元素
	heap.Push(h, 6)
	// 移除并返回堆顶元素
	value = heap.Pop(h)
	// 从指定位置i移除并返回该元素
	value = heap.Remove(h, 3)
	// 从i位置数据发生改变后，对堆再平衡，优先级队列使用到了该方法（修复堆）
	heap.Fix(h, 1)
}

// IntHeap 定义堆数据类型
type IntHeap []int

// Len 绑定len方法,返回长度
func (h IntHeap) Len() int {
	return len(h)
}

// Less 绑定less方法
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j] //如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}

// Swap 绑定swap方法，交换两个元素位置
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 绑定push方法，插入新元素
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 绑定pop方法，从最后拿出一个元素并返回
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
