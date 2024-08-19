package main

import (
	"fmt"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"time"

	"github.com/go-daq/crc8"
	"github.com/howeyc/crc16"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

func MeassureTime(f func()) time.Duration {
	begin := time.Now()
	f()
	return time.Since(begin)
}

type HashMap struct {
	m    map[string]interface{}
	hash hash.Hash
}

func (hm HashMap) Set(key string, value interface{}) {
	hashKey := hm.getKey(key)
	hm.m[hashKey] = value
}

func (hm HashMap) Get(key string) (interface{}, bool) {
	hashKey := hm.getKey(key)
	value, ok := hm.m[hashKey]
	return value, ok
}

func (hm HashMap) getKey(key string) string {
	hm.hash.Reset()
	hm.hash.Write([]byte(key))
	hashKey := hm.hash.Sum(nil)
	return string(hashKey)
}

func NewHashMap(size int, crc HashCRC) *HashMap {
	hashMap := &HashMap{m: make(map[string]interface{}, size)}
	crc(hashMap)
	return hashMap
}

func WithHashCRC32() HashCRC {
	return func(hm *HashMap) {
		hm.hash = crc32.New(crc32.IEEETable)
	}
}

func WithHashCRC64() HashCRC {
	return func(hm *HashMap) {
		hm.hash = crc64.New(crc64.MakeTable(crc64.ISO))
	}
}

func WithHashCRC16() HashCRC {
	return func(hm *HashMap) {
		hm.hash = crc16.NewCCITT()
	}
}

func WithHashCRC8() HashCRC {
	return func(hm *HashMap) {
		hm.hash = crc8.New(crc8.MakeTable(0x07))
	}
}

type HashCRC func(*HashMap)

func main() {
	m := NewHashMap(16, WithHashCRC64())
	since := MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC32())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC16())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC8())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)
}
