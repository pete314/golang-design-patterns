// Author: Peter Nagy <https://peternagy.ie>
// Since: Dec, 2017
// Package: flyweight
// Description: flyweight patter implementation

package flyweight

import (
	"sync"

	"github.com/fatih/color"
)

//Color - struct holds color for reuse
type Color struct {
	Name  string
	Color *color.Color
}

var (
	pool map[string]*Color
	once sync.Once
)

func init() {
	once.Do(func() {
		pool = make(map[string]*Color)
	})
}

//NewColor - create a new instance
func NewColor(name string) *Color {
	switch name {
	case "red":
		return &Color{Name: name, Color: color.New(color.FgRed)}
	case "blue":
		return &Color{Name: name, Color: color.New(color.FgBlue)}
	case "green":
		return &Color{Name: name, Color: color.New(color.FgGreen)}
	case "yellow":
		return &Color{Name: name, Color: color.New(color.FgYellow)}
	case "cyan":
		return &Color{Name: name, Color: color.New(color.FgCyan)}
	default:
		return &Color{Name: name, Color: color.New(color.FgWhite)}
	}
}

//ColorFactory - get an instance from the initiaslized pool
func ColorFactory(name string) *Color {
	if v, ok := pool[name]; ok {
		return v
	}

	c := NewColor(name)
	pool[name] = c

	return c
}

//PrintColoredln - print colored line ( to add some functionality)
func (c *Color) PrintColoredln(ln string) {
	c.Color.Println(ln)
}
