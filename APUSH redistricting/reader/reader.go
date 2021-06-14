package reader

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
)

//wooooo data structureeeeeeeeeeeeeeeeeeeeeeeeeeee
type State struct {
	StateName string
	Electors float64
	Pop float64
	Senators float64
}

func Readfile (filename string, row4 bool) []State {
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
	formSenate := make([]float64, len(unform[2]))
	if !row4{
		for i, _ := range formSenate {
			formSenate[i] = 2.0
		}
	} else {
		for i, _ := range formSenate {
			formtempsenate, _ := strconv.ParseFloat(unform[3][i], 64)
			formSenate[i] = formtempsenate
		} 
	}
	//slap into a []State format
	stats := make([]State, len(formElect))

		for i := 0; i < len(formElect); i++ {
			finalElect, _ := strconv.ParseFloat(formElect[i], 64)
			finalPop, _ := strconv.ParseFloat(formPop[i], 64)
			stats[i] = State{
				StateName : states[i],
				Electors : finalElect,
				Pop : finalPop,
				Senators: formSenate[i],
			}

		}
	//fmt.Println(stats)
	return stats
	
}