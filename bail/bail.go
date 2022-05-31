package bail

import (
	"fmt"
)

func Assert[X any](result X, err error) X {
	if err != nil {
		panic(fmt.Sprintf("Error encountered: %v", err))
	}
	return result
}
