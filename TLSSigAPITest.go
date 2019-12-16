package main

import (
	"./tencentyun"
	"fmt"
)

const (
	sdkappid = 1400000000
	key = "5bd2850fff3ecb11d7c805251c51ee463a25727bddc2385f3fa8bfee1bb93b5e"
)

func main()  {
	sig, err := tencentyun.GenSig(sdkappid, key, "xiaojun", 86400*180)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(sig)
	}
	var userbuf []byte = tencentyun.GetUserBuf("xiaojun",sdkappid,10000,86400*180,255,0);
	sig, err = tencentyun.GenSigWithUserBuf(sdkappid, key, "xiaojun", 86400*180, userbuf)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(sig)
	}
}
