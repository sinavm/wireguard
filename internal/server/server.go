package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/chromedp/chromedp"
)

type Subscription struct {
	ID      string `json:"id"`
	Username string `json:"username"`
}

func FetchSubscription(ctx context.Context) (string, error) {
	var subs []Subscription
	err := chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			var js string
			chromedp.Evaluate(`JSON.stringify(Array.from(document.querySelectorAll('.subscription-item')).map(item => ({id: item.dataset.id, username: item.textContent.trim()})))`, &js).Do(ctx)
			var data interface{}
			json.Unmarshal([]byte(js), &data)
			json.Unmarshal([]byte(js), &subs)
			return nil
		}),
	)
	if err != nil {
		return "", err
	}
	if len(subs) == 0 {
		return "", fmt.Errorf("no subscriptions found")
	}
	return subs[0].Username, nil
}

func Select(ctx context.Context, country, city string) error {
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://my.purevpn.com/v2/manual-wireguard"),
		chromedp.WaitVisible(`select[name="country"]`, chromedp.ByQuery),
		chromedp.SetValue(`select[name="country"]`, country, chromedp.ByQuery),
		chromedp.SetValue(`select[name="city"]`, city, chromedp.ByQuery),
		chromedp.Click(`button[id="select-server"]`, chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
	)
	return err
}