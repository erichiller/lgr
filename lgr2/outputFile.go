package lgr

import (
	"os"
	"ioutil"
)

type FileOutput struct {
	Output
	fileHandle		os.File
	filePath		string
}


// UseTempLogFile Creates a temporary file and sets the Log Handle to a io.writer created for it
// prefix is a string to be used as the filename prefix for the temporary file
func UseTempLogFile(prefix string) ( success bool, error err ) {
	file, err := ioutil.TempFile(os.TempDir(), prefix)
	if err != nil {
		return false, error.Errorf("Failed to open log file:%s\n%s", path, err)
	}
	output.fileHandle = file
	output.filePath = file.Name()
	return true, nil
}

// SetLogFile Sets the Log Handle to an io.writer
// takes a single string argument of `path` which is the path to be used as the log file
// This file will be appended to or created
func (output *FileHandle) SetLogFile(path string) ( success bool, error err ) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return false, error.Errorf("Failed to open log file:%s\n%s", path, err)
	}
	output.fileHandle = file
	output.filePath = path
	return true, nil
}

func (output *FileHandle) Writer(p []byte) (n int, err error) {
	
}
