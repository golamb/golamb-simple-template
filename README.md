# Go Lambda Simple Template for GoLambda

## Require
* [dep](https://github.com/golang/dep)
* [aws-lambda-go](https://github.com/aws/aws-lambda-go/)
## First Step
* Create function in AWS Lambda
* Choose Runtime to Go 1.x
* Create role or choose existing role

## Atfer create
* edit Handler to "main"
* save and deploy

## Dev
```sh
dep ensure
```

## But now not yet run lambda in locally. My way just normal run func Handle in main.