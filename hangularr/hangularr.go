package hangularr

import "fmt"

func example_array() {
	var (
		start = rune('가')
		end   = rune('힣' + 1)
	)

	fruits := []string{"수박", "멜론", "매실", "딸기", "사과", "참외"}
	var ends rune = 28
	for _, fruit := range fruits {
		word := []rune(fruit)
		word_end_idx := len(word) - 1
		end_chr := word[word_end_idx]
		if start <= end_chr && end_chr < end {
			index := end_chr - start
			chk_ends := index % ends
			if chk_ends != 0 {
				fmt.Printf("%s은 맛있다.\n", fruit)
			} else {
				fmt.Printf("%s가 맛있다.\n", fruit)
			}
		} else {
			fmt.Println("한글이 아닙니다")
		}
	}
}
