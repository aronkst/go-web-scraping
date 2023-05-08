package javascript

import (
	"context"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func GetHTML(url string) (string, error) {
	scriptFulllLoad := `
		function sleep(ms) {
			return new Promise(resolve => setTimeout(resolve, ms))
		}
		async function fullLoad() {
			await sleep(3000)
			for (let i = 0; i < 20; i++) {
				window.scrollTo(0, Math.floor(document.body.scrollHeight / 20) * i + 1)
				await sleep(200)
			}
			let divShowWaitReady = document.createElement("div")
			divShowWaitReady.setAttribute("id", "show_wait_ready")
			divShowWaitReady.appendChild(document.createTextNode("."))
			document.body.appendChild(divShowWaitReady)
		}
		fullLoad()
	`

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var html string

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(
			func(ctxJS context.Context) error {
				_, exp, err := runtime.Evaluate(scriptFulllLoad).Do(ctxJS)
				if err != nil {
					return err
				}
				if exp != nil {
					return exp
				}

				return nil
			}),
		chromedp.WaitVisible(`#show_wait_ready`),
		chromedp.OuterHTML(`html`, &html),
	)
	if err != nil {
		return "", err
	}

	return html, nil
}
