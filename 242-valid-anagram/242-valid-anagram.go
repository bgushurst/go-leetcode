package leetcode

func isAnagram(s string, t string) bool {
	// Compare the length of two strings
	if len(s) != len(t) {
		return false
	}
	// Create a map to store the characters of string s
	m := make(map[rune]int, len(s))

	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}

	// Loop through the map to check if the value is 0
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
  
	return true
}