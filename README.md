# Data Pipelines With Concurrency

For the project contained within this repository, we aim to learn more about how to construct image processing data pipelines that leverage concurrency using Golang.  We begin by exploring the work of Amrit Singh, who developed such a pipeline that reads images into Golang, resizes the images to 500 x 500 pixels, converts the color images to grayscale, and then writes the new versions of the images to external files (Singh 2024).  We then aim to build upon this work by: 

1. Setting up error checking for file input and outupt
2. Replacing the four input images with new images
3. Adding unit tests to the code repository
4. Developing benchmarking methods for capturing pipeline throughput times
5. Modifying the code to add enhancements
   
By studying Singh's work and enhancing his program's original functionality, we hope to gain a deeper understanding of image processing data pipelines in Golang.

To address our first objective - setting up error checking for file input and output - we created the "main_test.go" file and wrote the function "TestReadAndWriteImage).  This function first creates a test image, reads it, and then confirms that its dimensions match those expected. Subsequently, this file tests the file output writing functionality by confirming whether the output file exists.  As we see in "logTestingOutput.txt", both of these tests passed.

To address the second objective - replacing the four input images with new images - we load four images of my mom's dog, Snoopy, into the images folder in this repository. As we can see in the "images/output" subfolder within this repository, our data processing pipeline successfully resizes these images, converts them to grayscale, and writes them to new output JPEG files.

To address the third objective - adding unit tests to the code repository - ___________________

To address the fourth objective - developing benchmarking methods for capturing pipeline throughput times - ___________________________________

To address the fifth objective - modifying the code to add enhancements - ________________________

By addressing each of our five project objectives, we successfully gained a deeper understanding of how to implement and enhance image processing data pipelines that leverage concurrency in Golang. 

References

Singh, Amrit. 2024. "Episode #21: Concurrency in Go: Pipeline Pattern". Github. https://github.com/code-heim/go_21_goroutines_pipeline
