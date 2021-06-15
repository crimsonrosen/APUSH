package main

import (
	reader "./reader"
	"flag"
	functions "./functions"
	)

func main() {
	//conditions
	purpose := flag.String("purpose", "test", "test/redistrict/getmax function")
	electCount := flag.Int("count", 538, "total number of electors")
	row4 := flag.Bool("row4", false, "custom senators bool")
	flag.Parse() 

	//read csv data
	List := reader.Readfile("./ft.csv", *row4)


	//what function to run
	if *purpose == "test" {
		functions.Test(List, float64(*electCount))
	} else if *purpose == "redistrict" {
		functions.Redistrict(List, float64(*electCount))
	} else if *purpose == "getmax" {
		functions.Max(List)
	}
	
}