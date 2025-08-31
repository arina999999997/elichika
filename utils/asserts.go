package utils

import (
	"elichika/log"
)

func CheckErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func MustExist(exist bool) {
	if !exist {
		log.Panic("doesn't exist")
	}
}

func CheckErrMustExist(err error, exist bool) {
	CheckErr(err)
	MustExist(exist)
}
