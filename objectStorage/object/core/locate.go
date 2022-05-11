package core

import (
	"path/filepath"
)

var objects = make(map[string]int)

// GetObjectInfo 获取本地文件信息做缓存；
func GetObjectInfo() {
	matches, _ := filepath.Glob("./testData/data/*")
	for i := range matches {
		hash := filepath.Base(matches[i])
		objects[hash] = 1
	}
}

func Locate(hash string) bool {
	_, ok := objects[hash]
	return ok
}

func Add(hash string) {
	objects[hash] = 1
}

func Del(hash string) {
	delete(objects, hash)
}
