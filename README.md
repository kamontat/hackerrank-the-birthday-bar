# The Birthday Bar Performance Check

>> https://www.hackerrank.com/challenges/the-birthday-bar/problem

We have 2 implementation, 

1. `normal loop and check sum on each group`
2. `create total sum of each element on array and calculate diff on each group`

## Result

```
> go test -bench=. ./logic -benchmem -cpu 1,2,4
goos: darwin
goarch: amd64

pkg: hackerrank.kamontat.net/question/the-birthday-bar/logic

cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz

BenchmarkNormal/Normal_(size=83,_d=3,_m=3)     	 3869742	   290.7 ns/op	   0 B/op	   0 allocs/op
BenchmarkNormal/Normal_(size=83,_d=3,_m=3)-2   	 3953919	   288.9 ns/op	   0 B/op	   0 allocs/op
BenchmarkNormal/Normal_(size=83,_d=3,_m=3)-4   	 4015725	   289.3 ns/op	   0 B/op	   0 allocs/op
BenchmarkNormal/Sum_(size=83,_d=3,_m=3)        	 6725364	   177.2 ns/op	   0 B/op	   0 allocs/op
BenchmarkNormal/Sum_(size=83,_d=3,_m=3)-2      	 6805952	   176.3 ns/op	   0 B/op	   0 allocs/op
BenchmarkNormal/Sum_(size=83,_d=3,_m=3)-4      	 6825538	   177.4 ns/op	   0 B/op	   0 allocs/op
```

## Summary

**second** implementation (`Sum`) is ~40% faster
