// Package environment provides a set of functions for safely accessing the environment variables
package environment

import (
	"os"
	"strconv"
	"strings"
)

func GetVar(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func GetVarAsInt(name string, defaultVal int) int {
	valueStr := GetVar(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func GetVarAsBool(name string, defaultVal bool) bool {
	valStr := GetVar(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

func GetVarAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := GetVar(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
