package main

import "fmt"
import "os"
import "github.com/gillisandrew/attestations-demo/version"

func main() {
	switch os.Args[1] {
		case "hello":
			fmt.Println("Hello, World!")
		case "version":
			fmt.Println(version.BuildVersion())
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(1)
	}
}