package main

import "sync"

type Singleton struct {}

var (
	instance *Singleton
	once sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}