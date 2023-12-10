package main

func canConstruct(ransomNote string, magazine string) bool {
	tableA := make(map[byte]int)

	for i := range ransomNote {
		tableA[ransomNote[i]]++
	}

	for i := range magazine {
		if _, ok := tableA[magazine[i]]; ok {
			tableA[magazine[i]]--
		}
	}

	for i := range tableA {
		if tableA[i] > 0 {
			return false
		}
	}

	return true
}
