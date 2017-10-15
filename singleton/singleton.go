// Author: Peter Nagy <https://peternagy.ie>
// Since: Oct, 2017
// Package: singleton
// Description: holds interface and function implementation
package singleton

import "time"


type singleton struct{
	t time.Time
}

type timer interface{
	GetTime() *time.Time
	ToString() string
}

//GetTime - get instance creation time
func (s *singleton) GetTime() *time.Time{
	return &s.t
}

//ToString - get instance creation time as string
func (s *singleton) ToString() string{
	return s.t.String()
}