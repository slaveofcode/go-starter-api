package subscription

import (
	"encoding/json"
	"strconv"

	generator "github.com/angelodlfrtr/go-invoice-generator"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
)

type downloadBodyParam struct {
	Plan string `json:"plan"`
}

// DownloadInvoice will return invoice for the current plan
func (t Subscription) DownloadInvoice(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		var param downloadBodyParam
		err := json.Unmarshal(ctx.PostBody(), &param)
		if err != nil {
			httpresponse.JSONErr(ctx, "Wrong post data: "+err.Error(), fasthttp.StatusBadRequest)
			return
		}

		if IsValidPlan(param.Plan) {
			httpresponse.JSONErr(ctx, "Invalid Plan: "+param.Plan, fasthttp.StatusBadRequest)
			return
		}

		db := t.appCtx.DB

		var plans []models.Subscription
		if err := db.Where(&models.Subscription{
			UserID: sessionData.UserID,
		}).Find(&plans).Error; err != nil {
			httpresponse.JSONErr(ctx, "Error fetching subscription: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		doc, err := generator.New(generator.INVOICE, &generator.Options{})
		if err != nil {
			httpresponse.JSONErr(ctx, "Unable to create invoice: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		doc.SetHeader(&generator.HeaderFooter{
			Text: "Invoice",
		})

		doc.SetFooter(&generator.HeaderFooter{
			Text:       "<center>Go Starter Company</center>",
			Pagination: true,
		})

		doc.SetRef("0001")
		doc.SetVersion("1.0")
		doc.SetDescription("Invoice Letter")

		doc.SetCompany(&generator.Contact{
			Name: "Go Starter",
			Logo: &[]byte{},
			Address: &generator.Address{
				Address:    "Street No. 1",
				Address2:   "Street No. 2",
				PostalCode: "3243",
				City:       "Jakarta",
				Country:    "Indonesia",
			},
		})

		doc.SetCustomer(&generator.Contact{
			Name: sessionData.Name,
		})

		for _, item := range plans {
			doc.AppendItem(&generator.Item{
				Name:        PlanToHuman(item.PlanType),
				Description: PlanDescriptionToHuman(item.PlanType),
				Quantity:    "1",
				UnitCost:    "USD " + strconv.FormatInt(PlanToPriceUSD(item.PlanType), 10),
				Tax: &generator.Tax{
					Percent: "10",
				},
			})
		}

		pdf, err := doc.Build()

		if err != nil {
			httpresponse.JSONErr(ctx, "Unable to create invoice: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		// save to disk
		pdf.SetI
		err = pdf.OutputFileAndClose("./ext/files/invoive.pdf")
		if err != nil {
			httpresponse.JSONErr(ctx, "Unable to store invoice: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		// upload to S3
		// return the S3 url

		httpresponse.JSONOk(ctx, fasthttp.StatusOK)
		return
	}
}
