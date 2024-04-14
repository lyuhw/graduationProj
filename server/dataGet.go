package server

import (
	"math/rand"
)

var processData int

func ifFiftyAttack() {

	// 生成1-10的随机数
	randomNumber := rand.Intn(10) + 1

	if randomNumber%2 == 0 {
		processData = 10
	} else {
		processData = 0
	}
}

func ifGoodForDao() {

	rNumber := rand.Intn(10) + 1

	if rNumber%2 == 0 {
		processData++
	}
}

func GetProcessData() int {
	ifFiftyAttack()
	ifGoodForDao()
	return processData
}
