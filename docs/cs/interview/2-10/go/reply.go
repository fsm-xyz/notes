package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

}

const (
	workersNum   = 6
	maxQueueSize = 10000

	timeout       = 5 * time.Second
	replayTimeout = 1 * time.Second
)

var (
	replayQueue = make(chan replay, maxQueueSize)

	client = http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 100,
		},
	}
)

func init() {
	go runRelay()
}

// 1. 假定使用最基本的http库
// 2. 假定不使用太多的黑科技
// 3. 日志不用zerolog和zap
func forward(uri, host, rHost string) int {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	u, err := url.Parse(uri)
	if err != nil {
		log.Println(uri, " parse err", err)
		return http.StatusBadRequest
	}

	subUri := u.RequestURI()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, host+subUri, nil)
	if err != nil {
		log.Println("create request err: ", err)
		return http.StatusInternalServerError
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("http do request err: ", err)
		return http.StatusInternalServerError
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)

	sendToReplayQueue(replay{
		uri:  subUri,
		host: rHost,
	})
	return resp.StatusCode
}

type replay struct {
	uri  string
	host string
}

func runRelay() {
	for i := 0; i < workersNum; i++ {
		go func() {
			for v := range replayQueue {
				replayRequest(v)
			}
		}()
	}
}

func sendToReplayQueue(r replay) {
	select {
	case replayQueue <- r:
	default:
		log.Println("replyQueue is full drop request")
	}
}

func replayRequest(r replay) {
	ctx, cancel := context.WithTimeout(context.Background(), replayTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.host+r.uri, nil)
	if err != nil {
		log.Println("create request err: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("http do request err: ", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)
}
