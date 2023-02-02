# golang-datastruct-algorithm

This repo hosts code for datastruct and algorithm in golang. 

# Benchmark
- Go to folder that has benchmark file
- Run command
```
go test -bench=. -count=5 -run=^# -benchtime=10s
```
 - -count=5 run 5 times
 - -run=^# not running any test case, only run benchmark
 - -benchtime=10s the minimum amount of time that the benchmark should run

# Test
- Run command

```
go test ./dp/fib
```
