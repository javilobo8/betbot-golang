package util

import (
	"betbot/constants"
	"strings"
)

func GetCommandMessages(message string) []string {
	msg := strings.Replace(message, string(constants.CommandChar), "", 1)
	msg = strings.Trim(msg, " ")
	return strings.Split(msg, " ")
}
