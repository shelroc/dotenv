package utils

import (
	"os"
	"strings"
)

const FileName = ".env"

// Reads the passed env file and returns its contents as a string
func readEnvFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Takes a line from an env file and makes the part before "=" the key
// and the part after "=" the value of the hash map
func createEnvTable(line string, table map[string]string) {
	pairs := strings.Split(line, "=")
	if len(pairs) == 2 {
		key := strings.TrimSpace(pairs[0])
		value := strings.TrimSpace(pairs[1])
		table[key] = value
	}
}

func getLines(data string, table map[string]string) {
	lines := strings.Split(data, "\n")
	for i := range lines {
		if lines[i] != "" {
			createEnvTable(lines[i], table)
		}
	}
}

func Read(path string) (map[string]string, error) {
	table := make(map[string]string)
	data, err := readEnvFile(path)
	if err != nil {
		return nil, err
	}
	getLines(data, table)
	return table, nil
}
