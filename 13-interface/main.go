package main

import "fmt"

// 接口只定义方法，没有任何变量
// 接口的方法没有任何实现，只用定义 参数,名称，返回值 
type Usb interface {
	Start()
	Stop()
}
type Usb2 interface {
	Start()
	Stop()
	Runing()
}

type Phone struct {
}

// usb::Start
func (p Phone) Start() {
	fmt.Println("手机开启")
}

// usb::Stop
func (p Phone) Stop() {
	fmt.Println("手机关闭")
}

type Camera struct {
}

// usb::Start
func (c Camera) Start() {
	fmt.Println("相机开启")

}

// usb::Stop
func (c Camera) Stop() {
	fmt.Println("相机关闭")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}
func main() {
	fmt.Println()

	c := Computer{}
	phone := Phone{}
	camera := Camera{}

	c.Working(phone)
	c.Working(camera)

}
