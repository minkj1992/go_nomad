package main

import (
	"fmt"

	"github.com/minkj1992/go_nomad/array"
	"github.com/minkj1992/go_nomad/hello"
)

// 리터럴(축약형)은 func 밖에서는 사용 불가능하다.
// name := "minwook"
func main() {
	// 패키지의 func은 대문자로 시작
	name := "minwook"
	hello.SayHello(name)
	hello.SayBye(name)

	total := array.Add(1, 2, 3, 4, 5)
	fmt.Println(total)
}
