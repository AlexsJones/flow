# CDK in Golang

This project shows how you can write your CDK stacks in golang plus lambda.

![image](images/1.png)

## Structure

```
├── flow.go # <---- CDK stack
├── flow_test.go
├── go.mod
├── go.sum
├── images
│   └── 1.png
└── lambda  # <----- Golang lambda functions
    ├── consumer
    │   └── main.go 
    └── publisher
        └── main.go
```


## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests
