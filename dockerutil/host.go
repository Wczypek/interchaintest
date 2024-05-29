package dockerutil

import "os"

func GetHostAddress() string {
	if value, ok := os.LookupEnv("ICTEST_HOST"); ok {
		return value
	}
	return "docker"
}
