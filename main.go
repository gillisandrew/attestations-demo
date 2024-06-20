package main

import "fmt"
import "os"
import "github.com/gillisandrew/attestations-demo/version"

func printHelp() {
	fmt.Println("Usage: app <command>")
	fmt.Println("Commands:")
	fmt.Println("  hello    - Print 'Hello, World!'")
	fmt.Println("  version  - Print the version of the application")
}

func main() {

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "hello":
		fmt.Println("Hello, World!")
	case "version":
		fmt.Println(version.BuildVersion())
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
		printHelp()
		os.Exit(1)
	}
}
