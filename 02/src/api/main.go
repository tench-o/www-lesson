package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	Name    string
	Profile string
	Id      string
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if isAllowIpAddr(r.RemoteAddr) == false {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "forbidden:%s", r.RemoteAddr)
		return
	}

	user := User{
		Name:    "他人のユーザー",
		Profile: "天津飯大好き！",
		Id:      "2",
	}

	str, _ := json.Marshal(user)

	fmt.Fprintf(w, string(str))
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:    "テストタロウ",
		Profile: "ラーメン大好き！",
		Id:      "1",
	}

	if isCustomHostname(r.Host) {
		user = User{
			Name:    "独ドメ三太夫",
			Profile: "まむちゃん",
			Id:      "2",
		}
	}

	str, _ := json.Marshal(user)

	fmt.Fprintf(w, string(str))
}

func isCustomHostname(hostHeader string) bool {
	customHostname := getCustomHostName()
	if customHostname == "" {
		return false
	}

	if hostHeader == customHostname {
		return true
	}

	return false
}

func isAllowIpAddr(ipaddr string) bool {
	allowIpaddr := os.Getenv("ALLOW_IP_ADDR")

	if ipaddr == allowIpaddr {
		return true
	}

	return false
}

func getCustomHostName() string {
	return os.Getenv("CUSTOM_HOSTNAME")
}

func main() {
	http.HandleFunc("/current_user", GetCurrentUser)
	http.HandleFunc("/users", GetUser)
	http.ListenAndServe(":8081", nil)
}
