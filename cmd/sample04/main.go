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
	//验证开发商PIN码
	var nRemainCount int
	err = dongle_etarm.Dongle_VerifyPIN(hDongle, dongle_etarm.FLAG_ADMINPIN, dongle_etarm.CONST_ADMINPIN, &nRemainCount)
	if err != nil {
		panic(err)
	}

	var seed = make([]byte, 32)
	for i := range seed {
		seed[i] = 0x11
	}
	//唯一化一把锁，需要管理员权限
	cPid := make([]byte, 9)
	cAdminPin := make([]byte, 17)
	fmt.Println("cAdminPin", cAdminPin)
	err = dongle_etarm.Dongle_GenUniqueKey(hDongle, 32, &seed[0], &cPid[0], &cAdminPin[0])
	if err != nil {
		panic(err)
	}
	fmt.Println("cpid", cPid, "cAdminPin", cAdminPin)
	//验证用户PIN
	nRemainCount = 0
	userPin := "12345678"
	err = dongle_etarm.Dongle_VerifyPIN(hDongle, dongle_etarm.FLAG_USERPIN, userPin, &nRemainCount)
	if err != nil {
		panic(err)
	}
	//恢复到匿名状态
	err = dongle_etarm.Dongle_ResetState(hDongle)
	if err != nil {
		panic(err)
	}

	//验证开发商PIN
	nRemainCount = 0
	var cAdminPinStr = string(cAdminPin)
	err = dongle_etarm.Dongle_VerifyPIN(hDongle, dongle_etarm.FLAG_ADMINPIN, cAdminPinStr, &nRemainCount)
	if err != nil {
		panic(err)
	}
	fmt.Println("修改adminPin", cAdminPinStr)

	//更改开发商PIN
	nRemainCount = 100
	err = dongle_etarm.Dongle_ChangePIN(hDongle, dongle_etarm.FLAG_ADMINPIN, cAdminPinStr, dongle_etarm.CONST_ADMINPIN, nRemainCount)
	if err != nil {
		panic(err)
	}
	//一键恢复出厂设置
	err = dongle_etarm.Dongle_RFS(hDongle)
	if err != nil {
		panic(err)
	}

}
