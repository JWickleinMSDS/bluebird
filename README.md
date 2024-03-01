***Instructions***

1) Begin by getting the bluebird package by executing the command "go get github.com/JWickleinMSDS/bluebird" on the terminal.  This package contains the package I created to calculate the trimmed mean of integers and floats.  

2) Execute the test designed for the bluebird package with the command "go test".  This will create a sample of 100 integers and floats and test the trimmed mean function against the base mean function in R. 

***Additional notes***

- The bluebird.go file contains the custom built package that calculates the trimmed mean of integers and floats. 
- The bluebird.R package receives command line arguments from the bluebird_test.go package and outputs the trimmed mean in R.
- The bluebird_test.go package is the heavy hitter.  It receives inputs from both the bluebird.go and bluebird.R package and compares the results of the two to see if they match.