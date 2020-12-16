package test

import "time"

var CST = time.FixedZone("CST", 0)
/*now := time.Unix(time.Now().Unix(), 0).In(CST)
weekday := now.Weekday()
fmt.Println(int(weekday))
	now := time.Unix(time.Now().Unix(), 0)
	fmt.Println(now)
	weekday := int64(now.Weekday())
	timeNow := time.Now().Unix()
	nowSecond := timeNow % (24 * 60 * 60)
	lastEndTime := timeNow - nowSecond - ((weekday - 1) * int64(24*60*60))
	lastStartTime := lastEndTime - 7*int64(24*60*60)
	start := time.Unix(lastStartTime, 0).In(CST)
	end := time.Unix(lastEndTime-1, 0).In(CST)
	fmt.Println(start,end)
	endTime := time.Now().Unix()
	startTime := lastEndTime
	start = time.Unix(startTime, 0).In(CST)
	end = time.Unix(endTime, 0).In(CST)
	fmt.Println(start,end)fmt.Println(len("00000064bf96a01bca4af965dd30e159f755ed3ca71c63c8200c0c6c8bcf1919"))*/