package errors

import (
	// "strings"
	"log"
)

type Errs []error

// func (list Errs)error() string{
// 	if len(list) == 0{
// 		return ""
// 	}

// 	output := make([]string,len(list))

// 	for i := range list{
// 		output[i] = list[i].Error()
// 	}

// 	return strings.Join(output)
// }

func IsError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
