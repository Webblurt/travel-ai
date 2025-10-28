package utils

import "net/url"

func GetDefaultQueryValue(queryParams url.Values, key, defaultValue string) string {
	value := queryParams.Get(key)

	if value == "" {
		return defaultValue
	}

	return value
}
