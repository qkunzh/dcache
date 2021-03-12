package main

/*
func main() {
	cache := dcache.New(1024*1024)
	name := "zhang"
	value := "zhangshoulei"
	_,ok := cache.Get(name)
	fmt.Println(ok)
	if !ok {
		bytes := []byte(value)
		bv := dcache.ByteView{bytes}
		cache.Add(name,bv)
	}
	bv,ok := cache.Get(name)
	fmt.Println(ok)
	bytes := bv.Get()
	fmt.Println(string(bytes))
	cache.Delete(name)
	_,ok = cache.Get(name)
	fmt.Println(ok)

}
 */