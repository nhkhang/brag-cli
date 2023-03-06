package common

import (
	"fmt"
	"os"

	"github.com/goccy/go-json"
)

func LoadData(fileName string, data interface{}) error {
	f, err := os.Open(fmt.Sprintf("./data/%s.json", fileName))
	if err != nil {
		return err
	}

	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	return nil
}

func SaveData(fileName string, data interface{}) error {
	file, err := os.Create(fmt.Sprintf("./data/%s.json", fileName))
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&data); err != nil {
		return err
	}

	return nil
}
