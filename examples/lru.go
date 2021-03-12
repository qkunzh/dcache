package main

type user struct {
	Name string
	Age int
}
func (u *user) Len() int{
	return 8 + len(u.Name)
}
/*
func main() {
	cache := lru.New(1024*1024)
	name := "zhang"
	zsl := &user{"zhangshoulei",10}
	for i := 0; i < 2; i++ {
		value,ok := cache.Get(name)
		if ok {
			fmt.Println(*value.(*user))
		} else {
			cache.Add(name,zsl)
		}
	}

}
*/
