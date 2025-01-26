package utils

import (
	// "fmt"
	"os"
	"strings"
)

const file_name = ".env"

func read_env_file() (string, error) {
	data, err := os.ReadFile(file_name)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func create_env_map(line string, table map[string]string) {
	pairs := strings.Split(line, "=")
	if len(pairs) == 2 {
		table[pairs[0]] = pairs[1]
	}
}

func get_lines(data string, table map[string]string) {
	lines := strings.Split(data, "\n")
	for i := range lines {
		if lines[i] != "" {
			create_env_map(lines[i], table)
		}
	}
}

func Read() (map[string]string, error) {
	table := make(map[string]string)

	data, err := read_env_file()
	if err != nil {
		return nil, err
	}

	get_lines(data, table)
	return table, nil
}
