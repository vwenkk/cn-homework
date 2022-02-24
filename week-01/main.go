package main

import (
	"fmt"
	"time"
)

//课后练习 1.1
//编写一个小程序：
//给定一个字符串数组
//[“I”,“am”,“stupid”,“and”,“weak”]
//用 for 循环遍历该数组并修改为
//[“I”,“am”,“smart”,“and”,“strong”]
//
//课后练习 1.2
//基于 Channel 编写一个简单的单线程生产者消费者模型：
//
//队列：
//队列长度 10，队列元素类型为 int
//生产者：
//每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
//消费者：
//每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
func main() {
	arr := [...]string{"I", "am", "stupid", "and", "weak"}
	length := len(arr)
	for i := 0; i < length; i++ {
		if arr[i] == "stupid" {
			arr[i] = "smart"
		} else if arr[i] == "weak" {
			arr[i] = "strong"
		}
	}
	fmt.Println(arr)

	ch := make(chan int, 10)

	// 生产者
	go func(ch chan<- int) {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				ch <- int(time.Now().Unix())
			}
		}

	}(ch)

	// 消费者
	func(ch <-chan int) {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println(<-ch)
			}
		}
	}(ch)

}
