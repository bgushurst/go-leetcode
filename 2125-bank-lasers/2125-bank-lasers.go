package leetcode

// Runtime: 19ms
// Memory Usage: 6.9 MB

func numberOfBeams(bank []string) int {
	
	result := 0
	prc := 0
	rc := 0

	for _, r := range bank {
		for _, c := range r {
			if c == '1' {
				rc++
			}
		}
		if prc == 0 && rc > 0 {
			prc = rc
			rc = 0
			continue
		}

		if rc > 0 {
			result += prc * rc
			prc = rc
			rc = 0
		}
	}

	return result
}