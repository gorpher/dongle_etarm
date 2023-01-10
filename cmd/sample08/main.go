package main

import (
	"bytes"
	"github.com/gorpher/dongle_etarm"
	"os"
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
	var priAttr dongle_etarm.PRIKEY_FILE_ATTR
	//创建RSA私钥
	priAttr.Size = 1024 //也可以是2048
	priAttr.Type = dongle_etarm.FILE_PRIKEY_RSA
	priAttr.Lic.Count = 0xFFFFFFF //可调用次数
	priAttr.Lic.IsDecOnRAM = 0    //是否为在内存中递减
	priAttr.Lic.IsReset = 0       //调用完后是否恢复到匿名态
	priAttr.Lic.Priv = 0          //可调用的最小权限

	var wPriID uint16 = 0x00001
	err = dongle_etarm.Dongle_CreateFILE_PRIKEY_RSA(hDongle, wPriID, &priAttr)
	if err != nil {
		panic(err)
	}
	var rsaPub dongle_etarm.RSA_PUBLIC_KEY
	var rsaPri dongle_etarm.RSA_PRIVATE_KEY
	//产生RSA公私钥
	err = dongle_etarm.Dongle_RsaGenPubPriKey(hDongle, wPriID, &rsaPub, &rsaPri)
	if err != nil {
		panic(err)
	}
	//备份公私钥
	size := unsafe.Sizeof(&rsaPub) + 256
	body := (*byte)(unsafe.Pointer(&rsaPub))
	data := make([]byte, int(size))
	for i := 0; i < (int(size)); i++ {
		data[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(body)) + uintptr(i)))
	}

	err = os.WriteFile("mother.Rsapub", data, os.ModePerm)
	if err != nil {
		panic(err)
	}
	size = unsafe.Sizeof(&rsaPri) + 256 + 256
	body = (*byte)(unsafe.Pointer(&rsaPri))
	data = make([]byte, size)
	for i := 0; i < (int(size)); i++ {
		data[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(body)) + uintptr(i)))
	}
	err = os.WriteFile("mother.Rsapri", data, os.ModePerm)
	if err != nil {
		panic(err)
	}
	var buffer = make([]byte, 128, 128)
	var tmpbuf = make([]byte, 128, 128)
	//RSA私钥加密公钥解密
	//1.RSA私钥加密
	nInDataLen := 128 - 11
	nOutDataLen := 128
	for i := 0; i < 128; i++ {
		buffer[i] = uint8(i)
	}
	copy(tmpbuf, buffer)
	err = dongle_etarm.Dongle_RsaPri(hDongle, wPriID, dongle_etarm.FLAG_ENCODE, &buffer[0], nInDataLen, &buffer[0], &nOutDataLen)
	if err != nil {
		panic(err)
	}
	////2.RSA公钥解密
	nInDataLen = 128
	nOutDataLen = 128 - 11
	err = dongle_etarm.Dongle_RsaPub(hDongle, dongle_etarm.FLAG_DECODE, &rsaPub, &buffer[0], nInDataLen, &buffer[0], &nOutDataLen)
	if err != nil {
		panic(err)
	}
	compare := bytes.Compare(tmpbuf[:117], buffer[:117])
	if compare != 0 {
		panic("加密解密数据不一致")
	}
	//RSA公钥加密私钥解密
	//1.RSA公钥加密
	for i := 0; i < 128; i++ {
		buffer[i] = uint8(i)
	}
	nInDataLen = 128 - 11
	nOutDataLen = 128
	copy(tmpbuf, buffer)

	err = dongle_etarm.Dongle_RsaPub(hDongle, dongle_etarm.FLAG_ENCODE, &rsaPub, &buffer[0], nInDataLen, &buffer[0], &nOutDataLen)
	if err != nil {
		panic(err)
	}

	////2.RSA私钥解密
	nInDataLen = 128
	nOutDataLen = 128 - 11
	err = dongle_etarm.Dongle_RsaPri(hDongle, wPriID, dongle_etarm.FLAG_DECODE, &buffer[0], nInDataLen, &buffer[0], &nOutDataLen)
	if err != nil {
		panic(err)
	}

	compare = bytes.Compare(tmpbuf[:117], buffer[:117])
	if compare != 0 {
		panic("加密解密数据不一致")
	}

	//删除锁内私钥
	err = dongle_etarm.Dongle_DeleteFile(hDongle, dongle_etarm.FILE_PRIKEY_RSA, wPriID)
	if err != nil {
		panic(err)
	}

}

/*
// Sample08.cpp : Defines the entry point for the console application.
//

#include "stdafx.h"
#include "..\Lib\Dongle_API.h"
#pragma comment(lib, "..\\LIB\\Dongle_s.lib")

int main(int argc, char* argv[])
{
	DWORD dwRet = 0;
	int nCount = 0;
	int i = 0;
	char AdminPin [] = "FFFFFFFFFFFFFFFF";//默认开发商PIN码
	int nRemainCount = 0;

	DONGLE_HANDLE hDongle = NULL;
	RSA_PUBLIC_KEY  rsaPub;
	RSA_PRIVATE_KEY rsaPri;
	PRIKEY_FILE_ATTR priAttr;
	WORD wPriID = 0x1111;

	FILE *fp = NULL;
	int nLen = 0;
	BYTE buffer[128];
	BYTE tmpbuf[128];
	int  nInDataLen = 0;
	int  nOutDataLen = 0;

	for(i = 0; i < 128; i++) buffer[i] = i;

	//枚举锁
	dwRet = Dongle_Enum(NULL, &nCount);
	printf("Enum %d Dongle ARM. \n", nCount);

	//打开锁
	dwRet = Dongle_Open(&hDongle, 0);
	printf("Open Dongle ARM. Return : 0x%08X . \n", dwRet);

	//验证开发商PIN码
	dwRet = Dongle_VerifyPIN(hDongle, FLAG_ADMINPIN, AdminPin, &nRemainCount);
	printf("Verify Admin PIN. Return: 0x%08X\n", dwRet);

	//创建RSA私钥
	priAttr.m_Size = 1024;//也可以是2048
	priAttr.m_Type = FILE_PRIKEY_RSA;
	priAttr.m_Lic.m_Count =0xFFFFFFFF;//可调用次数
	priAttr.m_Lic.m_IsDecOnRAM = FALSE;//是否为在内存中递减
	priAttr.m_Lic.m_IsReset = FALSE;//调用完后是否恢复到匿名态
	priAttr.m_Lic.m_Priv = 0;//可调用的最小权限

	dwRet = Dongle_CreateFile(hDongle, FILE_PRIKEY_RSA, wPriID, (void*)&priAttr);
	printf("Create RSA private key file. Return: 0x%08X\n", dwRet);
	if (DONGLE_SUCCESS != dwRet)
	{
		return 0;
	}

	//产生RSA公私钥
	dwRet = Dongle_RsaGenPubPriKey(hDongle, wPriID, &rsaPub, &rsaPri);
	printf("Gen RSA Public key and private key. Return: 0x%08X\n", dwRet);
	if (DONGLE_SUCCESS != dwRet)
	{
		return 0;
	}

	//备份公私钥
	fp = fopen("1111.Rsapub", "wb");
	fwrite(&rsaPub, 1, sizeof(RSA_PUBLIC_KEY), fp);
	fclose(fp);

	fp = fopen("1111.Rsapri", "wb");
	fwrite(&rsaPri, 1, sizeof(RSA_PRIVATE_KEY), fp);
	fclose(fp);
	fp = NULL;


	//RSA私钥加密公钥解密
	//1.RSA私钥加密
	nInDataLen = (128 -11);
	nOutDataLen = 128;
	memcpy(tmpbuf, buffer, nInDataLen);
	dwRet = Dongle_RsaPri(hDongle, wPriID, FLAG_ENCODE, buffer, nInDataLen, buffer, &nOutDataLen);
	printf("RSA private key encode. Return: 0x%08X\n", dwRet);

	//2.RSA公钥解密
	nInDataLen = 128;
	nOutDataLen = (128-11);
	dwRet = Dongle_RsaPub(hDongle, FLAG_DECODE, &rsaPub, buffer, nInDataLen, buffer, &nOutDataLen);
	printf("RSA public key decode. Return: 0x%08X\n", dwRet);

	if (memcmp(tmpbuf, buffer, 117) == 0)
		printf("the private encode and public decode result is right. \n");
	else
		printf("the private encode and public decode result is wrong. \n");

	printf("\n");

	//RSA公钥加密私钥解密
	//1.RSA公钥加密
	for (i = 0 ; i < 128; i++) buffer[i] = i;
	nInDataLen  = (128-11);
	nOutDataLen = 128;
	memcpy(tmpbuf, buffer, nInDataLen);
	dwRet = Dongle_RsaPub(hDongle, FLAG_ENCODE, &rsaPub, buffer, nInDataLen, buffer, &nOutDataLen);
	printf("RSA public key encode. Return: 0x%08X\n", dwRet);

	//2.RSA私钥解密
	nInDataLen  = 128;
	nOutDataLen = (128 -11);
	dwRet = Dongle_RsaPri(hDongle, wPriID, FLAG_DECODE, buffer, nInDataLen, buffer, &nOutDataLen);
	printf("RSA private key decode. Return: 0x%08X\n", dwRet);

	if (memcmp(tmpbuf, buffer, 117) == 0)
		printf("the public encode and private decode result is right. \n");
	else
		printf("the public encode and private decode result is wrong. \n");

	//删除锁内私钥
	dwRet = Dongle_DeleteFile(hDongle, FILE_PRIKEY_RSA, wPriID);
	printf("Delete RSA private key file. Return: 0x%08X\n", dwRet);

	//关闭加密锁
	dwRet = Dongle_Close(hDongle);
	printf("Close Dongle ARM. Return: 0x%08X\n", dwRet);

	return 0;
}


*/
