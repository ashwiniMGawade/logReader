package service

import (
	"fmt"
	"logreader/pkg/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLogService(t *testing.T) {
	tests := []struct {
		name        string
		description string
		logData     string
		expected    []model.LogInfo
		err         error
	}{
		{
			name:        "TestNewLogServiceSuccess",
			description: "Test successful LogService creation",
			logData:     `177.71.128.21 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
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
			name:        "TestNewLogServiceSuccessWithEmptyLogEntriesDueToNoMatchOfData",
			description: "Test successful creation of LogService with empty ",
			logData:     `177.71.128.21 admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
			expected:    []model.LogInfo{},
			err:         nil,
		},
		{
			name:        "TestNewLogServiceFailsWhenPatternMatchesButDateFormatIsWrong",
			description: "InValid log entry with pattern match but wrong date format",
			logData:     `177.71.128.21 - admin [10/July/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
			expected:    []model.LogInfo{},
			err:         fmt.Errorf("Could not parse the date"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service, err := NewLogService(test.logData)

			if test.err != nil {
				assert.Equal(t, err.Error(), err.Error())
			} else {
				response := service.GetLogEntries()
				assert.Equal(t, test.err, err)
				assert.Equal(t, test.expected, response)
			}
		})
	}
}

func TestGetLogEntries(t *testing.T) {
	tests := []struct {
		name          string
		description   string
		logData       string
		expected      []model.LogInfo
		getLogEntries func() []model.LogInfo
	}{
		{
			name:        "TestGetLogEntriesSuccess",
			description: "Test successful Get log entries",
			logData:     `177.71.128.21 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
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
		},
		{
			name:        "TestGetLogEntriesSuccessWithNoResult",
			description: "Test successful Get log entries",
			logData:     `177.71.128.21 admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
			expected:    []model.LogInfo{},
		},
	}

	for _, test := range tests {
		service, _ := NewLogService(test.logData)
		t.Run(test.name, func(t *testing.T) {
			response := service.GetLogEntries()
			assert.Equal(t, test.expected, response)
		})
	}
}

func TestFindTopNEntries(t *testing.T) {
	tests := []struct {
		name          string
		description   string
		logData       string
		expected      []string
		field         string
		limit         int
		getLogEntries func() []model.LogInfo
	}{
		{
			name:        "TestFindTopNEntriesSuccessWhenFieldIsIpAndLimtIs3",
			description: "Test successful find top3 ip addressess",
			logData: `177.71.128.21 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.21 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.20 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.20 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.22 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.23 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
			expected: []string{"177.71.128.21", "177.71.128.20", "177.71.128.22"},
			field:    "ip",
			limit:    3,
		},
		{
			name:        "TestFindTopNEntriesSuccessWhenFieldIsUrlAndLimtIs2",
			description: "Test successful find top3 ip addressess",
			logData: `177.71.128.21 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.21 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.20 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.20 - admin [10/Jul/2018:22:21:28 +0200] "GET /test1/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.22 - admin [10/Jul/2018:22:21:28 +0200] "GET /test/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.23 - admin [10/Jul/2018:22:21:28 +0200] "GET /test1/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
			expected: []string{"/intranet-analytics/", "/test1/"},
			field:    "url",
			limit:    2,
		},
	}

	for _, test := range tests {
		service, _ := NewLogService(test.logData)
		t.Run(test.name, func(t *testing.T) {
			response := service.FindTopNEntries(test.limit, test.field)
			assert.Equal(t, test.expected, response)
		})
	}
}

func TestGetUniqueFields(t *testing.T) {
	tests := []struct {
		name          string
		description   string
		logData       string
		expected      []string
		field         string
		getLogEntries func() []model.LogInfo
	}{
		{
			name:        "TestGetUniqueFieldsSuccessWhenFieldIsIp",
			description: "Test successful find unique ip addressess",
			logData: `177.71.128.21 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.21 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.20 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.20 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.22 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.23 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
			expected: []string{"177.71.128.21", "177.71.128.20", "177.71.128.22", "177.71.128.23"},
			field:    "ip",
		},
		{
			name:        "TestFindTopNEntriesSuccessWhenFieldIsUrl",
			description: "Test successful find unique urls",
			logData: `177.71.128.21 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.21 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.20 - admin [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.20 - admin [10/Jul/2018:22:21:28 +0200] "GET /test1/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.22 - admin [10/Jul/2018:22:21:28 +0200] "GET /test/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
177.71.128.23 - admin [10/Jul/2018:22:21:28 +0200] "GET /test1/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
			expected: []string{"/intranet-analytics/", "/test1/", "/test/"},
			field:    "url",
		},
	}

	for _, test := range tests {
		service, _ := NewLogService(test.logData)
		t.Run(test.name, func(t *testing.T) {
			response := service.GetUniqueFields(test.field)
			assert.ElementsMatch(t, test.expected, response)
		})
	}
}
