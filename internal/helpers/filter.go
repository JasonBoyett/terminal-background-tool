package helpers

import "regexp"

func FilterByRegexp(data []string, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return data, err
	}

	result := []string{}

	// If there is no pattern to match against return and empty array
	if len(pattern) == 0 {
		return result, nil
	}

	for _, item := range data {
		if re.MatchString(item) {
			result = append(result, item)
		}
	}
	return result, nil
}
