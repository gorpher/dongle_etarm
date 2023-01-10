package dongle_etarm

import (
	"syscall"
	"unsafe"
)

var (
	modDongle = syscall.NewLazyDLL("Dongle_d.dll")
	//DWORD WINAPI Dongle_Enum(DONGLE_INFO * pDongleInfo, int * pCount);
	procDongleEnum = modDongle.NewProc("Dongle_Enum")
	//DWORD WINAPI Dongle_Open(DONGLE_HANDLE//phDongle, int nIndex);
	procDongleOpen = modDongle.NewProc("Dongle_Open")
	//DWORD WINAPI Dongle_Close(DONGLE_HANDLE hDongle);
	procDongleClose = modDongle.NewProc("Dongle_Close")
	//DWORD WINAPI Dongle_ResetState(DONGLE_HANDLE hDongle);
	procDongleResetstate = modDongle.NewProc("Dongle_ResetState")
	//	DWORD WINAPI Dongle_GenRandom(DONGLE_HANDLE hDongle, int nLen, BYTE * pRandom);
	procDongleGenrandom = modDongle.NewProc("Dongle_GenRandom")
	//	DWORD WINAPI Dongle_LEDControl(DONGLE_HANDLE hDongle, int nFlag);
	procDongleLedcontrol = modDongle.NewProc("Dongle_LEDControl")
	//DWORD WINAPI Dongle_SwitchProtocol(DONGLE_HANDLE hDongle, int nFlag);
	procDongleSwitchprotocol = modDongle.NewProc("Dongle_SwitchProtocol")
	//DWORD WINAPI Dongle_RFS(DONGLE_HANDLE hDongle);
	procDongleRfs = modDongle.NewProc("Dongle_RFS")
	//DWORD WINAPI Dongle_CreateFile(DONGLE_HANDLE hDongle, int nFileType, WORD wFileID, void * pFileAttr);
	procDongleCreatefile = modDongle.NewProc("Dongle_CreateFile")
	//DWORD WINAPI Dongle_WriteFile(DONGLE_HANDLE hDongle, int nFileType, WORD wFileID, WORD wOffset, BYTE * pInData, int nDataLen);
	procDongleWritefile = modDongle.NewProc("Dongle_WriteFile")
	//DWORD WINAPI Dongle_ReadFile(DONGLE_HANDLE hDongle, WORD wFileID, WORD wOffset, BYTE * pOutData, int nDataLen);
	procDongleReadfile = modDongle.NewProc("Dongle_ReadFile")
	//DWORD WINAPI Dongle_DownloadExeFile(DONGLE_HANDLE hDongle, EXE_FILE_INFO * pExeFileInfo, int nCount);
	procDongleDownloadexefile = modDongle.NewProc("Dongle_DownloadExeFile")
	//DWORD WINAPI Dongle_RunExeFile(DONGLE_HANDLE hDongle, WORD wFileID, BYTE * pInOutBuf, WORD wInOutBufLen, int * pMainRet);
	procDongleRunexefile = modDongle.NewProc("Dongle_RunExeFile")
	//DWORD WINAPI Dongle_DeleteFile(DONGLE_HANDLE hDongle, int nFileType, WORD wFileID);
	procDongleDeletefile = modDongle.NewProc("Dongle_DeleteFile")
	//DWORD WINAPI Dongle_ListFile(DONGLE_HANDLE hDongle, int nFileType, void* pFileList, int * pDataLen);
	procDongleListfile = modDongle.NewProc("Dongle_ListFile")
	//DWORD WINAPI Dongle_GenUniqueKey(DONGLE_HANDLE hDongle,int nSeedLen, BYTE * pSeed, char * pPIDstr, char * pAdminPINstr);
	procDongleGenuniquekey = modDongle.NewProc("Dongle_GenUniqueKey")
	//DWORD WINAPI Dongle_VerifyPIN(DONGLE_HANDLE hDongle, int nFlags, char * pPIN, int * pRemainCount);
	procDongleVerifypin = modDongle.NewProc("Dongle_VerifyPIN")
	//DWORD WINAPI Dongle_ChangePIN(DONGLE_HANDLE hDongle, int nFlags, char * pOldPIN, char * pNewPIN, int nTryCount);
	procDongleChangepin = modDongle.NewProc("Dongle_ChangePIN")
	//DWORD WINAPI Dongle_ResetUserPIN(DONGLE_HANDLE hDongle, char * pAdminPIN);
	procDongleResetuserpin = modDongle.NewProc("Dongle_ResetUserPIN")
	//DWORD WINAPI Dongle_SetUserID(DONGLE_HANDLE hDongle, DWORD dwUserID);
	procDongleSetuserid = modDongle.NewProc("Dongle_SetUserID")
	//DWORD WINAPI Dongle_GetDeadline(DONGLE_HANDLE hDongle, DWORD * pdwTime);
	procDongleGetdeadline = modDongle.NewProc("Dongle_GetDeadline")
	//DWORD WINAPI Dongle_SetDeadline(DONGLE_HANDLE hDongle, DWORD dwTime);
	procDongleSetdeadline = modDongle.NewProc("Dongle_SetDeadline")
	//DWORD WINAPI Dongle_GetUTCTime(DONGLE_HANDLE hDongle, DWORD * pdwUTCTime);
	procDongleGetutctime = modDongle.NewProc("Dongle_GetUTCTime")
	//DWORD WINAPI Dongle_ReadData(DONGLE_HANDLE hDongle, int nOffset, BYTE* pData, int nDataLen);
	procDongleReaddata = modDongle.NewProc("Dongle_ReadData")
	//DWORD WINAPI Dongle_WriteData(DONGLE_HANDLE hDongle, int nOffset, BYTE * pData, int nDataLen);
	procDongleWritedata = modDongle.NewProc("Dongle_WriteData")
	//DWORD WINAPI Dongle_ReadShareMemory(DONGLE_HANDLE hDongle, BYTE * pData);
	procDongleReadsharememory = modDongle.NewProc("Dongle_ReadShareMemory")
	//DWORD WINAPI Dongle_WriteShareMemory(DONGLE_HANDLE hDongle, BYTE * pData, int nDataLen);
	procDongleWritesharememory = modDongle.NewProc("Dongle_WriteShareMemory")
	//DWORD WINAPI Dongle_RsaGenPubPriKey(DONGLE_HANDLE hDongle, WORD wPriFileID, RSA_PUBLIC_KEY * pPubBakup, RSA_PRIVATE_KEY * pPriBakup);
	procDongleRsagenpubprikey = modDongle.NewProc("Dongle_RsaGenPubPriKey")
	//DWORD WINAPI Dongle_RsaPri(DONGLE_HANDLE hDongle, WORD wPriFileID, int nFlag, BYTE * pInData, int nInDataLen, BYTE * pOutData, int * pOutDataLen);
	procDongleRsapri = modDongle.NewProc("Dongle_RsaPri")
	//DWORD WINAPI Dongle_RsaPub(DONGLE_HANDLE hDongle, int nFlag, RSA_PUBLIC_KEY * pPubKey, BYTE * pInData, int nInDataLen, BYTE * pOutData, int * pOutDataLen);
	procDongleRsapub = modDongle.NewProc("Dongle_RsaPub")
	//DWORD WINAPI Dongle_EccGenPubPriKey(DONGLE_HANDLE hDongle, WORD wPriFileID, ECCSM2_PUBLIC_KEY * pPubBakup, ECCSM2_PRIVATE_KEY * pPriBakup);
	procDongleEccgenpubprikey = modDongle.NewProc("Dongle_EccGenPubPriKey")
	//DWORD WINAPI Dongle_EccSign(DONGLE_HANDLE hDongle, WORD wPriFileID, BYTE * pHashData, int nHashDataLen, BYTE * pOutData);
	procDongleEccsign = modDongle.NewProc("Dongle_EccSign")
	//DWORD WINAPI Dongle_EccVerify(DONGLE_HANDLE hDongle, ECCSM2_PUBLIC_KEY * pPubKey, BYTE * pHashData, int nHashDataLen, BYTE * pSign);
	procDongleEccverify = modDongle.NewProc("Dongle_EccVerify")
	//DWORD WINAPI Dongle_SM2GenPubPriKey(DONGLE_HANDLE hDongle, WORD wPriFileID, ECCSM2_PUBLIC_KEY * pPubBakup, ECCSM2_PRIVATE_KEY * pPriBakup);
	procDongleSm2genpubprikey = modDongle.NewProc("Dongle_SM2GenPubPriKey")
	//DWORD WINAPI Dongle_SM2Sign(DONGLE_HANDLE hDongle, WORD wPriFileID, BYTE * pHashData, int nHashDataLen, BYTE * pOutData);
	procDongleSm2sign = modDongle.NewProc("Dongle_SM2Sign")
	//DWORD WINAPI Dongle_SM2Verify(DONGLE_HANDLE hDongle, ECCSM2_PUBLIC_KEY * pPubKey, BYTE * pHashData, int nHashDataLen, BYTE * pSign);
	procDongleSm2verify = modDongle.NewProc("Dongle_SM2Verify")
	//DWORD WINAPI Dongle_TDES(DONGLE_HANDLE hDongle, WORD wKeyFileID, int nFlag, BYTE * pInData, BYTE * pOutData, int nDataLen);
	procDongleTdes = modDongle.NewProc("Dongle_TDES")
	//DWORD WINAPI Dongle_SM4(DONGLE_HANDLE hDongle, WORD wKeyFileID, int nFlag, BYTE * pInData, BYTE * pOutData, int nDataLen);
	procDongleSm4 = modDongle.NewProc("Dongle_SM4")
	//DWORD WINAPI Dongle_HASH(DONGLE_HANDLE hDongle, int nFlag, BYTE * pInData, int nDataLen, BYTE * pHash);
	procDongleHash = modDongle.NewProc("Dongle_HASH")
	//DWORD WINAPI Dongle_Seed(DONGLE_HANDLE hDongle, BYTE * pSeed, int nSeedLen, BYTE * pOutData);
	procDongleSeed = modDongle.NewProc("Dongle_Seed")
	//DWORD WINAPI Dongle_LimitSeedCount(DONGLE_HANDLE hDongle, int nCount);
	procDongleLimitseedcount = modDongle.NewProc("Dongle_LimitSeedCount")
	//DWORD WINAPI Dongle_GenMotherKey(DONGLE_HANDLE hDongle, MOTHER_DATA * pInData);
	procDongleGenmotherkey = modDongle.NewProc("Dongle_GenMotherKey")
	//DWORD WINAPI Dongle_RequestInit(DONGLE_HANDLE hDongle, BYTE * pRequest);
	procDongleRequestinit = modDongle.NewProc("Dongle_RequestInit")
	//DWORD WINAPI Dongle_GetInitDataFromMother(DONGLE_HANDLE hDongle, BYTE * pRequest, BYTE * pInitData, int * pDataLen);
	procDongleGetinitdatafrommother = modDongle.NewProc("Dongle_GetInitDataFromMother")
	//DWORD WINAPI Dongle_InitSon(DONGLE_HANDLE hDongle, BYTE * pInitData, int nDataLen);
	procDongleInitson = modDongle.NewProc("Dongle_InitSon")
	//DWORD WINAPI Dongle_SetUpdatePriKey(DONGLE_HANDLE hDongle, RSA_PRIVATE_KEY * pPriKey);
	procDongleSetupdateprikey = modDongle.NewProc("Dongle_SetUpdatePriKey")
	//DWORD WINAPI Dongle_MakeUpdatePacket(DONGLE_HANDLE hDongle, char * pHID, int nFunc, int nFileType, WORD wFileID, int nOffset, BYTE * pBuffer, int nBufferLen, RSA_PUBLIC_KEY * pUPubKey, BYTE * pOutData, int * pOutDataLen);
	procDongleMakeupdatepacket = modDongle.NewProc("Dongle_MakeUpdatePacket")
	//DWORD WINAPI Dongle_MakeUpdatePacketFromMother(DONGLE_HANDLE hDongle, char * pHID, int nFunc, int nFileType, WORD wFileID, int nOffset, BYTE * pBuffer, int nBufferLen, BYTE * pOutData, int * pOutDataLen);
	procDongleMakeupdatepacketfrommother = modDongle.NewProc("Dongle_MakeUpdatePacketFromMother")
	//DWORD WINAPI Dongle_Update(DONGLE_HANDLE hDongle, BYTE * pUpdateData, int nDataLen);
	procDongleUpdate = modDongle.NewProc("Dongle_Update")
)

func DongleEnum() ([]DONGLE_INFO, error) {
	var dongleInfo DONGLE_INFO
	var count int

	if err := Dongle_Enum(&dongleInfo, &count); err != nil {
		return nil, err
	}
	res := &dongleInfo
	result := make([]DONGLE_INFO, count)
	for i := 0; i < count; i++ {
		result[i] = *res
		res = (*DONGLE_INFO)(unsafe.Pointer(uintptr(unsafe.Pointer(res)) + unsafe.Sizeof(*res)))
	}
	return result, nil
}

// Dongle_Enum
// @brief  枚举加密锁。本函数最多会枚举出32个hid设备和32个ccid设备。
//
// @param  pDongleInfo     [out]     设备信息的数组。当此参数为NULL时, pCount返回找到的设备的数目。
// @param  pCount 	       [out]     设备数目。该函数最多可以同时枚举出32个HID设备和32个CCID设备。
//
// @return DONGLE_SUCCESS            执行成功。
func Dongle_Enum(info *DONGLE_INFO, pCount *int) error {
	r1, _, _ := procDongleEnum.Call(uintptr(unsafe.Pointer(info)),
		uintptr(unsafe.Pointer(pCount)))

	return DongleErrno(syscall.Errno(r1))
}

//Dongle_Open
//@brief  打开指定的加密锁。
//@param  phDongle     [out]     句柄指针。如果打开成功,会被填充。
//@param  nIndex       [in]      基于0的索引值。指示打开找到的第几把加密锁。
//@return DONGLE_SUCCESS         打开成功。
func Dongle_Open(handle *DONGLE_HANDLE, nIndex int) error {
	r1, _, _ := procDongleOpen.Call(uintptr(unsafe.Pointer(handle)), uintptr(nIndex))
	return DongleErrno(syscall.Errno(r1))
}

// Dongle_ResetState
// @brief  清除PIN码验证状态。将加密锁状态变为匿名。
//
// @param  hDongle     [in]     打开的加密锁句柄。
//
// @return DONGLE_SUCCESS       执行成功。
func Dongle_ResetState(handle DONGLE_HANDLE) error {
	r1, _, _ := procDongleResetstate.Call(uintptr(handle))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_Close
//@brief  关闭打开的加密锁。
//@param  hDongle     [in]     打开的加密锁句柄。
//@return DONGLE_SUCCESS       关闭成功。
func Dongle_Close(handle DONGLE_HANDLE) error {
	r1, _, _ := procDongleClose.Call(uintptr(handle))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_GenRandom
//@brief  产生随机数。匿名权限即可操作。
//
//@param  hDongle          [in]      打开的加密锁句柄。
//@param  nLen             [in]      要产生的随机数的长度。nLen的取值范围为 1~128。
//@param  pRandom          [out]     随机数缓冲区。
//
//@return DONGLE_SUCCESS             获取随机数成功。
func Dongle_GenRandom(hDongle DONGLE_HANDLE, nLen int, pRandom *byte) error {
	r1, _, _ := procDongleGenrandom.Call(uintptr(hDongle), uintptr(nLen), uintptr(unsafe.Pointer(pRandom)))
	return DongleErrno(syscall.Errno(r1))
}

// Dongle_LEDControl
//@brief  LED灯的控制操作。匿名权限即可操作。
//
//@param  hDongle     [in]     打开的加密锁句柄。
//@param  nFlag       [in]     控制类型。例如：nFlag = LED_ON，表示控制LED为亮的状态；
//                             nFlag = LED_OFF，表示控制LED为灭的状态；nFlag = LED_BLINK，
//                             表示控制LED为闪烁的状态。
//
//@return DONGLE_SUCCESS       命令执行成功。
func Dongle_LEDControl(hDongle DONGLE_HANDLE, nFlag int) error {
	r1, _, _ := procDongleLedcontrol.Call(uintptr(hDongle), uintptr(nFlag))
	return DongleErrno(syscall.Errno(r1))
}

// Dongle_SwitchProtocol
//@brief  切换通讯协议。调用执行成功后加密锁会自动重启,打开的句柄hDongle会无效,接下来执行关闭操作会返回
//        0xF0000002的错误码，这属正常。如需继续操作，请重新枚举并打开锁。该操作必须要验证开发商PIN码之
//	       后方可使用。
//
//@param  hDongle       [in]   打开的加密锁句柄。
//@param  nFlag         [in]   协议类型。例如：nFlag值为PROTOCOL_HID，表示将加密锁切换为HID设备；
//                             nFlag值为PROTOCOL_CCID，表示将加密锁切换为CCID设备
//
//@return DONGLE_SUCCESS       执行成功。
func Dongle_SwitchProtocol(hDongle DONGLE_HANDLE, nFlag int) error {
	r1, _, _ := procDongleSwitchprotocol.Call(uintptr(hDongle), uintptr(nFlag))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_RFS
//@brief  一键恢复。即返回出厂状态，加密锁的PID、用户PIN码、开发商PIN码等，全部恢复到出厂状态所有写入的
//      数据都将被清空。另外，调用执行成功后加密锁会自动重启,打开的句柄hDongle会无效,接下来执行关闭操
//      做会返回0xF0000002的错误码，这属正常。如需继续操作，请重新枚举并打开锁。该操作需要开发商权限。
//
//@param  hDongle       [in]     打开的加密锁句柄。
//
//@return DONGLE_SUCCESS         执行成功。
func Dongle_RFS(hDongle DONGLE_HANDLE) error {
	r1, _, _ := procDongleRfs.Call(uintptr(hDongle))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_CreateFile
//@brief  创建文件。该函数不支持可执行文件的创建。该操作需要开发商权限。
//
//@param  hDongle       [in]     打开的加密锁句柄。
//@param  nFileType     [in]     文件类型。
//                               nFileType = FILE_DATA，表示创建数据文件；对数据文件有以下说明：
//                                  1.文件大小设为252字节时,最多可创建54个文件,即占用空间13608字节
//                                  2.文件大小设为1024字节时,最多可创建31个文件，即占用空间31744字节
//                                  3.文件大小设为4096字节时，最多可创建9个文件,即占用空间36864字节
//                               nFileType = FILE_PRIKEY_RSA，表示创建RSA私钥文件；
//                               nFileType = FILE_PRIKEY_ECCSM2，表示创建ECCSM2私钥文件；
//
//                               nFileType = FILE_KEY，表示创建SM4和3DES密钥文件；
//                               不支持nFileType = FILE_EXE的文件类型。
//@param  wFileID       [in]     文件ID。
//@param  pFileAttr     [in]     文件的属性。参数的结构为：DATA_FILE_ATTR、PRIKEY_FILE_ATTR或KEY_FILE_ATTR。
//
//@return DONGLE_SUCCESS         创建文件成功。
func Dongle_CreateFile(hDongle DONGLE_HANDLE, nFileType int, wFileID uint16, pFukeAttr *byte) error {
	r1, _, _ := procDongleCreatefile.Call(uintptr(hDongle), uintptr(nFileType), uintptr(wFileID), uintptr(unsafe.Pointer(pFukeAttr)))
	return DongleErrno(syscall.Errno(r1))
}

func Dongle_CreateFILE_DATA(hDongle DONGLE_HANDLE, wFileID uint16, pFukeAttr *DATA_FILE_ATTR) error {
	r1, _, _ := procDongleCreatefile.Call(uintptr(hDongle), uintptr(FILE_DATA), uintptr(wFileID), uintptr(unsafe.Pointer(pFukeAttr)))
	return DongleErrno(syscall.Errno(r1))
}
func Dongle_CreateFILE_PRIKEY_RSA(hDongle DONGLE_HANDLE, wFileID uint16, pFukeAttr *PRIKEY_FILE_ATTR) error {
	r1, _, _ := procDongleCreatefile.Call(uintptr(hDongle), uintptr(FILE_PRIKEY_RSA), uintptr(wFileID), uintptr(unsafe.Pointer(pFukeAttr)))
	return DongleErrno(syscall.Errno(r1))
}

func Dongle_CreateFILE_PRIKEY_ECCSM2(hDongle DONGLE_HANDLE, wFileID uint16, pFukeAttr *KEY_FILE_ATTR) error {
	r1, _, _ := procDongleCreatefile.Call(uintptr(hDongle), uintptr(FILE_PRIKEY_ECCSM2), uintptr(wFileID), uintptr(unsafe.Pointer(pFukeAttr)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_WriteFile
//@brief  写文件。该函数不支持可执行文件的写入操作，且该操作需要开发商权限。
//
//@param  hDongle       [in]     打开的加密锁句柄。
//@param  nFileType     [in]     文件类型。例如，
//                               nFileType = FILE_DATA，表示创建数据文件；
//                               nFileType = FILE_PRIKEY_RSA，表示创建RSA私钥文件；
//                               nFileType = FILE_PRIKEY_ECCSM2，表示创建ECCSM2私钥文件；
//                               nFileType = FILE_KEY，表示创建SM4和3DES密钥文件；
//                               不支持nFileType = FILE_EXE的文件类型。
//@param  wFileID       [in]     文件ID。
//@param  wOffset       [in]     文件偏移。文件写入的起始偏移量。
//@param  pInData       [in]     准备写入的数据。
//@param  nDataLen      [in]     参数pInData的大小。
//
//@return DONGLE_SUCCESS         写入文件成功。
func Dongle_WriteFile(hDongle DONGLE_HANDLE, nFileType int, wFileID uint16, wOffset uint16, pInData *byte, nDataLen int) error {
	r1, _, _ := procDongleWritefile.Call(uintptr(hDongle),
		uintptr(nFileType),
		uintptr(wFileID),
		uintptr(wOffset),
		uintptr(unsafe.Pointer(pInData)),
		uintptr(nDataLen))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_ReadFile
//@brief  读取加密锁内的数据文件。数据文件的读取权限取决于创建时的设定。
//
//@param  hDongle      [in]         打开的加密锁句柄。
//@param  wFileID      [in]         文件ID。
//@param  wOffset      [in]         文件偏移量。
//@param  pOutData     [in]         数据缓冲区。
//@param  nDataLen     [out]        参数pOutData的长度。
//
//@return DONGLE_SUCCESS            读取数据文件成功
func Dongle_ReadFile(hDongle DONGLE_HANDLE, wFileID uint16, wOffset uint16, pOutData *byte, nDataLen int) error {
	r1, _, _ := procDongleReadfile.Call(uintptr(hDongle), uintptr(wFileID), uintptr(wOffset), uintptr(unsafe.Pointer(pOutData)), uintptr(nDataLen))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_DownloadExeFile
//@brief  批量下载可执行文件。锁内可执行文件的数量不能超过64个，可执行文件的总大小不能超过64K，
//        该操作需要验证管理员权限
//
//@param  hDongle          [in]     打开的加密锁句柄。
//@param  pExeFileInfo     [in]     可执行文件信息的数组。
//@param  nCount           [in]     即可执行文件的数量。
//
//@return DONGLE_SUCCESS            批量下载可执行文件成功。
func Dongle_DownloadExeFile(hDongle DONGLE_HANDLE, pExeFileInfo *EXE_FILE_INFO, nCount int) error {
	r1, _, _ := procDongleDownloadexefile.Call(uintptr(hDongle), uintptr(unsafe.Pointer(pExeFileInfo)), uintptr(nCount))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_RunExeFile
//@brief  运行指定的锁内可执行程序。运行可执行文件的权限，取决于批量下载时，每个可执行文件的设置，
//        即，EXE_FILE_INFO中的m_Priv参数。输入输出数据的最大长度不能超过1024字节，输入输出数据缓
//        冲区pInOutBuf对应锁内的InOutBuf。
//
//@param  hDongle            [in]             打开的加密锁句柄。
//@param  wFileID            [in]             可执行文件的文件ID。
//@param  pInOutBuf          [in,out]         输入输出数据缓冲区。
//@param  wInOutBufLen       [in]             输入输出数据缓冲区pInOutBuf的大小。
//@param  pMainRet           [out]            锁内可执行程序main函数的返回值，可以为NULL。
//
//@return DONGLE_SUCCESS                  运行指定的可执行文件成功。
func Dongle_RunExeFile(hDongle DONGLE_HANDLE, wFileID uint16, pInOutBuf *byte, wInOutBufLen uint16, pMainRet *int) error {
	r1, _, _ := procDongleRunexefile.Call(uintptr(hDongle), uintptr(wFileID), uintptr(unsafe.Pointer(pInOutBuf)), uintptr(wInOutBufLen), uintptr(unsafe.Pointer(pMainRet)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_DeleteFile
//@brief  删除文件。需要开发商权限。
//
//@param  hDongle       [in]      打开的加密锁句柄。
//@param  nFileType     [in]      文件类型。
//@param  wFileID       [in]      文件ID。
//
//@return DONGLE_SUCCESS          删除文件成功
func Dongle_DeleteFile(hDongle DONGLE_HANDLE, nFileType int, wFileID uint16) error {
	r1, _, _ := procDongleDeletefile.Call(uintptr(hDongle), uintptr(nFileType), uintptr(wFileID))
	return DongleErrno(syscall.Errno(r1))

}

//Dongle_ListFile
//@brief 列文件。需要开发商权限。
//
//@param  hDongle       [in]         打开的加密锁句柄。
//@param  nFileType     [in]         指示文件类型。例如，FILE_DATA等。
//@param  pFileList     [in]         pList：输出文件的列表 (此参数为NULL时, pLen中返回所需的缓冲区长度)
//                                   当nFileType = FILE_DATA时,       pFileList指向DATA_FILE_LIST结构；
//                                   当nFileType = FILE_PRIKEY_RSA时, pFileList指向PRIKEY_FILE_LIST结构；
//                                   当nFileType = FILE_PRIKEY_ECCSM2时, pFileList指向PRIKEY_FILE_LIST结构
//                                   当nFileType = FILE_KEY时,        pFileList指向KEY_FILE_LIST结构；
//	                                  当nFileType = FILE_EXE时,        pFileList指向EXE_FILE_LIST结构。
//@param  pDataLen      [in,out]     参数pFileList的输入长度，执行成功返回pFileList的字节长度。
//
//@return DONGLE_SUCCESS             列文件成功。
func Dongle_ListFile(hDongle DONGLE_HANDLE, nFileType int, pFileList *DATA_FILE_LIST, pDataLen *int) error {
	r1, _, _ := procDongleListfile.Call(uintptr(hDongle), uintptr(nFileType), uintptr(unsafe.Pointer(pFileList)), uintptr(unsafe.Pointer(pDataLen)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_GenUniqueKey
//@brief 唯一化锁。输入种子码产生PID和开发商PIN码，需要开发商权限，执行成功后加密锁自动回到匿名态。
//		  产生开发商PIN码的目的是为了使密码肯定不是弱密码，产生的开发商PIN可以借助Dongle_ChangePIN
//       进行更改。另外，种子码一定要牢记，否则任何人永远无法得知开发商PIN码。
//
//@param hDongle		   [in]       打开的加密锁句柄。
//@param nSeedLen		   [in]       参数pSeed的缓冲区长度。
//@param pSeed            [in]       种子码的缓冲区。
//@param pPIDStr          [out]      函数执行成功返回PID。该缓冲区大小至少应该为8字节，返回一个
//                                   8字节的以0终止的ansi字符串。
//@param pAdminPINstr     [out]      函数执行成功返回开发商PIN码。该缓冲区大小至少应该为16字节，
//                                   返回字符串长度为16字节的以0终止的ansi字符串。
//
//@return DONGLE_SUCCESS             唯一化锁成功。
func Dongle_GenUniqueKey(hDongle DONGLE_HANDLE, nSeedLen int, pSeed *byte, pPIDstr *byte, pAdminPINstr *byte) error {
	r1, _, _ := procDongleGenuniquekey.Call(uintptr(hDongle), uintptr(nSeedLen), uintptr(unsafe.Pointer(pSeed)), uintptr(unsafe.Pointer(pPIDstr)), uintptr(unsafe.Pointer(pAdminPINstr)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_VerifyPIN
//@brief  校验密码
//
//@param  hDongle		    [in]       打开的加密锁句柄。
//@param  nFlags           [in]       PIN码类型。参数取值为FLAG_USERPIN或者FLAG_ADMINPIN。
//@param  pPIN             [in]       PIN码，0终止的ansi字符串。
//@param  pRemainCount     [out]      剩余重试次数。返回0表示已锁死；1~253表示剩余次数；255表示不限制重试次数。
//
//@return DONGLE_SUCCESS              校验成功。如果校验失败，函数的返回值中也含有剩余的重试次数，
//                                    (错误码 & 0xFFFFFF00) == DONGLE_INCORRECT_PIN，即后两位表示剩余次数。
func Dongle_VerifyPIN(hDongle DONGLE_HANDLE, nFlags int, pPIN string, pRemainCount *int) error {
	r1, _, _ := procDongleVerifypin.Call(uintptr(hDongle), uintptr(nFlags), Lpstr(pPIN), uintptr(unsafe.Pointer(pRemainCount)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_ChangePIN
//@brief  更改密码
//
//@param  hDongle       [in]     打开的加密锁句柄。
//@param  nFlags        [in]     PIN码类型。参数取值为FLAG_USERPIN或者FLAG_ADMINPIN。
//@param  pOldPIN       [in]     旧的PIN码缓冲区。必须是一个字符串长度为16字节的0终止的ansi字符串,且可以是中文。
//@param  pNewPIN       [in]     新的PIN码缓冲区。必须是一个字符串长度为16字节的0终止的ansi字符串。
//@param  nTryCount     [in]     重试次数。该参数的取值范围为1~255,其中255表示不限制重试次数。
//
//@return DONGLE_SUCCESS         修改密码成功。
func Dongle_ChangePIN(hDongle DONGLE_HANDLE, nFlags int, pOldPIN string, pNewPIN string, nTryCount int) error {
	r1, _, _ := procDongleChangepin.Call(uintptr(hDongle), uintptr(nFlags), Lpstr(pOldPIN), Lpstr(pNewPIN), uintptr(nTryCount))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_ResetUserPIN
//@brief  重置用户PIN码。空锁(即PID=FFFFFFFF)不能重置密码。执行成功后用户密码恢复为出厂默认 CONST_USERPIN
//
//@param  hDongle       [in]     打开的加密锁句柄。
//@param  pAdminPIN     [in]     开发商PIN码缓冲区。长度为16字节的0终止的ansi字符串
//
//@return DONGLE_SUCCESS         重置用户PIN码成功。
func Dongle_ResetUserPIN(hDongle DONGLE_HANDLE, pAdminPIN string) error {
	r1, _, _ := procDongleResetuserpin.Call(uintptr(hDongle), Lpstr(pAdminPIN))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_SetUserID
//@brief  设置用户ID。需要开发商权限。
//
//@param  hDongle      [in]     打开的加密锁句柄。
//@param  dwUserID     [in]     用户ID。
//
//@return DONGLE_SUCCESS        重置用户PIN码成功。
func Dongle_SetUserID(hDongle DONGLE_HANDLE, dwUserID uint32) error {
	r1, _, _ := procDongleSetuserid.Call(uintptr(hDongle), uintptr(dwUserID))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_GetDeadline
//@brief  获取加密锁到期时间。匿名权限获取。
//
//@param  hDongle     [in]      打开的加密锁句柄。
//@param  pdwTime     [out]     获取的到期UTC时间值。
//                              若*pdwTime = 0XFFFFFFFF，表示不限制到期时间
//                              若(*pdwTime & 0XFFFF0000) == 0，值表示还剩余几小时
//                              若(*pdwTime & 0XFFFF0000) != 0，值表示到期的UTC的时间，可以通过gmtime等将该值进行转换。
//
//@return DONGLE_SUCCESS        获取加密锁到期时间成功。
func Dongle_GetDeadline(hDongle DONGLE_HANDLE, pdwTime *uint32) error {
	r1, _, _ := procDongleGetdeadline.Call(uintptr(hDongle), uintptr(unsafe.Pointer(pdwTime)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_SetDeadline
//@brief  设置加密锁的到期时间。该操作需要管理员权限。
//
//@param  hDongle     [in]     打开的加密锁句柄。
//@param  dwTime      [in]     时间值。说明：
//                             1.设置可用小时数，范围在1~65535，例如dwTime = 24。这种情况在校验了用户PIN码后开始计时。
//                             2.设置到期的年与日时分秒。可通过函数time或者mktime 取得即时的utc时间值(utc值都大于65535)；
//                             3.取消到期时间限制，此时dwTime的值只能为0xFFFFFFFF。
//
//@return DONGLE_SUCCESS       设置加密锁到期时间成功。
func Dongle_SetDeadline(hDongle DONGLE_HANDLE, dwTime uint32) error {
	r1, _, _ := procDongleSetdeadline.Call(uintptr(hDongle), uintptr(dwTime))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_GetUTCTime
//@brief  获取加密锁的UTC时间。
//
//@param  hDongle         [in]      打开的加密锁句柄。
//@param  pdwUTCTime      [out]     UTC时间值指针。
//
//@return DONGLE_SUCCESS       设置加密锁到期时间成功。
func Dongle_GetUTCTime(hDongle DONGLE_HANDLE, pdwUTCTime *uint32) error {
	r1, _, _ := procDongleGetutctime.Call(uintptr(hDongle), uintptr(unsafe.Pointer(pdwUTCTime)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_ReadData
//@brief  读取锁内数据区数据。数据区大小共8k，前4k(0~4095)的读写没有权限限制，后4k(4096~8191)任意权限可读，
//        但是只有开发商权限可写。
//
//@param  hDongle      [in]      打开的加密锁句柄。
//@param  nOffset      [in]      起始偏移。范围在0~8191
//@param  pData        [out]     读取的数据缓冲区。
//@param  nDataLen     [in]      参数pData的缓冲区大小。
//
//@return  DONGLE_SUCCESS        读取数据区数据成功。
func Dongle_ReadData(hDongle DONGLE_HANDLE, nOffset int, pData *byte, nDataLen int) error {
	r1, _, _ := procDongleReaddata.Call(uintptr(hDongle), uintptr(nOffset), uintptr(unsafe.Pointer(pData)), uintptr(nDataLen))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_WriteData
//@brief  写入锁内数据区数据。数据区大小共8k，前4k(0~4095)的读写没有权限限制，后4k(4096~8191)任意权限可读，
//        但是只有开发商权限可写。
//
//@param  hDongle      [in]      打开的加密锁句柄。
//@param  nOffset      [in]      起始偏移。范围在0~8191
//@param  pData        [in]      写入的数据缓冲区。
//@param  nDataLen     [in]      参数pData的缓冲区大小。
//
//@return  DONGLE_SUCCESS        写入数据区数据成功。
func Dongle_WriteData(hDongle DONGLE_HANDLE, nOffset int, pData *byte, nDataLen int) error {
	r1, _, _ := procDongleWritedata.Call(uintptr(hDongle), uintptr(nOffset), uintptr(unsafe.Pointer(pData)), uintptr(nDataLen))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_ReadShareMemory
//@brief  获取共享内存数据。共享内存总大小为32字节。没有权限限制，掉电后数据自动擦除。
//
//@param  hDongle     [in]      打开的加密锁句柄。
//@param  pData       [out]     输出的数据。输出共享内存的数据，固定为32个字节。
//
//@return  DONGLE_SUCCESS       获取共享内存数据成功。
func Dongle_ReadShareMemory(hDongle DONGLE_HANDLE, pData *byte) error {
	r1, _, _ := procDongleReadsharememory.Call(uintptr(hDongle), uintptr(unsafe.Pointer(pData)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_WriteShareMemory
//@brief  设置共享内存数据。没有权限限制，掉电后数据自动擦除。
//
//@param  hDongle      [in]     打开的加密锁句柄。
//@param  pData        [in]     输入数据。
//@param  nDataLen     [in]     参数pData的缓冲区大小。长度不能超过32。
//
//@return  DONGLE_SUCCESS       设置共享内存数据成功。
func Dongle_WriteShareMemory(hDongle DONGLE_HANDLE, pData *byte, nDataLen int) error {
	r1, _, _ := procDongleWritesharememory.Call(uintptr(hDongle), uintptr(unsafe.Pointer(pData)), uintptr(nDataLen))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_RsaGenPubPriKey
//@brief  产生RSA公钥和私钥。使用该函数之前需要先创建一个RSA私钥文件。需要开发商权限。成功后注意保存RSA公私钥数据。
//
//@param  hDongle        [in]      打开的加密锁句柄。
//@param  wPriFileID     [in]      RSA私钥文件ID。
//@param  pPubBakup      [out]     RSA公钥数据。
//@param  pPriBakup      [out]     RSA私钥数据。
//
//@return DONGLE_SUCCESS           产生RSA公私钥成功。
func Dongle_RsaGenPubPriKey(hDongle DONGLE_HANDLE, wPriFileID uint16, rsaPub *RSA_PUBLIC_KEY, pPriBakup *RSA_PRIVATE_KEY) error {
	pPubBakup := (*byte)(unsafe.Pointer(rsaPub))
	r1, _, _ := procDongleRsagenpubprikey.Call(uintptr(hDongle), uintptr(wPriFileID), uintptr(unsafe.Pointer(pPubBakup)), uintptr(unsafe.Pointer(pPriBakup)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_RsaPri
//@brief  RSA私钥运算。函数的使用权限取决于锁内RSA私钥文件的权限，在RSA私钥文件创建时设定。说明：
//        1.对于加密运算,输入数据长度必须小于私钥ID为wPriFileID的密钥长度除以8再减去11,以便在函数内部进行padding
//        2.对于解密运算,输入数据长度必须与wPriFileID中指示的密钥长度相一致(比如1024位密钥时为128，2048时为256)
//        3.加密时内部padding方式为:PKCS1_TYPE_1 (即第二个字节为0x01,空数据填充0XFF)
//
//@param  hDongle         [in]         打开的加密锁句柄。
//@param  wPriFileID      [in]         RSA私钥文件ID。
//@param  nFlag           [in]         运算类型。例如，FLAG_ENCODE表示加密运算；FLAG_DECODE表示解密运算。
//@param  pInData         [in]         输入数据。
//@param  nInDataLen      [in]         参数pInData的大小
//@param  pOutData        [out]        输出数据缓冲区。
//@param  pOutDataLen     [in,out]     参数pOutData的大小和返回的数据大小。
//
//@return DONGLE_SUCCESS               RSA私钥运算成功。
func Dongle_RsaPri(hDongle DONGLE_HANDLE, wPriFileID uint16, nFlag int, pInData *byte, nInDataLen int, pOutData *byte, pOutDataLen *int) error {
	r1, _, _ := procDongleRsapri.Call(uintptr(hDongle), uintptr(wPriFileID), uintptr(nFlag),
		uintptr(unsafe.Pointer(pInData)), uintptr(nInDataLen),
		uintptr(unsafe.Pointer(pOutData)), uintptr(unsafe.Pointer(pOutDataLen)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_RsaPub
//@brief  RSA公钥运算。匿名权限可调用。说明：
//        1.对于加密运算,输入数据长度必须小于pPubKey中指示的密钥长度除以8再减去11,以便在函数内部进行padding
//        2.对于解密运算,输入数据长度必须与pPubKey中指示的密钥长度相一致(比如1024位密钥时为128，2048时为256)
//        3.加密时内部padding方式为:PKCS1_TYPE_2 (即第二个字节为0x02,空数据填充随机数)
//
//@param  hDongle         [in]         打开的加密锁句柄。
//@param  nFlag           [in]         运算类型。例如，FLAG_ENCODE表示加密运算；FLAG_DECODE表示解密运算。
//@param  pPubKey         [in]         RSA公钥数据。该数据来源于生成RSA公私钥时的公钥数据。
//@param  pInData         [in]         输入数据。
//@param  nInDataLen      [in]         参数pInData的大小。
//@param  pOutData        [out]        输出数据缓冲区。
//@param  pOutDataLen     [in,out]     参数pOutData的大小和返回的数据大小。
//
//@return DONGLE_SUCCESS               RSA公钥运算成功。
func Dongle_RsaPub(hDongle DONGLE_HANDLE, nFlag int, pPubKey *RSA_PUBLIC_KEY, pInData *byte, nInDataLen int, pOutData *byte, pOutDataLen *int) error {
	r1, _, _ := procDongleRsapub.Call(uintptr(hDongle), uintptr(nFlag),
		uintptr(unsafe.Pointer(pPubKey)), uintptr(unsafe.Pointer(pInData)), uintptr(nInDataLen),
		uintptr(unsafe.Pointer(pOutData)), uintptr(unsafe.Pointer(pOutDataLen)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_EccGenPubPriKey
//@brief  产生ECC公钥和私钥。使用该函数之前需要先创建一个ECC私钥文件。需要开发商权限。成功后注意保存ECC公私钥数据。
//
//@param  hDongle        [in]      打开的加密锁句柄。
//@param  wPriFileID     [in]      ECC私钥文件ID。
//@param  pPubBakup      [out]     ECC公钥数据。
//@param  pPriBakup      [out]     ECC私钥数据。
//
//@return DONGLE_SUCCESS           产生ECC公私钥成功。
func Dongle_EccGenPubPriKey(hDongle DONGLE_HANDLE, wPriFileID uint16, pPubBakup *ECCSM2_PUBLIC_KEY, pPriBakup *ECCSM2_PRIVATE_KEY) error {
	r1, _, _ := procDongleEccgenpubprikey.Call(uintptr(hDongle), uintptr(wPriFileID),
		uintptr(unsafe.Pointer(pPubBakup)), uintptr(unsafe.Pointer(pPriBakup)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_EccSign
//@brief  ECC私钥签名。函数的使用权限取决于锁内ECC私钥文件的权限，在ECC私钥文件创建时设定。说明：
//        1.锁内签名算法为: ECDSA_Sign
//        2.输入的Hash值的长度与ECC私钥的密钥长度有关(如果密钥是192位的,则hash值长度不能超过24(192/8 = 24)字节)
//                                                   (如果密钥是256位的,则hash值长度不能超过32(256/8 = 32)字节)
//        3.曲线参数为:EC_NIST_PRIME_192及EC_NIST_PRIME_256
//
//@param  hDongle          [in]      打开的加密锁句柄。
//@param  wPriFileID       [in]      ECC私钥文件ID。
//@param  pHashData        [in]      Hash数据。
//@param  nHashDataLen     [in]      参数pHashData的大小。
//@param  pOutData         [out]     签名数据。大小固定为64字节(256位ECC时是正好,192位ECC时的位会补0)
//
//@return DONGLE_SUCCESS             表示签名成功。
func Dongle_EccSign(hDongle DONGLE_HANDLE, wPriFileID uint16, pHashData *byte, nHashDataLen int, pOutData *byte) error {
	r1, _, _ := procDongleEccsign.Call(uintptr(hDongle), uintptr(wPriFileID),
		uintptr(unsafe.Pointer(pHashData)), uintptr(nHashDataLen), uintptr(unsafe.Pointer(pOutData)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_EccVerify
//@brief  ECC公钥验签。函数的使用权限取决于锁内ECC私钥文件的权限，在ECC私钥文件创建时设定。说明：
//        1.锁内签名算法为: ECDSA_Verify
//        2.输入的Hash值的长度与ECC私钥的密钥长度有关(如果密钥是192位的,则hash值长度不能超过24(192/8 = 24)字节)
//                                                   (如果密钥是256位的,则hash值长度不能超过32(256/8 = 32)字节)
//        3.曲线参数为:EC_NIST_PRIME_192及EC_NIST_PRIME_256
//
//@param  hDongle          [in]      打开的加密锁句柄。
//@param  pPubKey          [in]      ECC公钥数据。
//@param  pHashData        [in]      Hash数据。
//@param  nHashDataLen     [in]      参数pHashData的大小。
//@param  pSign            [in]      签名数据。大小固定为64字节，为Dongle_EccSign函数返回的pOutData数据。
//
//@return DONGLE_SUCCESS             表示验签成功,否则表示验签失败。
func Dongle_EccVerify(hDongle DONGLE_HANDLE, pPubKey *ECCSM2_PUBLIC_KEY, pHashData *byte, nHashDataLen int, pSign *byte) error {
	r1, _, _ := procDongleEccverify.Call(uintptr(hDongle),
		uintptr(unsafe.Pointer(pPubKey)),
		uintptr(unsafe.Pointer(pHashData)), uintptr(nHashDataLen),
		uintptr(unsafe.Pointer(pSign)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_SM2GenPubPriKey
//@brief  产生SM2公钥和私钥。使用该函数之前需要先创建一个SM2私钥文件。需要开发商权限。成功后注意保存ECC公私钥数据。
//
//@param  hDongle        [in]      打开的加密锁句柄。
//@param  wPriFileID     [in]      SM2私钥文件ID。
//@param  pPubBakup      [out]     SM2公钥数据。
//@param  pPriBakup      [out]     SM2私钥数据。
//
//@return DONGLE_SUCCESS           产生ECC公私钥成功。
func Dongle_SM2GenPubPriKey(hDongle DONGLE_HANDLE, wPriFileID uint16, pPubBakup *ECCSM2_PUBLIC_KEY, pPriBakup *ECCSM2_PRIVATE_KEY) error {
	r1, _, _ := procDongleSm2genpubprikey.Call(uintptr(hDongle),
		uintptr(wPriFileID),
		uintptr(unsafe.Pointer(pPubBakup)),
		uintptr(unsafe.Pointer(pPriBakup)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_SM2Sign
//@brief  SM2私钥签名。函数的使用权限取决于锁内SM2私钥文件的权限，在SM2私钥文件创建时设定。
//
//@param  hDongle          [in]      打开的加密锁句柄。
//@param  wPriFileID       [in]      SM2私钥文件ID。
//@param  pHashData        [in]      Hash数据。
//@param  nHashDataLen     [in]      参数pHashData的大小。数据长度必须小于32个字节。
//@param  pOutData         [out]     签名数据。大小固定为64字节。
//
//@return DONGLE_SUCCESS             表示签名成功。
func Dongle_SM2Sign(hDongle DONGLE_HANDLE, wPriFileID uint16, pHashData *byte, nHashDataLen int, pOutData *byte) error {
	r1, _, _ := procDongleSm2sign.Call(uintptr(hDongle),
		uintptr(wPriFileID),
		uintptr(unsafe.Pointer(pHashData)),
		uintptr(nHashDataLen),
		uintptr(unsafe.Pointer(pOutData)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_SM2Verify
//@brief  SM2公钥验签。函数的使用权限取决于锁内SM2私钥文件的权限，在SM2私钥文件创建时设定。
//
//@param  hDongle          [in]      打开的加密锁句柄。
//@param  wPriFileID       [in]      SM2公钥数据。
//@param  pHashData        [in]      Hash数据。
//@param  nHashDataLen     [in]      参数pHashData的大小。
//@param  pSign            [in]      签名数据。大小固定为64字节，为Dongle_EccSign函数返回的pOutData数据。
//
//@return DONGLE_SUCCESS             表示验签成功,否则表示验签失败。
func Dongle_SM2Verify(hDongle DONGLE_HANDLE, pPubKey *ECCSM2_PUBLIC_KEY, pHashData *byte, nHashDataLen int, pSign *byte) error {
	r1, _, _ := procDongleSm2verify.Call(uintptr(hDongle),
		uintptr(unsafe.Pointer(pPubKey)),
		uintptr(unsafe.Pointer(pHashData)),
		uintptr(nHashDataLen),
		uintptr(unsafe.Pointer(pSign)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_TDES
//@brief  3DES加解密。解密运算匿名权限即可, 加密运算的权限取决于密钥文件的权限。
//
//@param  hDongle        [in]      打开的加密锁句柄。
//@param  wKeyFileID     [in]      密钥文件ID。
//@param  nFlag          [in]      运算类型。例如，FLAG_ENCODE表示加密运算；FLAG_DECODE表示解密运算。
//@param  pInData        [in]      输入数据缓冲区。
//@param  pOutData       [out]     输出数据缓冲区。大小至少要和输入数据缓冲区相同，输入和输出数据缓冲区可以为同一个。
//@param  nDataLen       [in]      参数pInData的数据大小。数据长度必须是16的整数倍,允许的最大值是1024。
//
//@return DONGLE_SUCCESS           3DES加密或解密运算成功。
func Dongle_TDES(hDongle DONGLE_HANDLE, wKeyFileID uint32, nFlag int, pInData *byte, pOutData *byte, nDataLen int) error {
	r1, _, _ := procDongleTdes.Call(uintptr(hDongle),
		uintptr(wKeyFileID),
		uintptr(nFlag),
		uintptr(unsafe.Pointer(pInData)),
		uintptr(unsafe.Pointer(pOutData)),
		uintptr(nDataLen))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_SM4
//@brief  SM4加解密。解密运算匿名权限即可, 加密运算的权限取决于密钥文件的权限。
//
//@param  hDongle        [in]      打开的加密锁句柄。
//@param  wKeyFileID     [in]      密钥文件ID。
//@param  nFlag          [in]      运算类型。例如，FLAG_ENCODE表示加密运算；FLAG_DECODE表示解密运算。
//@param  pInData        [in]      输入数据缓冲区。
//@param  pOutData       [out]     输出数据缓冲区。大小至少要和输入数据缓冲区相同，输入和输出数据缓冲区可以为同一个。
//@param  nDataLen       [in]      参数pInData的数据大小。数据长度必须是16的整数倍,允许的最大值是1024。
//
//@return DONGLE_SUCCESS           SM4加密或解密运算成功。
func Dongle_SM4(hDongle DONGLE_HANDLE, wKeyFileID uint32, nFlag int, pInData *byte, pOutData *byte, nDataLen int) error {
	r1, _, _ := procDongleSm4.Call(uintptr(hDongle),
		uintptr(wKeyFileID),
		uintptr(nFlag),
		uintptr(unsafe.Pointer(pInData)),
		uintptr(unsafe.Pointer(pOutData)),
		uintptr(nDataLen))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_HASH
//@brief  HASH运算。
//
//@param  hDongle        [in]      打开的加密锁句柄。
//@param  nFlag          [in]      Hash运算算法类型。
//                                 nFlag = FLAG_HASH_MD5，表示MD5运算，此时pHash的缓冲区大小为16字节。
//                                 nFlag = FLAG_HASH_SHA1，表示SHA1运算，此时pHash的缓冲区大小为20字节。
//                                 nFlag = FLAG_HASH_SM3，表示国密SM3运算，此时pHash的缓冲区大小为32字节。
//@param  pInData        [in]      输入数据缓冲区。
//@param  nDataLen       [in]      参数pInData的数据大小。SHA1、MD5为锁外运算，长度不限制；SM3为锁内运算，
//                                 最大不超过1024字节。
//@param  pHash          [out]     输出的Hash值。
//
//@return DONGLE_SUCCESS           HASH运算成功。
func Dongle_HASH(hDongle DONGLE_HANDLE, nFlag int, pInData *byte, nDataLen int, pHash *byte) error {
	r1, _, _ := procDongleHash.Call(uintptr(hDongle),
		uintptr(nFlag),
		uintptr(unsafe.Pointer(pInData)),
		uintptr(nDataLen),
		uintptr(unsafe.Pointer(pHash)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_Seed
//@brief  种子码算法。匿名权限可使用, 开发商可设置可运算次数。
//        1.种子码算法与PID有关，空锁(即PID=FFFFFFFF)不能进行种子码运算。
//        2.如果内部种子码可运算次数不为-1，当其递减到0后此函数将不能使用。
//
//@param  hDongle        [in]      打开的加密锁句柄。
//@param  pSeed          [in]      输入的种子码数据。
//@param  nSeedLen       [in]      种子码长度。取值范围为1~250字节。
//@param  pOutData       [out]     输出数据缓冲区。输出的大小固定为16字节。
//
//@return DONGLE_SUCCESS           种子码运算成功。
func Dongle_Seed(hDongle DONGLE_HANDLE, pSeed *byte, nSeedLen int, pOutData *byte) error {
	r1, _, _ := procDongleSeed.Call(uintptr(hDongle),
		uintptr(unsafe.Pointer(pSeed)),
		uintptr(nSeedLen),
		uintptr(unsafe.Pointer(pOutData)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_LimitSeedCount
//@brief 设置种子码算法可运算次数。需要开发商权限。
//
//@param  hDongle     [in]      打开的加密锁句柄。
//@param  nCount      [in]      可运算次数。如果此值设置为-1，表示不限制运算次数。
//
//@return DONGLE_SUCCESS        设置成功。
func Dongle_LimitSeedCount(hDongle DONGLE_HANDLE, nCount int) error {
	r1, _, _ := procDongleLimitseedcount.Call(uintptr(hDongle),
		uintptr(nCount))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_GenMotherKey
//@brief  制作一把母锁。子母锁的方式是是一种可选的初始化锁方式，安全又快速，推荐使用。需要开发商权限。
//        1.空锁(即PID=FFFFFFFF)不能写入母锁数据。
//	       2.出于安全考虑,MOTHER_DATA中的远程升级私钥不允许和母锁自身的远程升级私钥相同,否则会操作失败。
//
//@param  hDongle     [in]      打开的加密锁句柄。
//@param  pInData     [in]      输入数据。用于初始化母锁的结构为MOTHER_DATA的数据。
//
//@return DONGLE_SUCCESS        制作母锁成功。
func Dongle_GenMotherKey(hDongle DONGLE_HANDLE, pInData *MOTHER_DATA) error {
	r1, _, _ := procDongleGenmotherkey.Call(uintptr(hDongle),
		uintptr(unsafe.Pointer(pInData)))
	return DongleErrno(syscall.Errno(r1))
}
func Dongle_GenMotherKey2(hDongle DONGLE_HANDLE, pInData *MOTHER_DATA) error {
	// 由于字节数组导致传递的数据结构有误，这里转成byte发给dll
	k1 := int(unsafe.Sizeof(pInData.Count))
	k2 := int(unsafe.Sizeof(pInData.Son.SeedLen))
	k3 := int(unsafe.Sizeof(pInData.Son.UserTryCount))
	k4 := int(unsafe.Sizeof(pInData.Son.AdminTryCount))
	k5 := int(unsafe.Sizeof(pInData.Son.UpdatePriKey.Bits))
	k6 := int(unsafe.Sizeof(pInData.Son.UpdatePriKey.Modules))
	k7 := int(unsafe.Sizeof(pInData.Son.UserIDStart))
	var dataSize = k1 +
		k2 +
		len(pInData.Son.SeedForPID) +
		len(pInData.Son.UserPIN) +
		k3 +
		k4 +
		k5 +
		k6 +
		len(pInData.Son.UpdatePriKey.PublicExponent) +
		len(pInData.Son.UpdatePriKey.Exponent) +
		k7

	dataBytes := make([]byte, dataSize)
	offset := 0
	var nextIndex int
	for i := 0; i < k1; i++ {
		dataBytes[nextIndex] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(pInData)) + uintptr(i+offset)))
		nextIndex = nextIndex + 1
	}
	offset += k1

	for i := 0; i < k2; i++ {
		dataBytes[nextIndex] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(pInData)) + uintptr(i+offset)))
		nextIndex = nextIndex + 1
	}
	offset += k2

	for i := range pInData.Son.SeedForPID {
		dataBytes[nextIndex] = pInData.Son.SeedForPID[i]
		nextIndex = nextIndex + 1
	}
	for i := range pInData.Son.UserPIN {
		dataBytes[nextIndex] = pInData.Son.UserPIN[i]
		nextIndex = nextIndex + 1
	}

	for i := 0; i < k3; i++ {
		dataBytes[nextIndex] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(pInData)) + uintptr(i+offset)))
		nextIndex = nextIndex + 1
	}
	offset += k3

	for i := 0; i < k4; i++ {
		dataBytes[nextIndex] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(pInData)) + uintptr(i+offset)))
		nextIndex = nextIndex + 1
	}
	offset += k4

	for i := 0; i < k5; i++ {
		dataBytes[nextIndex] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(pInData)) + uintptr(i+offset)))
		nextIndex = nextIndex + 1
	}
	offset += k5

	for i := 0; i < k6; i++ {
		dataBytes[nextIndex] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(pInData)) + uintptr(i+offset)))
		nextIndex = nextIndex + 1
	}
	offset += k6

	publicExponentBy := make([]byte, 256)
	copy(publicExponentBy, pInData.Son.UpdatePriKey.PublicExponent[:])
	for i := range publicExponentBy {
		dataBytes[nextIndex] = publicExponentBy[i]
		nextIndex = nextIndex + 1
	}

	exponentBy := make([]byte, 256)
	copy(exponentBy, pInData.Son.UpdatePriKey.Exponent[:])
	for i := range exponentBy {
		dataBytes[nextIndex] = exponentBy[i]
		nextIndex = nextIndex + 1
	}
	for i := 0; i < k7; i++ {
		dataBytes[nextIndex] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(pInData)) + uintptr(i+offset)))
		nextIndex = nextIndex + 1
	}
	//offset += k7

	r1, _, _ := procDongleGenmotherkey.Call(uintptr(hDongle),
		uintptr(unsafe.Pointer(pInData)))
	return DongleErrno(syscall.Errno(r1))
}

//
//@brief  从空锁获取生产请求。此函数只对PID为FFFFFFFF的空锁有效。需要开发商权限。
//
//@param  hDongle      [in]       打开的加密锁句柄。
//@param  pRequest     [out]      输出数据。返回该数据的有效长度为16字节，因此需要至少16字节的空间。
//
//@return DONGLE_SUCCESS          获取生产请求数据成功。
//
func Dongle_RequestInit(hDongle DONGLE_HANDLE, pRequest *byte) error {
	r1, _, _ := procDongleRequestinit.Call(uintptr(hDongle),
		uintptr(unsafe.Pointer(pRequest)))
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_GetInitDataFromMother
//@brief  从母锁获取用于初始化子锁的数据，该函数只对母锁有效。匿名权限可使用。
//
//@param  hDongle       [in]          打开的加密锁句柄。
//@param  pRequest      [in]          请求数据。通过Dongle_RequestInit获取的请求数据。
//@param  pInitData     [out]         输出数据。函数执行成功返回用于初始化子锁的数据。
//@param  pDataLen      [int,out]     参数pInitData的有效长度。表示pInitData的缓冲区长度，函数执行成功
//                                    返回pInitData的有效长度。
//
//@return DONGLE_SUCCESS              从母锁获取生产数据成功。
func Dongle_GetInitDataFromMother(hDongle DONGLE_HANDLE, pRequest *byte, pInitData *byte, pDataLen *int) error {
	r1, _, _ := procDongleGetinitdatafrommother.Call(uintptr(hDongle),
		uintptr(unsafe.Pointer(pRequest)),
		uintptr(unsafe.Pointer(pInitData)),
		uintptr(unsafe.Pointer(pDataLen)),
	)
	return DongleErrno(syscall.Errno(r1))
}

// Dongle_InitSon
// @brief  生产子锁。用子母锁的方式制作子锁，匿名权限即可调用。
//
// @param  hDongle       [in]         打开的加密锁句柄。
// @param  pInitData     [in]         输入数据。函数Dongle_GetInitDataFromMother返回的用于初始化子锁的数据。
// @param  nDataLen      [in]         参数pInitData数据缓冲区的有效长度。
//
// @return DONGLE_SUCCESS             生产子锁成功。
func Dongle_InitSon(hDongle DONGLE_HANDLE, pInitData *byte, nDataLen int) error {
	r1, _, _ := procDongleInitson.Call(uintptr(hDongle),
		uintptr(unsafe.Pointer(pInitData)),
		uintptr(nDataLen),
	)
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_SetUpdatePriKey
//@brief  向锁内设置远程升级私钥。私钥长度为1024的RSA私钥。需要开发商权限。
//        出于安全考虑，如果锁是母锁的话，远程升级私钥不允许和母锁数据区中的子锁远程升级私钥相同，否则会操作失败。
//
//@param  hDongle       [in]          打开的加密锁句柄。
//@param  pPriKey       [in]          RSA私钥。
//
//@return DONGLE_SUCCESS              设置远程升级私钥成功。
func Dongle_SetUpdatePriKey(hDongle DONGLE_HANDLE, pPriKey *RSA_PRIVATE_KEY) error {
	r1, _, _ := procDongleSetupdateprikey.Call(uintptr(hDongle),
		uintptr(unsafe.Pointer(pPriKey)),
	)
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_MakeUpdatePacket
//@brief  制作远程升级数据包。匿名权限即可调用。
//
//@param  hDongle         [in]         打开的加密锁句柄。
//@param  pHID            [in]         硬件序列号。如果不需要绑定该参数可以为NULL。
//@param  nFunc           [in]         升级包类型。
//                                     nFunc = UPDATE_FUNC_CreateFile，表示创建文件。
//                                     nFunc = UPDATE_FUNC_WriteFile，写文件。只有锁内已有的文件才可升级写文件操作。
//                                     nFunc = UPDATE_FUNC_DeleteFile，删除文件。
//                                     nFunc = UPDATE_FUNC_FileLic，设置文件授权，不支持可执行文件授权升级。
//                                     nFunc = UPDATE_FUNC_SeedCount，设置种子码可运算次数。
//                                     nFunc = UPDATE_FUNC_DownloadExe，升级可执行文件。
//                                     nFunc = UPDATE_FUNC_UnlockUserPin，解锁用户PIN。出于安全考虑解锁用户PIN码必须绑定HID，
//                                     即pHID不能为空，只有这样才能升级成功。升级成功后用户PIN码恢复为"12345678"。
//                                     nFunc = UPDATE_FUNC_Deadline，时钟锁升级使用期限。
//
//@param  nFileType       [in]         文件类型。升级有关文件操作时的文件类型。其他升级类型该参数无效。
//@param  wFileID         [in]         文件ID。升级有关文件操作时的文件ID。其他升级类型该参数无效。
//                                     当wFileID=0xFFFF时，表示升级锁内数据区的数据。
//@param  nOffset         [in]         偏移量。升级有关文件操作时的文件偏移量。其他升级类型该参数无效。
//@param  pBuffer         [in]         输入数据。
//                                     当nFunc = UPDATE_FUNC_CreateFile时, pBuffer指向要文件的属性结构，例如KEY_FILE_ATTR。
//                                     当nFunc = UPDATE_FUNC_WriteFile时, pBuffer指向要写入的数据。
//                                     当nFunc = UPDATE_FUNC_FileLic时, pBuffer指向文件权限的数据结构，例如：DATA_LIC。
//                                     当nFunc = UPDATE_FUNC_SeedCount时, pBuffer指向long值，表示种子码可运算次数。
//                                     当nFunc = UPDATE_FUNC_DownloadExe时, pBuffer指向EXE_FILE_INFO结构，与Dongle_DownloadExeFile函数用法类似。
//                                     当nFunc = UPDATE_FUNC_Deadline时, pBuffer指向DWORD值，表示到期的时间。
//@param  nBufferLen      [in]         参数pBuffer的缓冲区大小。
//@param  pUPubKey        [in]         制作升级包的RSA公钥。与设置到锁内的远程升级私钥相对应。该值无论何种升级类型都必须填写。
//@param  pOutData        [out]        输出数据。制作的升级包数据。
//@param  pOutDataLen     [in,out]     参数pOutData输入大小，返回升级包的有效长度。
//
//@return DONGLE_SUCCESS               制作升级包成功。
func Dongle_MakeUpdatePacket(hDongle DONGLE_HANDLE, pHID *byte, nFunc int, nFileType int, wFileID uint16,
	nOffset int, pBuffer *byte, nBufferLen int, pUPubKey *RSA_PUBLIC_KEY, pOutData *byte, pOutDataLen *int) error {
	r1, _, _ := procDongleMakeupdatepacket.Call(uintptr(hDongle),
		uintptr(unsafe.Pointer(pHID)),
		uintptr(nFunc),
		uintptr(nFileType),
		uintptr(wFileID),
		uintptr(nOffset),
		uintptr(unsafe.Pointer(pBuffer)),
		uintptr(nBufferLen),
		uintptr(unsafe.Pointer(pUPubKey)),
		uintptr(unsafe.Pointer(pOutData)),
		uintptr(unsafe.Pointer(pOutDataLen)),
	)
	return DongleErrno(syscall.Errno(r1))

}

//Dongle_MakeUpdatePacketFromMother
//@brief  制作远程升级数据包。该函数采用母锁方式制作，与Dongle_MakeUpdatePacket相比少了远程升级公钥，其他相同。匿名权限即可调用。
//
//@param  hDongle         [in]         打开的加密锁句柄。
//@param  pHID            [in]         硬件序列号。如果不需要绑定该参数可以为NULL。
//@param  nFunc           [in]         升级包类型。
//                                     nFunc = UPDATE_FUNC_CreateFile，表示创建文件。
//                                     nFunc = UPDATE_FUNC_WriteFile，写文件。只有锁内已有的文件才可升级写文件操作。
//                                     nFunc = UPDATE_FUNC_DeleteFile，删除文件。
//                                     nFunc = UPDATE_FUNC_FileLic，设置文件授权，不支持可执行文件授权升级。
//                                     nFunc = UPDATE_FUNC_SeedCount，设置种子码可运算次数。
//                                     nFunc = UPDATE_FUNC_DownloadExe，升级可执行文件。
//                                     nFunc = UPDATE_FUNC_UnlockUserPin，解锁用户PIN。出于安全考虑解锁用户PIN码必须绑定HID，
//                                     即pHID不能为空，只有这样才能升级成功。升级成功后用户PIN码恢复为"12345678"。
//                                     nFunc = UPDATE_FUNC_Deadline，时钟锁升级使用期限。
//
//@param  nFileType       [in]         文件类型。升级有关文件操作时的文件类型。其他升级类型该参数无效。
//@param  wFileID         [in]         文件ID。升级有关文件操作时的文件ID。其他升级类型该参数无效。
//                                     当wFileID=0xFFFF时，表示升级锁内数据区的数据。
//@param  nOffset         [in]         偏移量。升级有关文件操作时的文件偏移量。其他升级类型该参数无效。
//@param  pBuffer         [in]         输入数据。
//                                     当nFunc = UPDATE_FUNC_CreateFile时, pBuffer指向要文件的属性结构，例如KEY_FILE_ATTR。
//                                     当nFunc = UPDATE_FUNC_WriteFile时, pBuffer指向要写入的数据。
//                                     当nFunc = UPDATE_FUNC_FileLic时, pBuffer指向文件权限的数据结构，例如：DATA_LIC。
//                                     当nFunc = UPDATE_FUNC_SeedCount时, pBuffer指向long值，表示种子码可运算次数。
//                                     当nFunc = UPDATE_FUNC_DownloadExe时, pBuffer指向EXE_FILE_INFO结构，与Dongle_DownloadExeFile函数用法类似。
//                                     当nFunc = UPDATE_FUNC_Deadline时, pBuffer指向DWORD值，表示到期的时间。
//@param  nBufferLen      [in]         参数pBuffer的缓冲区大小。
//@param  pOutData        [out]        输出数据。制作的升级包数据。
//@param  pOutDataLen     [in,out]     参数pOutData输入大小，返回升级包的有效长度。
//
//@return DONGLE_SUCCESS               制作升级包成功。
func Dongle_MakeUpdatePacketFromMother(hDongle DONGLE_HANDLE, pHID string, nFunc UPDATE_FUNC, nFileType int,
	wFileID uint16, nOffset int, pBuffer *byte, nBufferLen int,
	pOutData *byte, pOutDataLen *int) error {
	r1, _, _ := procDongleMakeupdatepacketfrommother.Call(uintptr(hDongle),
		Lpstr(pHID),
		uintptr(nFunc),
		uintptr(nFileType),
		uintptr(wFileID),
		uintptr(nOffset),
		uintptr(unsafe.Pointer(pBuffer)),
		uintptr(nBufferLen),
		uintptr(unsafe.Pointer(pOutData)),
		uintptr(unsafe.Pointer(pOutDataLen)),
	)
	return DongleErrno(syscall.Errno(r1))
}

//Dongle_Update
//@brief  远程升级子锁中的数据。匿名权限即可。升级数据pUpdateData对一把锁只能使用一次。
//        1.本函数内部是按1024字节的分块机制发送,如遇返回值不是DONGLE_SUCCESS会立即中断发送并返回。
//        2.如果需要进行流程控制,可由调用方来分块(每块1024字节)发送，并保证数据块的顺序不被打乱，根据返回的错误码来控制流程。
//
//@param  hDongle         [in]     打开的加密锁句柄。
//@param  pUpdateData     [in]     输入数据。升级数据，由Dongle_MakeUpdatePacket或者Dongle_MakeUpdatePacketFromMother产生
//@param  nDataLen        [in]     参数pUpdateData的大小。必须为1024的整数倍。
//
//@return DONGLE_SUCCESS           升级成功。
func Dongle_Update(hDongle DONGLE_HANDLE, pUpdateData *byte, nDataLen int) error {
	r1, _, _ := procDongleUpdate.Call(uintptr(hDongle),
		uintptr(unsafe.Pointer(pUpdateData)),
		uintptr(nDataLen),
	)
	return DongleErrno(syscall.Errno(r1))
}
