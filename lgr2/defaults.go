package lgr2


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
