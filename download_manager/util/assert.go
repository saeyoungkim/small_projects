package util

import (
	"errors"
	"log"
)

func AssertNoError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func AssertAcceptableError(err error, acceptableErrors ...error) {
	for _, ae := range acceptableErrors {
		if errors.Is(err, ae) {
			return
		}
	}
	log.Fatal(err.Error())
}
