package errors

import (
	"Validators/constants"
	"fmt"
)

var defaultMsgs = map[string]string{
	constants.Req: "is required.",
}

func SetDefaultMsg(field, rule string) string {

	if defaultMsg, ok := defaultMsgs[rule]; ok {

		return fmt.Sprintf("%s %s", field, defaultMsg)
	}

	return ""
}
