package util

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

// FileExist 判断文件是否存在
func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// FileRemove 文件删除
func FileRemove(file string)  {
	err := os.Remove(file)
	if err != nil {
		log.Printf("删除失败")
	} else {
		log.Printf("删除成功")
	}
}


// FileFolder  文件所在目录 returns the directory for the current filename.
func FileFolder(name string) string {
	return filepath.Dir(filepath.Base(name))
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString 返回指定值的随机字符串
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
