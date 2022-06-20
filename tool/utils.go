package tool

import (
	"io/ioutil"
	"os"
)

func CreateDir(path string) error {
	// Check if folder exists
	_, err := os.Stat(path)

	// Create directory if not exists
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModeDir|0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateFile(filename string) error {
	// Check if folder exists
	_, err := os.Stat(filename)

	// Create directory if not exists
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}

		defer file.Close()
	}

	return nil
}

func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFileOrDirectory(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(filepath string) string {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return ""
	}

	return string(data)
}

func WriteFileLineByLine(filepath string, data []string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, line := range data {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteFileAll(filepath string, data string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	file.Sync()

	return nil
}
