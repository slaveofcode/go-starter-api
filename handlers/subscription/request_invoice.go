package subscription

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
	gowk "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
)

type downloadBodyParam struct {
	Plan string `json:"plan"`
}

// DownloadInvoice will return invoice for the current plan
func (t Subscription) RequestInvoice(sessionData *session.Data) func(*fasthttp.RequestCtx) {
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

		wkPdf, err := gowk.NewPDFGenerator()
		if err != nil {
			httpresponse.JSONErr(ctx, "Error generate invoice: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		wkPdf.Dpi.Set(300)
		wkPdf.Orientation.Set(gowk.OrientationPortrait)
		wkPdf.Grayscale.Set(true)

		pwd, _ := os.Getwd()
		tpl, err := ioutil.ReadFile(pwd + "/ext/files/invoice_tpl.html")
		if err != nil {
			httpresponse.JSONErr(ctx, "Error generate invoice: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(tpl))
		if err != nil {
			httpresponse.JSONErr(ctx, "Error generate invoice: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		doc.Find(".information table td:first-child").Each(func(i int, s *goquery.Selection) {
			s.SetHtml("Go Starter <br>Mega Plaza, Kuningan<br>Jakarta, Indonesia 13234")
		})

		res, err := doc.Html()

		page := gowk.NewPageReader(bytes.NewReader([]byte(res)))
		page.FooterRight.Set("[page]")
		page.FooterFontSize.Set(10)
		page.Zoom.Set(0.95)

		wkPdf.AddPage(page)

		err = wkPdf.Create()
		if err != nil {
			httpresponse.JSONErr(ctx, "Error generate invoice: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		err = wkPdf.WriteFile(pwd + "/ext/files/invoice.pdf")
		if err != nil {
			httpresponse.JSONErr(ctx, "Error generate invoice: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		// for _, item := range plans {
		// doc.AppendItem(&generator.Item{
		// 	Name:        PlanToHuman(item.PlanType),
		// 	Description: PlanDescriptionToHuman(item.PlanType),
		// 	Quantity:    "1",
		// 	UnitCost:    "USD " + strconv.FormatInt(PlanToPriceUSD(item.PlanType), 10),
		// 	Tax: &generator.Tax{
		// 		Percent: "10",
		// 	},
		// })
		// }

		// save to disk

		// upload to S3
		// return the S3 url

		httpresponse.JSONOk(ctx, fasthttp.StatusOK)
		return
	}
}
