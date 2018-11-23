package main

import (
	"io/ioutil"
	"os"
)

func Read(path string) ([]byte, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func Write(path string, data []byte) error {
	// Delete file if exits.
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		// Ignore error.
		os.Remove(path)
	}

	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return err
	}

	return nil
}

func Delete(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

// GetDirFiles returns all file names in dir.
func GetDirFiles(path string) ([]string, error) {
	files := make([]string, 0)

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, nil
	}

	for _, f := range dir {
		if f.IsDir() {
			continue
		}

		files = append(files, f.Name())
	}

	return files, nil
}
