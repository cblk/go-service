package catch

import (
	"strconv"
	"testing"
)

func TestCatch(t *testing.T) {
	defer func() {
		testTx := func(params ...interface{}) {
			t.Log("testTx")
		}

		Finally(recover(), testTx, "")
	}()

	var i int
	var j int
	j = 10
	w := j / i

	t.Log(strconv.Itoa(w))
}
