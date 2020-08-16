package store

import (
	"github.com/peterbourgon/diskv/v3"
)

var dkv *diskv.Diskv

// ReadyDiskv initializes the diskv object for access.
func ReadyDiskv() {

	transform := func(s string) []string { return []string{} }

	dkv = diskv.New(diskv.Options{
		BasePath:     "/usr/local/hselect",
		Transform:    transform,
		CacheSizeMax: 10 * 1024 * 1024, // 10 MB
	})
}

// ReadDKV reads content from the diskv file marked with key 'key'.
func ReadDKV(key string) []byte {

	key += "_hs.log"
	value, err := dkv.Read(key)
	if err != nil {
		// ignore
		// fmt.Fprint(os.Stderr, "unable to read from diskv "+err.Error()+"\n")
		return nil
	}
	return value
}

// WriteDKV appends value to the diskv file marked with key 'key'.
func WriteDKV(key string, value []byte) {

	key += "_hs.log"
	err := dkv.Write(key, value)
	if err != nil {
		// ignore
		// fmt.Fprint(os.Stderr, "unable to write to diskv "+err.Error()+"\n")
	}
}
