package _interface

type IInputStream interface {
	Begin(packet []byte) (data []byte)
	End(data []byte)
}
