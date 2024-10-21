- 
    - ## Triangle

        - **Request Body**:
            ```json
            {
                "shape":"Triangle",
                "sideB":20,
                "sideA": 10,
                "sideC":15    
            }
            ```
        - **Response**:
            ```json
                {
                    "success": true,
                    "data": {
                        "shape": "Triangle",
                        "dimensions": {
                            "sideA": 10,
                            "sideB": 20,
                            "sideC": 15
                        },
                        "area": "72.618",
                        "perimeter": "45.000"
                    },
                    "timestamp": "21 Oct 24 12:16 IST"
                }
            ```
    
    - ## **Square**
        - Request Body:
            ```json
                {
                    "shape":"square",
                    "side": 20
                }
            ```

        - **Response**:
            ```json
                {
                    "success": true,
                    "data": {
                        "shape": "square",
                        "dimensions": {
                        "side": 20
                        },
                        "area": "400.000",
                        "perimeter": "80.000"
                    },
                    "timestamp": "21 Oct 24 15:24 IST"
                }
            ```
    - ## **Circle**
        - Request Body:
            ```json
                {
                    "shape":"circle",
                    "radius": 10
                }
            ```

        - **Response**:
            ```json
                {
                    "success": true,
                    "data": {
                        "shape": "circle",
                        "dimensions": {
                        "radius": 10
                        },
                        "area": "314.159",
                        "perimeter": "62.832"
                    },
                    "timestamp": "21 Oct 24 15:25 IST"
                }
            ```
    - ## **Rectangle**
        - Request Body:
            ```json
                {
                    "shape":"rectangle",
                    "length": 10,
                    "breadth":20
                }
            ```

        - **Response**:
            ```json
                {
                    "success": true,
                    "data": {
                        "shape": "rectangle",
                        "dimensions": {
                        "length": 10,
                        "breadth": 20
                        },
                        "area": "200.000",
                        "perimeter": "60.000"
                    },
                    "timestamp": "21 Oct 24 15:26 IST"
                }
            ```