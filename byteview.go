package dcache

type ByteView struct{
	Data []byte
}
func (this ByteView) Len() int{
	return len(this.Data)
}
func (this ByteView) Get() []byte{
	copyData := make([]byte,len(this.Data))
	copy(copyData,this.Data)
	return copyData
}

