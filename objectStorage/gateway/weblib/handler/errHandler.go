package handler

import (
	"errors"
	"log"
)

// PanicIfFileError 包装错误
func PanicIfFileError(err error) {
	if err != nil {
		err = errors.New("fileService--" + err.Error())
		log.Println("fileService err: ", err)
		panic(err)
	}
}

// PanicIfMetaError 包装错误
func PanicIfMetaError(err error) {
	if err != nil {
		err = errors.New("metaService--" + err.Error())
		log.Println("metaService err: ", err)
		panic(err)
	}
}
