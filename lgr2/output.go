package lgr

import "runtime"
import "errors"
import "log"

type Output struct {
	Name			string
	Filters			Filters
	outputThreshold	Level
}

type OutputI interface {
	Writer(p []byte) (n int, err error)
}

// getCallerInformation retrieves information about the point in code which logged this message
// callerName is a string containing the calling functions name
// this will be printed in the log message
func getCallerInformation() (fileName string, lineNumber int, callerName string, err error) {
	// incase we encounter some panic here, let's try to exit with grace
	defer func(){
		if r := recover(); r != nil {
			fileName = ""
			lineNumber = 0
			callerName = ""
			err = errors.New("Error while trying to discover caller information, 1 or more lines may be missing from the log.")
		}
	}()
		
	// lvl is the number of levels to go up the call tree
	var lvl int = 2

	// get function http://moazzam-khan.com/blog/golang-get-the-function-callers-name/
	// get calling function
	// callStack is an array of calling entities
	callStack := make([]uintptr, 1)
	// Skip 2 levels to get the caller
	if runtime.Callers(lvl, callStack) == 0 {
			// No caller found
			callerName = "****NOT*FOUND****"
	}

	caller := runtime.FuncForPC(callStack[0]-1)
	if caller == nil {
			// caller was nil
			callerName = "nil"
	}

	// Print the file name and line number
	// https://golang.org/pkg/runtime/#Func.FileLine
	fileName, lineNumber = caller.FileLine(callStack[0]-1)

	// Print the name of the function
	callerName = caller.Name()

	return fileName, lineNumber, callerName, nil
}

func (output *Output) SetOutputThreshold(level *log.Logger){
	output.outputThreshold = level.
}
