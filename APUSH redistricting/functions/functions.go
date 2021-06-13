package functions

import (
    reader "../reader"
    "fmt"
    "math"
)

func Test (list []reader.ChiSquare, count float64) {
    var chiSquaredStat float64
    expvalues := make([]float64,len(list))
    obsvalues := make([]float64,len(list))

    var totalElectors float64
    var totalPop float64
    for i := 0; i < len(list); i++ {
        totalElectors += list[i].Electors
        totalPop += list[i].Pop
    } 
    fmt.Println(totalElectors)
    fmt.Println(totalPop)

    for i := 0; i < len(list); i++ {
        expvalues[i] = (list[i].Pop/totalPop)*count
        obsvalues[i] = (list[i].Electors/totalElectors)*count
    }
	for i := 0; i < len(list); i++ {
		o := obsvalues[i]
		e := expvalues[i]
		chiSquaredStat = chiSquaredStat + ((math.Pow((o-e),2))/e)
	}
	fmt.Println(chiSquaredStat)
}

func Redistrict (list []reader.ChiSquare, count float64) {
    expvalues := make([]float64,len(list))
    obsvalues := make([]float64,len(list))

    var totalElectors float64
    var totalPop float64
    for i := 0; i < len(list); i++ {
        totalElectors += list[i].Electors
        totalPop += list[i].Pop
    } 
    fmt.Println(totalElectors)
    fmt.Println(totalPop)

    for i := 0; i < len(list); i++ {
        expvalues[i] = (list[i].Pop/totalPop)*count
        obsvalues[i] = (list[i].Electors/totalElectors)*count
    }
     for i, value := range expvalues {
         fmt.Printf("%v : %v Electors\n", list[i].StateName, value)
     }
}

func Max (list []reader.ChiSquare, senatorCount int) {
    var totalElectors int
    var totalReps int
    for _, state := range list {
        reps := (int(state.Pop)-(int(state.Pop) % 30000))/30000
        electors :=(reps + senatorCount)
        totalReps += reps
        totalElectors += electors
        fmt.Printf("%v : %v Electors and %v Representatives\n", state.StateName, electors, reps)
    }
    fmt.Println("max reps:", totalReps)
    fmt.Println("max electors:", totalElectors)
}