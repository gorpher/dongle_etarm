package dongle_etarm

//RSA_PUBLIC_KEY RSA公钥格式(兼容1024,2048)
type RSA_PUBLIC_KEY struct {
	Bits     uint32    // length in bits of modulus
	Modules  uint32    // modulus
	Exponent [256]byte // public exponent
}

//RSA_PRIVATE_KEY RSA私钥格式(兼容1024,2048)
type RSA_PRIVATE_KEY struct {
	Bits           uint32     // length in bits of modulus
	Modules        uint32     // modulus
	PublicExponent [256]uint8 // public exponent
	Exponent       [256]uint8 // private exponent
}

//ECCSM2_PUBLIC_KEY 外部ECCSM2公钥格式 ECC(支持bits为192或256)和SM2的(bits为固定值0x8100)公钥格式
type ECCSM2_PUBLIC_KEY struct {
	Bits        uint32   // length in bits of modulus
	XCoordinate [8]uint8 // 曲线上点的X坐标
	YCoordinate [8]uint8 // 曲线上点的Y坐标
}

//ECCSM2_PRIVATE_KEY 外部ECCSM2私钥格式 ECC(支持bits为192或256)和SM2的(bits为固定值0x8100)私钥格式
type ECCSM2_PRIVATE_KEY struct {
	Bits       uint32   // length in bits of modulus
	PrivateKey [8]uint8 // 私钥
}

// DONGLE_INFO 加密锁信息
type DONGLE_INFO struct {
	Ver      uint16   //COS版本,比如:0x0201,表示2.01版
	Type     uint16   //产品类型: 0xFF表示标准版, 0x00为标准时钟锁,0x02为标准U盘锁
	BirthDay [8]uint8 //出厂日期
	Agent    uint32   //代理商编号,比如:默认的0xFFFFFFFF
	PID      uint32   //产品ID
	UserID   uint32   //用户ID
	HID      [8]uint8 //8字节的硬件ID
	IsMother uint32   //母锁标志: 0x01表示是母锁, 0x00表示不是母锁
	DevType  uint32   //设备类型(PROTOCOL_HID或者PROTOCOL_CCID)
}

/**
 *   锁内文件说明
 *   1.RSA私钥文件允许创建的最大数量为8个
 *   2.ECCSM2私钥文件允许创建的最大数量为16个
 *   3.3DES/SM4密钥文件允许创建的最大数量为32个
 *   4.可执行文件允许创建的最大数量为64个,总大小不能超过64K
 *   5.数据文件创建个数受锁内空间大小和文件系统其他因素的影响，最大个数不超过54个。
 *   6.文件ID取值范围为0x0000~0xFFFF之间，其中ID：0x0000、0xFFFF、0x3F00被锁内系统占用，用户不能使用。
 */

/*************************文件授权结构***********************************/

//DATA_LIC 数据文件授权结构
type DATA_LIC struct {
	ReadPriv  uint16 //读权限: 0为最小匿名权限，1为最小用户权限，2为最小开发商权限
	WritePriv uint16 //写权限: 0为最小匿名权限，1为最小用户权限，2为最小开发商权限
}

//PRIKEY_LIC 私钥文件授权结构
type PRIKEY_LIC struct {
	Count      int32 //可调次数: 0xFFFFFFFF表示不限制, 递减到0表示已不可调用
	Priv       uint8 //调用权限: 0为最小匿名权限，1为最小用户权限，2为最小开发商权限
	IsDecOnRAM uint8 //是否是在内存中递减: 1为在内存中递减，0为在FLASH中递减
	IsReset    uint8 //用户态调用后是否自动回到匿名态: TRUE为调后回到匿名态 (开发商态不受此限制)
	Reserve    uint8 //保留,用于4字节对齐
}

//KEY_LIC 对称加密算法(SM4/TDES)密钥文件授权结构
type KEY_LIC struct {
	PrivEnc uint32 //加密时的调用权限: 0为最小匿名权限，1为最小用户权限，2为最小开发商权限
}

//EXE_LIC 可执行文件授权结构
type EXE_LIC struct {
	PrivExe uint32 //运行的权限: 0为最小匿名权限，1为最小用户权限，2为最小开发商权限
}

/****************************文件属性结构********************************/

//DATA_FILE_ATTR 数据文件属性数据结构
type DATA_FILE_ATTR struct {
	Size uint32   //数据文件长度，该值最大为4096
	Lic  DATA_LIC //授权
}

//PRIKEY_FILE_ATTR ECCSM2/RSA私钥文件属性数据结构
type PRIKEY_FILE_ATTR struct {
	Type uint16     //数据类型:ECCSM2私钥 或 RSA私钥
	Size uint16     //数据长度:RSA该值为1024或2048, ECC该值为192或256, SM2该值为0x8100
	Lic  PRIKEY_LIC //授权
}

//KEY_FILE_ATTR 对称加密算法(SM4/TDES)密钥文件属性数据结构
type KEY_FILE_ATTR struct {
	Size uint32  //密钥数据长度=16
	Lic  KEY_LIC //授权
}

//EXE_FILE_ATTR 可执行文件属性数据结构
type EXE_FILE_ATTR struct {
	Lic EXE_LIC //授权
	Len uint16  //文件长度
}

/*************************文件列表结构***********************************/

//PRIKEY_FILE_LIST 获取私钥文件列表时返回的数据结构
type PRIKEY_FILE_LIST struct {
	FILEID  uint16           //文件ID
	Reserve uint16           //保留,用于4字节对齐
	Attr    PRIKEY_FILE_ATTR //文件属性
}

//KEY_FILE_LIST 获取SM4及TDES密钥文件列表时返回的数据结构
type KEY_FILE_LIST struct {
	FILEID  uint16        //文件ID
	Reserve uint16        //保留,用于4字节对齐
	Attr    KEY_FILE_ATTR //文件属性
}

//DATA_FILE_LIST 获取数据文件列表时返回的数据结构
type DATA_FILE_LIST struct {
	FILEID  uint16         //文件ID
	Reserve uint16         //保留,用于4字节对齐
	Attr    DATA_FILE_ATTR //文件属性
}

//EXE_FILE_LIST 获取可执行文件列表时返回的数据结构
type EXE_FILE_LIST struct {
	FILEID  uint16
	Attr    EXE_FILE_ATTR
	Reserve uint16
}

//EXE_FILE_INFO 下载和列可执行文件时填充的数据结构
type EXE_FILE_INFO struct {
	DwSize  uint16 //可执行文件大小
	WFileID uint16 //可执行文件ID
	Priv    uint8  //调用权限: 0为最小匿名权限，1为最小用户权限，2为最小开发商权限
	PData   *byte  //可执行文件数据
}

//SON_DATA 需要发给空锁的初始化数据
type SON_DATA struct {
	SeedLen       int             //种子码长度
	SeedForPID    [256]byte       //产生产品ID和开发商密码的种子码 (最长250个字节)
	UserPIN       [18]byte        //用户密码(16个字符的0终止字符串)
	UserTryCount  uint8           //用户密码允许的最大错误重试次数
	AdminTryCount uint8           //开发商密码允许的最大错误重试次数
	UpdatePriKey  RSA_PRIVATE_KEY //远程升级私钥
	UserIDStart   uint32          //起始用户ID
}

//MOTHER_DATA 母锁数据
type MOTHER_DATA struct {
	Son   SON_DATA //子锁初始化数据
	Count int32    //可产生子锁初始化数据的次数 (-1表示不限制次数, 递减到0时会受限)
}
