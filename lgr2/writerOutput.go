
import "runtime"

type Output struct {
	Name		string
	Writer		io.Writer
	Filters		[]Filter
}


func getCallerInformation() (fileNameLine string, callerName string, err) {


	defer recover( return "" , "", error.Error("Error while trying to discover caller information, 1 or more lines may be missing from the log.") )
	// lvl is the number of levels to go up the call tree
	var lvl int = 2

	// get function http://moazzam-khan.com/blog/golang-get-the-function-callers-name/
	// get calling function
	// callStack is an array of calling entities
	callStack := make([]uintptr, 1)
	// callerName is a string containing the calling functions name
	// this will be printed in the log message
	var callerName string
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
	fileNameLine := caller.FileLine(callStack[0]-1)

	// Print the name of the function
	callerName = caller.Name()

	return fileNameLine, callerName, nil
}
