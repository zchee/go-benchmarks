// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/dgraph-io/badger"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	testKey       = []byte("testKey")
	testValue     = []byte("testValue")
	testValueSize = int64(len(testValue))
)

func BenchmarkGetBadger(b *testing.B) {
	tmp, err := ioutil.TempDir(os.TempDir(), "badger")
	if err != nil {
		b.Fatal(err)
	}

	opt := badger.DefaultOptions
	opt.Dir = tmp
	opt.ValueDir = tmp
	db, err := badger.NewKV(&opt)
	if err != nil {
		b.Fatal(err)
	}

	if err := db.Set(testKey, testValue); err != nil {
		b.Fatal(err)
	}
	item := new(badger.KVItem)

	b.SetBytes(testValueSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		db.Get(testKey, item)
	}

	b.StopTimer()
	db.Close()
	os.RemoveAll(tmp)
}

func BenchmarkGetLevelDB(b *testing.B) {
	tmp, err := ioutil.TempDir(os.TempDir(), "leveldb")
	if err != nil {
		b.Fatal(err)
	}

	db, err := leveldb.OpenFile(tmp, nil)
	if err != nil {
		b.Fatal(err)
	}

	if err := db.Put(testKey, testValue, nil); err != nil {
		b.Fatal(err)
	}

	b.SetBytes(testValueSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		db.Get(testKey, nil)
	}

	b.StopTimer()
	db.Close()
	os.RemoveAll(tmp)
}

func BenchmarkGetBolt(b *testing.B) {
	tmp, err := ioutil.TempDir(os.TempDir(), "bolt")
	if err != nil {
		b.Fatal(err)
	}

	db, err := bolt.Open(filepath.Join(tmp, "test.db"), 0600, nil)
	if err != nil {
		b.Fatal(err)
	}

	bucketName := []byte("testBucket")
	updateFn := func(tx *bolt.Tx) error {
		bc, err := tx.CreateBucket(bucketName)
		if err != nil {
			b.Fatal(err)
		}
		bc.Put(testKey, testValue)
		return nil
	}
	if err := db.Update(updateFn); err != nil {
		b.Fatal(err)
	}

	b.SetBytes(testValueSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		db.View(func(tx *bolt.Tx) error {
			bc := tx.Bucket(bucketName)
			bc.Get(testKey)
			return nil
		})
	}

	b.StopTimer()
	db.Close()
	os.RemoveAll(tmp)
}
