package handler

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/kunalmemane9150/AreaCalculator/internal/model"
	"github.com/kunalmemane9150/AreaCalculator/internal/service"
	"github.com/kunalmemane9150/AreaCalculator/internal/validator"
	log "github.com/kunalmemane9150/AreaCalculator/pkg/logger"
)

func GetAreaHandler(w http.ResponseWriter, r *http.Request) {

	//check method
	if r.Method != http.MethodPost {

		log.New().Error().Println("invalid method " + r.Method)

		response := model.ErrorResponse{
			Success:    false,
			Error:      "Method not allowed " + r.Method,
			StatusCode: http.StatusMethodNotAllowed,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer r.Body.Close()

	//decode and map body with request struct
	var req model.ShapeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		log.New().Error().Println(err)

		response := model.ErrorResponse{
			Success:    false,
			Error:      "Invalid request body",
			StatusCode: http.StatusBadRequest,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.New().Info().Println("Request: ", req)

	// build shape struct based on requested shape,
	var shape model.IShape
	var shapeDimensions model.ShapeDimensions
	switch strings.ToLower(req.Shape) {
	case "square":
		shape = model.NewSquare(req.Side)
		shapeDimensions = model.ShapeDimensions{
			Side: req.Side,
		}
	case "circle":
		shape = model.NewCircle(req.Radius)
		shapeDimensions = model.ShapeDimensions{
			Radius: req.Radius,
		}
	case "rectangle":
		shape = model.NewRectangle(req.Length, req.Breadth)
		shapeDimensions = model.ShapeDimensions{
			Length:  req.Length,
			Breadth: req.Breadth,
		}
	case "triangle":
		shape = model.NewTriangle(req.SideA, req.SideB, req.SideC)
		shapeDimensions = model.ShapeDimensions{
			SideA: req.SideA,
			SideB: req.SideB,
			SideC: req.SideC,
		}
	default:
		response := model.ErrorResponse{
			Success:    false,
			Error:      "Invalid Shape",
			StatusCode: http.StatusBadRequest,
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.New().Debug().Println("Before Validation: ", reflect.TypeOf(shape))

	// validate requested shape field values
	shape, err := validator.Validator(shape)
	if err != nil {

		response := model.ErrorResponse{
			Success:    false,
			Error:      err.Error(),
			StatusCode: http.StatusExpectationFailed,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.New().Debug().Println("After Validation: ", reflect.TypeOf(shape))

	// service - get area and perimeter
	area, perimeter := service.Calculate(shape)

	// Build response
	response := model.APIResponse{
		Success: true,
		Shape: model.ShapeResponse{
			Shape:      req.Shape,
			Dimensions: shapeDimensions,
			Area:       area,
			Perimeter:  perimeter,
		},

		Timestamp: time.Now().Format(time.RFC822),
	}

	//log response on server
	resp, _ := json.Marshal(response)
	log.New().Info().Println("Response: ", string(resp))

	//return response with relevant header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
