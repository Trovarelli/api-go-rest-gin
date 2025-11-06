package helpers

import "regexp"

func NormalizeString(s string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		return "", err
	}

	return reg.ReplaceAllLiteralString(s, ""), nil
}
