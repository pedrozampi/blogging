package main

import (
	"sync"
	"time"
)

type autoInc struct {
	sync.Mutex
	id int
}

func (a *autoInc) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++
	return
}

var ai autoInc

type BlogPost struct {
	ID        int       `json:"ID"`
	Content   string    `json:"Content"`
	Category  string    `json:"Category"`
	Tags      []string  `json:"Tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var PostList = []BlogPost{}
