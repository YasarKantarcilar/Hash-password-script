package hashPassword

var Alph string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func HashPassword(password string) string {
	passwordBytes := []byte(password)
	for i := 0; i < len(passwordBytes); i++ {
		for j := 0; j < len(Alph); j++ {
			if passwordBytes[i] == Alph[j] {
				passwordBytes[i] = Alph[j+1]
				break
			}
		}
	}

	return string(passwordBytes)
}
