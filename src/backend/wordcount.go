package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/binary"
	"runtime"
	"strings"
	"sync"

	"github.com/cnnrznn/gomr"
	. "github.com/cnnrznn/wordcounter/util"
)

type WordCountJob struct{}

func (w *WordCountJob) Map(in <-chan interface{}, out chan<- interface{}) {
	counts := make(map[string]int)

	for elem := range in {
		for _, word := range strings.Split(elem.(string), " ") {
			counts[word]++
		}
	}

	for k, v := range counts {
		out <- Count{k, v}
	}

	close(out)
}

func (w *WordCountJob) Partition(in <-chan interface{}, outs []chan interface{}, wg *sync.WaitGroup) {
	for elem := range in {
		key := elem.(Count).Key

		h := sha1.New()
		h.Write([]byte(key))
		hash := int(binary.BigEndian.Uint64(h.Sum(nil)))
		if hash < 0 {
			hash = hash * -1
		}

		outs[hash%len(outs)] <- elem
	}

	wg.Done()
}

func (w *WordCountJob) Reduce(in <-chan interface{}, out chan<- interface{}, wg *sync.WaitGroup) {
	counts := make(map[string]int)

	for elem := range in {
		ct := elem.(Count)
		counts[ct.Key] += ct.Val
	}

	for k, v := range counts {
		out <- Count{k, v}
	}

	wg.Done()
}

func wordcount(text string) (counts []Count) {
	nCPU := runtime.NumCPU()
	wcj := &WordCountJob{}

	ins, out := gomr.RunLocal(nCPU, nCPU, wcj)

	go func() {
		scanner := bufio.NewScanner(strings.NewReader(text))
		for i := 0; scanner.Scan(); i = (i + 1) % len(ins) {
			ins[i] <- scanner.Text()
		}

		for _, ch := range ins {
			close(ch)
		}
	}()

	for item := range out {
		count := item.(Count)
		counts = append(counts, count)
	}

	return
}
