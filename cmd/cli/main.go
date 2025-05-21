package main

import (
	"flag"
	"fmt"
)

func main() {
	user := flag.String("user", "RupenderSinghRathore", "Specifies user")
	flag.Parse()

	activities := apiCall(*user)
	fmt.Println(activities)
}
