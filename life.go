package main

import "fmt"

var WhatIsThe int

func AnswerToLife() int {
	return 42
}

func init() {
	WhatIsThe = 0
}

func main() {
	WhatIsThe = AnswerToLife()
	if WhatIsThe == 42 {
		fmt.Println("It's all a lie.")
	}
}
