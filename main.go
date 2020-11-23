package main

import "github.com/minkj1992/go_nomad/hello"

func main() {
	// 패키지의 func은 대문자로 시작
	name := "minwook"
	hello.SayHello(name)
	hello.SayBye(name)
}
