package cmd

import (
	"fmt"

	"github.com/mathsant/go-arch-hexagonal/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var productPrice float64
var productQuantity int

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Products CRUD",
	Long:  `Just a CRUD of products using CLI and Arch Hexagonal`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(&productService, action, productId, productName, productPrice, productQuantity)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "Enable/Disable a product")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&productName, "name", "n", "", "Product Name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Product Price")
	cliCmd.Flags().IntVarP(&productQuantity, "quantity", "q", 0, "Product Quantity")
}
