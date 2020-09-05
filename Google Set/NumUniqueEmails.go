package main

func processEmail(input string) string {
	inputByteString := []byte(input)
	foundPlus := false
	i := 0
	for true {
		v := inputByteString[i]
		if v == '@' {
			break
		} else if foundPlus && v != '@' {
			inputByteString = append(inputByteString[:i], inputByteString[i+1:]...)
			continue
		} else if v == '.' {
			inputByteString = append(inputByteString[:i], inputByteString[i+1:]...)
		} else if v == '+' {
			inputByteString = append(inputByteString[:i], inputByteString[i+1:]...)
			foundPlus = true
		} else {
			i++
		}
	}
	return string(inputByteString)
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)
	for _, v := range emails {
		str := processEmail(v)
		_, ok := m[str]
		if !ok {
			m[str] = true
		}
	}

	return len(m)
}