package client

import "context"

func (rc *RateLimitedClient) Do(uri string, method string, queryParams any, headers any, reqBody any) (*Response, error) {
	ctx := context.Background()
	err := rc.RateLimited.Wait(ctx) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}
	return rc.FHClient.Do(uri, method, queryParams, headers, reqBody)
}
