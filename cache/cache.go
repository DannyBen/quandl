// Package cache is a sample cache Handler for Quandl
package cache

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
)

// Type Handler implements quandl.Cacher interface
type Handler struct{}

// Get returns content from a dummy cache
func (h Handler) Get(key string) []byte {
	filename := os.TempDir() + "/quandl" + getHash(key)
	result, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}
	return result
}

// Set saves content to a dummy cache
func (h Handler) Set(key string, data []byte) error {
	filename := os.TempDir() + "/quandl" + getHash(key)
	err := ioutil.WriteFile(filename, data, 0644)
	return err
}

// getHash returns checksub of a string
func getHash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
