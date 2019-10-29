// Use env variables to set program parameters, often used in cloud infrastructure, db connection configuration, and config in general to avoid having the developer change application code, and to set defaults
package main

import (
	"log"
	"os"
)

func main() {
	key := "DB_CONN"
	// Set the environment variable
	os.Setenv(key, "postgres://as:as@example.com/pg sslmode=verify-full")

	val := GetEnvDefault(key, "postgres://as:as@localhost/pg?sslmode=verify-full")
	log.Println("The value is: ", val)
	os.Unsetenv(key)
	val = GetEnvDefault(key, "postgres://as:as@127.0.0.1/pg? sslmode=verify-full")
	log.Println("The value is: ", val)
}

// Get the value of previously set env variable, else return defaultVal
func GetEnvDefault(key, defaultVal string) string {
	val, found := os.LookupEnv(key)
	if !found {
		return defaultVal
	}
	return val
}
