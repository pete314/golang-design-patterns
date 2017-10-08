// Author: Peter Nagy <https://peternagy.ie>
// Since: Oct, 2017
// Package: singleton
// Description: eager singleton example with pre initialized instance

package singleton

import "time"

var (
	eagerInstance *singleton
)

func init(){
	//initialize static instance on load
	eagerInstance = &singleton{t: time.Now()}
}

//GetInstanceEager - get singleton instance pre-initialized
func GetInstanceEager() *singleton{
	return eagerInstance
}
