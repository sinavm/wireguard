package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chromedp/chromedp"
	"github.com/spf13/viper"
	"yourusername.com/wireguard-improved/internal/login"
	"yourusername.com/wireguard-improved/internal/server"
	"yourusername.com/wireguard-improved/internal/wg"
)

func Full() error {
	username := viper.GetString("username")
	password := viper.GetString("password")
	if username == "" || password == "" {
		return fmt.Errorf("username and password required")
	}

	ctx, cancel := chromedp.NewContext(
		chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	if err := login.Perform(ctx, username, password); err != nil {
		return err
	}

	subUsername := viper.GetString("subscription_username")
	if subUsername == "" {
		subUsername, err := server.FetchSubscription(ctx)
		if err != nil {
			return err
		}
		viper.Set("subscription_username", subUsername)
	}

	country := viper.GetString("server_country")
	city := viper.GetString("server_city")
	if country == "" || city == "" {
		return fmt.Errorf("server country and city required")
	}

	if err := server.Select(ctx, country, city); err != nil {
		return err
	}

	device := viper.GetString("device")
	if device == "" {
		device = "linux"
	}

	configPath := viper.GetString("wireguard_file")
	if configPath == "" {
		configPath = "wg0.conf"
	}

	config, err := wg.Generate(ctx, device)
	if err != nil {
		return err
	}

	if err := os.WriteFile(configPath, []byte(config), 0644); err != nil {
		return err
	}

	fmt.Printf("WireGuard config saved to %s\n", configPath)
	return nil
}