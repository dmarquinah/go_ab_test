# Bayesian A/B Testing
This repository is a Golang implementation of a Bayesian A/B testing showcase.

## Data Creation
In order to make the test, we have to create a CSV file simulating multiple entries from a use case.
For this case, we want to decide which treatment (A or B) is better given the business data and conditions.
We define our `treatments` as different message templates called Highly Structured Message or HSM.

## Generate the CSV file
The main file will attempt to create the `output.csv` file based on some logic to create them.
By default we obtain `10000` rows of data, but that number can be modified by using the `-size=number` flag when executing it.
We can obtain the file by either:

* Build and run
```console
go build
./go_ab_test -size=1000
```

* Run the main file
```console
go run main.go
```
