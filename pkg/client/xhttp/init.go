package xhttp

import (
	"github.com/go-resty/resty/v2"
	"time"
)

type Client struct {
	Host string
	r *resty.Client
}

func New(host string) Client {
	r := resty.New()
	
	r.SetBaseURL(host)
	r.SetTimeout(10 * time.Second)
	
	r.SetDebug(true)
	r.OnRequestLog(requestLogCallback)
	r.OnResponseLog(responseLogCallback)
	
	return Client{r: r}
}

func (c *Client) GetClient() *resty.Client {
	return c.r
}

func (c *Client) R() *resty.Request {
	req := c.r.R()

	req.SetHeader("X-USER-HEADER", "1")

	return req
}