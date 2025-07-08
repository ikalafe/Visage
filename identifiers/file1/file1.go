package main

import (
	"fmt"

	"github.com/ikalafe/Visage/identifiers/file2"
)

var ExportedVariable = "Hello, World"

func main() {
	fmt.Println(ExportedVariable)
	fmt.Println(file2.AnotherExportVariable)
}