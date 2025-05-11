package main

import (
	"fmt"
	"log"
	"logreader/pkg/constants"
	"logreader/pkg/service"
	"logreader/pkg/util"
)

func main() {
	data, err := util.ReadFile("programming-task-example-data.log")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	// Create new logService with imput data
	logService, err := service.NewLogService(data)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	// The number of unique IP addresses
	uniqueIps := logService.GetUniqueFields(constants.IP_FIELD_NAME)
	fmt.Println(fmt.Sprintf("Unique IPs are : %v", uniqueIps))

	// The top 3 most visited URLs
	top3Urls := logService.FindTopNEntries(3, constants.URL_FIELD_NAME)
	fmt.Println(fmt.Sprintf("Top 3 URLs are : %v", top3Urls))

	// The top 3 most active IP addresses
	top3Ips := logService.FindTopNEntries(3, constants.IP_FIELD_NAME)
	fmt.Println(fmt.Sprintf("Top 3 active IPs are : %v", top3Ips))

}
