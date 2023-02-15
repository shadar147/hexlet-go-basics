package ex19

func MostPopularWord(words []string) string {
	wordCounter := make(map[string]int)

	for _, word := range words {
		wordCounter[word]++
	}

	var mostPopularWord string
	maxCnt := 0
	for _, word := range words {
		if wordCounter[word] > maxCnt {
			maxCnt = wordCounter[word]
			mostPopularWord = word
		}
	}

	return mostPopularWord
}
