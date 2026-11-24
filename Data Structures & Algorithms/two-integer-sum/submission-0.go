func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ { // 标准写法通常是从 i+1 往后找
            if nums[i]+nums[j] == target {
                return []int{i, j} // 修正：返回切片
            }
        }
    }
    return nil // 修正：返回 nil 或 []int{}
}