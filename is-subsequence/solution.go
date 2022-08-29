import "unicode"
func isSubsequence(s string, t string) bool {
  if (!isLower(s)) || (!isLower(t)) {
    return false
  }
  
  var i, j = 0, 0
  for i < len(s) && j < len(t) {
    if s[i] == t[j] {
      i += 1
    }
    j += 1
  }
  return (i == len(s))
}

func isLower(s string) bool {
  for _, i := range s {
    if !unicode.IsLower(i) && !unicode.IsLetter(i) {
      return false
    }
  }
  return true
}
