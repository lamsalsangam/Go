package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	u1 := user{
		First: "I need",
		Last:  "Someone to love.",
		Age:   25,
	}

	u2 := user{
		First: "Can Someone",
		Last:  "please find me.",
		Age:   25,
	}

	users := []user{u1, u2}
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(b)
	os.Stdout.Write(b)
}
