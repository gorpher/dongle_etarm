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
	var exeFileInfo = make([]dongle_etarm.EXE_FILE_INFO, 2)

	//设置批量下载的可执行程序
	//设置第一个文件
	exeFileInfo[0].DwSize = EXE_SIZE_1 //文件大小
	exeFileInfo[0].Priv = 0            //匿名可调用
	exeFileInfo[0].WFileID = 0x1111    //文件ID
	exeFileInfo[0].PData = &g_progExeFile1[0]

	//设置第二个文件
	exeFileInfo[1].DwSize = EXE_SIZE_2 //文件大小
	exeFileInfo[1].Priv = 0            //匿名可调用
	exeFileInfo[1].WFileID = 0x2222    //文件ID
	exeFileInfo[1].PData = &g_progExeFile2[0]

	//批量下载可执行文件
	err = dongle_etarm.Dongle_DownloadExeFile(hDongle, &exeFileInfo[0], 2)
	if err != nil {
		panic(err)
	}

	//运行可执行文件
	var wInOufBufLen uint16 = 20
	buffer := make([]byte, 1020)
	var nMainRet = 0

	err = dongle_etarm.Dongle_RunExeFile(hDongle, 0x1111, &buffer[0], wInOufBufLen, &nMainRet)
	if err != nil {
		panic(err)
	}

	fmt.Println("nMainRet", nMainRet)
	//运行可执行文件
	wInOufBufLen = 32
	err = dongle_etarm.Dongle_RunExeFile(hDongle, 0x2222, &buffer[1], wInOufBufLen, &nMainRet)
	if err != nil {
		panic(err)
	}
	fmt.Println("nMainRet", nMainRet)

}
