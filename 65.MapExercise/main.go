package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pws"] = "88888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pws"] = "88888888"
		users[name]["nickname"] = "昵称·" + name
	}
}

func main() {
	users := make(map[string]map[string]string, 10)
	users["李六"] = make(map[string]string, 2)
	users["李六"]["pws"] = "99999999"
	users["李六"]["nickname"] = "老六·李六"
	modifyUser(users, "张三")
	modifyUser(users, "李四")
	modifyUser(users, "王五")
	modifyUser(users, "李六")
	fmt.Println(users)
}
