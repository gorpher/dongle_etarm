package main

import (
	"fmt"
	"github.com/gorpher/dongle_etarm"
)

// 恢复出厂设置

// 慎重使用此函数
var hid = "6E92743A7B06157D"

func main() {
	list, err := dongle_etarm.DongleEnum()
	if err != nil {
		panic(err)
	}

	var index = -1
	for i := range list {
		fmt.Printf("%#v\n", list[i])
		fmt.Printf("HID:%X\n", list[i].HID)
		if fmt.Sprintf("%X", list[i].HID) == hid {
			index = i
		}
	}
	if index == -1 {
		panic("没找到对应的加密狗")
	}

	var hDongle dongle_etarm.DONGLE_HANDLE
	//打开第1把锁
	err = dongle_etarm.Dongle_Open(&hDongle, 0)
	if err != nil {
		panic(err)
	}
	//关闭加密锁
	defer dongle_etarm.Dongle_Close(hDongle)

	//验证开发商PIN码
	var nRemainCount int
	err = dongle_etarm.Dongle_VerifyPIN(hDongle, dongle_etarm.FLAG_ADMINPIN, dongle_etarm.CONST_ADMINPIN, &nRemainCount)
	if err != nil {
		panic(err)
	}

	//一键恢复出厂设置
	err = dongle_etarm.Dongle_RFS(hDongle)
	if err != nil {
		panic(err)
	}
}
