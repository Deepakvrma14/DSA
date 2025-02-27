func twoSum(nums []int, target int) []int {
    // for i:= 0 ; i< len(nums) ; i++ {
    //     for j:= i+1 ; j<len(nums) ; j++ {
    //         if (nums[i] + nums[j] == target){
    //             return []int{i,j}
    //         }
    //     }
    // }
    // return nil

   indexMap := make(map[int]int)

   for i:=0;i<len(nums);i++{
        
        diff:= target-nums[i]

        if idx,ok :=indexMap[(diff)]; ok{
            return []int{i,idx}
        }
        indexMap[(nums[i])] = i;
   }
   return nil;
}