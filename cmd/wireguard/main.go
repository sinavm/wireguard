package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sinavm/wireguard/internal/app"
)

var rootCmd = &cobra.Command{
	Use:   "wireguard",
	Short: "PureVPN WireGuard fetcher",
}

var fullCmd = &cobra.Command{
	Use:   "full",
	Short: "Fetch and generate WireGuard config",
	Run: func(cmd *cobra.Command, args []string) {
		if err := app.Full(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("purevpn")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	rootCmd.AddCommand(fullCmd)
	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
