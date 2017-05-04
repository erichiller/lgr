package lgr

import "io"

import "github.com/fatih/color"



type ConsoleColorWriter struct {
	Filter
	color									*color.Color
}


func (log *LoggerConfig) ConsoleColorWriter(p []byte) (n int, err error){

}

