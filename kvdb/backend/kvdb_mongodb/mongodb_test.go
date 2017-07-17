package kvdb_mongo

import (
	"math/rand"
	"strconv"
	"testing"

	"io"
)

func TestMongoKVDB_Set(t *testing.T) {
	kvdb, err := OpenMongoKVDB("mongodb://127.0.0.1:27017/goworld", "goworld", "__kv__")
	if err != nil {
		t.Fatal(err)
	}
	val, err := kvdb.Get("__key_not_exists__")
	if err != nil || val != "" {
		t.Fatal(err)
	}

	for i := 0; i < 10000; i++ {
		key := strconv.Itoa(rand.Intn(10000))
		val := strconv.Itoa(rand.Intn(10000))
		err = kvdb.Put(key, val)
		if err != nil {
			t.Fatal(err)
		}
		var verifyVal string
		verifyVal, err = kvdb.Get(key)
		if err != nil {
			t.Fatal(err)
		}

		if verifyVal != val {
			t.Errorf("%s != %s", val, verifyVal)
		}
	}

}

func TestMongoKVDB_Find(t *testing.T) {
	kvdb, err := OpenMongoKVDB("mongodb://127.0.0.1:27017/goworld", "goworld", "__kv__")
	if err != nil {
		t.Fatal(err)
	}

	beginKey := "1000"
	endKey := "9999"
	it := kvdb.Find(beginKey, endKey)
	oldKey := ""
	for {
		item, err := it.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Error(err)
			break
		}

		if item.Key <= oldKey { // the keys should be increasing
			t.Errorf("old key is %s, new key is %s, should be increasing", oldKey, item.Key)
		}

		println(item.Key, item.Val)
		oldKey = item.Key
	}
}