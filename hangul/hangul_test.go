package hangul

import (
	"fmt"
)

func ExampleHasConsonanSuffix() {
	fmt.Println(HasConsonanSuffix("고 언어"))
	fmt.Println(HasConsonanSuffix("그럼"))
	fmt.Println(HasConsonanSuffix("우리 밥 먹고 합시다."))
	// Output:
	// false
	//	true
	// false
}
