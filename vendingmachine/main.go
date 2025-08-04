package main

import (
	"fmt"

	"github.com/monikalilhare01/GO-LLD/vendingmachine/models"
)

func main() {
	vm := models.NewVendingMachine()

	chips := models.NewProduct("Chips", 0.5)
	coke := models.NewProduct("Coke", 0.10)
	pepsi := models.NewProduct("pepsi", 0.20)

	vm.Inventory.AddProduct(chips, 10)
	vm.Inventory.AddProduct(coke, 10)
	vm.Inventory.AddProduct(pepsi, 10)

	// Demonstrate a transaction
	fmt.Println("Starting Vending Machine Demo")

	fmt.Println("\nSelecting Coke")
	vm.SelectProduct(coke)

	fmt.Println("Inserting coins")
	vm.InsertCoin(models.QUARTER)
	vm.InsertCoin(models.QUARTER)
	vm.InsertCoin(models.QUARTER)
	vm.InsertCoin(models.QUARTER)

	fmt.Println("Dispensing product")
	vm.DispenseProduct()

	fmt.Println("Returning change")
	vm.ReturnChange()

	// Another example with insufficient funds
	fmt.Println("\nSelecting Pepsi with insufficient funds")
	vm.SelectProduct(pepsi)
	vm.InsertCoin(models.QUARTER)

	fmt.Println("Trying to dispense Pepsi")
	vm.DispenseProduct()

	fmt.Println("Adding more coins for Pepsi")
	vm.InsertCoin(models.QUARTER)
	vm.InsertCoin(models.QUARTER)
	vm.InsertCoin(models.QUARTER)
	vm.InsertCoin(models.QUARTER)

	fmt.Println("Dispensing product")
	vm.DispenseProduct()

	fmt.Println("Returning change")
	vm.ReturnChange()

}
