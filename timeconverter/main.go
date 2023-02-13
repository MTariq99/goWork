package main

import (
	"fmt"
	"time"
)

// 979352580
// time.Now -- NOW(unix)

func main() {
	now := time.Now()
	fmt.Println("current time ", now)
	unixtime := now.Unix()
	fmt.Println("unixTime is ", unixtime)
	timeT := time.Unix(unixtime, 0)
	fmt.Println("converting Unix into time ", timeT)

}
