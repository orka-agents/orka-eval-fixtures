package cache

import (
	"testing"
	"time"
)

func TestEvictsLeastRecentlyUsed(t *testing.T) {
	c := New(2, time.Minute)
	c.Set("a", 1)
	c.Set("b", 2)
	if _, ok := c.Get("a"); !ok {
		t.Fatal("a is missing")
	}
	c.Set("c", 3)
	if _, ok := c.Get("b"); ok {
		t.Fatal("b should have been evicted")
	}
	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Fatalf("Get(a) = %v, %v", v, ok)
	}
}

func TestEntriesExpire(t *testing.T) {
	now := time.Unix(0, 0)
	c := New(2, time.Second)
	c.now = func() time.Time { return now }
	c.Set("a", 1)
	now = now.Add(time.Second)
	if _, ok := c.Get("a"); ok {
		t.Fatal("a should have expired")
	}
}

func BenchmarkParallelGet(b *testing.B) {
	c := New(128, time.Minute)
	c.Set("hot", 1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get("hot")
		}
	})
}
