func lengthOfLongestSubstring(s string) int {
    window := make(map[uint8]int)
    var left, right, result int

    for right < len(s) {
        var r = s[right]
        window[r] += 1
        for window[r] > 1 {
            l := s[left]
            window[l] -= 1
            left += 1
        }
        result = max(result, right-left+1)
        right += 1
    }
    
    return result
}
