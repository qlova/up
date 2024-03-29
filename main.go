package main

import "fmt"
import "os"
import "os/exec"

func exists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func getgo() {
	fmt.Println("Please install Go (https://golang.org/)")
	os.Exit(1)
}

func gofmt() {
	err := exec.Command("gofmt", "-s", "-w", ".").Run()
	if err != nil {
		fmt.Println("The project didn't format, please fix any syntax errors before running up again.")
		os.Exit(1)
	}
}

func gobuild() {
	err := exec.Command("go", "build").Run()
	if err != nil {
		fmt.Println("The project didn't build, please fix any compiler errors before running up again.")
		os.Exit(1)
	}
}

func goget(pkg string) {
	err := exec.Command("go", "get", pkg).Run()
	if err != nil {
		fmt.Println("There was an error trying to get ", pkg)
		fmt.Println("Please get it manually!")
		os.Exit(1)
	}
}

func gitcola() {
	err := exec.Command("git-cola").Run()
	if err != nil {
		fmt.Println("git-cola encountered an error")
		fmt.Println("If not, maybe you should report it... (https://github.com/git-cola/git-cola/issues)")
	}
}

func lazygit() {
	if !exists("lazygit") {
		fmt.Println("Please add $GOPATH/bin to your $PATH!")
		os.Exit(1)
	}

	err := exec.Command("lazygit").Run()
	if err != nil {
		fmt.Println("lazygit encountered an error")
		fmt.Println("Maybe you should report it... (https://github.com/jesseduffield/lazygit/issues)")
	}
}

func main() {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		fmt.Println("This is not a git repo, please create one with 'git init'")
		return
	}

	if !exists("go") {
		getgo()
	}

	gofmt()
	gobuild()

	//Open the best available git client.

	if exists("git-cola") {
		gitcola()
		return
	}

	if !exists("lazygit") {
		goget("github.com/jesseduffield/lazygit")
	}
	lazygit()
}
