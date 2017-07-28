package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"syscall"

	"github.com/glyif/go-arscine/builtin"
	"github.com/glyif/go-arscine/parser"
)

type BUILT_IN_FUNC func([] string)

type BUILT_IN struct {
	s string
	f BUILT_IN_FUNC
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	//env := os.Environ()

	var builtin_list [] BUILT_IN
	builtin_list = append(builtin_list, BUILT_IN{s:"exit", f:builtin.Exit})
	builtin_list = append(builtin_list, BUILT_IN{s:"hi", f:builtin.Hi})

	for true {
		fmt.Print("$ ")
		text, _ := reader.ReadString('\n')

		args := parser.Parser(text)

		for _, elem := range builtin_list {
			if strings.Compare(args[0], elem.s) == 0 {
				elem.f(args)
				break
			}
		}

		//execErr := syscall.Exec(args[0], args, env)
		//
		//if execErr != nil {
		//	panic(execErr)
		//}
	}
}
