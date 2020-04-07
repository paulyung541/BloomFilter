package bloomfilter

import (
	"errors"
	"github.com/R0L/BloomFilter/bloomfilter/calculate"
	"github.com/R0L/BloomFilter/bloomfilter/storage"
)

const (
	ErrAddCalculateErr   = "add calculates is not empty"
	ErrCheckCalculateErr = "check calculates is not empty"
)

type BloomFilter struct {
	length     uint32
	storage    storage.Storager
	calculates []calculate.Calculater
}

func NewBloomFilter(storage storage.Storager, adapters ...calculate.Calculater) *BloomFilter {
	return &BloomFilter{
		length:     storage.Length(),
		storage:    storage,
		calculates: adapters,
	}
}

func (bf *BloomFilter) Add(bytes []byte) error {
	if len(bf.calculates) == 0 {
		return errors.New(ErrAddCalculateErr)
	}

	for _, bfCalculate := range bf.calculates {
		index, err := bfCalculate.Calculate(bytes, bf.length)
		if nil != err {
			return err
		}
		if _, err = bf.storage.Mark(index); nil != err {
			return err
		}
	}
	return nil
}

func (bf *BloomFilter) Check(bytes []byte) error {
	if len(bf.calculates) == 0 {
		return errors.New(ErrCheckCalculateErr)
	}

	for _, bfCalculate := range bf.calculates {
		index, err := bfCalculate.Calculate(bytes, bf.length)
		if nil != err {
			return err
		}
		flag, err := bf.storage.Find(index)
		if nil != err {
			return err
		}
		if !flag {
			return errors.New("not exists")
		}
	}
	return nil
}
