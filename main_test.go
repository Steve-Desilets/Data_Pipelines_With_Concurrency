package main

import (
	// Import necessary packages
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"testing"
)

func logTestingMessage(message string) {
	// Function to log messages to a log file
	fmt.Println(message)

	logfile, err := os.OpenFile("logTestingOutput.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()

	log.SetOutput(logfile)
	log.Println(message)
}

func TestReadAndWriteImage(t *testing.T) {
	// Function to test the ReadImage and WriteImage functions

	// Test image file path
	testImagePath := "images/test_image.jpeg"
	// Test output image path
	testOutputPath := "images/test_output.jpeg"

	// Create a test image
	testImage := createTestImage()
	// Write the test image to file
	WriteImage(testImagePath, testImage)

	// Read the image from file
	readImage := ReadImage(testImagePath)

	// Compare the dimensions of the original and read images
	if testImage.Bounds().Dx() != readImage.Bounds().Dx() || testImage.Bounds().Dy() != readImage.Bounds().Dy() {
		t.Errorf("Dimensions of read image do not match original image")
		logTestingMessage("Error Reading Image: Dimensions of read image do not match original image")
	} else {
		logTestingMessage("Dimensions of read image correctly match original image")
	}

	// Remove the test output file after the test is complete
	defer func() {
		err := os.Remove(testImagePath)
		if err != nil {
			t.Errorf("Error removing test image file: %v", err)
			logTestingMessage("Error removing test image file")
		} else {
			logTestingMessage("Test image file removed successfully")
		}
	}()

	// Test writing the image to file
	WriteImage(testOutputPath, testImage)

	// Check if the output file exists
	_, err := os.Stat(testOutputPath)
	if err != nil {
		t.Errorf("Error writing image to file: %v", err)
		logTestingMessage("Error writing image to file")
	} else {
		logTestingMessage("Image written to file successfully")
	}
}

func createTestImage() image.Image {
	// Helper function to create a test image

	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// Fill the image with a color
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			img.Set(x, y, image.NewUniform(color.RGBA{255, 0, 0, 255}))
		}
	}
	return img
}

func TestResize(t *testing.T) {
	// Function to unit test the Resize function

	// Create a test image with known dimensions
	testImage := createTestImage() // Width: 100, Height: 50

	// Resize the test image
	resizedImage := Resize(testImage)

	// Check if the resized image dimensions are correct
	expectedWidth := uint(500)
	expectedHeight := uint(500)
	if resizedImage.Bounds().Dx() != int(expectedWidth) || resizedImage.Bounds().Dy() != int(expectedHeight) {
		t.Errorf("Expected resized image dimensions: %dx%d, got %dx%d",
			expectedWidth, expectedHeight, resizedImage.Bounds().Dx(), resizedImage.Bounds().Dy())
		logTestingMessage("Error resizing image: Resized image dimensions do not match expected dimensions")
	} else {
		logTestingMessage("Image resized successfully")
	}
}

func TestGrayscale(t *testing.T) {
	// Function to unit test the Grayscale function

	// Create a test image with known dimensions
	testImage := createTestImage() // Width: 100, Height: 100

	// Convert the test image to grayscale
	grayImage := Grayscale(testImage)

	// Check if the grayscale image dimensions are correct
	if grayImage.Bounds().Dx() != testImage.Bounds().Dx() || grayImage.Bounds().Dy() != testImage.Bounds().Dy() {
		t.Errorf("Grayscale image dimensions do not match original image")
		logTestingMessage("Error converting to grayscale: Grayscale image dimensions do not match original image")
	} else {
		logTestingMessage("Success: Grayscale image dimensions match original image dimensions")
	}
}
