
```
❯ go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/bellwood4486/sandbox-go/slice-allocate
BenchmarkAllocateOnce-8         1000000000               0.0425 ns/op          1 B/op          0 allocs/op
BenchmarkAllocateEverytime-8    1000000000               0.118 ns/op           1 B/op          0 allocs/op
PASS
ok      github.com/bellwood4486/sandbox-go/slice-allocate       2.419s
```
