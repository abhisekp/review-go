package utils

import (
	"fmt"
	"strings"
)

type DividerOption struct {
	Title string
}

func Divider(options ...DividerOption) {
	option := TakeFirst(DividerOption{}, options)
	if option.Title != "" {
		titleLen := len(option.Title)
		hyphenCount := (50 - titleLen - 2) / 2 // 6 hyphens total (3 on each side)
		if hyphenCount < 0 {
			hyphenCount = 0 // In case the title is too long
		}
		divider := ""
		if hyphenCount == 0 {
			divider = fmt.Sprintf("%s", strings.Repeat("-", 50))
			fmt.Println(divider) // Ensure the length is exactly 50
			fmt.Printf(" %s \n", option.Title)
			divider = fmt.Sprintf("%s", strings.Repeat("-", 50))
			fmt.Println(divider) // Ensure the length is exactly 50
		} else {
			divider = fmt.Sprintf("%s %s %s", strings.Repeat("-", hyphenCount), option.Title, strings.Repeat("-", hyphenCount))
		}
		fmt.Println(divider) // Ensure the length is exactly 50
		return
	}

	fmt.Println(strings.Repeat("-", 50))
}
