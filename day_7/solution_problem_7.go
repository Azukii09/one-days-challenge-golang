package day_7

import "slices"

func FilterUnique(developers []Developer) []string {
	var output []string
	for _, developer := range developers {
		if !slices.Contains(output, developer.Name) {
			output = append(output, developer.Name)
		}
	}
	return output
}
