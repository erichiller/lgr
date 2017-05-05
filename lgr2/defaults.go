package lgr

import "log"
import "io"
import "io/ioutil"
import "github.com/fatih/color"


// these are the numeric values used to define levels
// the actual Loggers fall into these levels,
// and are defined in their LoggerT s
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

var defaultLog = Log{
	TRACE: &LoggerT{
        Level: levelTrace, 
        Name:   "TRACE",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgCyan),
        printDebug: true,
        Logger: &TRACE,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    },
	DEBUG: &LoggerT{
        Level: levelDebug, 
        Name:   "DEBUG",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgMagenta),
        printDebug: true,
        Logger: &DEBUG,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    },
	INFO: &LoggerT{
        Level: levelInfo, 
        Name:   "INFO",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgBlue),
        printDebug: false,
        Logger: &INFO,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    },
	MSG: &LoggerT{
        Level: levelMsg, 
        Name:   "MSG",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgWhite),
        printDebug: false,
        Logger: &MSG,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    },
	WARN: &LoggerT{
        Level: levelWarn,
        Name:   "WARN",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgYellow).Add(color.Underline),
        printDebug: true,
        Logger: &WARN,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    },
	ERROR: &LoggerT{
        Level: levelError,
        Name:   "ERROR",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgRed),
        printDebug: true,
        Logger: &ERROR,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    },
	CRITICAL: &LoggerT{
        Level: levelCritical,
        Name:   "CRITICAL",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgRed).Add(color.Underline),
        printDebug: true,
        Logger: &CRITICAL,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    },
	FATAL: &LoggerT{
        Level: levelFatal,
        Name:   "FATAL",
		PrefixName: defaultPrefixName,
        Prefix: defaultPrefix,
        color:  color.New(color.FgRed).Add(color.Underline).Add(color.Bold),
        printDebug: true,
        Logger: &FATAL,
        Flags: defaultFlags,
        Outputs: defaultOutputs,
    },
}

func init(){
	NewLogger(defaultLog)
}
