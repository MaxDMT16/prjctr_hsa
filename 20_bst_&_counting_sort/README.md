## Run benchmarks
```bash
go test -bench=.
```

## AVL tree (balanced binary search tree)

| operation |items count| ops/ sec|time|
|-----------|-----------|---------|----|
| insert | 10 | 1237154 | 961.9 ns/op |
| insert | 100 | 80379 | 13776 ns/op |
| insert | 1000 | 6758 | 179777 ns/op |
| insert | 10000 | 477 | 2394143 ns/op |
| insert | 100000 | 31 | 35982921 ns/op |
| insert* | 1000000 | 12 | 951821782 ns/op |
| insert (same input) | 10 | 9543699 | 110.1 ns/op |
| insert (same input) | 100 | 1646738 | 721.7 ns/op |
| insert (same input) | 1000 | 163234 | 6802 ns/op |
| insert (same input) | 10000 | 17694 | 67285 ns/op |
| insert (same input) | 100000 | 1546 | 698540 ns/op |
| insert (same input)* | 1000000 | 1713 | 6743177 ns/op |
| find | 10 | 298964486 | 3.808 ns/op |
| find | 100 | 305052837 | 3.804 ns/op |
| find | 1000 | 307928382 | 3.773 ns/op |
| find | 10000 | 304441972 | 3.820 ns/op |
| find | 100000 | 305117412 | 3.796 ns/op |
| find* | 1000000 | 1000000000 | 3.807 ns/op |
| find (same input) | 10 | 42420254 | 28.89 ns/op |
| find (same input) | 100 | 31248565 | 39.23 ns/op |
| find (same input) | 1000 | 22474227 | 47.85 ns/op |
| find (same input) | 10000 | 18328552 | 59.79 ns/op |
| find (same input) | 100000 | 14822144 | 72.85 ns/op |
| find (same input)* | 1000000 | 152889236 | 76.93 ns/op |
| remove | 10 | 28278637 | 40.99 ns/op |
| remove | 100 | 20618539 | 55.23 ns/op |
| remove | 1000 | 16050721 | 69.37 ns/op |
| remove | 10000 | 11622096 | 92.50 ns/op |
| remove | 100000 | 9515421 | 118.1 ns/op |
| remove* | 1000000 | 85635900 | 130.3 ns/op |
| remove (same input) | 10 | 85232859 | 12.97 ns/op |
| remove (same input) | 100 | 86658385 | 13.10 ns/op |
| remove (same input) | 1000 | 84102798 | 13.02 ns/op |
| remove (same input) | 10000 | 88649030 | 13.04 ns/op |
| remove (same input) | 100000 | 81596511 | 13.03 ns/op |
| remove (same input)* | 1000000 | 877455028 | 13.03 ns/op |


* for 1000000 items count, benchmark run for 10 seconds, to more results for statistics

## Counting sort

For random input where input values are in range [0, 'items_count']
| items_count | ops/ sec | time |
|-----------|-----------|---------|
| 1 | 16863336 |  64.81 ns/op |
| 10 | 3915160 | 298.4 ns/op |
| 100 | 412136 |     3454 ns/op |
| 1000 | 39264 |     29935 ns/op |
| 10000 | 3722 |     375684 ns/op |
| 100000 | 342 |     3474316 ns/op |
| 1000000 | 12 |     92562114 ns/op |

**Counting sort has poor performance for arrays of small size with big values inside.Under the hood it creates a *count array* that has length of the biggest input array value. Thus, it's risky to use this algorithm to sort dynamic input.**