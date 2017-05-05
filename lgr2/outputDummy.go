package lgr

import "ioutil"


type DiscardOutput struct {
	Output
	fileHandle		os.File
}


// DiscardLogging Disables logging
func (output *Output) DiscardLogging() {
	output.fileHandle = ioutil.Discard
}

//  func (output *FileHandle) 
func (output *Output) Writer(p []byte) (n int, err error) {
	
}
