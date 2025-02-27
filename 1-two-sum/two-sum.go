func twoSum(nums []int, target int) []int {
    for i:= 0 ; i< len(nums) ; i++ {
        for j:= i+1 ; j<len(nums) ; j++ {
            if (nums[i] + nums[j] == target){
                return []int{i,j}
            }
        }
    }
    return nil

    // numToIndexMap := make(map[int]int)

    // for i, num := range nums {
    //     diff := target - num
    //     if idx, found := numToIndexMap[diff]; found {
    //         return []int{i, idx}
    //     }
    //     numToIndexMap[num] = i
    // }

    // return nil
}