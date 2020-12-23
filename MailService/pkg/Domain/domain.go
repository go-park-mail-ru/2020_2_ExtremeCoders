package Domain

func GetMailDomain(email string) string {
	flag := false
	var domail string
	for _, char := range email {
		if char == '@' {
			flag = true
			continue
		}
		if flag {
			domail += string(char)
		}
	}
	return domail
}