package application

import (
	"io"
	"testing"
)

func TestTapeWrite(t *testing.T) {
	t.Run("should write to file correctly", func(t *testing.T) {
		tempFile, clean := createTempFile(t, "12345")
		defer clean()

		content := []byte("abc")

		tape := tape{tempFile}

		tape.Write(content)

		tempFile.Seek(0, io.SeekStart)

		newFileContents, _ := io.ReadAll(tempFile)

		got := string(newFileContents)

		want := "abc"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
