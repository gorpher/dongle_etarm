package main

import (
	"fmt"
	"github.com/gorpher/dongle_etarm"
)

type DongleInfo struct {
	Pid      string `json:"pid" pg:"pid"`
	UserId   string `json:"user_id" pg:"user_id"`
	DevType  string `json:"dev_type" pg:"dev_type"` //设备类型(PROTOCOL_HID或者PROTOCOL_CCID)
	Type     string `json:"type" pg:"type"`         //1表示标准版, 2为标准时钟锁,3为标准U盘锁
	Ver      string `json:"ver" pg:"ver"`
	Agent    string `json:"agent" pg:"agent"`
	Hid      string `json:"hid" pg:"hid"`
	Birthday string `json:"birthday" pg:"birthday"`
	IsMother int    `json:"is_mother" pg:"is_mother"` // 1表示是母锁,  0表示不是母锁
}

func main() {
	list, err := dongle_etarm.DongleEnum()
	if err != nil {
		panic(err)
	}
	for i, _ := range list {
		info := DongleInfo{
			Pid:    fmt.Sprintf("%08X", list[i].PID),
			UserId: fmt.Sprintf("%08X", list[i].UserID),
			Ver:    fmt.Sprintf("%d.%02d", list[i].Ver>>8, list[i].Ver&0x0F),
			Agent:  fmt.Sprintf("%08X", list[i].Agent),
			Hid:    fmt.Sprintf("%X", list[i].HID),
			Birthday: fmt.Sprintf("20%02X-%02X-%02X %02X:%02X:%02X",
				list[i].BirthDay[0],
				list[i].BirthDay[1],
				list[i].BirthDay[2],
				list[i].BirthDay[3],
				list[i].BirthDay[4],
				list[i].BirthDay[5],
			),
			IsMother: int(list[i].IsMother),
		}
		//time.Parse("206-1-2 15:04:05", fmt.Sprintf("%s-%s-%s %s:%s:%s", year, month, date, hour, minute, second))
		if list[i].DevType == 0 {
			info.DevType = "HID"
		} else {
			info.DevType = "CCID"
		}
		if list[i].Type == 0 {
			info.Type = "标准时钟锁"
		}
		if list[i].Type == 2 {
			info.Type = "标准U盘锁"
		}
		if list[i].Type != 0 && list[i].Type != 2 {
			info.Type = "标准锁"
		}

		fmt.Printf("%#v\n", info)
	}
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
	err = dongle_etarm.Dongle_Close(hDongle)
	if err != nil {
		panic(err)
	}
}
