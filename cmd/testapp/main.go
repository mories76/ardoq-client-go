package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	ardoq "github.com/mories76/ardoq-client-go/pkg"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

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

	// Test search component
	if false {
		workspace := "d85f9d74393dd8cb053e7e09"
		name := "myTerraformComponent2"
		cmps, err := a.Components().Search(context.TODO(), &ardoq.ComponentSearchQuery{Workspace: workspace, Name: name})
		if err != nil {
			fmt.Printf("error during component search: %s", err)
		}
		// fmt.Printf("result of search component \n%v\n", cmps)

		for _, cmp := range *cmps {
			// fields := make(map[string]string)
			fields := cmp.GetConvertedFields()
			fmt.Printf("\nfields:\n %s", fields)
		}
	}

	// Test the Model type
	if false {
		rootworkspace := "12ef1c9b57a1e67dcf9c7fe1"

		// get workspace by id
		workspace, err := a.Workspaces().Get(context.TODO(), rootworkspace)
		if err != nil {
			fmt.Printf("error during get workspace: %s", err)
		}
		// set componentModel to the componentModel from the found workspace
		componentModel := workspace.ComponentModel

		model, err := a.Models().Read(context.TODO(), componentModel)
		if err != nil {
			fmt.Printf("error during get model: %s", err)
		}
		fmt.Printf("result of get model \n%s\n", prettyPrint(model))

		cmpTypes := model.GetComponentTypeID()
		fmt.Printf("componentTypes \n%s\n", prettyPrint(cmpTypes))

		// for _, cmp := range *cmps {
		// 	// fields := make(map[string]string)
		// 	fields := cmp.GetConvertedFields()
		// 	fmt.Printf("\nfields:\n %s", fields)
		// }
	}

	// Test the Fields type
	if true {
		// rootworkspace := "12ef1c9b57a1e67dcf9c7fe1"

		// get fields
		fields, err := a.Fields().GetAll(context.TODO())
		if err != nil {
			fmt.Printf("error during get fields: %s", err)
		}
		// fmt.Printf("result of get all fields \n%s\n", prettyPrint(fields))

		for _, field := range *fields {
			if len(field.Fields) > 0 {
				fmt.Printf("field found with non empty Fields \n%s\n", prettyPrint(field))
			}
		}

		// fieldID := "7391a71024c8f0480fb998a5"
		fieldID := "9e38c3b6b82e6cfefdb915db"
		field, err := a.Fields().Read(context.TODO(), fieldID)
		if err != nil {
			fmt.Printf("error during get field: %s", err)
		}
		fmt.Printf("result of get field \n%s\n", prettyPrint(field))

		// cmpTypes := model.GetComponentTypeID()
		// fmt.Printf("componentTypes \n%s\n", prettyPrint(cmpTypes))

		// for _, cmp := range *cmps {
		// 	// fields := make(map[string]string)
		// 	fields := cmp.GetConvertedFields()
		// 	fmt.Printf("\nfields:\n %s", fields)
		// }
	}
}
