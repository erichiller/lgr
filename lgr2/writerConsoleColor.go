package lgr

import "io"

import "github.com/fatih/color"



type ConsoleColorOutput struct {
	Output
	color									*color.Color
}

// ConsoleColorWriter acts as a modifier pre-output for the logs.
// Here we can add additional information (such the function the log is in)
// or styling, such as coloration
func (output Output) ConsoleColorWriter(p []byte) (n int, err error) {

// NOTE:
// THIS CURRENTLY ONLY IS DOING STDOUT!!
// I THINK YOU WANT MORE THAN THAT
//
// BEWARE OF COLOR, ON MULTIPLE-THREADS, CAN BE BAD NEWS
//  when multiple writers are writing to the same output
//  for example, the console, the output can become corrupt
//  most likely because of ANSI injection sequences conflicting
//  but I havent had time to investigate this yet



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

func (output *Output) SetFilters(filters []Filters){
	//set filters only for this Output
}

func (log *LoggerConfig) SetFilters(filters []Filters){
	//set filters for all Writers
}
