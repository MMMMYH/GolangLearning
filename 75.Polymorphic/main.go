package main

import "fmt"

type Usb interface {
	start()
	stop()
}

type Phone struct {
	name string
}

func (p Phone) start() {
	fmt.Println("手机开始工作")
}

func (p Phone) stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
	name string
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
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"索尼"}
	fmt.Println(usbArr)

}
