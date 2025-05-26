package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/RupenderSinghRathore/github-Activity/internal/models"
)

func apiCall(user string) *[]models.Activity {
	url := fmt.Sprintf("https://api.github.com/users/%v/events?per_page=10", user)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	activities := &[]models.Activity{}
	if err = json.Unmarshal(data, activities); err != nil {
		log.Fatal(err)
	}
	return activities
}
