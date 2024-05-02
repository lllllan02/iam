package main

import "github.com/lllllan02/iam/cmd/wire"

func main() {
	wire.Init().ListenAndServe()
}
