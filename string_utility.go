package golib

func LowerCamelCase(str string) string {
	if isLower := 'a' <= str[0] && str[0] <= 'z'; isLower {
		return str
	}

	newstr := make([]byte, 0, len(str))

	for i := 0; i < len(str); i++ {
		c := str[i]

		if i == 0 {
			c += 'a' - 'A'

			newstr = append(newstr, c)
		} else {
			newstr = append(newstr, c)
		}
	}

	return BytesToString(newstr)
}

func UpperCamelCase(str string) string {
	if isUpper := 'A' <= str[0] && str[0] <= 'Z'; isUpper {
		return str
	}

	newstr := make([]byte, 0, len(str))

	for i := 0; i < len(str); i++ {
		c := str[i]

		if i == 0 {
			c -= 'a' - 'A'

			newstr = append(newstr, c)
		} else {
			newstr = append(newstr, c)
		}
	}

	return BytesToString(newstr)
}

func SnakeCase(str string) string {
	newstr := make([]byte, 0, len(str)+1)

	for i := 0; i < len(str); i++ {
		c := str[i]

		if isUpper := 'A' <= c && c <= 'Z'; isUpper {
			if i > 0 {
				newstr = append(newstr, '_')
			}

			c += 'a' - 'A'
		}

		newstr = append(newstr, c)
	}

	return BytesToString(newstr)
}

func SpinalCase(str string) string {
	newstr := make([]byte, 0, len(str)+1)

	for i := 0; i < len(str); i++ {
		c := str[i]

		if isUpper := 'A' <= c && c <= 'Z'; isUpper {
			if i > 0 {
				newstr = append(newstr, '-')
			}

			c += 'a' - 'A'
		}

		newstr = append(newstr, c)
	}

	return BytesToString(newstr)
}
