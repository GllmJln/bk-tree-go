package distance

import "testing"

type sample struct {
	distance int
	a, b     string
}

var samples = []sample{
	{0, "", ""},
	{6, "123456", ""},
	{6, "", "654321"},
	{1, "foobar", "fobobar"},
	{4, "fo", "foobar"},
	{2, "foobar", "foob"},
	{2, "obar", "foobar"},
	{4, "dsfal", "adkl"},
	{3, "close", "clothes"},
	{5, "thsthist", "tsh"},
	{3, "tihsissth", "thisisths"},
	{1, "flahs", "flash"},
	{1, "lash", "flash"},
	{2, "las", "flash"},
	{2, "flash", "lahs"},
}

func TestCalculateDistance(t *testing.T) {
	for _, v := range samples {
		calc := CalculateDistance(v.a, v.b)
		if calc != v.distance {
			t.Errorf("Error for %v and %v, expected %v but got %v", v.a, v.b, v.distance, calc)
		}
	}
}

func BenchmarkCalculateDistanceLongString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CalculateDistance("this is quite a long string", "this is also quite a long string")
	}
}

func BenchmarkCalculateDistanceVeryLongString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CalculateDistance(
			"this is very much a super long string, it is actually quite slow to do this because of the operation time",
			"this is another super long string, testing this benchmark is actually quite a lot to process in terms of the matrix",
		)
	}
}

func BenchmarkCalculateDistanceShort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CalculateDistance("flash", "lash")
	}
}
