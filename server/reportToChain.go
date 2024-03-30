package server

func CoinAgeAttack() string {
	if coinAgeCompare() == true {
		return "CoinAgeAttack!"
	}
	return ""
}

func BadForDao() string {
	return "BadForDao!"
}
