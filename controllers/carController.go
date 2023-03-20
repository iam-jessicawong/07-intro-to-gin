package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car

	err := ctx.ShouldBindJSON(&newCar) //cara untuk mendapatkan input data dari request dalam bentuk json
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// untuk membuat car id dalam bentuk string jadinya c1 atau c2 atau c3
	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	// write response dalam bentuk json
	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false

	var updateCar Car

	err := ctx.ShouldBindJSON(&updateCar)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updateCar.CarID = carID

	// cek id apakah ada dalam list data
	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			CarDatas[i] = updateCar
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "data not found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully updated", carID),
	})
}

func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false

	var carData Car

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "data not found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})

}
func GetCars(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"cars": CarDatas,
	})

}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false

	var carIndex int

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "data not found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", carID),
	})

}
