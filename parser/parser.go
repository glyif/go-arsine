package parser


func Parser(s string) [] string {
	var args [] string

	for i, j, c := 0, 0, 0; i <= len(s) - 1; i++ {
		if s[i] == ' ' || s[i] == '\n' {
			args = append(args, s[c:i])
			c = i
			j = j+1
		}
	}

	return args
}