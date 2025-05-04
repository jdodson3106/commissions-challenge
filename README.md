# Challenge Benchamrking

> I still think there are some memory reduction that could be had to be a bit more optimal. I ran a few benchmarks and this was the average of the result.

To run benchmark execute the following command in the project root
```bash
go test -bench=. -benchmem
```

Benchmarking against 125,000 visits total and ~10% of the visits converting to an order
| Metric        | `BenchmarkOriginal` | `BenchmarkMoreOptimal` | Difference           |
| ------------- | ------------------- | ---------------------- | -------------------- |
| Time per op   | \~954 ms            | \~3 ms                 | \~314x faster        |
| Bytes per op  | 192 B               | ~109 MB                | \~567x more memory   |
| Allocs per op | 2                   | 22                     | 11x more allocations |
