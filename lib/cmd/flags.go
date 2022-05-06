package cmd

var Flags = map[string]bool{"a": false}

func ParseFlags(args []string) []string {
	j := 0
	for _, v := range args {
		if string(v[0]) == "-" && len(v) > 1 {
			if _, ok := Flags[string(v[1])]; ok {
				Flags[string(v[1])] = true
			}
		} else {
			args[j] = v
			j++
		}
	}
	return args[:j]
}
