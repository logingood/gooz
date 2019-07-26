package zfile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

type Store struct {
	FilePath string
	FD       afero.File
	GoozFS   afero.Fs
}

// Create empty Store object with Filesystem and Path
func New(fs afero.Fs, filePath string) *Store {
	return &Store{
		GoozFS:   fs,
		FilePath: filePath,
	}
}

// Open FD for reading
func (f *Store) Open() error {
	goozFD, err := f.GoozFS.Open(f.FilePath)
	if err != nil {
		log.Errorf("Error has occured when opening a file %s, error: %s", f.FilePath, err)
	}

	f.FD = goozFD

	return err
}

// Marshal everything into map[string]
func (f *Store) Read() (data []map[string]interface{}, err error) {
	if f.FD == nil {
		return data, fmt.Errorf("File descriptor is nil")
	}

	byteData, err := ioutil.ReadAll(f.FD)
	if err != nil {
		log.Errorf("Can't read file content %s", err)
	}
	// Test file is closed
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		log.Errorf("Can't parse json %s", err)
	}

	return data, err
}

// Close files correctly
func (f *Store) Close() error {
	if f.FD != nil {
		f.FD.Close()
		return nil
	}

	return fmt.Errorf("File descriptor is nil")
}
