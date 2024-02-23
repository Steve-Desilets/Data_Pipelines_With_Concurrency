package main

import (
	// Import necessary packages
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	resizePackage "github.com/nfnt/resize"
)

type Job struct {
	// Create job struct to hold input and output paths
	InputPath string
	Image     image.Image
	OutPath   string
}

// Create InitLogger and LogInfo functions to log messages to a log file
// InitLogger initializes the logger with the given log file path.
var logger *log.Logger

func InitLogger(logFilePath string) {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo writes an information log message to the logger.
func LogInfo(message string) {
	logger.Println(message)
}

func loadImage(paths []string) <-chan Job {
	// Load images from the given paths into Go channels

	//Initialize logger
	InitLogger("logOutput.txt")

	out := make(chan Job)
	go func() {
		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			startTimeLoad := time.Now()

			job := Job{InputPath: p,
				OutPath: strings.Replace(p, "images/", "images/output/", 1)}
			job.Image = ReadImage(p)
			out <- job

			// Log the time taken to load the image
			endTimeLoad := time.Now()
			elapsedTimeLoad := endTimeLoad.Sub(startTimeLoad)
			LogInfo("Image loaded in: " + elapsedTimeLoad.String() + " for path: " + p)
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	// Resize images from the input channel and add them to the output channel

	//Initialize logger
	InitLogger("logOutput.txt")

	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			startTimeResize := time.Now()

			job.Image = Resize(job.Image)
			out <- job

			// Log the time taken to resize the image
			endTimeResize := time.Now()
			elapsedTimeResize := endTimeResize.Sub(startTimeResize)
			LogInfo("Image resized in: " + elapsedTimeResize.String())
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	// Convert images to grayscale from the input channel and add them to the output channel

	//Initialize logger
	InitLogger("logOutput.txt")

	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			startTimeGreyscale := time.Now()

			job.Image = Grayscale(job.Image)
			out <- job

			// Log the time taken to convert the image to greyscale
			endTimeGreyscale := time.Now()
			elapsedTimeGreyscale := endTimeGreyscale.Sub(startTimeGreyscale)
			LogInfo("Image converted to greyscale in: " + elapsedTimeGreyscale.String())
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	// Save images from the input channel and add a boolean to the output channel

	//Initialize logger
	InitLogger("logOutput.txt")

	out := make(chan bool)
	go func() {
		for job := range input { // Read from the channel
			startTimeSave := time.Now()

			WriteImage(job.OutPath, job.Image)
			out <- true

			// Log the time taken to save the image
			endTimeSave := time.Now()
			elapsedTimeSave := endTimeSave.Sub(startTimeSave)
			LogInfo("Image loaded in: " + elapsedTimeSave.String())
		}
		close(out)
	}()
	return out
}

func main() {
	// Main function to run the image processing pipeline

	//Initialize logger
	InitLogger("logOutput.txt")

	// Log that the application has started
	LogInfo("Application has started")

	// Starting software profiling
	var memStatsBefore, memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	// List of image paths
	imagePaths := []string{"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	//Create text file for program outputs
	outputFile, err := os.Create("dataPipelinesWithConcurrencyOutput.txt")
	if err != nil {
		LogInfo("Error creating output file")
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Start the benchmarking timer
	startTime := time.Now()

	channel1 := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	// Stop the benchmarking timer and calculate the elapsed time
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	elapsedMicroseconds := elapsedTime.Microseconds()

	// Print the elapsed time
	fmt.Fprintf(outputFile, "Total Pipeline Throughput Time: %d %s\n", elapsedMicroseconds, " microseconds")

	// Log whether the application completed successfully or failed
	for success := range writeResults {
		if success {
			LogInfo("Success! Image processing completed.")
		} else {
			LogInfo("Failed! Image processing failed.")
		}
	}

	// Calculate the memory usage and write memory usage information output file
	runtime.ReadMemStats(&memStatsAfter)
	memUsed := memStatsAfter.TotalAlloc - memStatsBefore.TotalAlloc

	fmt.Fprintf(outputFile, "Total Memory Used: %d %s\n", memUsed, " bytes")
}

func ReadImage(path string) image.Image {
	// Open the image file and decode it

	// Open the image file
	inputFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println(path)
		panic(err)
	}
	return img
}

func WriteImage(path string, img image.Image) {
	// Create a new file and encode the image to the file

	// Create a new file
	outputFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Encode the image to the new file
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		panic(err)
	}
}

func Grayscale(img image.Image) image.Image {
	// Convert an image to grayscale

	// Create a new grayscale image
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	// Convert each pixel to grayscale
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg
}

func Resize(img image.Image) image.Image {
	// Resize an image to a new width and height
	newWidth := uint(500)
	newHeight := uint(500)
	resizedImg := resizePackage.Resize(newWidth, newHeight, img, resizePackage.Lanczos3)
	return resizedImg
}
