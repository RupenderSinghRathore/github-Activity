package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/RupenderSinghRathore/github-Activities/internals/models"
)

func apiCall(user string) *[]models.Activity {
	url := fmt.Sprintf("https://api.github.com/users/%v/events?per_page=10", user)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var buffer bytes.Buffer
	data := make([]byte, 1024)
	for {
		n, err := res.Body.Read(data)
		if n > 0 {
			_, writeErr := buffer.Write(data[:n])
			if writeErr != nil {
				log.Fatal(writeErr)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	activities := &[]models.Activity{}
	if err = json.Unmarshal(buffer.Bytes(), activities); err != nil {
		log.Fatal(err)
	}
	return activities
}
