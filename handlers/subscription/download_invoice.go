package subscription

import (
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/valyala/fasthttp"
)

// DownloadInvoice will return invoice for the current plan
func (t Subscription) DownloadInvoice(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		// return pdf invoice
	}
}
