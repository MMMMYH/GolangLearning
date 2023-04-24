package main

import "fmt"

type Usb interface {
	start()
	stop()
}

type Phone struct {
}

func (p Phone) start() {
	fmt.Println("手机开始工作")
}

func (p Phone) stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
}

func (c Camera) start() {
	fmt.Println("相机开始工作")
}

func (c Camera) stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.start()
	usb.stop()
}

func main() {
	phone := Phone{}
	camera := Camera{}
	computer := Computer{}
	computer.Working(phone)
	computer.Working(camera)
}
