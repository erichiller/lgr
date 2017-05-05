
package lgr

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"

)

import "strings"

// Level describes the chosen log level between
// debug and critical.
type Level int

type LogType struct {
	Level           Level
	Name            string
	Prefix          string
	Handle          io.Writer
	Logger          **log.Logger
	color           *color.Color
	PrintDebug      bool
	Flags           int
}


// Write acts a modifier pre-output for the logs.
// Here we can add additional information (such the function the log is in)
// or styling, such as coloration
func (lt LogType) Write(p []byte) (n int, err error) {

// NOTE:
// THIS CURRENTLY ONLY IS DOING STDOUT!!
// I THINK YOU WANT MORE THAN THAT
//
// BEWARE OF COLOR, ON MULTIPLE-THREADS, CAN BE BAD NEWS
//  when multiple writers are writing to the same output
//  for example, the console, the output can become corrupt
//  most likely because of ANSI injection sequences conflicting
//  but I havent had time to investigate this yet

		// get function http://moazzam-khan.com/blog/golang-get-the-function-callers-name/
	// get calling function
	// callStack is an array of calling entities
	callStack := make([]uintptr, 1)
	// callerName is a string containing the calling functions name
	// this will be printed in the log message
	var callerName string
	// Skip 2 levels to get the caller
	if runtime.Callers(2, callStack) == 0 {
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


		


	// split the received message on colons
	var str string = string(p[:])
	var strs []string = strings.SplitN(str,":",6)
	var msg string = str
	// the first 5 of which are time, etc in a normal message
	if len(strs) >= 6 {
		// but the 6th is the message type, ie. DEBUG
		// which can be used to map back to the logger
		msg = strs[5]
	} 
	// and now we can check if it should be colorized, etc.
	if !lt.PrintDebug { 
		lt.color.Print(msg)
	} else {
		lt.color.Print(str)
	}
	return len(p), nil
}


// init will setup the standard approach of providing the user
// some feedback and logging a potentially different amount based on independent log and output thresholds.
// By default the output has a lower threshold than logged
// Don't use if you have manually set the Handles of the different levels as it will overwrite them.
func init() {
	SetStdoutThreshold(DefaultStdoutThreshold)
	SetLogThreshold(DefaultStdoutThreshold)
	
}


func refreshLogTypes(){
	// see log flag constants
	// https://golang.org/pkg/log/#pkg-constants
	for _, n := range LogTypes {
	
		// if the log level is less than the outputThreshold (stdout)
		// and less than logThreshold (file output)
		// than don't log anything
		if n.Level < outputThreshold && n.Level < logThreshold {
			n.Handle = ioutil.Discard
		} else if n.Level >= outputThreshold && n.Level >= logThreshold {
			// if greater than or equal to both, log to both
			n.Handle = io.MultiWriter(FileHandle, n)
		} else if n.Level >= outputThreshold && n.Level < logThreshold {
			// if only outputThreshold is greater, only log to console
			n.Handle = n
		} else {
			// else (the only option remaining is logThreshold is greater)
			// log to FileLogger only
			n.Handle = FileHandle
		}

		*n.Logger = log.New(n.Handle, n.Prefix, n.Flags)

	}

}

// LogThreshold returns the current global log threshold.
// Level is the current Log Level ( file output level )
func LogThreshold() Level {
	return logThreshold
}

// StdoutThreshold returns the current global output threshold.
// Level is the current Stdout ( terminal output level )
func StdoutThreshold() Level {
	return outputThreshold
}

// levelCheck Ensures that the level provided is within the bounds of available levels
func levelCheck(level Level) Level {
	switch {
		case level <= LevelTrace:
			return LevelTrace
		case level >= LevelFatal:
			return LevelFatal
		default:
			return level
	}
}

// SetLogFlags runs log.SetFlags on all of the log handles contained within LogTypes
func SetLogFlags(flags int) {
	for _, n := range LogTypes {
		n.Flags = flags
	}
	refreshLogTypes()
	INFO.Printf("DefaultFlags(%+v)",flags)
}

// SetLogThreshold Establishes a threshold where anything matching or above will be logged
func SetLogThreshold(level Level) {
	logThreshold = levelCheck(level)
	refreshLogTypes()
	INFO.Printf("SetLogThreshold(%+v/%+v)",level,logThreshold)
}

// SetStdoutThreshold Establishes a threshold where anything matching or above will be output
func SetStdoutThreshold(level Level) {
	outputThreshold = levelCheck(level)
	refreshLogTypes()
	INFO.Printf("SetStdoutThreshold(%+v/%+v)",level,outputThreshold)
}




// StringToLevel returns the level which has the name levelName: 
// , TRACE 
// , DEBUG 
// , INFO 
// , MSG 
// , WARN 
// , ERROR 
// , CRITICAL 
// , FATAL 
func StringToLevel(levelName string) Level {
	for _, n := range LogTypes {
		if strings.ToLower(n.Name) == strings.ToLower(levelName) {
			return n.Level
		}
	}
	return DefaultLogThreshold
}

// LevelToString takes type level and converts it to a string readable representation
func LevelToString(level Level) string {
	for _, n := range LogTypes {
		if n.Level == level {
			return n.Name
		}
	}
	return "<unknown level name>"
}

