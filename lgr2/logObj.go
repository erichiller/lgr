package lgr


import "io"
import "log"

import "github.com/fatih/color"

type LogI interface {
	GetFilters() []Filters
}

type LoggerT struct {
	Level            Level
	Name             string
	Prefix           PrefixList
	PrefixName       bool
	Writer           *io.Writer				// Writer is a pointer to a Writer that already exists in Output .Writer
	Logger           **log.Logger
	color            *color.Color
	printDebug       bool					// this is for internal debugging of lgr.
	Flags            int
	Outputs          []OutputI
	AllowableFilters Filters
	HighlightFilters Filters
}

type PrefixList []interface{}

type Filters []Filter

type Filter struct {
	Keywords	[]string
	Level		int
}

type Log map[*log.Logger]*LoggerT


func NewLog(log Log){
	for _, n := range log {
			
		// if the log level is less than the outputThreshold (stdout)
		// and less than LoggerThreshold (file output)
		// than don't log anything
		if n.Level < outputThreshold && n.Level < LoggerThreshold {
			n.Handle = ioutil.Discard
		} else if n.Level >= outputThreshold && n.Level >= LoggerThreshold {
			// if greater than or equal to both, log to both
			n.Handle = io.MultiWriter(FileHandle, n)
		} else if n.Level >= outputThreshold && n.Level < LoggerThreshold {
			// if only outputThreshold is greater, only log to console
			n.Handle = n
		} else {
			// else (the only option remaining is LoggerThreshold is greater)
			// log to FileLogger only
			n.Handle = FileHandle
		}

		*n.Logger = log.New(n.Handle, n.Prefix, n.Flags)
	}
}
		

// SetPrefix allows for changing the prefixes of ALL logs in lgr.
func SetPrefix(prefix string){
	for _, n := range LoggerTs {
        n.Prefix = prefix
    }
	refreshLoggerTs()
    INFO.Printf("NewPrefix(%+v)",prefix)
}
//                             50c  |                       |        |                    |
//<time> <filename.ext>(30char, left) line ##(4 char, left) (pid,ppid)MESSAGE_TYPE(8, left) function() ?****EVENT**** message ?Value

// SetPrefix allows for changing the prefix of a specific log.
func (log *LoggerT) SetPrefix(prefix string){
    log.Prefix = prefix
    refreshLoggerTs()
}

// AppendPrefix allows for appending to the prefixes of ALL lgr logs 
func AppendPrefix(prefix string){
	for _, n := range LoggerTs {
        n.Prefix = prefix + n.Prefix
    }
	refreshLoggerTs()
    INFO.Printf("NewPrefix(%+v)",prefix)
}

// AppendPrefix allows for appending to the prefix of a specific log.
func (log *LoggerT) AppendPrefix(prefix string){
    log.Prefix = prefix + log.Prefix
    refreshLoggerTs()
}

//Filter lets you add Terms to the Filter
func (log *lgr) Filter() {

}

func refreshLoggerTs(){
	// see log flag constants
	// https://golang.org/pkg/log/#pkg-constants
	for _, n := range LoggerTs {
				
				// if the log level is less than the outputThreshold (stdout)
				// and less than LoggerThreshold (file output)
				// than don't log anything
		if n.Level < outputThreshold && n.Level < LoggerThreshold {
			n.Handle = ioutil.Discard
		} else if n.Level >= outputThreshold && n.Level >= LoggerThreshold {
			// if greater than or equal to both, log to both
			n.Handle = io.MultiWriter(FileHandle, n)
		} else if n.Level >= outputThreshold && n.Level < LoggerThreshold {
			// if only outputThreshold is greater, only log to console
			n.Handle = n
		} else {
			// else (the only option remaining is LoggerThreshold is greater)
			// log to FileLogger only
			n.Handle = FileHandle
		}
		*n.Logger = log.New(n.Handle, n.Prefix, n.Flags)
	}
}


func New() (logger LoggerT) {



	// logger.Writer = 
}








/**


3 log functions

Log -> internal
Printf -> Console , meant for user, web etc,
(Error? -> User, non crashing)
Critical -> System Crash

**/
