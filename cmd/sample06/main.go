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

	var nIndex int = -1
	//0xFF表示标准版, 0x00为时钟锁,0x01为带时钟的U盘锁,0x02为标准U盘锁
	for i, info := range list {
		if info.Type == 0 || info.Type == 1 {
			nIndex = i
		}
	}

	if nIndex == -1 { //没有找到时钟锁
		fmt.Println("Can't Find Time Dongle ARM.\n")
		return
	}
	var hDongle dongle_etarm.DONGLE_HANDLE
	//打开第1把锁
	err = dongle_etarm.Dongle_Open(&hDongle, nIndex)
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

	var dwTime uint32 = 0

	//获取锁内时间
	err = dongle_etarm.Dongle_GetUTCTime(hDongle, &dwTime)
	if err != nil {
		panic(err)
	}
	//	 TODO 栗子未完成，请移步vc源码

}
