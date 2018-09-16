// Package core hold models and interfaces of the whole application
package core

//Pagination represent paging request
type Pagination struct {
	Offset int64
	Limit  int64
}
