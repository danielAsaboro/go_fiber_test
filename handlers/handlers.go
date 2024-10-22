package handlers

import (
	"fiber_sample/data"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ReadData handles reading all data
func ReadData(ctx *fiber.Ctx) error {
	dataService := data.InitData()  // Initialize your data service
	return ctx.JSON(dataService.GetData())  // Return JSON response
}

// DeleteData handles deletingx a record by ID
func DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))  // Get the ID from URL params
	dataService := data.InitData()
	return ctx.JSON(dataService.DeleteData(id))  // Return the response after deletion
}

// InsertData handles inserting new data
func InsertData(ctx *fiber.Ctx) error {
	dataService := data.InitData()

	// Use the correct data.UserModel type
	user := new(data.UserModel)  // Correctly refer to the model from the data package
	fmt.Println([]byte(ctx.Body()))
	if err := ctx.BodyParser(user); err != nil {
		return err  // Handle body parsing error
	}

	// Insert the parsed user data
	users := dataService.InsertData(data.UserModel{
		Name:   user.Name,
		Gender: user.Gender,
		Age:    user.Age,
	})

	return ctx.JSON(users)  // Return JSON response
}

// ReadDataById handles fetching a record by ID
func ReadDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	dataService := data.InitData()
	return ctx.JSON(dataService.GetDataById(id))
}

// PatchData handles updating a record by ID
func PatchData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	dataService := data.InitData()

	// Use the correct data.UserModel type
	user := new(data.UserModel)

	if err := ctx.BodyParser(user); err != nil {
		return err  // Handle body parsing error
	}

	// Update the record by ID
	users := dataService.UpdateDataById(
		id,
		data.UserModel{
			Name:   user.Name,
			Gender: user.Gender,
			Age:    user.Age,
		},
	)

	return ctx.JSON(users)  // Return JSON response
}
