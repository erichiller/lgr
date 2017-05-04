package lgr

import "io"

import "github.com/fatih/color"



type ConsoleColorWriter struct {
	Filter
	color									*color.Color
}


func (log *LoggerConfig) ConsoleColorWriter(p []byte) (n int, err error){

}



func (Output *Output) SetFilters(filters []Filters){
	//set filters only for this Output
}

func (log *LoggerConfig) SetFilters(filters []Filters){
	//set filters for all Writers
}
