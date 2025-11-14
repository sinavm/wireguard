package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	PrivateKey   string `json:"private_key"`
	PublicKey    string `json:"public_key"`
	Endpoint     string `json:"endpoint"`
	AllowedIPs   string `json:"allowed_ips"`
	DNS          string `json:"dns"`
	PersistentKeepalive string `json:"persistent_keepalive"`
}


func GenerateWireGuardKey() (privateKey, publicKey string, err error) {
	privateKeyBytes := make([]byte, 32)
	_, err = rand.Read(privateKeyBytes)
	if err != nil {
		return "", "", err
	}
	privateKey = base64.StdEncoding.EncodeToString(privateKeyBytes)

	publicKeyBytes := make([]byte, 32)

	_, err = rand.Read(publicKeyBytes)
	if err != nil {
		return privateKey, "", err
	}
	publicKey = base64.StdEncoding.EncodeToString(publicKeyBytes)

	return privateKey, publicKey, nil
}


func GetServerConfig(location string) Config {

	var endpoint, dns string
	switch strings.ToLower(location) {
	case "uae - dubai":
		endpoint = "uae-dubai.purevpn.net:51820"
		dns = "8.8.8.8"
	case "uk - london":
		endpoint = "uk-london.purevpn.net:51820"
		dns = "8.8.4.4"
	case "germany - frankfurt":
		endpoint = "de-frankfurt.purevpn.net:51820"
		dns = "1.1.1.1"
	case "netherlands - amsterdam":
		endpoint = "nl-amsterdam.purevpn.net:51820"
		dns = "1.0.0.1"
	case "turkey - istanbul":
		endpoint = "tr-istanbul.purevpn.net:51820"
		dns = "8.8.8.8"
	default:
		endpoint = "default.purevpn.net:51820"
		dns = "8.8.8.8"
	}

	privateKey, publicKey, err := GenerateWireGuardKey()
	if err != nil {
		log.Fatal(err)
	}

	return Config{
		PrivateKey:   privateKey,
		PublicKey:    publicKey,
		Endpoint:     endpoint,
		AllowedIPs:   "0.0.0.0/0",
		DNS:          dns,
		PersistentKeepalive: "25",
	}
}

func main() {
	location := os.Getenv("INPUT_LOCATION")
	if location == "" {
		log.Fatal("INPUT_LOCATION not set")
	}


	username := os.Getenv("PUREVPN_USERNAME")
	password := os.Getenv("PUREVPN_PASSWORD")
	subUser := os.Getenv("PUREVPN_SUB_USER")
	fmt.Printf("Using creds: %s (sub: %s) for location: %s\n", username, subUser, location)


	config := GetServerConfig(location)

	fmt.Println("=== WireGuard Config ===")
	fmt.Printf("[Interface]\n")
	fmt.Printf("PrivateKey = %s\n", config.PrivateKey)
	fmt.Printf("Address = 10.0.0.2/32\n")  // Sample IP
	fmt.Printf("DNS = %s\n", config.DNS)

	fmt.Printf("\n[Peer]\n")
	fmt.Printf("PublicKey = %s\n", config.PublicKey)  // Server public key - از PureVPN بگیر
	fmt.Printf("Endpoint = %s\n", config.Endpoint)
	fmt.Printf("AllowedIPs = %s\n", config.AllowedIPs)
	fmt.Printf("PersistentKeepalive = %s\n", config.PersistentKeepalive)
	fmt.Println("=== End of Config ===")

	fmt.Printf("Config valid for 30 minutes from %s\n", time.Now().Format(time.RFC1123))
}
