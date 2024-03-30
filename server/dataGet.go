package server

import (
	"math/rand"
	"time"
)

var processData = 00

func coinAgeCompare() bool {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成1-10的随机数
	randomNumber := rand.Intn(10) + 1

	if randomNumber%2 == 0 {
		processData = 10
		return true
	} else {
		processData = 00
		return false
	}
}

func ifGoodForDao(gb bool) {
	if !gb {
		processData++
	}
}

func GetProcessData(gOrb bool) int {
	coinAgeCompare()
	ifGoodForDao(gOrb)
	return processData
}
