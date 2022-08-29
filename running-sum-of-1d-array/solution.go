import "math"
func runningSum(nums []int) []int {
  if len(nums) <= 1 || len(nums) >= 1000 {
    return nil
  }
  sum := 0
  runningSum := make([]int, len(nums))
  for i := 0; i < len(nums); i++ {
    if nums[i] <= int(-1*(math.Pow(10, 6))) || nums[i] >= int(math.Pow(10, 6)) {
      return nil
    }
    sum += nums[i]
    runningSum[i] = sum
  }
  return runningSum
}
