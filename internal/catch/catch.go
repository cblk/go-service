package catch

import (
	"fmt"

	logy "github.com/sirupsen/logrus"
)

type CatchFunc func(params ...interface{})

func Finally(err interface{}, fc CatchFunc, params ...interface{}) {
	if err == nil {
		return
	}

	realErr, ok := err.(error)
	if !ok {
		realErr = fmt.Errorf("%+v", err)
	}

	if fc != nil {
		go fc(params)
	}

	logy.Errorf("Finally: %v", realErr)

	return
}
