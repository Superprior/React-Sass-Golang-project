package parsers

func LinesArrayToPostgresArray(initial []string) string {
	result := "["

	for i, line := range initial {
		result += string("'" + line + "'")
		if i != len(initial)-1 {
			result += ","
		}
	}

	result += "]"
	return result
}
