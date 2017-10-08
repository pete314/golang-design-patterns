// Author: Peter Nagy <https://peternagy.ie>
// Since: Oct, 2017
// Package: singleton
// Description: use the sync package to lock while init
package singleton

import (
	"time"
	"sync"
)

var (
	threadSafeInstance *singleton
	once sync.Once
)

//GetInstanceThreadSafe - get singleton instance, with thread safe initialization
func GetInstanceThreadSafe() *singleton{
	once.Do(func() {
		if threadSafeInstance == nil{
			threadSafeInstance = &singleton{t: time.Now()}
		}
	})


	return threadSafeInstance
}