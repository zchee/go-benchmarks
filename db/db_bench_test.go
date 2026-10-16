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
	tmp := b.TempDir()

	db, err := badger.Open(badger.DefaultOptions(tmp).WithLogger(nil))
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	if err := db.Update(func(txn *badger.Txn) error {
		return txn.Set(testKey, testValue)
	}); err != nil {
		b.Fatal(err)
	}

	b.SetBytes(testValueSize)

	for b.Loop() {
		db.View(func(txn *badger.Txn) error {
			_, err := txn.Get(testKey)
			return err
		})
	}
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

	for b.Loop() {
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

	db, err := bolt.Open(filepath.Join(tmp, "test.db"), 0o600, nil)
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

	for b.Loop() {
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
