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
