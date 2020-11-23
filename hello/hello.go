package hello

import "fmt"

// import 할 func은 대문자로 시작 (public)
func SayHello(name string) {
	fmt.Println("Hello " + name)
}

func SayBye(name string) {
	fmt.Println("Bye " + name)
}

// private
func sayBye(name string) {
	fmt.Println("Bye " + name)
}
