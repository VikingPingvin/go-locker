package fileutils

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func EnsurePathExists(elements ...string) error {
	fullPath := filepath.Join(elements...)

	err := os.MkdirAll(fullPath, os.ModeDir)
	if err != nil {
		log.Err(err).Msg("Cannot create folder structure")
	}
	return err
}

func CreateFileAt() {

}

func MoveFile(oldPath string, newPath string) error {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		log.Err(err).Msg("File move failed!")
	}
	return err
}
