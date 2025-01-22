[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kunalmemane/AreaCalculator?logo=go&color=%2300ADD8)](https://github.com/kunalmemane/AreaCalculator/blob/main/go.mod) [![license](https://img.shields.io/github/license/kunalmemane/AreaCalculator)](https://github.com/kunalmemane/AreaCalculator/blob/main/LICENSE) ![GitHub Repo stars](https://img.shields.io/github/stars/kunalmemane/AreaCalculator?style=flat&logo=github)  ![GitHub Issues or Pull Requests](https://img.shields.io/github/issues-pr/kunalmemane/AreaCalculator?logo=github)


# Area Calculator

The Area Calculator API provides a simple and efficient way to compute the area of various geometric shapes, including square, circles, rectangles, triangles, and more.

<img src="https://www.shutterstock.com/image-vector/perimeter-area-square-formula-mathematics-600nw-2325901341.jpg" style="width:100%">
credits: shutterstock.com

## Table of Contents

- [Features](#features)
- [Prerequisite](#prerequisite)
- [Installation](#installation)
- [Supported Shapes](#supported-shapes)
- [Usage](#usage)
- [Test](#test)
- [Makefile](#makefile)
- [Contribute](#contribute)
- [License](#license)

## Features

- Calculate the area of common shapes: Circle, Rectangle, Triangle, and Square.
- Fast and efficient calculations.
- Formatted response
- Efficient error handling

## Prerequisite
Before you begin, ensure you have met the following requirements:

- **Go**: You need to have Go **v1.21+** installed on your machine. Download it from the [official Go website](https://golang.org/dl/).
  
  To check if Go is installed, run:
  ```bash
  go version
    ```
## Installation

To get started with the Area Calculator, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/kunalmemane/AreaCalculator.git
   cd AreaCalculator
    ```
2. Build the application:
    ``` bash
    go build -o ./bin/main ./cmd/main.go
    ```
3. Run the application:

    ```bash
    ./bin/main
    ```

The API will be available at http://localhost:8080/getArea


## Supported Shapes

- **Circle:** Requires the radius.
- **Rectangle:** Requires the length and breadth.
- **Triangle:** Requires the sideA, sideB and sideC.
- **Square:** Requires the side length.

## Usage

Once the Application is running, you can make HTTP requests to calculate areas of supported shapes using the `/getArea` endpoint.

### Endpoint

    POST /getArea

 ### **[See Demo Request and Expected Responses](USAGE.md)**

### Example Request

To calculate the area of multiple shapes, use the following curl command:

- Triangle

    ```bash
    curl -X POST http://localhost:8080/getArea -d '{
            "shape":"Triangle",
            "sideB":20,
            "sideA": 10,
            "sideC":15    
        }' -H "Content-Type: application/json"
    ```

## Test

- To test the application run below commands

    ```bash
        make test

        or

        go test ./...
    ``` 

## Makefile
This project includes a `Makefile` to simplify common tasks. You can use the following commands:


### Available Commands

- **`make build`**: Compile the project.
- **`make run`**: Build and run the application.
- **`make test`**: Run the tests for the project.
- **`make clean`**: Remove build artifacts.
- **`make all`**: To do clean, test, build, run tasks in one command.
- **`make podman-build`**: Build Application Image using Podman.
- **`make podman-run`**: Run Application Container Image.
- **`make podman-push`**: Push Application Image to quay.io _**- podman login required.**_
- **`make oc-deploy`**: Deploy application to openshift using docker strategy _**- oc login required.**_
- **`make oc-delete`**: Delete all resources in openshift project _**- oc login required.**_

## Contribute
Contributions are welcome! Please feel free to submit a Pull Request or open an Issue for discussion.

1. Fork the repository.
2. Create your feature branch (git checkout -b feature/AmazingFeature).
3. Commit your changes (git commit -m 'Add some AmazingFeature').
4. Push to the branch (git push origin feature/AmazingFeature).
5. Open a Pull Request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
