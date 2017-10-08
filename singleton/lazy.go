// Author: Peter Nagy <https://peternagy.ie>
// Since: Oct, 2017
// Package: singleton
// Description: lazy singleton staring with object
package singleton

import "time"

var (
	lazyInstance *singleton
)

//GetInstanceLazy - get singleton instance initialized on call
func GetInstanceLazy() *singleton{
	if lazyInstance == nil{
		lazyInstance = &singleton{t: time.Now()}
	}

	return lazyInstance
}