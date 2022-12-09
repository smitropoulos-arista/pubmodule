package pubmodule

import json "github.com/goccy/go-json"

func Hello() string {
	return "Hello, You!"
}

func Goodbye() string {
	return "Goodbye for now!"
}

func GoodEvening() string {
	return "Good evening sir!"
}

func GetJson() []byte {
	res, _ := json.Marshal(struct {
		X int    `json:"x"`
		Y string `json:"y"`
	}{X: 1, Y: "hello"})

	return res
}
