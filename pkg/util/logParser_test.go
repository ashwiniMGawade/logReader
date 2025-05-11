package util

import (
	"fmt"
	"logreader/pkg/model"
	"testing"
	"time"
)

func TestParseLog(t *testing.T) {
	testCases := []struct {
		name     string
		logData  string
		expected []model.LogInfo
		err      error
	}{
		{
			name:    "Valid log entry for guest user",
			logData: `177.71.128.21 - - [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
			expected: []model.LogInfo{
				{
					IpAddress:    "177.71.128.21",
					User:         "guest",
					TimeStamp:    time.Date(2018, time.July, 10, 22, 21, 28, 0, time.FixedZone("", 7200)),
					RequestType:  "GET",
					RequestUrl:   "/intranet-analytics/",
					ResponseCode: 200,
					ResponseSize: 3574,
					Platform:     "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7",
				},
			},
			err: nil,
		},
		{
			name:    "Valid log entry for admin user",
			logData: `177.71.128.21 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
			expected: []model.LogInfo{
				{
					IpAddress:    "177.71.128.21",
					User:         "admin",
					TimeStamp:    time.Date(2018, time.July, 10, 22, 21, 28, 0, time.FixedZone("", 7200)),
					RequestType:  "GET",
					RequestUrl:   "/intranet-analytics/",
					ResponseCode: 200,
					ResponseSize: 3574,
					Platform:     "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7",
				},
			},
			err: nil,
		},
		{
			name:     "InValid log entry with no pattern match",
			logData:  `test 177.71.128.21 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
			expected: []model.LogInfo{},
			err:      nil,
		},
		{
			name:     "InValid log entry with pattern match but wrong date format",
			logData:  `177.71.128.21 - admin [10/July/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
			expected: []model.LogInfo{},
			err:      fmt.Errorf("Could not parse the date"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result, err := ParseLog(test.logData)

			if err != nil && test.err == nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if len(result) != len(test.expected) {
				t.Errorf("Expected %d log entries, got %d", len(test.expected), len(result))
			}

			for i := range result {
				if result[i].IpAddress != test.expected[i].IpAddress ||
					result[i].User != test.expected[i].User ||
					result[i].TimeStamp != test.expected[i].TimeStamp ||
					result[i].RequestType != test.expected[i].RequestType ||
					result[i].RequestUrl != test.expected[i].RequestUrl ||
					result[i].ResponseCode != test.expected[i].ResponseCode ||
					result[i].ResponseSize != test.expected[i].ResponseSize ||
					result[i].Platform != test.expected[i].Platform {
					t.Errorf("Mismatch in log entry at index %d", i)
				}
			}
		})
	}
}
