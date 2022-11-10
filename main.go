package main

import (
	"fmt"
)

func main() {
	fmt.Println("Which type of project you would like to create?")
	fmt.Println("[ 1 ] Go")
	fmt.Println("[ 2 ] C#")
	fmt.Println("[ 3 ] C")
	var sUserInput string
	fmt.Scanln(&sUserInput)
	projectType(sUserInput)
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
	var sInput string
	var sName string

	switch ProgrammingLang {
	case "1":
		fmt.Println("Go Options")
		fmt.Println("[ 1 ]  Console Application")
		fmt.Scanln(&sInput)
		fmt.Println("Please enter the Name of the Project (ex. net.hakuryuu/hello)")
		fmt.Scanln(&sName)
		projectGO(sInput, sName)
		break
	case "2":
		fmt.Println("C# Options")
		fmt.Println("[ 1 ]  Console Application")
		fmt.Println("[ 2 ]  ASP.NET Core")
		fmt.Println("[ 3 ]  ASP.NET Core Blazor Server")
		fmt.Println("[ 4 ]  ASP.NET Core Blazor WebAssembly")
		fmt.Println("[ 5 ]  ASP.NET Core MVC")
		fmt.Scanln(&sInput)
		fmt.Println("Please enter the Name of the Project ")
		fmt.Scanln(&sName)
		projectCS(sInput, sName)
		break
	case "3":
		fmt.Println("C Options")
		fmt.Println("[ 1 ]  Console Application")
		fmt.Scanln(&sInput)
		fmt.Println("Please enter the Name of the Project ")
		fmt.Scanln(&sName)
		projectC(sInput, sName)
		break
	default:
		log(false, "Invalid selection", INFO)
		break

	}
}

// C# Projects
func projectCS(sProjectType string, sProjectName string) {

	switch sProjectType {
	case "1":
		createCsConsole(sProjectName)
		break
	case "2":
		createCsAspNetCore(sProjectName)
		break
	case "3":
		createCsAspNetCoreBlazorServer(sProjectName)
		break
	case "4":
		createCsAspNetCoreBlazorWASM(sProjectName)
		break
	case "5":
		createCsAspNetCoreMvc(sProjectName)
		break
	default:
		log(false, "Invalid selection.", INFO)
		break

	}
}

// Go Projects
func projectGO(sProjectType string, sProjectName string) {

	switch sProjectType {
	case "1":
		createGoConsole(sProjectName)
		break
	default:
		log(false, "Invalid selection.", INFO)
		break

	}
}

// C Projects
func projectC(sProjectType string, sProjectName string) {
	switch sProjectType {
	case "1":
		createCConsole(sProjectName)
		break
	default:
		log(false, "Invalid selection.", INFO)
		break

	}
}
