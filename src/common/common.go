package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func PathExists(path string) bool {
	fmt.Println(path)
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString returns a random string with a fixed length
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func ToJson(data interface{}) ([]byte, error) {
	if data == nil {
		return nil, errors.New("错误,data 是空")
	}
	switch data.(type) {
	case string:
		return []byte(data.(string)), nil
	}
	if r, err := json.Marshal(data); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func IsEmpty(value ...interface{}) bool {
	for _, val := range value {
		if val == "" || val == nil {
			return true
		}
	}
	return false
}

// "1path","/2path","3path/","/4path/"
func PathConcat(paths ...string) string {
	var p bytes.Buffer
	for _, path := range paths {
		if !(strings.Index(path, "/") == 0) {
			p.WriteString("/")
		}
		var idxLast = strings.LastIndex(path, "/")
		if idxLast == len(path)-1 {
			p.WriteString(path[:idxLast])
		} else {
			p.WriteString(path)
		}
	}
	return p.String()
}
