package builtin

import (
	"strings"
	"os"
)

func Exit (s []string){
	if strings.Compare("exit", s[0]) == 0 {
		os.Exit(0)
	}
}
