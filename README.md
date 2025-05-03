
| Metric        | `BenchmarkCopy` | `BenchmarkMoreOptimal` | Difference           |
| ------------- | --------------- | ---------------------- | -------------------- |
| Time per op   | \~954 ms        | \~3 ms                 | \~314x faster        |
| Bytes per op  | 192 B           | 108,952 B              | \~567x more memory   |
| Allocs per op | 2               | 22                     | 11x more allocations |
