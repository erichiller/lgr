package lgr




type lgrI interface {
	GetFilters() []Filters

}

type lgr struct {
	Prefix []interface{}
	BlockFilters Filters
	HighlightFilters Filters


}

type Filters []Filter

type Filter struct {
	Keywords []string
}


func (log *lgr) Filter () {

}




/**


3 log functions

Log -> internal
Printf -> Console , meant for user, web etc,
(Error? -> User, non crashing)
Critical -> System Crash

**/
