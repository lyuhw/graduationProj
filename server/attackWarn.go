package server

func Warn(userSetAttackThreshold int) string {
	if DataDetect() >= userSetAttackThreshold {
		return "ATTACK!"
	}
	return ""
}
