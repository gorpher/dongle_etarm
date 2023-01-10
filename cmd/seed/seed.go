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

	var seed = make([]byte, 32)
	for i := range seed {
		seed[i] = 0x11
	}
	seedData := make([]byte, 8)
	err = dongle_etarm.Dongle_Seed(hDongle, &seed[0], 32, &seedData[0])
	if err != nil {
		panic(err)
	}
	for i := range seedData {
		fmt.Printf("%X-", seedData[i])
	}
	fmt.Println()
	fmt.Printf("%X\n", seedData)
	fmt.Println(string(seedData))
	//os.WriteFile("seed.bin", seed[:], os.ModePerm)
	//var nRemainCount int
	//err = dongle_etarm.Dongle_VerifyPIN(hDongle, dongle_etarm.FLAG_ADMINPIN, string(fmt.Sprintf("%X", seedData)), &nRemainCount)
	//if err != nil {
	//	panic(err)
	//}

}
