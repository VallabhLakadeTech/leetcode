package firstoccurenceindex28

func StrStr(haystack string, needle string) int {

	needleLen := len(needle)
	haystackLen := len(haystack)
	if needleLen > haystackLen {
		return -1
	}
	for i := 0; i <= haystackLen-needleLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}

	return -1
}
