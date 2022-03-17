package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"sort"
	"sync"
	"time"
)

type singleHandler func() (data string, n int, err error)

// handlerConfig 代表处理流程配置的类型。
type handlerConfig struct {
	handler   singleHandler // 单次处理函数。
	goNum     int           // 需要启用的goroutine的数量。
	number    int           // 单个goroutine中的处理次数。
	interval  time.Duration // 单个goroutine中的处理间隔时间。
	counter   int           // 数据量计数器，以字节为单位。
	counterMu sync.Mutex    // 数据量计数器专用的互斥锁。

}

func (hc *handlerConfig) count(increment int) int {
	hc.counterMu.Lock()
	defer hc.counterMu.Unlock()
	hc.counter += increment
	return hc.counter
}

func test(writer io.Writer) int {
	return 0
}

func main() {
	var mu sync.Mutex

	genWriter := func(writer io.Writer) singleHandler {
		return func() (data string, n int, err error) {
			// 准备数据。
			data = fmt.Sprintf("%s\t",
				time.Now().Format(time.StampNano))
			mu.Lock()
			defer mu.Unlock()
			n, err = writer.Write([]byte(data))
			return data, n, err
		}
	}
	// genReader 代表的是用于生成读取函数的函数。
	genReader := func(reader io.Reader) singleHandler {
		return func() (data string, n int, err error) {
			buffer, ok := reader.(*bytes.Buffer)
			if !ok {
				err = errors.New("unsupported reader")
				return
			}
			// 读取数据。
			mu.Lock()
			defer mu.Unlock()
			data, err = buffer.ReadString('\t')
			n = len(data)
			return
		}
	}

	// buffer 代表缓冲区。
	var buffer bytes.Buffer

	// 数据写入配置。
	writingConfig := handlerConfig{
		handler:  genWriter(&buffer),
		goNum:    5,
		number:   4,
		interval: time.Millisecond * 100,
	}
	// 数据读取配置。
	readingConfig := handlerConfig{
		handler:  genReader(&buffer),
		goNum:    10,
		number:   2,
		interval: time.Millisecond * 100,
	}

	sign := make(chan struct{}, writingConfig.goNum+readingConfig.goNum)
	// 启用多个goroutine对缓冲区进行多次数据写入。
	for i := 1; i <= writingConfig.goNum; i++ {
		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j <= writingConfig.number; j++ {
				time.Sleep(writingConfig.interval)
				data, n, err := writingConfig.handler()
				if err != nil {
					log.Printf("writer [%d-%d]: error: %s",
						i, j, err)
					continue
				}
				total := writingConfig.count(n)
				log.Printf("writer [%d-%d]: %s (total: %d)",
					i, j, data, total)
			}
		}(i)
	}

	// 启用多个goroutine对缓冲区进行多次数据读取。
	for i := 1; i <= readingConfig.goNum; i++ {

		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j <= readingConfig.number; j++ {
				time.Sleep(readingConfig.interval)
				var data string
				var n int
				var err error
				for {
					data, n, err = readingConfig.handler()
					if err == nil || err != io.EOF {
						break
					}
					// 如果读比写快（读时会发生EOF错误），那就等一会儿再读。
					time.Sleep(readingConfig.interval)
				}
				if err != nil {
					log.Printf("reader [%d-%d]: error: %s",
						i, j, err)
					continue
				}
				total := readingConfig.count(n)
				log.Printf("reader [%d-%d]: %s (total: %d)",
					i, j, data, total)
			}
		}(i)
	}

	// signNumber 代表需要接收的信号的数量。
	signNumber := writingConfig.goNum + readingConfig.goNum
	// 等待上面启用的所有goroutine的运行全部结束。
	for j := 0; j < signNumber; j++ {
		<-sign
	}
}

func longestWord(words []string) (longest string) {
	sort.Slice(words, func(i, j int) bool {
		s, t := words[i], words[j]
		return len(s) < len(t) || len(s) == len(t) && s > t
	})
	candidates := map[string]struct {
		a string
	}{"": {"f"}}
	for _, word := range words {
		if _, ok := candidates[word[:len(word)-1]]; ok {
			longest = word
			candidates[word] = struct{ a string }{a: ""}
		}
	}
	return longest
}
