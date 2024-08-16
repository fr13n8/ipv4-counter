package counter

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sync"

	"github.com/bits-and-blooms/bitset"
	"golang.org/x/exp/mmap"
)

var myBitSet = bitset.New(math.MaxInt)

// var mu sync.Mutex

func IPV4CountFromFile(input string, gWorkers, BUFFER_SIZE int, isMmap bool) (uint, error) {
	if isMmap {
		if err := readFileLineByLineWithMmap(input, gWorkers); err != nil {
			return 0, fmt.Errorf("counting run: %w", err)
		}
	} else {
		if err := readFileLineByLine(input, gWorkers, BUFFER_SIZE); err != nil {
			return 0, fmt.Errorf("counting run: %w", err)
		}
	}

	return myBitSet.Count(), nil
}

var bufferPool *sync.Pool

func newBufferPool(bufferSizeMB int) *sync.Pool {
	return &sync.Pool{
		New: func() any {
			b := make([]byte, bufferSizeMB)
			return &b
		},
	}
}

func readFileLineByLineWithMmap(filepath string, gWorkers int) error {
	file, err := mmap.Open(filepath)
	if err != nil {
		return fmt.Errorf("open file error: %w", err)
	}
	defer file.Close()

	chunkStream := make(chan []byte, gWorkers)
	chunkSize := file.Len() / gWorkers
	var wg sync.WaitGroup

	bufferPool = newBufferPool(chunkSize)

	wg.Add(gWorkers)
	for i := 0; i < gWorkers; i++ {
		go func() {
			defer wg.Done()
			for chunk := range chunkStream {
				processReadChunk(&chunk)
				bufferPool.Put(&chunk)
			}
		}()
	}

	var innErr error
	go func() {
		leftover := make([]byte, 0, chunkSize)
		defer close(chunkStream)
		offset := int64(0)
		for {
			buf := *(bufferPool.Get().(*[]byte))
			readTotal, err := file.ReadAt(buf, offset)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				innErr = fmt.Errorf("read from file error: %w", err)
				return
			}
			buf = buf[:readTotal]
			offset += int64(readTotal)

			// "\n" == 10 ASCII
			lastNewLineIndex := bytes.LastIndex(buf, []byte{10})

			// This line combines any leftover data from the previous chunk (if there was an incomplete line) with the current chunk up to the last newline.
			// includes everything from the start of the buffer up to (and including) the last newline character.
			// toSend contains only complete lines.
			toSend := append(leftover, buf[:lastNewLineIndex+1]...)

			// There might be leftover data after the last newline character. This could be an incomplete line that needs to be preserved for the next chunk.
			// For that we create a slice with the correct size for the leftover data.
			leftover = make([]byte, len(buf[lastNewLineIndex+1:]))
			// Copies the leftover data from the current chunk into the new slice. This data will be appended to the next chunk's beginning.
			copy(leftover, buf[lastNewLineIndex+1:])

			chunkStream <- toSend
		}
	}()

	wg.Wait()

	return innErr
}

func readFileLineByLine(filepath string, gWorkers, BUFFER_SIZE int) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("open file error: %w", err)
	}
	defer file.Close()

	chunkStream := make(chan []byte, gWorkers)
	chunkSize := BUFFER_SIZE * 1024 * 1024
	var wg sync.WaitGroup

	bufferPool = newBufferPool(chunkSize)

	wg.Add(gWorkers)
	for i := 0; i < gWorkers; i++ {
		go func() {
			defer wg.Done()
			for chunk := range chunkStream {
				processReadChunk(&chunk)
				bufferPool.Put(&chunk)
			}
		}()
	}

	var innErr error
	go func() {
		leftover := make([]byte, 0, chunkSize)
		defer close(chunkStream)
		for {
			buf := *(bufferPool.Get().(*[]byte))
			readTotal, err := file.Read(buf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				innErr = fmt.Errorf("read from file error: %w", err)
				return
			}
			buf = buf[:readTotal]

			// "\n" == 10 ASCII
			lastNewLineIndex := bytes.LastIndex(buf, []byte{10})

			// This line combines any leftover data from the previous chunk (if there was an incomplete line) with the current chunk up to the last newline.
			// includes everything from the start of the buffer up to (and including) the last newline character.
			// toSend contains only complete lines.
			toSend := append(leftover, buf[:lastNewLineIndex+1]...)

			// There might be leftover data after the last newline character. This could be an incomplete line that needs to be preserved for the next chunk.
			// For that we create a slice with the correct size for the leftover data.
			leftover = make([]byte, len(buf[lastNewLineIndex+1:]))
			// Copies the leftover data from the current chunk into the new slice. This data will be appended to the next chunk's beginning.
			copy(leftover, buf[lastNewLineIndex+1:])

			chunkStream <- toSend
		}
	}()

	wg.Wait()

	return innErr
}

func processReadChunk(buf *[]byte) {
	start := 0
	for index, char := range *buf {
		switch char {
		// "\n" == 10 ASCII
		case 10:
			ip := (*buf)[start:index]
			start = index + 1

			num := IPv4toDec(&ip)

			// mu.Lock()
			if !myBitSet.Test(num) {
				myBitSet.Set(num)
			}
			// mu.Unlock()
		}
	}
}

// https://interlir.com/2024/02/19/converting-ipv4-addresses-to-decimal-a-step-by-step-guide/
func IPv4toDec(ipAddress *[]byte) uint {
	var result, part uint
	// Decimal IP = (A x 256^3) + (B x 256^2) + (C x 256^1) + (D x 256^0)
	// Decimal IP = A << (8*3) | B << (8*2) | C << (8*1) | D << (8*0)
	// "." == 46 ASCII
	// "0" == 48 ASCII
	for _, symbol := range *ipAddress {
		if symbol == 46 {
			// result * 256 | part [When the first dot is encountered, 192 is shifted left by 8 bits, leaving space for the next octet (168) to be added.]
			// 2) 0 * 256 | 192 => result = 192
			// 4) 192 * 256	| 168 = 49152 + 168 => result = 49320
			// 6) 49320 * 256 | 0 => result = 12625920
			result = (result << 8) | part
			part = 0
		} else {
			// 1) '1' - '0' = 49 - 48 = 1 => part = 1
			// 1) '9' - '0' = 57 - 48 = 9 => part = 10 + 9 = 19
			// 1) '2' - '0' = 50 - 48 = 2 => part = 190 + 8 = 192
			// 3) '1' - '0' = 49 - 48 = 1 => part = 1
			// 3) '6' - '0' = 54 - 48 = 6 => part = 16
			// 3) '8' - '0' = 56 - 48 = 8 => part = 168
			// 5) '0' - '0' = 48 - 48 = 0 => part = 0
			// 7) '1' - '0' = 49 - 48 = 1 => part = 1
			part = part*10 + uint(symbol-48)
		}
	}
	// 8) 12625920 * 256 | 1 = 3232235521
	return (result << 8) | part
}
