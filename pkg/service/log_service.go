package service

import (
	"logreader/pkg/constants"
	"logreader/pkg/model"
	"logreader/pkg/util"
	"sort"
)

//go:generate moq -pkg mocks -out ./mocks/log_service_mock.go . LogService
type LogService interface {
	GetLogEntries() []model.LogInfo
	FindTopNEntries(limit int, field string) []string
	GetUniqueFields(field string) []string
}

type logService struct {
	LogEntries []model.LogInfo
}

func NewLogService(data string) (LogService, error) {
	l := logService{}
	err := l.ParseLog(data)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (l *logService) ParseLog(data string) error {
	logEntries, err := util.ParseLog(data)
	if err != nil {
		return err
	}
	l.LogEntries = logEntries
	return nil
}

func (l logService) GetLogEntries() []model.LogInfo {
	if len(l.LogEntries) > 0 {
		return l.LogEntries
	}
	return []model.LogInfo{}
}

func (l logService) FindTopNEntries(limit int, field string) []string {
	fieldCount := make(map[string]int)

	// Count occurrences of the specified field
	for _, entry := range l.LogEntries {
		switch field {
		case constants.IP_FIELD_NAME:
			fieldCount[entry.IpAddress]++
		case constants.URL_FIELD_NAME:
			fieldCount[entry.RequestUrl]++
		}
	}

	// Sort the field values by frequency
	type fieldFrequency struct {
		Field     string
		Frequency int
	}
	var frequencies []fieldFrequency
	for f, count := range fieldCount {
		frequencies = append(frequencies, fieldFrequency{Field: f, Frequency: count})
	}
	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].Frequency > frequencies[j].Frequency
	})

	// Get the top N field values
	var topN []string
	for i := 0; i < limit && i < len(frequencies); i++ {
		topN = append(topN, frequencies[i].Field)
	}

	return topN
}

func (l logService) GetUniqueFields(field string) []string {
	// The struct{} value is essentially a placeholder that occupies zero bytes of memory, making it efficient for this purpose.
	uniqueEntires := make(map[string]struct{})

	// Collect unique fields
	for _, entry := range l.LogEntries {
		switch field {
		case constants.IP_FIELD_NAME:
			uniqueEntires[entry.IpAddress] = struct{}{}
		case constants.URL_FIELD_NAME:
			uniqueEntires[entry.RequestUrl] = struct{}{}
		}
	}

	// Convert unique IP addresses to a slice
	var uniqueDataList []string
	for ip := range uniqueEntires {
		uniqueDataList = append(uniqueDataList, ip)
	}

	// Sort the unique IP addresses
	sort.Strings(uniqueDataList)

	return uniqueDataList
}
