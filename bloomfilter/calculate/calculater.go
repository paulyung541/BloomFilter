package calculate

type Calculater interface {
	Calculate(bytes []byte, length uint32) (uint32, error)
}
