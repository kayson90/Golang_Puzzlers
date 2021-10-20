package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

//鸭子类型，无侵入式的接口实现方式
type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1。
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	// 示例2。
	//指针类型实现了Pet接口
	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}
