package main

import (
	"crypto/rand"
	"fmt"
	"github.com/gorpher/dongle_etarm"
	"unsafe"
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
	//创建数据文件
	var dataAttr dongle_etarm.DATA_FILE_ATTR
	dataAttr.Size = 1024
	dataAttr.Lic.ReadPriv = 0  //最小匿名权限
	dataAttr.Lic.WritePriv = 0 //最小匿名权限
	err = dongle_etarm.Dongle_CreateFILE_DATA(hDongle, 0x0001, &dataAttr)
	if err != nil {
		panic(err)
	}
	dataAttr.Size = 1024

	err = dongle_etarm.Dongle_CreateFILE_DATA(hDongle, 0x0002, &dataAttr)
	if err != nil {
		panic(err)
	}
	var writeLen = 1024
	//写数据文件
	buffer := make([]byte, writeLen)
	_, err = rand.Read(buffer)
	if err != nil {
		panic(err)
	}

	err = dongle_etarm.Dongle_WriteFile(hDongle, dongle_etarm.FILE_DATA, 0x0001, 0, &buffer[0], writeLen)
	if err != nil {
		panic(err)
	}
	//读取数据文件
	buffer2 := make([]byte, writeLen)

	err = dongle_etarm.Dongle_ReadFile(hDongle, 0x0001, 0, &(buffer2[0]), writeLen)

	//列文件 需要管理员权限
	var dataFileList = make([]dongle_etarm.DATA_FILE_LIST, 30)
	var dataLen = 32 * int(unsafe.Sizeof(dataFileList[0]))

	err = dongle_etarm.Dongle_ListFile(hDongle, dongle_etarm.FILE_DATA, &dataFileList[0], &dataLen)
	if err != nil {
		panic(err)
	}
	fmt.Println(dataFileList)
	//删除文件
	err = dongle_etarm.Dongle_DeleteFile(hDongle, dongle_etarm.FILE_DATA, 0x0001)
	if err != nil {
		panic(err)
	}
	err = dongle_etarm.Dongle_DeleteFile(hDongle, dongle_etarm.FILE_DATA, 0x0002)
	if err != nil {
		panic(err)
	}

}
