package main

import (
	reader "./reader"
	"flag"
	functions "./functions"
	)
func main() {
	purpose := flag.String("purpose", "test", "test/redistrict/maxelectors function")
	electCount := flag.Int("count", 538, "total number of electors")
	senators := flag.Int("senators", 2, "how many senators per state")
	flag.Parse() 
	List := reader.Readfile("./ft.csv")

	if *purpose == "test" {
		functions.Test(List, float64(*electCount))
	} else if *purpose == "redistrict" {
		functions.Redistrict(List, float64(*electCount))
	} else if *purpose == "maxelectors" {
		functions.Max(List, *senators)
	}
	
	
}