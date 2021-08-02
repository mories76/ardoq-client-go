package main

import (
	"context"
	"fmt"
	"os"

	ardoq "github.com/mories76/ardoq-client-go/pkg"
)

func main() {
	baseUri := os.Getenv("ARDOQ_BASEURI")
	apiKey := os.Getenv("ARDOQ_APIKEY")
	org := os.Getenv("ARDOQ_ORG")
	// do nothing
	a, err := ardoq.NewRestClient(baseUri, apiKey, org, "v0.0.0")
	if err != nil {
		//return nil, errors.Wrap(err, "cannot create new restclient")
		fmt.Printf("cannot create new restclient %s", err)
		os.Exit(1)
	}

	if true {
		workspace := "1551274144b0a61af921549a"
		name := "myComponent"
		cmp, err := a.Components().Search(context.TODO(), &ardoq.ComponentSearchQuery{Workspace: workspace, Name: name})
		if err != nil {
			fmt.Printf("error during component search: %s", err)
		}
		fmt.Printf("result of search component %v", cmp)
	}

}
