package hashids

type Hashids struct {
	options *HashidsConfig
}

func NewHashidsObjectSimple() *Hashids {
	return NewHashidsObject(NewDefaultHashidsConfig())
}

func NewHashidsObject(options *HashidsConfig) *Hashids {
	return &Hashids{
		options: options,
	}
}

//func (this *Hashids) Encode(numbers []int) (string, error) {
//
//}
//
//func (this *Hashids) EncodeHex(hex string) (string, error) {
//
//}
//
//func (this *Hashids) Decode(hash string) ([]int, error) {
//
//}
//
//func (this *Hashids) DecodeHex(hash string) ([]int, error) {
//
//}

//func (this *Hashids) EncodeInt64(numbers []int64) (string, error) {
//	if len(numbers) == 0 {
//		return "", errors.New("can not encoding empty array of numbers")
//	}
//	for _, n := range numbers {
//		if n < 0 {
//			return "", errors.New("negative number not supported")
//		}
//	}
//}

//func hash(input int64, alphabet []rune) []rune {
//}
//
//func unhash(input string, alphabet string) int64 {
//}
