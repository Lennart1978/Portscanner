package misc

import (
	"bufio"
	"os"
	"os/user"
	"strings"
	"unicode"
)

func Input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
func GetUserName() string {
	currentUser, err := user.Current()
	if err != nil {
		return "User"
	} else {
		runes := []rune(currentUser.Username)
		runes[0] = unicode.ToUpper(runes[0])
		return string(runes)
	}
}
