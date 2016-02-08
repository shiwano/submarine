package main

func main() {
	server := NewServer()
	server.Run(":5000")
}
