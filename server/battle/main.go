// +build !debug

package main

import "github.com/shiwano/submarine/server/battle/src"

func main() {
	s := server.New(":5000")
	s.ListenAndServe()
}
