package constants

const (

	/* Regular expression pattern used to match log entry against various component in go struct:
	* - IP address: (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
	* - User: (\S+)?
	* - Timestamp Info: \[(.*?)\]\
	* - Request Type/ HTTP Method: (\w+)
	* - URL: (.*?)
	* - HTTP Protocol: (HTTP\/[\d.]+)
	* - Status Code: (\d+)
	* - Response Size: (\d+)
	* - Platfrom : (.*)*/
	LOG_REGEX_PATTERN = `^(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) (\S+)? (\S+)? \[(.*?)\]\ "(\w+) (.*?) (HTTP\/[\d.]+)" (\d+) (\d+) "-" "(.*)"`

	IP_FIELD_NAME   = "ip"
	URL_FIELD_NAME  = "url"
	DEFAULT_USER    = "guest"
	LOG_DATE_FORMAT = "02/Jan/2006:15:04:05 -0700"
)
