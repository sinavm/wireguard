package wg

import (
	"context"
	"io/ioutil"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/k8s"
)

func Generate(ctx context.Context, device string) (string, error) {
	var config string
	err := chromedp.Run(ctx,
		chromedp.Click(`button[value="`+device+`"]`, chromedp.ByQuery),
		chromedp.Click(`button[id="generate-config"]`, chromedp.ByQuery),
		chromedp.WaitVisible(`#config-download`, chromedp.ByID),
		chromedp.AttributeValue(`#config-textarea`, "value", &config, nil, chromedp.ByID),
	)
	if err != nil {
		return "", err
	}
	return config, nil
}