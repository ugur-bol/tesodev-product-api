package controllers

import (
	"context"
	"net/http"
	"tesodev-product-api/config"
	"tesodev-product-api/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ListProducts - Get all products
func ListProducts(c echo.Context) error {
	products := []models.Product{}
	collection := config.DB.Database("tesodev").Collection("products")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product models.Product
		cursor.Decode(&product)
		products = append(products, product)
	}

	return c.JSON(http.StatusOK, products)
}

// CreateProduct - Create a new product
func CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	collection := config.DB.Database("tesodev").Collection("products")
	result, err := collection.InsertOne(context.Background(), product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	product.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return c.JSON(http.StatusCreated, product)
}

// GetProductByID - Retrieve a product by ID
func GetProductByID(c echo.Context) error {
	id := c.Param("id")
	collection := config.DB.Database("tesodev").Collection("products")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var product models.Product
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, product)
}

// UpdateProduct - Update an existing product
func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	collection := config.DB.Database("tesodev").Collection("products")
	objectID, _ := primitive.ObjectIDFromHex(id)

	update := bson.M{"$set": bson.M{"name": product.Name, "price": product.Price}}
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Product updated successfully")
}

// DeleteProduct - Delete a product
func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	collection := config.DB.Database("tesodev").Collection("products")
	objectID, _ := primitive.ObjectIDFromHex(id)

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Product deleted successfully")
}
