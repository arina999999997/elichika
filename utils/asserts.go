package utils

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func MustExist(exist bool) {
	if !exist {
		panic("doesn't exist")
	}
}

func CheckErrMustExist(err error, exist bool) {
	CheckErr(err)
	MustExist(exist)
}
