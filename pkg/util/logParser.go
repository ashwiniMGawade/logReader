package util

import (
	"fmt"
	"logreader/pkg/constants"
	"logreader/pkg/model"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/**
* Method for parsing log entries
* Responsible for transforming unprocessed string log data into structured Golang structs by accurately parsing each individual field.
   * @param logData - Raw data to parse
   * @returns slice of LogInfo object if parsing succeeds, nil if the data doesn't match expected format
			  error if parsing fails at any point of time for wrong data
*/

func ParseLog(logData string) ([]model.LogInfo, error) {
	var logEntries []model.LogInfo

	// Compile the Regular expression pattern
	re := regexp.MustCompile(constants.LOG_REGEX_PATTERN)
	logLines := strings.Split(logData, "\n")

	for _, line := range logLines {
		/* Matching line will store the matching parts data in following format, with key and value
		1 50.112.00.11
		2 -
		3 admin
		4 11/Jul/2018:17:33:01 +0200
		5 GET
		6 /asset.css
		7 HTTP/1.1
		8 200
		9 3574
		10 Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6
		*/
		parts := re.FindStringSubmatch(line)

		// check if logLine follows the log regex, if not continue with next line
		if len(parts) == 0 {
			continue
		}

		var logEntry model.LogInfo
		logEntry.IpAddress = parts[1]

		// Assign user with default value if not specified
		logEntry.User = constants.DEFAULT_USER
		if parts[3] != "-" {
			logEntry.User = parts[3]
		}

		// Assign Date by parsing against the date format
		parsedDate, err := time.Parse(constants.LOG_DATE_FORMAT, parts[4])
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("Could not parse the date") // Return error if date parsing fails
		}
		logEntry.TimeStamp = parsedDate

		// Assign Request Type ex. GET, POST, PUT
		logEntry.RequestType = parts[5]

		// Assign Request URL
		logEntry.RequestUrl = parts[6]

		// Assign Response code, convert string to integer
		// we can ignore the error as Regex will always match integer values for response code
		logEntry.ResponseCode, _ = strconv.Atoi(parts[8])

		// Assign Response Size, convert string to int64
		// we can ignore the error as Regex will always match integer values for response code
		logEntry.ResponseSize, _ = strconv.ParseInt(parts[9], 10, 64)

		// Assign Platfrom and OS related info
		logEntry.Platform = parts[10]

		logEntries = append(logEntries, logEntry)
	}

	return logEntries, nil
}
