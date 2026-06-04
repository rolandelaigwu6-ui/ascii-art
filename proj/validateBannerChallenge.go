package main 

// import (
// 	"fmt"
// )

// func ValidateBannerChallenge(banner map[rune][]string)error {
// 	const (
// 		firstprintable = 32
// 		lastprintable = 126
// 		glyphHeigth = 8

// 	)

// 	if banner == nil {
// 		return fmt.Errorf("invalid banner: banner is nil")
// 	}

// 	if len(banner) != int(lastprintable-firstprintable+1) {
// 		return fmt.Errorf("invalid banner: character out-of ASCII Printable range expected 95 unit got %d", len(banner))
// 	}

//}