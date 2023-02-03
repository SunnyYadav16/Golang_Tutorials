package main

func main() {
	go greeter("Hello")
	greeter("World")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		println(s)
	}
}
