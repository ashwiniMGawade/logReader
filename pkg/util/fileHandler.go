package util

import (
	"bufio"
	"os"
)

/*
*
* Method for reading the file from filePath
* Responsible for reading the file from specified path
  - @param filePath - Path of the file to read
  - @returns content of the file if file found at the path
    error if could not Open the file
*/
func ReadFile(filePath string) (content string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return content, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return content, err
	}

	return content, nil
}
