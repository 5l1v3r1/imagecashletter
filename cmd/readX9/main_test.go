package main

import "testing"

func TestFileRead(t *testing.T) {
	testFileRead(t)
}

/*//BenchmarkTestFileCreate benchmarks creating an X9 File
func BenchmarkTestFileRead(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileRead(b)
	}
}*/

// FileCreate creates an X9 File
func testFileRead(t testing.TB) {
	main()
}
