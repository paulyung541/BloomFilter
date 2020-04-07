package calculate

import (
	"github.com/pkg/errors"
	"hash/fnv"
)

const (
	ErrFnvHashWriteErr = "calculate hash err: hash write err"
)

type fnvCalculate struct {
}

func NewFnvCalculate() *fnvCalculate {
	return &fnvCalculate{}
}

func (fnvC *fnvCalculate) Calculate(bytes []byte, length uint32) (uint32, error) {
	fHash := fnv.New32()
	if _, err := fHash.Write(bytes); nil != err {
		return 0, errors.Wrap(err, ErrFnvHashWriteErr)
	}

	return fHash.Sum32() % length, nil
}
