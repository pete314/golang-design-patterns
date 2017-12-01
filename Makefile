# Author: Peter Nagy <https://peternagy.ie>
# Since: Dec, 2017
# Description: dependacy update 

DEPEND=github.com/fatih/color

gdp: 
	go get -u $(DEPEND)