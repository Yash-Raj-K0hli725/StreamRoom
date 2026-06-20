package util

import "regexp"

func SanitizeFilename(name string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9\.]`)
	processed := reg.ReplaceAllString(name, "_")
	return processed
}
