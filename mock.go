package mgo

import (
	"errors"
	// "time"
)

var Mock = false
var MckC = map[string]int{}
var mckc = map[string]int{}
var MckV = map[string]interface{}{}
var MckE = errors.New("Mock Error")

func SetMckC(key string, c int) {
	MckC[key] = c
}

func SetMckV(key string, c int, v interface{}) {
	MckC[key] = c
	MckV[key] = v
}

func chk_mock(key string) bool {
	tc, ok := MckC[key]
	if ok {
		eq := tc == mckc[key]
		mckc[key] += 1
		return eq
	} else {
		return false
	}
}

func ClearMock() {
	MckC = map[string]int{}
	mckc = map[string]int{}
	MckV = map[string]interface{}{}
}
