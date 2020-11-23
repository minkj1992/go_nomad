package main

import (
	"fmt"

	"github.com/minkj1992/go_nomad/accounts"
)

func main() {
	account := accounts.NewAccount("leoo")
	fmt.Println(account)
}
