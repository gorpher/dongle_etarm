package main

import (
	"fmt"
	"github.com/gorpher/dongle_etarm"
)

//说明：应该先唯一化锁之后才能运行该示例，否则无法修改用户PIN码
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

	var cUserPin = "87654321"

	//修改用户PIN码为 “87654321”
	err = dongle_etarm.Dongle_ChangePIN(hDongle, dongle_etarm.FLAG_USERPIN, dongle_etarm.CONST_USERPIN, cUserPin, 0xFF)
	if err != nil {
		panic(err)
	}

	//清除开发商权限状态
	err = dongle_etarm.Dongle_ResetState(hDongle)
	if err != nil {
		panic(err)
	}

	//验证旧的用户PIN => 失败
	err = dongle_etarm.Dongle_VerifyPIN(hDongle, dongle_etarm.FLAG_USERPIN, dongle_etarm.CONST_USERPIN, &nRemainCount)
	if err != dongle_etarm.DONGLE_INCORRECT_PIN {
		panic(err)
	}
	fmt.Println("验证旧密码说明：", err.Error())

	//验证新的用户PIN => 成功
	err = dongle_etarm.Dongle_VerifyPIN(hDongle, dongle_etarm.FLAG_USERPIN, cUserPin, &nRemainCount)
	if err != nil {
		panic(err)
	}

	//设置用户ID,需要开发商权限
	err = dongle_etarm.Dongle_VerifyPIN(hDongle, dongle_etarm.FLAG_ADMINPIN, dongle_etarm.CONST_ADMINPIN, &nRemainCount)
	if err != nil {
		panic(err)
	}
	var dwUserID uint32 = 0x11111111
	err = dongle_etarm.Dongle_SetUserID(hDongle, dwUserID)
	if err != nil {
		panic(err)
	}

	//重设用户PIN码
	err = dongle_etarm.Dongle_ResetUserPIN(hDongle, dongle_etarm.CONST_ADMINPIN)
	if err != nil {
		panic(err)
	}

}

/*// Sample05.cpp : Defines the entry point for the console application.
//

#include "stdafx.h"
#include "..\Lib\Dongle_API.h"
#pragma comment(lib, "..\\LIB\\Dongle_s.lib")

//唯一化锁的示例程序，请参考Sample04
int main(int argc, char* argv[])
{
	DWORD dwRet = 0;
	int nCount = 0;
	BYTE seed[32];
	char cPid[9];
	char cAdminPin[17];
	char AdminPin [] = "FFFFFFFFFFFFFFFF";//默认开发商PIN码
	char UserPin [] = "12345678";
	char cUserPin [] = "87654321";
	DWORD dwUserID = 0x11111111;
	int nRemainCount = 0;

	DONGLE_HANDLE hDongle = NULL;

	//枚举锁
	dwRet = Dongle_Enum(NULL, &nCount);
	printf("Enum %d Dongle ARM. \n", nCount);

	//打开锁
	dwRet = Dongle_Open(&hDongle, 0);
	printf("Open Dongle ARM. Return : 0x%08X . \n", dwRet);

	//验证开发商PIN码
	dwRet = Dongle_VerifyPIN(hDongle, FLAG_ADMINPIN, AdminPin, &nRemainCount);
	printf("Verify Admin PIN. Return: 0x%08X\n", dwRet);

	//修改用户PIN码为 “87654321”
	dwRet = Dongle_ChangePIN(hDongle, FLAG_USERPIN, UserPin, cUserPin, 0xFF);
	printf("Change User PIN. Return: 0x%08X\n", dwRet);

	//清除开发商权限状态
	dwRet = Dongle_ResetState(hDongle);
	printf("Reset COS state. Return: 0x%08X\n", dwRet);

	//验证旧的用户PIN => 失败
	dwRet = Dongle_VerifyPIN(hDongle, FLAG_USERPIN, UserPin, &nRemainCount);
	printf("Verify old user PIN. Return: 0x%08X\n", dwRet);


	//验证新的用户PIN => 成功
	dwRet = Dongle_VerifyPIN(hDongle, FLAG_USERPIN, cUserPin, &nRemainCount);
	printf("Verify new user PIN. Return: 0x%08X\n", dwRet);


	//设置用户ID,需要开发商权限
	dwRet = Dongle_VerifyPIN(hDongle, FLAG_ADMINPIN, AdminPin, &nRemainCount);
	dwRet = Dongle_SetUserID(hDongle, dwUserID);
	printf("Set User ID. Return: 0x%08X\n", dwRet);

	//重设用户PIN码
	dwRet = Dongle_ResetUserPIN(hDongle, AdminPin);
	printf("Reset User PIN. Return: 0x%08X\n", dwRet);

	//关闭加密锁
	dwRet = Dongle_Close(hDongle);
	printf("Close Dongle ARM. Return: 0x%08X\n", dwRet);

	return 0;
}

*/
