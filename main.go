package dotenv

import (
	utils "github.com/shelroc/dotenv/src"
)

func Read() (map[string]string, error) {
	return utils.Read()
}
