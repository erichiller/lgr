package lgr

import (
	"image/color"
	"io"
	"log"
)


type LoggerConfigI interface {
	GetFilters() []Filters
}

type LoggerConfig struct {
	Level Level
	Name  string
	//Prefix     string
	Prefix           []interface{}
	Writer           *io.Writer				// Writer is a pointer to a Writer that already exists in Output .Writer
	Logger           **log.Logger
	color            *color.Color
	PrintDebug       bool
	Flags            int
	BlockFilters     Filters
	HighlightFilters Filters
}

type Filters []Filter

type Filter struct {
	Keywords	[]string
	Level			int
}



// defaults // move these //

var (
	defaultOutputs = []Output{

	}
)
		

// SetPrefix allows for changing the prefixes of ALL logs in lgr.
func SetPrefix(prefix string){
	for _, n := range LogTypes {
        n.Prefix = prefix
    }
	refreshLogTypes()
    INFO.Printf("NewPrefix(%+v)",prefix)
}



// SetPrefix allows for changing the prefix of a specific log.
func (log *LogType) SetPrefix(prefix string){
    log.Prefix = prefix
    refreshLogTypes()
}

// AppendPrefix allows for appending to the prefixes of ALL lgr logs 
func AppendPrefix(prefix string){
	for _, n := range LogTypes {
        n.Prefix = prefix + n.Prefix
    }
	refreshLogTypes()
    INFO.Printf("NewPrefix(%+v)",prefix)
}

// AppendPrefix allows for appending to the prefix of a specific log.
func (log *LogType) AppendPrefix(prefix string){
    log.Prefix = prefix + log.Prefix
    refreshLogTypes()
}

//Filter lets you add Terms to the Filter
func (log *lgr) Filter() {

}

func refreshLoggerConfigs(){
	// see log flag constants
	// https://golang.org/pkg/log/#pkg-constants
	for _, n := range LoggerConfigs {
				
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


func New() (logger LoggerConfig) {



	// logger.Writer = 
}








/**


3 log functions

Log -> internal
Printf -> Console , meant for user, web etc,
(Error? -> User, non crashing)
Critical -> System Crash

**/
