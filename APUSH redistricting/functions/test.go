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
