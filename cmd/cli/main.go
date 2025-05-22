package main

import (
	"flag"
)

func main() {
	user := flag.String("user", "RupenderSinghRathore", "Specifies user")
	flag.Parse()

	activities := apiCall(*user)
	eventsEva(*activities)
}
