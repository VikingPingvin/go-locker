package fileutils

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

const (
	ErrorEmptyArgument = "Input is empty"
)

func EnsurePathExists(elements ...string) error {
	fullPath := filepath.Join(elements...)

	if elements == nil || len(elements) == 0 || fullPath == "" {
		emptyErr := errors.New(ErrorEmptyArgument)
		log.Err(emptyErr).Msg("Cannot create folder structure")
		return emptyErr
	}

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
