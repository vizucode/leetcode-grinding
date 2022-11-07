func subtractProductAndSum(n int) int {
    sums := 0
    multiplies := 1
    
    result := strconv.Itoa(n)
    
    for _, value := range result{
        i, _ := strconv.Atoi(string(value))
        sums += i
        multiplies *= i
    }
    
    return multiplies - sums
}