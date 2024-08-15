# Count unique IPv4 addresses from large file

There is a simple text file with IPv4 addresses. One line is one address, line by line:

```txt
145.67.23.4
8.34.5.23
89.54.3.124
89.54.3.124
3.45.71.5
...
```

**The file is unlimited in size and can occupy tens and hundreds of gigabytes.**

---
Download sample file from [here](https://ecwid-vgv-storage.s3.eu-central-1.amazonaws.com/ip_addresses.zip). Attention - the file weighs about 20Gb, and unzips to about 120Gb.

## Run app

```bash
go mod tidy
go run main.go -input="file path -cf="cpu.prof"" -tf="trace.out" -mf="mem.prof"
```

## Generate random IPv4 addresse with nmap

```bash
nmap -n -iR 100000 --exclude 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16,224.0.0.0/4 -sL | grep "Nmap scan report for" | awk '{print $NF}' > RANDOM-IPS.txt
```

`-iR 100000` : Count of random IPs

## Benchmarking and Tests

### Benchmarking

```bash
 go test -v ./... -count 10 -run=^$ -benchmem -bench=Benchmark | benchstat -
```

### Tests

```bash
go test -v ./...
```

## Some benchmarks

IPv4 addresses file size ~120mb

```bash
goos: windows
goarch: amd64
pkg: github.com/fr13n8/ipv4-counter/counter
cpu: AMD Ryzen 9 5900HX with Radeon Graphics
BenchmarkIPV4CountFromFile-16                  2         555930650 ns/op        2147490348 B/op       50 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_64_goroutines_count_16-16                   15          75494160 ns/op        143800361 B/op        56 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_64_goroutines_count_32-16                   10         108068020 ns/op        160543268 B/op        54 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_64_goroutines_count_64-16                    8         139382450 ns/op        194104992 B/op       110 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_64_goroutines_count_128-16                   5         215772440 ns/op        268446987 B/op       162 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_512_goroutines_count_16-16                  13          83985162 ns/op        143796261 B/op        42 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_512_goroutines_count_32-16                   9         114715133 ns/op        160542059 B/op        57 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_512_goroutines_count_64-16                   8         157169550 ns/op        194101018 B/op        98 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_512_goroutines_count_128-16                  5         225482480 ns/op        268445118 B/op       164 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_1024_goroutines_count_16-16                 13          86178100 ns/op        143796063 B/op        40 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_1024_goroutines_count_32-16                  9         120553178 ns/op        160541227 B/op        48 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_1024_goroutines_count_64-16                  8         156844838 ns/op        194099246 B/op        96 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_1024_goroutines_count_128-16                 5         221888440 ns/op        268444734 B/op       160 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_2048_goroutines_count_16-16                 13          84940031 ns/op        143796150 B/op        41 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_2048_goroutines_count_32-16                 10         119551090 ns/op        160541486 B/op        48 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_2048_goroutines_count_64-16                  7         152220743 ns/op        194097925 B/op        83 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_2048_goroutines_count_128-16                 5         217606180 ns/op        268445521 B/op       168 allocs/op
PASS
```

IPv4 addresses file size ~120gb

```bash
goos: windows
goarch: amd64
pkg: github.com/fr13n8/ipv4-counter/counter
cpu: AMD Ryzen 9 5900HX with Radeon Graphics
BenchmarkIPV4CountFromFile-16                  1        94009243200 ns/op       116825339728 B/op            351 allocs/op
IPV4CountFromFile-16   93.54 ± ∞ ¹
PASS
```

IPv4 addresses file size ~120mb

```bash
$ go test -v ./... -count 10 -run=^$ -benchmem -bench=Benchmark | benchstat -
goos: windows
goarch: amd64
pkg: github.com/fr13n8/ipv4-counter/counter
cpu: AMD Ryzen 9 5900HX with Radeon Graphics
                                                              │      -       │
                                                              │    sec/op    │
IPV4CountFromFile-16                                            417.2m ± 30%
IPV4CountFromFileOpts/input_size_64_goroutines_count_16-16      89.75m ± 25%
IPV4CountFromFileOpts/input_size_64_goroutines_count_32-16      120.5m ± 24%
IPV4CountFromFileOpts/input_size_64_goroutines_count_64-16      151.1m ±  3%
IPV4CountFromFileOpts/input_size_64_goroutines_count_128-16     282.3m ± 35%
IPV4CountFromFileOpts/input_size_512_goroutines_count_16-16     104.2m ± 22%
IPV4CountFromFileOpts/input_size_512_goroutines_count_32-16     116.6m ±  5%
IPV4CountFromFileOpts/input_size_512_goroutines_count_64-16     153.4m ± 10%
IPV4CountFromFileOpts/input_size_512_goroutines_count_128-16    226.1m ±  5%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_16-16    86.27m ±  2%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_32-16    119.0m ±  3%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_64-16    150.9m ±  8%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_128-16   218.4m ±  1%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_16-16    85.66m ±  2%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_32-16    116.8m ±  3%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_64-16    148.9m ±  6%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_128-16   223.1m ±  5%
geomean                                                         149.3m

                                                              │      -       │
                                                              │     B/op     │
IPV4CountFromFile-16                                            2.000Gi ± 0%
IPV4CountFromFileOpts/input_size_64_goroutines_count_16-16      137.1Mi ± 0%
IPV4CountFromFileOpts/input_size_64_goroutines_count_32-16      153.1Mi ± 0%
IPV4CountFromFileOpts/input_size_64_goroutines_count_64-16      185.1Mi ± 0%
IPV4CountFromFileOpts/input_size_64_goroutines_count_128-16     256.0Mi ± 0%
IPV4CountFromFileOpts/input_size_512_goroutines_count_16-16     137.1Mi ± 0%
IPV4CountFromFileOpts/input_size_512_goroutines_count_32-16     153.1Mi ± 0%
IPV4CountFromFileOpts/input_size_512_goroutines_count_64-16     185.1Mi ± 0%
IPV4CountFromFileOpts/input_size_512_goroutines_count_128-16    256.0Mi ± 0%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_16-16    137.1Mi ± 0%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_32-16    153.1Mi ± 0%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_64-16    185.1Mi ± 0%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_128-16   256.0Mi ± 0%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_16-16    137.1Mi ± 0%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_32-16    153.1Mi ± 0%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_64-16    185.1Mi ± 0%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_128-16   256.0Mi ± 0%
geomean                                                         205.1Mi

                                                              │      -      │
                                                              │  allocs/op  │
IPV4CountFromFile-16                                            31.50 ± 27%
IPV4CountFromFileOpts/input_size_64_goroutines_count_16-16      43.50 ±  6%
IPV4CountFromFileOpts/input_size_64_goroutines_count_32-16      53.50 ± 14%
IPV4CountFromFileOpts/input_size_64_goroutines_count_64-16      89.00 ± 11%
IPV4CountFromFileOpts/input_size_64_goroutines_count_128-16     154.5 ±  9%
IPV4CountFromFileOpts/input_size_512_goroutines_count_16-16     40.00 ±  2%
IPV4CountFromFileOpts/input_size_512_goroutines_count_32-16     50.50 ±  5%
IPV4CountFromFileOpts/input_size_512_goroutines_count_64-16     85.00 ± 11%
IPV4CountFromFileOpts/input_size_512_goroutines_count_128-16    162.0 ± 12%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_16-16    40.00 ±  2%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_32-16    52.50 ± 10%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_64-16    83.50 ±  4%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_128-16   157.0 ± 10%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_16-16    40.00 ±  0%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_32-16    51.00 ±  6%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_64-16    83.50 ±  8%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_128-16   157.5 ±  9%
geomean                                                         69.52
```
