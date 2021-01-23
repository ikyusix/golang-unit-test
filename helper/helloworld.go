package helper

import (
	"strconv"
)

func HelloWorld(name string) string {
	return "Hello " + name
}

func UserGradeA(username string, cnt int) []string {
	var users []string
	var user string
	for i := 0; i <= cnt; i++ {
		j := strconv.Itoa(i)
		user = username + j
		users = append(users, user)
	}
	return users
}
