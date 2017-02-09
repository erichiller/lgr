// jww mod from spf13

package lgr

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

// Level describes the chosen log level between
// debug and critical.
type Level int

type LogType struct {
	Level           Level
	Prefix          string
	FileLogger      *log.Logger
	PrintLogger	    *log.Logger
    color           *color.Color
	Loggers         []log.Logger
	PrintDebug      bool
}

const (
    // LevelTrace Excessive User Output
	LevelTrace Level = iota
    // LevelDebug Detailed User Output
	LevelDebug
    // LevelInfo Elevated User Output
	LevelInfo			
    // LevelMsg Standard User Output
	LevelMsg
    // LevelWarn Non-Critical Errors
	LevelWarn
    // LevelError Important Errors
	LevelError
    // LevelCritical Disrupting Errors
	LevelCritical
    // LevelFatal System destroying, flee the building errors
	LevelFatal
	DefaultLogThreshold    = LevelInfo
	DefaultStdoutThreshold = LevelMsg
)

var (
	TRACE *LogType = &LogType{
        Level: LevelTrace, 
        Prefix: "TRACE: ",
        color:  color.New(color.FgCyan),
        PrintDebug: true,
    }
	DEBUG *LogType = &LogType{
        Level: LevelDebug, 
        Prefix: "DEBUG: ",
        color:  color.New(color.FgMagenta),
        PrintDebug: true,
    }
	INFO *LogType = &LogType{
        Level: LevelInfo, 
        Prefix: "INFO: ",
        color:  color.New(color.FgBlue),
        PrintDebug: false,
    }
	MSG *LogType = &LogType{
        Level: LevelMsg, 
        Prefix: "MSG: ",
        color:  color.New(color.FgWhite),
        PrintDebug: false,
    }
	WARN *LogType = &LogType{
        Level: LevelWarn,
        Prefix: "WARN: ",
        color:  color.New(color.FgYellow).Add(color.Underline),
        PrintDebug: true,
    }
	ERROR *LogType = &LogType{
        Level: LevelError,
        Prefix: "ERROR: ",
        color:  color.New(color.FgRed),
        PrintDebug: true,
    }
	CRITICAL *LogType = &LogType{
        Level: LevelCritical,
        Prefix: "CRITICAL: ",
        color:  color.New(color.FgRed).Add(color.Underline),
        PrintDebug: true,
    }
	FATAL *LogType = &LogType{
        Level: LevelFatal,
        Prefix: "FATAL: ",
        color:  color.New(color.FgRed).Add(color.Underline).Add(color.Bold),
        PrintDebug: true,
    }
	logThreshold    Level    = DefaultLogThreshold
	outputThreshold Level    = DefaultStdoutThreshold

    // FileHandle is the handle for the log file to write to
	FileHandle      io.Writer  = ioutil.Discard

    LogTypes        []*LogType = []*LogType{TRACE, DEBUG, INFO, MSG, WARN, ERROR, CRITICAL, FATAL}
)


// init will setup the standard approach of providing the user
// some feedback and logging a potentially different amount based on independent log and output thresholds.
// By default the output has a lower threshold than logged
// Don't use if you have manually set the Handles of the different levels as it will overwrite them.
func init() {
	SetStdoutThreshold(DefaultStdoutThreshold)
}

type eWriterI interface {
    Write(p []byte) (n int, err error)

}

type eWriter struct {
    eWriterI
    color *color.Color

}

// Write replacement function for io.writer
func Write(bt []byte) (n int, err error) {
    return 3,nil
}

func refreshLogTypes(){
	// see log flag constants
	// https://golang.org/pkg/log/#pkg-constants
	for _, n := range LogTypes {

        // this is an actual file on the filesystem to be logged to
		n.FileLogger = log.New(FileHandle,n.Prefix, log.Ldate|log.Ltime|log.Lshortfile)
        // this is a custom writer for the purpose of coloration
        colorOut := eWriter{ color:  n.color }
        
        // this is output to the console - going to be displayed to the user - ideally colorized
        // if the log level has PrintDebug set, the detailed log information will be included
        if n.PrintDebug {
            n.PrintLogger = log.New(colorOut,n.Prefix,log.Ldate|log.Ltime|log.Lshortfile)
        } else {
            // else if not, it will be a plain message, with only the prefix (which can be blank at the log level)
            n.PrintLogger = log.New(colorOut,n.Prefix,0)
        }
		
        // if the log level is less than the ouputThreshold (stdout)
        // and less than logThreshold (file output)
        // than don't log anything
		if n.Level < outputThreshold && n.Level < logThreshold {
			n.Loggers = []log.Logger{}
		} else if n.Level >= outputThreshold && n.Level >= logThreshold {
            // if greater than or equal to both, log to both
			n.Loggers = []log.Logger{*n.FileLogger,*n.PrintLogger}
		} else if n.Level >= outputThreshold && n.Level < logThreshold {
            // if only outputThreshold is greater, only log to PrintLogger
			n.Loggers = []log.Logger{*n.PrintLogger}
		} else {
            // else (the only option remaining is logThreshold is greater)
            // log to FileLogger only
			n.Loggers = []log.Logger{*n.FileLogger}
		}
	}

	// for _, n := range LogTypes {
	// 	*n.Logger = log.New(n.Handle, n.Prefix, log.Ldate|log.Ltime|log.Lshortfile)
	// }
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

// SetLogThreshold Establishes a threshold where anything matching or above will be logged
func SetLogThreshold(level Level) {
	logThreshold = levelCheck(level)
	refreshLogTypes()
}

// SetStdoutThreshold Establishes a threshold where anything matching or above will be output
func SetStdoutThreshold(level Level) {
	outputThreshold = levelCheck(level)
	refreshLogTypes()
}

// SetLogFile Sets the Log Handle to an io.writer
// takes a single string argument of `path` which is the path to be used as the log file
// This file will be appended to or created
func SetLogFile(path string) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		CRITICAL.Println("Failed to open log file:", path, err)
		os.Exit(-1)
	}

	INFO.Println("Logging to", file.Name())

	FileHandle = file
	refreshLogTypes()
}

// UseTempLogFile Creates a temporary file and sets the Log Handle to a io.writer created for it
func UseTempLogFile(prefix string) {
	file, err := ioutil.TempFile(os.TempDir(), prefix)
	if err != nil {
		CRITICAL.Println(err)
	}

	INFO.Println("Logging to", file.Name())

	FileHandle = file
	refreshLogTypes()
}

// DiscardLogging Disables logging
func DiscardLogging() {
	FileHandle = ioutil.Discard
	refreshLogTypes()
}

// Feedback is special. It writes plainly to the output while
// logging with the standard extra information (date, file, etc)
// Only Println and Printf are currently provided for this
func (log *LogType) Print(v ...interface{}) {
	log.color.Print(v...)
	log.FileLogger.Print(v...)
}

// Feedback is special. It writes plainly to the output while
// logging with the standard extra information (date, file, etc)
// Only Println and Printf are currently provided for this
func (log *LogType) Printf(format string, v ...interface{}) {
	log.color.Printf(format,v...)
	log.FileLogger.Printf(format,v...)
}

// Feedback is special. It writes plainly to the output while
// logging with the standard extra information (date, file, etc)
// Only Println and Printf are currently provided for this
func (log *LogType) Println(v ...interface{}) {
    log.color.Set()
	log.PrintLogger.Println(v...)
    color.Unset()
	log.FileLogger.Println(v...)
}

