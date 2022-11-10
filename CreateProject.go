package main

import (
	"os/exec"
)

func createCsConsole(sName string) {
	executeShell("dotnet new console -o " + sName)
}

func createCsAspNetCore(sName string) {
	executeShell("dotnet new web -o " + sName)
}

func createCsAspNetCoreBlazorServer(sName string) {
	executeShell("dotnet new blazorserver -o " + sName)
}

func createCsAspNetCoreBlazorWASM(sName string) {
	executeShell("dotnet new blazorwasm -o " + sName)
}

func createCsAspNetCoreMvc(sName string) {
	executeShell("dotnet new mvc -o " + sName)
}

func createGoConsole(sName string) {
	sStandardMain := "package main\n\n	 import \"fmt\"\n	  func main() {\n		 fmt.Println(\"Hello, Alcina.\")\n	 }"
	executeShell("mkdir goproject && cd goproject && type " + sStandardMain + " > main.go && go mod init " + sName)
}

func createCConsole(sName string) {

}

func executeShell(sCommand string) {
	app := "cmd"

	arg0 := "/C"
	arg1 := sCommand

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		log(true, err.Error(), ERROR)
		return
	}

	// Print the output
	log(false, string(stdout)+"Done.", INFO)
}
