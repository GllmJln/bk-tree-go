package bktree

import (
	"strconv"
	"testing"

	"github.com/gllmjln/bk-tree-go/distance"
)

type bktSample struct {
	matches   int
	values    []string
	input     string
	tolerance int
}

var bktSamples = []bktSample{
	{1, []string{"123456"}, "123456", 0},
	{2, []string{"123456", "1234567"}, "123456", 1},
	{3, []string{"123456", "1234567", "12345"}, "123456", 1},
	{4, []string{"123456", "1234567", "12345", "123465"}, "123456", 1},
	{3, []string{"123456", "1234567", "12345", "654789"}, "123456", 1},
	{3, []string{"123456", "1234567", "12345", "654789"}, "123456", 1},
	{3, []string{"123456", "12345678", "12345", "654789"}, "123456", 2},
	{4, []string{"123456", "12345678", "12345", "654789", "1255"}, "123456", 3},
	{5, []string{"123456", "12345678", "12345", "654789", "1255"}, "123456", 10},
}

func TestAddItem(t *testing.T) {
	for _, v := range bktSamples {
		bkt := NewBKTree(distance.CalculateDistance)

		for _, w := range v.values {
			bkt.AddWord(w)
		}

		matches := bkt.Search(v.input, v.tolerance)
		if len(matches) != v.matches {
			t.Errorf("Error for %v and %v, expected %v but got %v. Matches were %v", v.input, v.values, v.matches, len(matches), matches)
		}
	}
}

func BenchmarkSearching(b *testing.B) {
	bkt := NewBKTree(distance.CalculateDistance)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bkt.AddWord(strconv.Itoa(i))
	}
	bkt.Search("foo", 2)
}
