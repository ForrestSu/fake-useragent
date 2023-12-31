package useragent

import (
	"math/rand"
	"sync"
	"time"
)

type useragent struct {
	data map[string][]string
	lock sync.Mutex
}

var (
	UA = useragent{data: make(map[string][]string)}
	r  = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func (u *useragent) Get(key string) []string {
	return u.data[key]
}

func (u *useragent) GetAll() map[string][]string {
	return u.data
}

func (u *useragent) GetRandom(key string) string {
	browser := u.Get(key)
	if len(browser) == 0 {
		return ""
	}
	n := r.Intn(len(browser))
	return browser[n]
}

func (u *useragent) GetAllRandom() string {
	browsers := u.GetAll()
	var datas []string
	for _, uas := range browsers {
		datas = append(datas, uas...)
	}
	if len(datas) == 0 {
		return ""
	}
	n := r.Intn(len(datas))
	return datas[n]
}

func (u *useragent) Set(key, value string) {
	u.lock.Lock()
	defer u.lock.Unlock()
	u.data[key] = append(u.data[key], value)
}

func (u *useragent) SetData(data map[string][]string) {
	u.data = data
}
