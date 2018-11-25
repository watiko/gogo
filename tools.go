// +build tools

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	_ "github.com/codegangsta/gin/lib" // これは微妙な書き方かもしれない
)

func main() {
	install_bin("gin", "github.com/codegangsta/gin")
	teardown()
}

func install_bin(bin_name string, package_name string) {
	bin_location := fmt.Sprintf("bin/%s", bin_name)

	cmd := exec.Command("go", "build", "-o", bin_location, package_name)
	out, err := cmd.CombinedOutput()

	fmt.Printf("cmd: %s\n", strings.Join(cmd.Args, " "))
	fmt.Printf("out: %s\n", out)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %s\n", err)
	}
}

func teardown() {
	// go.mod に indirect な依存関係が追加されてしまうので綺麗にしている
	// ref: https://github.com/golang/go/issues/26474#issuecomment-407519043
	_, err := exec.Command("go", "mod", "tidy").CombinedOutput()

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %s\n", err)
	}
}
