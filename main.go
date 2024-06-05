package main

import (
	"fmt"
	"sync"
)

// Singleton 结构体定义
type Singleton struct {
	Instance *Singleton
	Data     string
}

// 实例变量和锁
var instance *Singleton
var mutex sync.Mutex

// GetInstance 方法，用于获取 Singleton 实例的指针
func GetInstance() *Singleton {
	// 锁定以确保线程安全
	mutex.Lock()
	defer mutex.Unlock()

	// 检查实例是否已经创建
	if instance == nil {
		// 如果没有，则创建一个新的实例
		instance = &Singleton{
			Data: "Singleton Data",
		}
		instance.Instance = instance
	}
	return instance
}

func main() {
	// 测试代码
	s1 := GetInstance()
	fmt.Println("First instance:", s1)

	s2 := GetInstance()
	fmt.Println("Second instance:", s2)

	// 验证是否是同一个实例
	if s1 == s2 {
		fmt.Println("Both instances are the same.")
	} else {
		fmt.Println("Error: Instances are different.")
	}
}
