package mpccode

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/errors/gerror"
)

func Test_err_detail(t *testing.T) {
	e := CodeInternalError(map[string]string{"err": "testerr"})
	code := gerror.Code(e)
	d := code.Detail()
	fmt.Println(d)
	////
	er := CodeUnknown(d)
	er.Error()
}
