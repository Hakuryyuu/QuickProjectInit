package main

import (
	"fmt"
)

func main() {
	fmt.Println("Which type of project you would like to create?")
	fmt.Println("[1] Go")
	fmt.Println("[2] C#")
	fmt.Println("[3] C")
	var inp string
	fmt.Scanln(&inp)
	projectType(inp)
}

/*func validateInput() string {

	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString(byte(reader.Size()))
		if err != nil {
			fmt.Println(text)
			return text
		} else {
			// Add logger
			fmt.Println(err)
		}
		return err.Error()
	}
}*/

// Determite Programming Language
func projectType(ProgrammingLang string) {

	fmt.Print("\033[H\033[2J") // -> Clean Console
	var inp string

	switch ProgrammingLang {
	case "1":
		fmt.Println("Go Options")
		fmt.Println("[ 1 ]  Console Application")
		fmt.Scanln(&inp)
		projectGO(inp)
		break
	case "2":
		fmt.Println("C# Options")
		fmt.Println("[ 1 ]  Console Application")
		fmt.Println("[ 2 ]  ASP.NET Core")
		fmt.Println("[ 3 ]  ASP.NET Core Blazor Server")
		fmt.Println("[ 4 ]  ASP.NET Core Blazor WebAssembly")
		fmt.Scanln(&inp)
		projectCS(inp)
		break
	case "3":
		fmt.Println("C Options")
		fmt.Println("[ 1 ]  Console Application")
		fmt.Scanln(&inp)
		projectC(inp)
		break
	default:
		//Log
		break

	}
}

// C# Projects
func projectCS(ProjectType string) {

	switch ProjectType {
	case "1":
		CreateCSConsole()
		break
	case "2":
		break
	case "3":
		break
	case "4":
		break
	default:
		break

	}
}

// Go Projects
func projectGO(ProjectType string) {

	switch ProjectType {
	case "1":
		CreateGoConsole()
		break
	default:
		break

	}
}

// C Projects
func projectC(ProjectType string) {
	switch ProjectType {
	case "1":
		CreateCConsole()
		break
	default:
		break

	}
}
