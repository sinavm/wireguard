package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type cfg struct {
	PrivateKey string
	PublicKey  string
	Endpoint   string
	DNS        string
}

func genKey() (priv, pub string, err error) {
	b := make([]byte, 32)
	if _, err = rand.Read(b); err != nil {
		return
	}
	priv = base64.StdEncoding.EncodeToString(b)

	// برای نمونه، public key را هم random می‌کنیم (در عمل از PureVPN می‌گیرید)
	if _, err = rand.Read(b); err != nil {
		return
	}
	pub = base64.StdEncoding.EncodeToString(b)
	return
}

func serverCfg(loc string) cfg {
	// TODO: اینجا API PureVPN را با credentials صدا کنید
	m := map[string]struct{ ep, dns string }{
		"uae - dubai":          {"uae-dubai.purevpn.net:51820", "8.8.8.8"},
		"uk - london":          {"uk-london.purevpn.net:51820", "8.8.4.4"},
		"germany - frankfurt":  {"de-frankfurt.purevpn.net:51820", "1.1.1.1"},
		"netherlands - amsterdam": {"nl-amsterdam.purevpn.net:51820", "1.0.0.1"},
		"turkey - istanbul":    {"tr-istanbul.purevpn.net:51820", "8.8.8.8"},
	}
	l := strings.ToLower(loc)
	if v, ok := m[l]; ok {
		return cfg{Endpoint: v.ep, DNS: v.dns}
	}
	return cfg{Endpoint: "default.purevpn.net:51820", DNS: "8.8.8.8"}
}

func main() {
	loc := os.Getenv("INPUT_LOCATION")
	if loc == "" {
		log.Fatal("INPUT_LOCATION not set")
	}

	// PureVPN credentials (فقط برای API واقعی)
	_ = os.Getenv("PUREVPN_USERNAME")
	_ = os.Getenv("PUREVPN_PASSWORD")
	_ = os.Getenv("PUREVPN_SUB_USER")

	priv, pub, err := genKey()
	if err != nil {
		log.Fatal(err)
	}
	c := serverCfg(loc)

	fmt.Println("=== WireGuard Config ===")
	fmt.Printf("[Interface]\nPrivateKey = %s\nAddress = 10.0.0.2/32\nDNS = %s\n\n", priv, c.DNS)
	fmt.Printf("[Peer]\nPublicKey = %s\nEndpoint = %s\nAllowedIPs = 0.0.0.0/0\nPersistentKeepalive = 25\n", pub, c.Endpoint)
	fmt.Println("=== End of Config ===")
	fmt.Printf("Config valid until %s\n", time.Now().Add(30*time.Minute).Format(time.RFC1123))
}
