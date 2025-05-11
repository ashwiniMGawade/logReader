# HTTP Log File Reader with Analysis

A Golang application that parses logs and provides useful information about IP addresses, URL, Request data and platform Info. Application is capable of giving analytical infomration using logs.

## Functionalities and Features

- Functionality to read the log files and Handle standard HTTP access log format 
- Written in Golang using go tools
- Includes unit tests for each component
- Structure is divided into small managable packages
- Counts unique IP addresses in the log file
- Identifies top 3 most visited URLs
- Identifies top 3 most active IP addresses

## Prerequisites

- Go (v1.23.4 or higher)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/ashwiniMGawade/logReader.git
cd logReader
git checkout master
```

2. Install dependencies:
```bash
go mod download
```


## Usage

1. Build the project:
```bash
go build
```

2. Run the project:
```bash
./logReader
```
or 
```bash
go run main.go
```

The analyzer will process the included sample log file (`programming-task-example-data.log`) and output:
- Number of unique IP addresses
- Top 3 most visited URLs 
- Top 3 most active IP addresses


### Project Structure
```bash
├── pkg/
│   ├── constants/          # Constants
│   ├── helpers/            # helpers
│   ├── model/              # Model or data structs declaration 
│   ├── service/            # Business logic
│   └── util                # Common utility functions
├── main.go                 # Entrypoint to application
└── programming-task-example-data.log    # Sample log file
```

### Commands 

- `go build` - Compiles go application and generate executable
- `./logReader` - Runs the compiled application
- `go run main.go` - Runs the application in development without generating the build
- `go test -v -skip=TestMain -cover ./pkg/... -coverprofile=profile.cov` - Runs the test suite and return the code coverage

### Running Tests
```bash
go test -v -skip=TestMain -cover ./pkg/... -coverprofile=profile.cov
```

## Log File Format

The application expects log files in the following format:

Example 1 : Guest access
```
50.112.00.28 - - [11/Jul/2018:15:49:46 +0200] "GET /faq/how-to-install/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; ca-ad) AppleWebKit/531.2+ (KHTML, like Gecko) Safari/531.2+ Epiphany/2.30.6"
```

Example 1 : admin access
```
50.112.00.11 - admin [11/Jul/2018:17:31:56 +0200] "GET /asset.js HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6"
```
### LogInfo Struct

The `LogInfo` struct represents information about a log entry in a log file.

## Fields

- `IpAddress` (string): The IP address associated with the log entry.
- `User` (string): The user associated with the log entry.
- `TimeStamp` (time.Time): The timestamp of the log entry.
- `RequestType` (string): The type of HTTP request made.
- `RequestUrl` (string): The URL of the request.
- `ResponseCode` (int): The HTTP response code returned.
- `ResponseSize` (int64): The size of the response in bytes.
- `Platform` (string): Information about the platform and user agent.

## Constants

- `LOG_REGEX_PATTERN` : Stores the regex pattern to parse the log line
- `IP_FIELD_NAME`  : Name of the ip-address field
- `URL_FIELD_NAME` : Name of the url field
- `DEFAULT_USER` : Default value of the user "guest"
- `LOG_DATE_FORMAT`: "02/Jan/2006:15:04:05 -0700"

### LogParser utility function

Handles parsing of all log lines and generate the slice of logInfo struct.
Throws the error if file not found.
Uses scanner to read one file at a time and joins them using '\n'\
Returns all the data in the file

### FileHandler utility function

Handles reading the files content.
It iterates through each line of the content by splitting the data on '\n'.
For each log line, it matches with the standard regex. The line would be converted to string only if it matches the pattern.

### LogService Interface

This interface incapsulates the most needed functionalities in the service. It includes following functions:
- `GetLogEntries() []model.LogInfo` : Returns all log entries
- `FindTopNEntries(limit int, field string) []string` : Returns top `limit` entries of the `field`, `field` can accept either 'ip' or 'url'
- `GetUniqueFields(field string) []string` : Returns unique values of `field`

### logService struct

Provides implementation of the LogService Interface. It has all the three methods implemented in it.
It also has its own field `LogEntries []model.LogInfo` which acts as the data store and stores all logEntry data to commonly accessed among functions
Apart from interface implementation, It has following its own functions
- `NewLogService(data string) (LogService, error) `: Returns new LogService Interface with logService implementation
- `(l *logService) ParseLog(data string) error `: Accepts the logService object as reference, parses the data and initialize the LogEntries

## Error Handling

- If regex does not match with any of the line in log, that line is skipped during parsing
- File read errors are caught and reported
- All the other comilation errors are handled in go using static typing

## Assumptions

1. Log file format follows the standard pattern as mentioned
2. Anything other than the standard format of log would be skipped and not parsed
3. Dates need to be in the format `"02/Jan/2006:15:04:05 -0700"` else parsing error would be thrown
4. URLs are case-sensitive
5. IP addresses are well-formed

## License

This project is licensed under the MIT License - see the LICENSE file for details.