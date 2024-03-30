package server

var temp = true

func DataDetect() int {
	defer reserve()
	return GetProcessData(temp)
}

func reserve() {
	temp = false
}
