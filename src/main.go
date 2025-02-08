package main

import (
	"fmt"

	"github.com/patrickishaf/lema-be/src/db"
)

func main() {
	fmt.Println("Hello, World!")
	db.InitializeDb()
	fmt.Println("Connected to DB")
}
