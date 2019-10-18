package cli

import (
	// "strings"
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