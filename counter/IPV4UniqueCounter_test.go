package counter

import (
	"fmt"
	"log"
	"runtime"
	"testing"
)

var (
	inputFile = "../ip-addr.txt"
)

func TestIPv4Converter(t *testing.T) {
	type TestCase struct {
		name  string
		input []byte
		want  uint
	}

	var cases = []TestCase{}

	ips := map[uint][]byte{
		8:          []byte("0.0.0.8"),
		255:        []byte("0.0.0.255"),
		256:        []byte("0.0.1.0"),
		300:        []byte("0.0.1.44"),
		65536:      []byte("0.1.0.0"),
		65792:      []byte("0.1.1.0"),
		2147549680: []byte("128.1.1.240"),
		2130772464: []byte("127.1.1.240"),
		1:          []byte("0.0.0.1"),
		0:          []byte("0.0.0.0"),
		4294967295: []byte("255.255.255.255"),
		2147483647: []byte("127.255.255.255"),
		2147483648: []byte("128.0.0.0"),
	}

	for res, ip := range ips {
		cases = append(cases, TestCase{
			input: ip,
			want:  res,
			name:  "Ip to dec",
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ans := IPv4toDec(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkIPV4CountFromFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if _, err := IPV4CountFromFile(inputFile, runtime.NumCPU(), 1024); err != nil {
			log.Fatalln(err)
		}
	}
}

var (
	sizes   = []int{64, 512, 1024, 2048} // mb
	workers = []int{1, 2, 4, 8}
)

func BenchmarkIPV4CountFromFileOpts(b *testing.B) {
	type TestCase struct {
		BufferSize int
		GCount     int
	}

	cases := []TestCase{}

	for _, s := range sizes {
		for _, w := range workers {
			cases = append(cases, TestCase{
				BufferSize: s,
				GCount:     w * runtime.NumCPU(),
			})
		}
	}

	for i := range cases {
		b.Run(fmt.Sprintf("input_size_%d_goroutines_count_%d", cases[i].BufferSize, cases[i].GCount), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				IPV4CountFromFile(inputFile, cases[i].GCount, cases[i].GCount)
			}
		})
	}
}
