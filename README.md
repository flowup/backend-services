# FlowUp Services

FlowUp services is Golang project focused on reusable backend services. This project offers services that are repeatedly required in a new projects such as emailing, hashing passwords, zipping etc. Every service contains it's own README where **usage** of the package is **detailedly described**.

## List of the services

1. Age      - calculate age of the events
2. Email    - sending of the emails 
3. Encrypt  - password hashing and validating
4. File     - file service such as upload and download to the server
5. Time     - mock time especially useful in tests
6. Token    - token service
7. Tracker  - event tracking
8. Zip      - compressing and decompressing

## Getting Started

Following paragraphs shows how to install FlowUp Services. 

### Prerequisities

As long as this is Golang project there has to be Go installed.

### Installing

For installing the packages there can be import of the specific service package added (as shown in examples in README for the given service) Whole project with all the services can be added by:

```
import "flowdock.eu/flowup/services"
```
    
Then by writting `go get ./...` in the root folder of your project services should be imported.

## Running the tests

Every service has got tests for it's functionality and every method that services offers is tested.

All tests can be executed by running commad in the folder with test file: 

    goconvey 

Alternatively you can use command: 

    go test

## Authors

FlowUp develepers.

## License

