package client

import (
	"sync"
)

func MakeConcurrentRequests(requests []*Request, client ClientInterface) []*ConcurrentResponse {
	var wg sync.WaitGroup
	responsesCh := make(chan *ConcurrentResponse, len(requests))

	for _, request := range requests {
		wg.Add(1)
		go func(request *Request) {
			defer wg.Done()
			resp, err := client.Do(request.Uri, request.Method, request.QueryParams, request.Headers, request.Body)
			responsesCh <- NewConcurrentResponse(resp, err)
		}(request)
	}

	go func() {
		wg.Wait()
		close(responsesCh)
	}()

	var responses []*ConcurrentResponse
	for resp := range responsesCh {
		responses = append(responses, resp)
	}

	return responses

}
