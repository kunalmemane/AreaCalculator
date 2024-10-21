[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kunalmemane/AreaCalculator)](https://github.com/kunalmemane/AreaCalculator/blob/main/go.mod) [![license](https://img.shields.io/github/license/kunalmemane/AreaCalculator)](https://github.com/kunalmemane/AreaCalculator/blob/main/LICENSE) ![GitHub Repo stars](https://img.shields.io/github/stars/kunalmemane/AreaCalculator?style=flat&logo=github)


# Area Calculator

A simple Application Programming interface  (API) written in Go to calculate the area and perimeter of various geometric shapes.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Supported Shapes](#supported-shapes)
- [Usage](#usage)
- [Contribute](#contribute)
- [Test](#test)
- [License](#license)

## Features

- Calculate the area of common shapes: Circle, Rectangle, Triangle, and Square.
- Fast and efficient calculations.
- Formatted response
- Efficient error handling

## Installation

To get started with the Area Calculator, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/kunalmemane/AreaCalculator.git
   cd area-calculator
    ```
2. Build the application:
    ``` bash
    go build -o ./bin/main ./cmd/main.go
    ```
3. Run the application:

    ```bash
    ./bin/main
    ```

The API will be available at http://localhost:8080/


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
## Contribute
Contributions are welcome! Please feel free to submit a Pull Request or open an Issue for discussion.

1. Fork the repository.
2. Create your feature branch (git checkout -b feature/AmazingFeature).
3. Commit your changes (git commit -m 'Add some AmazingFeature').
4. Push to the branch (git push origin feature/AmazingFeature).
5. Open a Pull Request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.