package client

import (
	"os"
	"sync"

	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/term"
)

func MakeConcurrentRequests(requests []*Request, client ClientInterface) []*ConcurrentResponse {
	var wg sync.WaitGroup
	responsesCh := make(chan *ConcurrentResponse, len(requests))

	// Get the terminal size
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // Default width
	}

	// Adjust the progress bar width based on terminal size
	barWidth := width - 40 // Subtract 40 to account for other UI elements

	bar := progressbar.NewOptions(
		len(requests),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(barWidth),
		progressbar.OptionSetTheme(
			progressbar.Theme{
				Saucer:        "[green]█[reset]",
				SaucerHead:    "[green]▓[reset]",
				SaucerPadding: "░",
				BarStart:      "╢",
				BarEnd:        "╟",
			},
		),
	)

	for _, request := range requests {
		wg.Add(1)
		go func(request *Request) {
			defer wg.Done()
			resp, err := client.Do(request.Uri, request.Method, request.QueryParams, request.Headers, request.Body)
			responsesCh <- NewConcurrentResponse(resp, err)
			bar.Add(1)
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
