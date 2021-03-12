package dcache
import (
	"github.com/qkunzh/dcache/lru"
	"sync"
)

type Cache struct {
	cache *lru.Cache
	cap int
	lock sync.Mutex
}
func New(cap int) *Cache {
	return &Cache{
		cache:lru.New(cap),
		cap:cap,
	}
}
func (this *Cache) Get(key string)(ByteView,bool) {
	this.lock.Lock()
	defer this.lock.Unlock()
	value,ok := this.cache.Get(key)
	if ok {
		bv,suc := value.(ByteView)
		if suc {
			return bv,true
		}
	}
	return ByteView{},false
}
func (this *Cache) Add(key string, bv ByteView) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.cache.Add(key,bv)
}
func (this *Cache) Delete(key string) {
	_, ok := this.cache.Get(key)
	if !ok {
		return
	}
	this.lock.Lock()
	defer this.lock.Unlock()
	this.cache.Delete(key)
}