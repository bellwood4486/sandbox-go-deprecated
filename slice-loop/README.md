1回目
```
❯ go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/bellwood4486/sandbox-go/slice-loop
BenchmarkLoopWithIndex_100-8            100000000               16.4 ns/op
BenchmarkLoopWithValue_100-8            94014164                12.6 ns/op
BenchmarkLoopWithIndex_1000-8           87538519               497 ns/op
BenchmarkLoopWithValue_1000-8           16121455               538 ns/op
```

2回目
```
❯ go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/bellwood4486/sandbox-go/slice-loop
BenchmarkLoopWithIndex_100-8            97494472                10.4 ns/op
BenchmarkLoopWithValue_100-8            100000000               10.2 ns/op
BenchmarkLoopWithIndex_1000-8           94538271               500 ns/op
BenchmarkLoopWithValue_1000-8           14534280               536 ns/op
```

3回目
```
❯ go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/bellwood4486/sandbox-go/slice-loop
BenchmarkLoopWithIndex_100-8            96248690                23.1 ns/op
BenchmarkLoopWithValue_100-8            100000000               17.5 ns/op
BenchmarkLoopWithIndex_1000-8           98743707               488 ns/op
BenchmarkLoopWithValue_1000-8           16727176               536 ns/op
```
