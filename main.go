package main

import (
	"fmt"
)

type GitlabsParser struct {
	UID         string     `json:"before_sha"`
	BuildID     int        `json:"build_id"`
	DateCreated string     `json:"build_started_at"`
	ProjectName string     `json:"project_name"`
	Repository  string `json:"repository"`
	Commit      string     `json:"commit"`
}

func main(){
	var G GitlabsParser
	G.UID = "Yo"
	fmt.Print(G)
}