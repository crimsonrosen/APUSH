package reader

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
)
type ChiSquare struct {
	StateName string
	Electors float64
	Pop float64
}

func Readfile (filename string) []ChiSquare {
	//open the file
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil{fmt.Println("owie1!")}
	
	//slurp all contents
	r := csv.NewReader(f)
	unform, err := r.ReadAll()
	if err != nil {fmt.Println("owie2!")}
	
	//trimming everything
	states := unform[0]
	formElect := unform[2]
	formPop := unform[1]

	//slap into a []Chisquare format
	stats := make([]ChiSquare, len(formElect))
	for i := 0; i < len(formElect); i++ {
			finalElect, _ := strconv.ParseFloat(formElect[i], 64)
			finalPop, _ := strconv.ParseFloat(formPop[i], 64)
		stats[i] = ChiSquare{
			StateName : states[i],
			Electors : finalElect,
			Pop : finalPop,
		}

	}
	return stats
}