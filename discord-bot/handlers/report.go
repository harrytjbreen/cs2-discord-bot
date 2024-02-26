package handlers

import "strings"

func Report(words []string) string {
	if len(words) < 2 {
		return "You need to specify the user to report"
	}

	//TODO get Steam User by ID

	//TODO check if user exists

	//TODO check is user is already reported

	//TODO add to DynamoDB

	return "Reported " + strings.Join(words[1], " ")
}
