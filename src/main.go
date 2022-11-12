package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Read in the file.
	data, err := ioutil.ReadFile("terraform.tfstate")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	// Convert []byte to string and print to screen
	text := string(data)
	//fmt.Println(text)

	// split on resources
	resources := strings.Split(text, "resources")

	// split on each resource
	for _, resource := range resources {
		//fmt.Println(resource)
		resourceLines := strings.Split(resource, "

")
		//fmt.Println(resourceLines)
		//fmt.Println("")

		// find resource type
		resourceType := ""
		for _, line := range resourceLines {
			if strings.Contains(line, "type") {
				resourceType = strings.Split(line, "\"")[3]
				break
			}
		}
		//fmt.Println(resourceType)

		// find resource name
		resourceName := ""
		for _, line := range resourceLines {
			if strings.Contains(line, "name") {
				resourceName = strings.Split(line, "\"")[3]
				break
			}
		}
		//fmt.Println(resourceName)

		// find resource id
		resourceID := ""
		for _, line := range resourceLines {
			if strings.Contains(line, "id") {
				resourceID = strings.Split(line, "\"")[3]
				break
			}
		}
		//fmt.Println(resourceID)

		// find resource location
		resourceLocation := ""
		for _, line := range resourceLines {
			if strings.Contains(line, "location") {
				resourceLocation = strings.Split(line, "\"")[3]
				break
			}
		}
		//fmt.Println(resourceLocation)

		// find resource tags
		resourceTags := ""
		for _, line := range resourceLines {
			if strings.Contains(line, "tags") {
				resourceTags = strings.Split(line, "\"")[3]
				break
			}
		}
		//fmt.Println(resourceTags)

		// find resource properties
		resourceProperties := ""
		for _, line := range resourceLines {
			if strings.Contains(line, "properties") {
				resourceProperties = strings.Split(line, "\"")[3]
				break
			}
		}
		//fmt.Println(resourceProperties)
	}
}
