package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/famz/SetLocale"
)

const DICT string = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`

var break6bit = regexp.MustCompilePOSIX(`.{1,6}`)
var break8bit = regexp.MustCompilePOSIX(`.{1,8}`)

func parseByte(str string) int64 {
	parsed, _ := strconv.ParseInt(str, 2, 8)
	return parsed
}

func main() {
	SetLocale.SetLocale(SetLocale.LC_ALL, "C")

	payload := os.Args[1]
	block := ""
	encoded := ""

	// for _, char := range []byte(payload) {
	// 	block += fmt.Sprintf("%08s", strconv.FormatInt(int64(char), 2))
	// 	if len(block) < 24 { continue }
	// 	fmt.Printf("%s", block)
	// 	block = ""
	// 	fmt.Println("")
	// }

	for i := 0; i < len(payload); i++ {
		block += fmt.Sprintf("%08s", strconv.FormatInt(int64(payload[i]), 2))
		if len(block) < 24 { continue }

		match := break6bit.FindAllString(block, -1)
		encoded += fmt.Sprintf("%c%c%c%c",
			DICT[parseByte(match[0])],
			DICT[parseByte(match[1])],
			DICT[parseByte(match[2])],
			DICT[parseByte(match[3])],
		)
		block = ""
	}

	if len(block) > 0 {
		drain := ""
		match := break6bit.FindAllString(block, -1)
    if len(match) > 0 { drain += fmt.Sprintf("%c", DICT[parseByte(match[0])]) }
    if len(match) > 1 { drain += fmt.Sprintf("%c", DICT[parseByte(match[1])]) }
    if len(match) > 2 { drain += fmt.Sprintf("%c", DICT[parseByte(match[2])]) }
    if len(match) > 3 { drain += fmt.Sprintf("%c", DICT[parseByte(match[3])]) }
		encoded += strings.ReplaceAll(fmt.Sprintf("%-4s", drain), " ", "=")
	}

	fmt.Printf("%s\n", encoded)
}
