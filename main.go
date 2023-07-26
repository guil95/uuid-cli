package main

import (
	"github.com/atotto/clipboard"
	"github.com/google/uuid"
)

func main() {
	_ = clipboard.WriteAll(uuid.NewString())
}
