package lgr

import "log"
import "io"
import "io/ioutil"
import "github.com/fatih/color"


// these are the numeric values used to define levels
// the actual Loggers fall into these levels,
// and are defined in their LoggerT s
// these are the 8 logging levels listed in their intended order of urgency 
// TRACE being for detailed reporting whereas FATAL is for _total_ failure
const (
    // LevelTrace Excessive User Output
	levelTrace Level = iota
    // LevelDebug Detailed User Output
	levelDebug
    // LevelInfo Elevated User Output
	levelInfo			
    // LevelMsg Standard User Output
	levelMsg
    // LevelWarn Non-Critical Errors
	levelWarn
    // LevelError Important Errors
	levelError
    // LevelCritical Disrupting Errors
	levelCritical
    // LevelFatal System destroying, flee the building errors
	levelFatal
	defaultLogThreshold    = levelInfo
	defaultStdoutThreshold = levelMsg
	defaultFlags           = log.Ldate|log.Ltime|log.Lshortfile
)

// local defaults
var (
	// don't export these, they should be set 
	logThreshold    Level    = defaultLogThreshold
	outputThreshold Level    = defaultStdoutThreshold

	// FileHandle is the handle for the log file to write to
	fileHandle        = ioutil.Discard

	// default Outputs to write to
	// Outputs should shadow (aka implement io.Writer)
	defaultOutputs = []io.Writer{
		fileHandle,
	}
	
	defaultPrefixName = true
	defaultPrefix PrefixList = nil

)

//default logset
var (
	TRACE = &LoggerT{
        Level: levelTrace, 
        Name:   "TRACE",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgCyan),
        printDebug: true,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    }
	DEBUG = &LoggerT{
        Level: levelDebug, 
        Name:   "DEBUG",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgMagenta),
        printDebug: true,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    }
	INFO = &LoggerT{
        Level: levelInfo, 
        Name:   "INFO",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgBlue),
        printDebug: false,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    }
	MSG = &LoggerT{
        Level: levelMsg, 
        Name:   "MSG",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgWhite),
        printDebug: false,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    }
	WARN = &LoggerT{
        Level: levelWarn,
        Name:   "WARN",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgYellow).Add(color.Underline),
        printDebug: true,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    }
	ERROR = &LoggerT{
        Level: levelError,
        Name:   "ERROR",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgRed),
        printDebug: true,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    }
	CRITICAL = &LoggerT{
        Level: levelCritical,
        Name:   "CRITICAL",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgRed).Add(color.Underline),
        printDebug: true,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    }
	FATAL = &LoggerT{
        Level: levelFatal,
        Name:   "FATAL",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgRed).Add(color.Underline).Add(color.Bold),
        printDebug: true,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    }
)

func init(){
	NewLogger(TRACE,DEBUG,INFO,MSG,WARN,ERROR,CRITICAL,FATAL)
}
