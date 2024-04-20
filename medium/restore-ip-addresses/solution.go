func restoreIpAddresses(s string) []string {
  var result []string
  if len(s) > 0 && len(s) < 21 {
    generateIP(&result, s, 0, 0, "")
  }
  return result  
}

func generateIP(result *[]string, s string, start int, step int, ip string) {
    if start == len(s) && step == 4 {
        *result = append(*result, ip[:len(ip)-1])
        return
    }
    if start == len(s) || step == 4 {
		return
    }
    for i := start; i < start + 3 && i < len(s); i++ {
		part := s[start:i+1]
		if isValid(part) {
			generateIP(result, s, i + 1, step + 1, ip + part + ".")
		}
	}
}

func isValid(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	num, _ := strconv.Atoi(s)
	return num <= 255
}
