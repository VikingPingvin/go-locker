package fileutils

import (
	"os"
	"testing"

	"github.com/rs/zerolog/log"
)

// Test EnsurePathexists with 0, 1 and more []string
func TestEnsurePath(t *testing.T) {
	t.Run("with empty input", func(t *testing.T) {
		want := ErrorEmptyArgument
		got := EnsurePathExists("")

		if got.Error() != want {
			t.Errorf("got %s, want %s", got.Error(), want)
		}
	})

	// TODO: Figure out how to test this
	t.Run("With 1 input", func(t *testing.T) {
		cwd, _ := os.Getwd()
		got := EnsurePathExists(cwd)

		log.Logger.Info().Msgf("%v", got)
	})

	// TODO: test with more inputs
}
