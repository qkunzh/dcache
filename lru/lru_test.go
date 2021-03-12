package lru

import (
	"testing"
)

type user struct {
	Name string
	age  int
}

func (u *user) Len() int {
	return 8 + len(u.Name)
}
func TestCache(t *testing.T) {

}
