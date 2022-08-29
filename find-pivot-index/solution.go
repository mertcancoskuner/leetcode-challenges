func pivotIndex(nums []int) int {
  var sum = 0
  for i := 0; i < len(nums); i++ {
    sum += nums[i]
  }
  var leftSum = 0
  for i := 0; i < len(nums); i++ {
    leftSum += nums[i]
    if leftSum - nums[i] == sum - leftSum {
      return i
    }
  }
  return -1
}
