package leetcode

import (
	"strconv"
)

func groupAnagrams(strs []string) [][]string {

	// Create map of anagrams
	hm := map[string][]string{}

	// Iterate through strs to generate a key and store in map
	for _, str := range strs {
		key := makeKey(str)
		println(str, key)
		hm[key] = append(hm[key], str)
	}

  result := [][]string{}

	// Iterate through map and append to result
	for _, v := range hm {
		result = append(result, v)
	}

	return result
}

func makeKey(str string) string {
	key := make([]int, 26)

	for _, char := range str {
		keyIndex := int(char) - int('a')
		println(keyIndex)
		key[keyIndex]++
	}

	println(key)

	result := ""

	for _, v := range key {
		result += strconv.Itoa(v)
	}

	return result

}