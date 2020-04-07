package storage

type Storager interface {
	Length() uint32
	Mark(index uint32) (bool, error)
	Find(index uint32) (bool, error)
}
