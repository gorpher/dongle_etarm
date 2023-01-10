package dongle_etarm

import (
	"errors"
	"syscall"
)

const (
	DONGLE_SUCCESS             syscall.Errno = 0x00000000 // 操作成功
	DONGLE_NOT_FOUND           syscall.Errno = 0xF0000001 // 未找到指定的设备
	DONGLE_INVALID_HANDLE      syscall.Errno = 0xF0000002 // 无效的句柄
	DONGLE_INVALID_PARAMETER   syscall.Errno = 0xF0000003 // 参数错误
	DONGLE_COMM_ERROR          syscall.Errno = 0xF0000004 // 通讯错误
	DONGLE_INSUFFICIENT_BUFFER syscall.Errno = 0xF0000005 // 缓冲区空间不足
	DONGLE_NOT_INITIALIZED     syscall.Errno = 0xF0000006 // 产品尚未初始化 (即没设置PID)
	DONGLE_ALREADY_INITIALIZED syscall.Errno = 0xF0000007 // 产品已经初始化 (即已设置PID)
	DONGLE_ADMINPIN_NOT_CHECK  syscall.Errno = 0xF0000008 // 开发商密码没有验证
	DONGLE_USERPIN_NOT_CHECK   syscall.Errno = 0xF0000009 // 用户密码没有验证
	DONGLE_INCORRECT_PIN       syscall.Errno = 0xF000FF00 // 密码不正确 (后2位指示剩余次数)
	DONGLE_INCORRECT_ADMIN_PIN syscall.Errno = 0xF000FFFF // 密码不正确 (后2位指示剩余次数) //4026597375
	DONGLE_PIN_BLOCKED         syscall.Errno = 0xF000000A // PIN码已锁死
	DONGLE_ACCESS_DENIED       syscall.Errno = 0xF000000B // 访问被拒绝
	DONGLE_FILE_EXIST          syscall.Errno = 0xF000000E // 文件已存在
	DONGLE_FILE_NOT_FOUND      syscall.Errno = 0xF000000F // 未找到指定的文件
	DONGLE_READ_ERROR          syscall.Errno = 0xF0000010 // 读取数据错误
	DONGLE_WRITE_ERROR         syscall.Errno = 0xF0000011 // 写入数据错误
	DONGLE_FILE_CREATE_ERROR   syscall.Errno = 0xF0000012 // 创建文件错误
	DONGLE_FILE_READ_ERROR     syscall.Errno = 0xF0000013 // 读取文件错误
	DONGLE_FILE_WRITE_ERROR    syscall.Errno = 0xF0000014 // 写入文件错误
	DONGLE_FILE_DEL_ERROR      syscall.Errno = 0xF0000015 // 删除文件错误
	DONGLE_FAILED              syscall.Errno = 0xF0000016 // 操作失败
	DONGLE_CLOCK_EXPIRE        syscall.Errno = 0xF0000017 // 加密锁时钟到期
	DONGLE_ERROR_UNKNOWN       syscall.Errno = 0xFFFFFFFF // 未知的错误

)

func DongleErrno(errno syscall.Errno) error {
	switch errno {
	case DONGLE_SUCCESS:
		return nil
	case DONGLE_NOT_FOUND:
		return errors.New("未找到指定的设备")
	case DONGLE_INVALID_HANDLE:
		return errors.New("无效的句柄")
	case DONGLE_INVALID_PARAMETER:
		return errors.New("参数错误")
	case DONGLE_COMM_ERROR:
		return errors.New("通讯错误")
	case DONGLE_INSUFFICIENT_BUFFER:
		return errors.New("缓冲区空间不足")
	case DONGLE_NOT_INITIALIZED:
		return errors.New("产品尚未初始化")
	case DONGLE_ALREADY_INITIALIZED:
		return errors.New("产品已经初始化")
	case DONGLE_ADMINPIN_NOT_CHECK:
		return errors.New("开发商密码没有验证")
	case DONGLE_USERPIN_NOT_CHECK:
		return errors.New("用户密码没有验证")
	case DONGLE_INCORRECT_PIN:
		return errors.New("密码不正确")
	case DONGLE_INCORRECT_ADMIN_PIN:
		return errors.New("开发商密码不正确")
	case DONGLE_PIN_BLOCKED:
		return errors.New("PIN码已锁死")
	case DONGLE_ACCESS_DENIED:
		return errors.New("访问被拒绝")
	case DONGLE_FILE_EXIST:
		return errors.New("文件已存在")
	case DONGLE_FILE_NOT_FOUND:
		return errors.New("未找到指定的文件")
	case DONGLE_READ_ERROR:
		return errors.New("读取数据错误")
	case DONGLE_WRITE_ERROR:
		return errors.New("写入数据错误")
	case DONGLE_FILE_CREATE_ERROR:
		return errors.New("创建文件错误")
	case DONGLE_FILE_READ_ERROR:
		return errors.New("读取文件错误")
	case DONGLE_FILE_WRITE_ERROR:
		return errors.New("写入文件错误")
	case DONGLE_FILE_DEL_ERROR:
		return errors.New("删除文件错误")
	case DONGLE_FAILED:
		return errors.New("操作失败")
	case DONGLE_CLOCK_EXPIRE:
		return errors.New("加密锁时钟到期")
	case DONGLE_ERROR_UNKNOWN:
		return errors.New("未知的错误")
	}
	return errors.New("调用动态链接库失败")

}

//加密锁句柄定义
type DONGLE_HANDLE syscall.Handle

//默认的PIN码重试次数为无限制
//根据种子码初始化锁时会同时初始化PID和ADMINPIN (PID不可更改, ADMINPIN可更改)

const CONST_PID = 0xFFFFFFFF              //出厂时默认的PID
const CONST_USERPIN = "12345678"          //出厂时默认的USERPIN
const CONST_ADMINPIN = "FFFFFFFFFFFFFFFF" //出厂时默认的ADMINPIN

//通讯协议类型定义
const PROTOCOL_HID = 0  //hid协议
const PROTOCOL_CCID = 1 //ccid协议

//文件类型定义

const FILE_DATA = 1          //普通数据文件
const FILE_PRIKEY_RSA = 2    //RSA私钥文件
const FILE_PRIKEY_ECCSM2 = 3 //ECC或者SM2私钥文件(SM2私钥文件和ECC私钥文件结构相同，属相同文件类型)
const FILE_KEY = 4           //SM4和3DES密钥文件
const FILE_EXE = 5           //可执行文件

//LED灯状态定义
const LED_OFF = 0   //灯灭
const LED_ON = 1    //灯亮
const LED_BLINK = 2 //灯闪

//PIN码类型
const FLAG_USERPIN = 0  //用户PIN
const FLAG_ADMINPIN = 1 //开发商PIN

//加解密标志
const FLAG_ENCODE = 0 //加密
const FLAG_DECODE = 1 //解密

//HASH算法类型
const FLAG_HASH_MD5 = 0  //MD5     运算结果16字节
const FLAG_HASH_SHA1 = 1 //SHA1    运算结果20字节
const FLAG_HASH_SM3 = 2  //SM3     运算结果32字节

type UPDATE_FUNC int

//远程升级的功能号
const (
	UPDATE_FUNC_CreateFile    UPDATE_FUNC = iota + 1 //创建文件
	UPDATE_FUNC_WriteFile                            //写文件
	UPDATE_FUNC_DeleteFile                           //删除文件
	UPDATE_FUNC_FileLic                              //设置文件授权
	UPDATE_FUNC_SeedCount                            //设置种子码可运算次数
	UPDATE_FUNC_DownloadExe                          //升级可执行文件
	UPDATE_FUNC_UnlockUserPin                        //解锁用户PIN
	UPDATE_FUNC_Deadline                             //时钟锁升级使用期限
)
