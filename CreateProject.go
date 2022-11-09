package main

import (
	"fmt"
	"os/exec"
)

func CreateCSConsole() {
	executeShell("dotnet new console \"testapp\"")
}

func CreateCSASPNETCore() {

}

func CreateCSASPNETCoreBlazorServer() {

}

func CreateCSASPNETCoreBlazorWASM() {

}

func CreateGoConsole() {

}

func CreateCConsole() {

}

func executeShell(command string) {
	app := "cmd"

	arg0 := "/C"
	arg1 := command

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
