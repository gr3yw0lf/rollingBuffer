package rollingBuffer

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var buf *RollingBuf
	itemsToAdd := [...]string{"one", "two", "three", "four", "five", "six"}

	buf = New(5)
	count := 0

	// test initial add of items

	buf.Add(itemsToAdd[count])
	current := buf.Current()
	if current != itemsToAdd[count] {
		t.Error("current fails to return the first added item: ", itemsToAdd[count], "!=", current)
	}
	if itemsToAdd[count] != buf.Get(0) {
		t.Error("first item fails to return the first added item: ", itemsToAdd[count], buf.Get(0))
	}
	count++

	// test full collection of 5 items
	buf.Add(itemsToAdd[count])
	count++
	buf.Add(itemsToAdd[count])
	count++
	buf.Add(itemsToAdd[count])
	count++
	buf.Add(itemsToAdd[count])
	count++
	newBuf := buf.All()
	if len(newBuf[:]) != 5 {
		t.Error("For length of newBuf expected 5 but got", len(newBuf[:]))
	}
	if buf.Len() != 5 {
		t.Error("For length of rollingBuffer, expected 5 but got", buf.Len())
	}
	for i, item := range newBuf[:] {
		if item != itemsToAdd[i] {
			t.Error("For item", i, "of newBuf, expected ", itemsToAdd[i], " but got", item)
		}
	}

	// test 6th addition, should result in 5 items
	buf.Add(itemsToAdd[count])
	count++
	newBuf = buf.All()
	if len(newBuf[:]) != 5 {
		t.Error("For length of newBuf expected 5 but got", len(newBuf[:]))
	}
	if buf.Len() != 5 {
		t.Error("For length of rollingBuffer, expected 5 but got", buf.Len())
	}
	for i, item := range newBuf[:] {
		if item != itemsToAdd[i+1] {
			t.Error("Second Test: For item", i, "of newBuf, expected ", itemsToAdd[i+1], " but got", item)
		}
	}

}
