# BSON Serialization fail with uint64

```go
package main

import (
	"math"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestMarshal_Bson_UInt64(t *testing.T) {
	type testObj struct {
		Value uint64 `bson:"value"`
	}
	obj := testObj{
		Value: math.MaxInt64 + 1,
	}
	_, err := bson.Marshal(obj)
	if err != nil {
		t.Errorf("bson marshal failed: %v", err)
	}
}
```

```bash
$ go test ./...
--- FAIL: TestMarshal_Bson_UInt54 (0.00s)
    main_test.go:19: bson marshal failed: 9223372036854775808 overflows int64
FAIL
FAIL	bson-uint64	0.001s
FAIL
```
