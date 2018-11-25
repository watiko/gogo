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
