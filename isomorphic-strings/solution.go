import (
  "unicode"
  "math"
)

func isIsomorphic(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }
  if len(s) < 1 || len(s) > int(5 * math.Pow(10, 4)) {
    return false
  }
  if (!isASCII(s)) {
    return false
  }
  if (!isASCII(t)) {
    return false
  }
  sMap, tMap := map[uint8] int{}, map[uint8] int{}
  for i := range s {
    if sMap[s[i]] != tMap[t[i]] {
      return false
    } else {
      sMap[s[i]] = i + 1
      tMap[t[i]] = i + 1
    }
  }
  return true
}

func isASCII (s string) bool {
  for i := 0; i < len(s); i++ {
    if s[i] > unicode.MaxASCII {
      return false
    }
  }
  return true
}
