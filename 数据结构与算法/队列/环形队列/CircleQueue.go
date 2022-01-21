package main

import (
	"errors"
	"fmt"
)

// Queue 环形队列结构体
type Queue struct {
	MaxLen int           //队列最大长度
	List   []interface{} //模拟队列
	Head   int           //指向队列头部 0
	Tail   int           //指向队列尾部 0
}

// IsEmpty 判断队列是否为空（队列尾部元素索引和队列头部元素相等)
func (q *Queue) IsEmpty() bool {
	return q.Tail == q.Head
}

// IsFull 判断队列是否已满 (容量最后一位用来做约定计算队列是否已满，所以队列真实长度为MaxLen-1) (简单来说就是看头部是否为0，只有当头部为0且尾部为maxLen-1时队列才无法继续追加)
func (q *Queue) isFull() bool {
	return (q.Tail+1)%q.MaxLen == q.Head
}

// Size 获取队列长度
func (q *Queue) Size() int {
	return (q.Tail + q.MaxLen - q.Head) % q.MaxLen
}

// Push 向队列中添加元素
func (q *Queue) Push(data ...interface{}) (err error) {
	//往队列中添加元素
	for _, val := range data {
		fmt.Println(val)
		//每次追加循环判断  当数量不足时直至添加满队列
		if q.isFull() {
			fmt.Println(1)
			return errors.New("queue full")
		} else {
			fmt.Println(2)
			//因为尾部初始是0，所以先赋值再自增

			//添加元素
			q.List[q.Tail] = val

			//Tail加1后%队列长度，如果等于零则队列尾部已加满，重新再加头部
			q.Tail = (q.Tail + 1) % q.MaxLen
		}
	}

	return
}

// Pop 从队列中取出元素
func (q *Queue) Pop() (data interface{}, err error) {
	//队列是否为空
	if q.IsEmpty() {
		return "", errors.New("queue empty")
	}

	//取出队列头部元素
	data = q.List[q.Head]

	//Head加1后%队列长度，如果等于零则队列尾部已取完，重新再取头部
	q.Head = (q.Head + 1) % q.MaxLen

	return
}

// Range 显示队列
func (q *Queue) Range() {
	//取出当前队列的长度
	size := q.Size()
	if size == 0 {
		fmt.Println("当前队列为空")
	} else {
		//保存一个变量来存储队列头部
		tempHead := q.Head
		for i := 0; i < size; i++ {
			fmt.Printf("索引：%d\t值：%s\n", tempHead, q.List[tempHead])

			//头部索引自增，如果等于零则队列尾部已取完，重新再取头部
			tempHead = (tempHead + 1) % q.MaxLen
		}
	}
}

func main() {
	//环形队列结构体实例
	queue := &Queue{
		MaxLen: 7,
		List:   make([]interface{}, 7),
		Head:   0,
		Tail:   0,
	}

	//插入数据
	err := queue.Push("tom", "jerry", "jom", "lori", "jack", "angus")
	if err != nil {
		fmt.Println(err)
	}

	//显示队列数据结构
	queue.Range()

	//取出数据
	data, err := queue.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("从队列中取出：%s\n", data)
	}
}
