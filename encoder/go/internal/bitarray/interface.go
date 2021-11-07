package bitarray

type BitArray interface {
	Shl(num uint64)
	Shr(num uint64)
	Or(other bitArray) bitArray
	InsertUInt64At(pos, value uint64)
	String() string
	ToByteArray() []byte
}
