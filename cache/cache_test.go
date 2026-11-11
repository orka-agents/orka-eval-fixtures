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

func TestPurge(t *testing.T) {
	now := time.Unix(0, 0)
	c := New(3, time.Second)
	c.now = func() time.Time { return now }
	c.Set("a", 1)
	c.Set("b", 2)
	now = now.Add(time.Second)
	c.Set("c", 3)
	if removed := c.Purge(); removed != 2 {
		t.Fatalf("Purge() = %d, want 2", removed)
	}
	if _, ok := c.Get("c"); !ok {
		t.Fatal("c is missing")
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
