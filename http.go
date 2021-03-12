package dcache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)
type CacheServer struct {
	cache *Cache
	ip string
	port int
	cap int
}
func NewCacheServer(ip string, port,cap int) *CacheServer{
	return &CacheServer{
		New(cap),
		ip,
		port,
		cap,
	}
}
func (this *CacheServer) Start() {
	address := fmt.Sprintf("%s:%d",this.ip,this.port)
	err := http.ListenAndServe(address,this)
	log.Fatal(err)
}
func (this *CacheServer) Stop() {

}

func(this *CacheServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.URL.Path[1:]
	params := r.URL.RawQuery
	switch method {
	case "delete": {
		key := strings.Split(params,"=")[1]
		this.cache.Delete(key)
		_,err := w.Write([]byte(fmt.Sprintf("key:%s has been deleted",key)))
		if err != nil {
			log.Fatal("internet error")
		}
	}
	case "get": {
		key := strings.Split(params,"=")[1]
		bv,ok := this.cache.Get(key)
		if ok {
			_, err := w.Write(bv.Get())
			if err != nil {
				log.Fatal("internet error")
			}
		} else {
			_,err := w.Write([]byte(fmt.Sprintf("key:%s not cached\n",key)))
			if err != nil {
				log.Fatal("internet error")
			}
		}
	}
	case "add": {
		items := strings.Split(params,"&")
		key,value := items[0],items[1]
		bv := ByteView{[]byte(strings.Split(value,"=")[1])}
		this.cache.Add(strings.Split(key,"=")[1],bv)
		_,err := w.Write([]byte("success"))
		if err != nil {
			log.Fatal(err)
		}
	}
	default:
		log.Fatal("method error")
	}
}