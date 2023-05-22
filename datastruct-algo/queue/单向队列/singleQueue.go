package main

import (
	"errors"
	"fmt"
)

// Queue 单向队列结构体
type Queue struct {
	MaxLen int           //队列最大长度
	List   []interface{} //模拟队列
	Front  int           //指向队列头部 -1
	Rear   int           //指向队列尾部 -1
}

// Method 设计接口：
type Method interface {
	IsEmpty() bool                       // 判断队列是否为空
	isFull() bool                        // 判断队列是否已满
	Set(data ...interface{}) (err error) // 向队列中添加元素
	Get() (data interface{}, err error)  // 从队列中取出元素
	Show()                               // 显示队列
}

// IsEmpty 判断队列是否为空（队列尾部元素索引和队列头部元素相等)
func (q *Queue) IsEmpty() bool {
	return q.Rear == q.Front
}

// IsFull 判断队列是否已满（队列尾部元素索引和队列最大长度相等)
func (q *Queue) isFull() bool {
	return q.Rear+1 == q.MaxLen
}

// Set 向队列中添加元素
func (q *Queue) Set(data ...interface{}) (err error) {
	//往队列中添加元素
	for _, val := range data {
		//每次追加循环判断  当数量不足时直至添加满队列
		if q.isFull() {
			return errors.New("queue full")
		} else {
			//因为尾部初始是-1，所以先自增再赋值

			//Rear后移
			q.Rear++

			//添加元素
			q.List[q.Rear] = val
		}
	}

	return
}

// Get 从队列中取出元素
func (q *Queue) Get() (data interface{}, err error) {
	//队列是否为空
	if q.IsEmpty() {
		return "", errors.New("queue empty")
	}

	//Front后移
	q.Front++

	//取出队列头部元素
	data = q.List[q.Front]

	return
}

// Show 显示队列
func (q *Queue) Show() {
	fmt.Printf("队列最大长度：%d\n", q.MaxLen)

	fmt.Printf("队列当前长度：%d\n", len(q.List))

	fmt.Printf("队列当前头部：%d\n", q.Front)

	fmt.Printf("队列当前尾部：%d\n", q.Rear)

	//因为队列头部默认是-1，所有取出队列则+1（不包含头部对应索引的元素）
	for i := q.Front + 1; i <= q.Rear; i++ {
		fmt.Printf("索引：%d\t值：%s\n", i, q.List[i])
	}
}

func main() {
	//单向队列结构体实例
	queue := &Queue{
		MaxLen: 4,
		List:   make([]interface{}, 4),
		Front:  -1,
		Rear:   -1,
	}

	//实现接口
	var m Method = queue

	//插入数据
	err := m.Set("tom", "jerry", "jom", "lori")
	if err != nil {
		fmt.Println(err)
	}

	//显示队列数据结构
	m.Show()

	//取出数据
	data, err := m.Get()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("从队列中取出：%s\n", data)
	}
}
