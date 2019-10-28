package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Custom type needs to implement flag.Value interface to be able to use it in flag.Var function
type ArrayValue []string

func (av *ArrayValue) String() string {
	return fmt.Sprintf("%v", *av)
}

func (av *ArrayValue) Set(s string) error {
	*av = strings.Split(s, ",")
	return nil
}

func main() {
	// Extracting flag values with methods returning pointers
	// flag.[Type]() returns pointer of type Type
	retry := flag.Int("retry", -1, "Defines max retry count")

	// Read the flag using the XXXVar function
	// In this case, the variable must be defined prior to the flag
	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "Logger prefix")

	var arr ArrayValue
	flag.Var(&arr, "array", "Input array to iterate through")

	// Execute the flag. Parser function to read the flags to defined variables.
	// Without this call the flag variables remain empty
	// Must be called after all flags are defined,
	// and before the flags are accessed
	flag.Parse()

	// Sample logic not related to flags
	logger := log.New(os.Stdout, logPrefix, log.Ldate)

	retryCount := 0
	for retryCount < *retry {
		logger.Println("Retrying connection")
		logger.Printf("Sending array %v\n", arr)
		retryCount++
	}
}

/* NOTES
- go flag pkg doesn't support flag combining like -ll
- --flag and -flag are equivalent
- no difference between long and short options
- flag pkg defines 2 types of funcs
	1. simple name of flag, like Int. Value will be an int, and this func returns pointer to integer variable where the value of the parsed flag is
	2. XXXVar functions Provide same functionality, but you provide the pointer to the variable. After parsing, the input flag value will be
	   stored there
- Custom types are allowed. They must implement the Value interface from the flag package
- Flag subsets are also possible. See FlagSet
`myFlagset.Parse(os.Args[2:])`
*/
