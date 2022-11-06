func twoSum(nums []int, target int) []int {
    list := []int{}
    
    out:
    for index,value := range nums{
        for index1, value1 := range nums {
            if index != index1 {
                sum := value + value1
                if sum == target {
                    list = append(list, index, index1)
                    break out
                }
            }
        }
    }
    
    return list
}