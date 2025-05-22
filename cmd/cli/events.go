package main

import (
	"fmt"

	"github.com/RupenderSinghRathore/github-Activity/internal/models"
)

func eventsEva(activities []models.Activity) {
	for _, event := range activities {
		repo := event.Repo.Name
		date := event.Created_at
		fmt.Println()
		if event.Type == "PushEvent" {
			fmt.Printf("> Pushed %v commits to %v\n", len(event.Payload.Commits), repo)
			for _, commit := range event.Payload.Commits {
				msg := commit.Message
				fmt.Printf("  >> \"%v\"  at:%v\n", msg, date)
			}
		} else if event.Type == "CreateEvent" {
			if event.Payload.Ref_type == "repository" {
				fmt.Printf("< Created Repository:%v at:%v\n", repo, date)
			}
		}
	}
}
