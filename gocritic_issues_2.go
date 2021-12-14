package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

// PipelineLogEntry represents one log entry (record) read from log file.
type PipelineLogEntry struct {
	Level    string `json:"levelname"`
	Time     string `json:"asctime"`
	Name     string `json:"name"`
	Filename string `json:"filename"`
	Message  string `json:"message"`
}

func readPipelineLogFile(filename string) ([]PipelineLogEntry, error) {
	entries := []PipelineLogEntry{}

	file, err := os.Open(filename)
	if err != nil {
		return entries, err
	}

	defer func() {
		err := file.Close()

		if err != nil {
			log.Println(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := PipelineLogEntry{}
		err = json.Unmarshal([]byte(scanner.Text()), &entry)
		if err != nil {
			log.Fatal(err)
		} else {
			entries = append(entries, entry)
		}
	}

	if err := scanner.Err(); err != nil {
		return entries, err
	}

	return entries, nil
}

func main() {
	readPipelineLogFile("foobar")
}
