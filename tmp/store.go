package tmp

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Store struct {
	FileName string // json file name in tmp folder
}

func (s *Store) Get(data *map[string]string) error {

	if tmp_file, err := getFilePath(s.FileName); err == nil {
		if file, err := ioutil.ReadFile(tmp_file); err == nil {
			_ = json.Unmarshal([]byte(file), data)
		} else {
			return err
		}
	} else {
		return err
	}

	return nil
}

func (s *Store) Set(data map[string]string) {
	if tmp_file, err := getFilePath(s.FileName); err == nil {
		file, _ := json.MarshalIndent(data, "", "  ")
		_ = ioutil.WriteFile(tmp_file, file, 0666)
	}
}

func getFilePath(jsonFileName string) (string, error) {

	tmp_folder := "/tmp"
	if _, err := os.Stat(tmp_folder); errors.Is(err, os.ErrNotExist) {
		home, _ := os.UserHomeDir()
		tmp_folder = filepath.Join(home, "tmp")
	}

	tmp_file := filepath.Join(tmp_folder, jsonFileName)
	if _, err := os.Stat(tmp_file); errors.Is(err, os.ErrNotExist) {
		d1 := []byte("[]")
		err := os.WriteFile(tmp_file, d1, 0666)

		return tmp_file, err
	}

	return tmp_file, nil
}
