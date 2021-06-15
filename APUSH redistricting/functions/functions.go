package functions

import (
    reader "../reader"
    "fmt"
    "math"
)

//perform a chisquare gof test on the current distribution
func Test (list []reader.State, count float64) {
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


//print out population balenced redistrictings
func Redistrict (list []reader.State, count float64) {
    expvalues := make([]float64,len(list))
    obsvalues := make([]float64,len(list))

    var totalReps float64
    var totalPop float64
    for i := 0; i < len(list); i++ {
        totalReps += list[i].Reps
        totalPop += list[i].Pop
    } 
    fmt.Println(totalReps)
    fmt.Println(totalPop)

    for i := 0; i < len(list); i++ {
        expvalues[i] = (list[i].Pop/totalPop)*count
        obsvalues[i] = (list[i].Reps/totalReps)*count
    }
     for i, value := range expvalues {
         fmt.Printf("%v : %v Representatives\n", list[i].StateName, value)
     }
}

//find the max possible electors and representatives per state
func Max (list []reader.State) {
    var totalElectors int
    var totalReps int
    for _, state := range list {
        reps := (int(state.Pop)-(int(state.Pop) % 30000))/(30000*19)
        electors :=(reps + int(state.Senators))
        totalReps += reps
        totalElectors += electors
        fmt.Printf("%v : %v Electors, %v Senators, and %v Representatives\n", state.StateName, electors, state.Senators, reps)
    }
    fmt.Println("max reps:", totalReps)
    fmt.Println("max electors:", totalElectors)
}