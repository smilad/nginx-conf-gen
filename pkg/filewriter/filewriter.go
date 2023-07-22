package filewriter

import (
	"io"
	"log"
	"os"
)

type file struct {
	file *os.File
}

func NewFile(filename string, flag int) *file {
	f, err := os.OpenFile(filename, flag, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return &file{file: f}
}

func (f *file) Write(p []byte) (int, error) {
	n, err := f.file.Write(p)
	if err != nil {
		return 0, err
	}

	return n, err
}

func (f *file) ReadAll() (string, error) {
	// Seek to the beginning of the file
	_, err := f.file.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	// Read the entire file into a byte slice
	data, err := io.ReadAll(f.file)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string and return it
	return string(data), nil
}
