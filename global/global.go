package global

import "log"

func HandleErr(err error) bool {
	if err != nil {
		log.Fatalln(err)
		return true
	}
	return false
}
