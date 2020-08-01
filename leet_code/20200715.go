package main

func lengthOf(s string) int {
	tmpMap := make(map[string]bool)
	i := 0
	j := i + 1
	res := 0
	tmpMap[string(s[i])] = true
	for j < len(s) {
		if _, ok := tmpMap[string(s[j])]; ok {
			i++
		} else {
			tmpMap[string(s[j])] = true
			l := j - i + 1
			if l > res {
				res = l
			}
		}
		j++
	}
	return res
}

func main() {
	lengthOf("abcabcbb")
}
