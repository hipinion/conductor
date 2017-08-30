package conductor

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func GenerateKey(n int, title string) string {

	title = strings.ToLower(title)
	title = strings.Trim(title, " ")

	puncRe := regexp.MustCompile("[,!?'.\"@#$%&]+")
	title = puncRe.ReplaceAllString(title, "")

	re := regexp.MustCompile("[^A-Za-z0-9]+")
	title = re.ReplaceAllString(title, "-")
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b) + `-` + title
}
