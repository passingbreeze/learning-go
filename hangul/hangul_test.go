package hangul

import "fmt"

func exampleHasConsonantSuffix() {
	fmt.Println(HanConsonantSuffix("Go 언어"))
	fmt.Println(HanConsonantSuffix("그럼"))
	fmt.Println(HanConsonantSuffix("우리 밥 먹고 합시다."))
	// Order:
	// false
	// true
	// false
}
