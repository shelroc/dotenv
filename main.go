package dotenv

import (
	utils "github.com/shelroc/dotenv/src"
)

func Read(path string) (map[string]string, error) {
	return utils.Read(path)
}

func ReadDefault() (map[string]string, error) {
	return utils.Read(utils.FileName)
}
