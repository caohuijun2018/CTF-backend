package client

import "CTF-backend/internal/mysql"

func IsExisted(userId string) (bool, error) {
	var user mysql.User
	if mysql.IsMissing(mysql.DB.Where(mysql.User{UserID: userId}).First(&user)) {
		return false, nil
	}
	if user.UserID != "" {
		return true, nil
	} else {
		return false, nil
	}
}

func GetPage(length, size int) int {
	if length%size == 0 {
		return length / size
	}
	return length/size + 1
}
