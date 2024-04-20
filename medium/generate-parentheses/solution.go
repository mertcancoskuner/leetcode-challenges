func generateParenthesis(n int) []string {
    var result []string
    backtrack(&result, n, 0, 0, "")
    return result
}

func backtrack(result *[]string, n, open int, close int, curSequence string) {
    if len(curSequence) == n * 2 {
        *result = append(*result, curSequence)
        return
    }

    if open < n {
        curSequence += "("
        backtrack(result, n, open + 1, close, curSequence)
        curSequence = curSequence[:len(curSequence)-1]
    }

    if close < open {
        curSequence += ")"
        backtrack(result, n, open, close + 1, curSequence)
        curSequence = curSequence[:len(curSequence)-1]
    }
}
