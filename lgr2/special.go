package lgr

import "time"
import "strconv"
import "runtime"


var Title string = "lgr logger\nEric D Hiller\n"
var Created string = strconv.Itoa(time.Now().Year())

func buildHeader() ( header string ) {
	header += Title
	header += Created + "\n"
	header += runtime.Compiler + "\n"
	header += sys.GOOS + " - " + sys.GOARCH + "\n"
	
}

func getStatistics() ( stats string ) {
	stats += sprintf("# CGO Calls:    %d \n",runtime.NumCgoCall())
	stats += sprintf("# GO Routines:  %d \n",runtime.NumGoroutine())
	stats += sprintf("# GO Routines:  %d \n",runtime.NumGoroutine())
	stats += sprintf("# GO Routines:  %d \n",runtime.NumGoroutine())

	
}
