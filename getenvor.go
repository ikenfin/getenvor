package getenvor

import "os"

// GetenvOr shortcat function that returns requested
// environment var value or default value
func GetenvOr(name string, defaults string) string {
	if variable := os.Getenv(name); len(variable) > 0 {
		return variable
	}
	return defaults
}
