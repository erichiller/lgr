// jww mod from spf13

package lgr

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
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
    DefaultFlags           = log.Ldate|log.Ltime|log.Lshortfile
)

// LOG LEVELS
// these are the 8 logging levels listed in their intended order of urgency 
// TRACE being for detailed reporting whereas FATAL is for _total_ failure
var (
    TRACE *log.Logger
	DEBUG *log.Logger
	INFO *log.Logger
	MSG *log.Logger
	WARN *log.Logger
	ERROR *log.Logger
	CRITICAL *log.Logger
	FATAL *log.Logger
)

var (
	Trace *LogType = &LogType{
        Level: LevelTrace, 
        Name:   "TRACE",
        Prefix: "TRACE: ",
        color:  color.New(color.FgCyan),
        PrintDebug: true,
        Logger: &TRACE,
        Flags: DefaultFlags,
    }
	Debug *LogType = &LogType{
        Level: LevelDebug, 
        Name:   "DEBUG",
        Prefix: "DEBUG: ",
        color:  color.New(color.FgMagenta),
        PrintDebug: true,
        Logger: &DEBUG,
        Flags: DefaultFlags,
    }
	Info *LogType = &LogType{
        Level: LevelInfo, 
        Name:   "INFO",
        Prefix: "INFO: ",
        color:  color.New(color.FgBlue),
        PrintDebug: false,
        Logger: &INFO,
        Flags: DefaultFlags,
    }
	Msg *LogType = &LogType{
        Level: LevelMsg, 
        Name:   "MSG",
        Prefix: "MSG: ",
        color:  color.New(color.FgWhite),
        PrintDebug: false,
        Logger: &MSG,
        Flags: DefaultFlags,
    }
	Warn *LogType = &LogType{
        Level: LevelWarn,
        Name:   "WARN",
        Prefix: "WARN: ",
        color:  color.New(color.FgYellow).Add(color.Underline),
        PrintDebug: true,
        Logger: &WARN,
        Flags: DefaultFlags,
    }
	Error *LogType = &LogType{
        Level: LevelError,
        Name:   "ERROR",
        Prefix: "ERROR: ",
        color:  color.New(color.FgRed),
        PrintDebug: true,
        Logger: &ERROR,
        Flags: DefaultFlags,
    }
	Critical *LogType = &LogType{
        Level: LevelCritical,
        Name:   "CRITICAL",
        Prefix: "CRITICAL: ",
        color:  color.New(color.FgRed).Add(color.Underline),
        PrintDebug: true,
        Logger: &CRITICAL,
        Flags: DefaultFlags,
    }
	Fatal *LogType = &LogType{
        Level: LevelFatal,
        Name:   "FATAL",
        Prefix: "FATAL: ",
        color:  color.New(color.FgRed).Add(color.Underline).Add(color.Bold),
        PrintDebug: true,
        Logger: &FATAL,
        Flags: DefaultFlags,
    }
	logThreshold    Level    = DefaultLogThreshold
	outputThreshold Level    = DefaultStdoutThreshold

    // FileHandle is the handle for the log file to write to
	FileHandle      io.Writer  = ioutil.Discard

    LogTypes        []*LogType = []*LogType{Trace, Debug, Info, Msg, Warn, Error, Critical, Fatal}
)

func (lt LogType) Write(p []byte) (n int, err error) {
    var str string = string(p[:])
    var strs []string = strings.SplitN(str,":",6)
    var msg string = str
    if len(strs) >= 6 {
        msg = strs[5]
    } 
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

