package handlers

import "log"

func HandleError(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}
