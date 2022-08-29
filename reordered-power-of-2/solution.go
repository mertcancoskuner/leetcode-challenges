import "math"
func reorderedPowerOf2(n int) bool {
  if n < 1 || n >= int(math.Pow(10, 9)) {
    return false
  }
  ct := count(n)
  for i := 0; i < 31; i++ {
    if ct == count(1<<i) {
      return true
    }
  }
  return false
}

func count(n int) [10]byte {
  var result [10]byte
  
  for n > 0 {
    result[n % 10]++
    n /= 10
  }
  return result
}
