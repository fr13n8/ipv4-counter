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
go run main.go -input="file path" -cf="cpu.prof"" -tf="trace.out" -mf="mem.prof"
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
$ go test -v ./... --count 10 -run=^$ -benchmem -bench=BenchmarkIPV4CountFromFileOpts | benchstat -

goos: windows
goarch: amd64
pkg: github.com/fr13n8/ipv4-counter/counter
cpu: AMD Ryzen 9 5900HX with Radeon Graphics
                                                              │ .\BenchmarkIPV4CountFromFileOpts.out │
                                                              │                sec/op                │
IPV4CountFromFileOpts/input_size_64_goroutines_count_16-16                              97.48m ±  5%
IPV4CountFromFileOpts/input_size_64_goroutines_count_32-16                              120.0m ±  2%
IPV4CountFromFileOpts/input_size_64_goroutines_count_64-16                              154.6m ±  4%
IPV4CountFromFileOpts/input_size_64_goroutines_count_128-16                             210.1m ±  3%
IPV4CountFromFileOpts/input_size_64_goroutines_count_160-16                             215.5m ±  1%
IPV4CountFromFileOpts/input_size_512_goroutines_count_16-16                             101.9m ±  4%
IPV4CountFromFileOpts/input_size_512_goroutines_count_32-16                             120.3m ±  2%
IPV4CountFromFileOpts/input_size_512_goroutines_count_64-16                             152.2m ±  6%
IPV4CountFromFileOpts/input_size_512_goroutines_count_128-16                            211.6m ±  2%
IPV4CountFromFileOpts/input_size_512_goroutines_count_160-16                            216.9m ±  2%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_16-16                            102.4m ± 15%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_32-16                            127.4m ± 15%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_64-16                            154.2m ±  4%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_128-16                           211.6m ±  2%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_160-16                           216.1m ±  1%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_16-16                            101.4m ±  5%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_32-16                            118.6m ±  2%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_64-16                            158.3m ±  7%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_128-16                           209.4m ±  1%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_160-16                           215.9m ±  1%
geomean                                                                                 153.9m

                                                              │ .\BenchmarkIPV4CountFromFileOpts.out │
                                                              │                 B/op                 │
IPV4CountFromFileOpts/input_size_64_goroutines_count_16-16                              262.5Mi ± 2%
IPV4CountFromFileOpts/input_size_64_goroutines_count_32-16                              281.1Mi ± 0%
IPV4CountFromFileOpts/input_size_64_goroutines_count_64-16                              313.1Mi ± 0%
IPV4CountFromFileOpts/input_size_64_goroutines_count_128-16                             384.0Mi ± 0%
IPV4CountFromFileOpts/input_size_64_goroutines_count_160-16                             480.0Mi ± 0%
IPV4CountFromFileOpts/input_size_512_goroutines_count_16-16                             262.5Mi ± 1%
IPV4CountFromFileOpts/input_size_512_goroutines_count_32-16                             281.1Mi ± 0%
IPV4CountFromFileOpts/input_size_512_goroutines_count_64-16                             313.1Mi ± 0%
IPV4CountFromFileOpts/input_size_512_goroutines_count_128-16                            384.0Mi ± 0%
IPV4CountFromFileOpts/input_size_512_goroutines_count_160-16                            480.0Mi ± 0%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_16-16                            262.7Mi ± 1%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_32-16                            281.1Mi ± 0%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_64-16                            313.1Mi ± 0%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_128-16                           384.0Mi ± 0%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_160-16                           480.0Mi ± 0%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_16-16                            263.5Mi ± 1%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_32-16                            281.1Mi ± 0%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_64-16                            313.1Mi ± 0%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_128-16                           384.0Mi ± 0%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_160-16                           480.0Mi ± 0%
geomean                                                                                 335.7Mi

                                                              │ .\BenchmarkIPV4CountFromFileOpts.out │
                                                              │              allocs/op               │
IPV4CountFromFileOpts/input_size_64_goroutines_count_16-16                               75.00 ±  7%
IPV4CountFromFileOpts/input_size_64_goroutines_count_32-16                               66.50 ±  8%
IPV4CountFromFileOpts/input_size_64_goroutines_count_64-16                               96.00 ± 14%
IPV4CountFromFileOpts/input_size_64_goroutines_count_128-16                              172.5 ± 12%
IPV4CountFromFileOpts/input_size_64_goroutines_count_160-16                              200.0 ± 11%
IPV4CountFromFileOpts/input_size_512_goroutines_count_16-16                              72.00 ±  4%
IPV4CountFromFileOpts/input_size_512_goroutines_count_32-16                              66.50 ± 11%
IPV4CountFromFileOpts/input_size_512_goroutines_count_64-16                              97.50 ±  5%
IPV4CountFromFileOpts/input_size_512_goroutines_count_128-16                             164.5 ± 11%
IPV4CountFromFileOpts/input_size_512_goroutines_count_160-16                             198.0 ±  4%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_16-16                             71.00 ±  3%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_32-16                             66.00 ± 11%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_64-16                             94.00 ±  9%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_128-16                            161.0 ±  8%
IPV4CountFromFileOpts/input_size_1024_goroutines_count_160-16                            200.0 ±  6%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_16-16                             71.00 ±  3%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_32-16                             66.00 ± 11%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_64-16                             95.00 ±  5%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_128-16                            171.0 ± 10%
IPV4CountFromFileOpts/input_size_2048_goroutines_count_160-16                            207.0 ±  7%
geomean                                                                                  109.0
```

IPv4 addresses file size ~120gb

```bash
$ go test -v ./... --count 2 -run=^$ -benchmem -bench=BenchmarkIPV4CountFromFileOpts > BenchmarkIPV4CountFromFileOpts120GB.out

goos: windows
goarch: amd64
pkg: github.com/fr13n8/ipv4-counter/counter
cpu: AMD Ryzen 9 5900HX with Radeon Graphics        
BenchmarkIPV4CountFromFileOpts
BenchmarkIPV4CountFromFileOpts/input_size_2048_goroutines_count_16
BenchmarkIPV4CountFromFileOpts/input_size_2048_goroutines_count_16-16                 1 63489548900 ns/op 122355327208 B/op    22532 allocs/op
BenchmarkIPV4CountFromFileOpts/input_size_2048_goroutines_count_16-16                 1 62358866300 ns/op 121013232680 B/op    22474 allocs/op
PASS
ok   github.com/fr13n8/ipv4-counter/counter  126.095s
```

```bash
$ benchstat BenchmarkIPV4CountFromFileOpts120GB.out

goos: windows
goarch: amd64
pkg: github.com/fr13n8/ipv4-counter/counter
cpu: AMD Ryzen 9 5900HX with Radeon Graphics
                                                             │ .\BenchmarkIPV4CountFromFileOpts120GB.out │
                                                             │                  sec/op                   │
IPV4CountFromFileOpts/input_size_2048_goroutines_count_16-16                                 62.92 ± ∞ ¹
¹ need >= 6 samples for confidence interval at level 0.95

                                                             │ .\BenchmarkIPV4CountFromFileOpts120GB.out │
                                                             │                   B/op                    │
IPV4CountFromFileOpts/input_size_2048_goroutines_count_16-16                               113.3Gi ± ∞ ¹
¹ need >= 6 samples for confidence interval at level 0.95

                                                             │ .\BenchmarkIPV4CountFromFileOpts120GB.out │
                                                             │                 allocs/op                 │
IPV4CountFromFileOpts/input_size_2048_goroutines_count_16-16                                22.50k ± ∞ ¹
¹ need >= 6 samples for confidence interval at level 0.95
```
