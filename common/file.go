package common

import (
	"fmt"
	"os"

	"github.com/goccy/go-json"
)

const DATA_PATH = "./data/"

func GetDataFilePath(name string) string {
	return fmt.Sprintf("%s%s.json", DATA_PATH, name)
}

func LoadData(fileName string, data interface{}) error {
	filePath := fmt.Sprintf("%s%s.json", DATA_PATH, fileName)
	f, err := os.Open(filePath)
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
	filePath := GetDataFilePath(fileName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
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
