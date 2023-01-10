package main

import (
	"fmt"
	"github.com/gorpher/dongle_etarm"
)

func main() {
	list, err := dongle_etarm.DongleEnum()
	if err != nil {
		panic(err)
	}
	//枚举锁
	if len(list) == 0 {
		return
	}

	var hDongle dongle_etarm.DONGLE_HANDLE
	//打开第1把锁
	err = dongle_etarm.Dongle_Open(&hDongle, 0)
	if err != nil {
		panic(err)
	}
	//关闭加密锁
	defer dongle_etarm.Dongle_Close(hDongle)

	//写共享内存区
	bShareMemory := make([]byte, 32)
	for i := range bShareMemory {
		bShareMemory[i] = 0x22
	}

	err = dongle_etarm.Dongle_WriteShareMemory(hDongle, &bShareMemory[0], 32)
	if err != nil {
		panic(
			err)
	}

	//读共享内存区，总大小为32个字节
	err = dongle_etarm.Dongle_WriteShareMemory(hDongle, &bShareMemory[0], 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Share Memory[0~31]: %x\n", bShareMemory)

	var nRemainCount int
	//验证开发商PIN码
	err = dongle_etarm.Dongle_VerifyPIN(hDongle, dongle_etarm.FLAG_ADMINPIN, dongle_etarm.CONST_ADMINPIN, &nRemainCount)
	if err != nil {
		panic(err)
	}

	//写数据区，匿名和用户权限可写前4k(0~4095),开发商有所有8k的写权限
	bDataSec := make([]byte, 8192)
	for i := range bDataSec {
		bDataSec[i] = 0x33
	}
	err = dongle_etarm.Dongle_WriteData(hDongle, 0, &bDataSec[0], 8192)
	if err != nil {
		panic(err)
	}

	//读数据区
	bDataSec = make([]byte, 8192)
	err = dongle_etarm.Dongle_ReadData(hDongle, 0, &bDataSec[0], 8192)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data Section[4096~8191]: %#X", bDataSec)

}
