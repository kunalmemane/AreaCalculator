package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/kunalmemane9150/AreaCalculator/internal/handler"
)

func TestGetAreaHandler(t *testing.T) {

	tests := []struct {
		name               string
		requestBody        string
		method             string
		expectedResponse   string
		expectedStatusCode int
	}{
		//Invalid method
		{name: "Invalid Method", method: "GET", requestBody: `{"shape":"Circle","radius":20}`, expectedResponse: `{"success":false,"message":"Method not allowed GET","statusCode":405}` + "\n", expectedStatusCode: http.StatusMethodNotAllowed},

		//Empty request body
		{name: "Empty request body", method: "POST", requestBody: ``, expectedResponse: `{"success":false,"message":"Invalid request body","statusCode":400}` + "\n", expectedStatusCode: http.StatusBadRequest},

		//Invalid shape
		{name: "Invalid Shape", requestBody: `{"shape":"Pentagon","radius":20}`, method: "POST", expectedResponse: `{"success":false,"message":"Invalid Shape","statusCode":400}` + "\n", expectedStatusCode: http.StatusBadRequest},

		//Invalid request body
		{name: "Invalid request body - Circle", method: "POST", requestBody: `{"shape":"Circle","radius":20`, expectedResponse: `{"success":false,"message":"Invalid request body","statusCode":400}` + "\n", expectedStatusCode: http.StatusBadRequest},

		//Circle
		{name: "Valid Request body - Circle", requestBody: `{"shape":"Circle","radius":20}`, method: "POST", expectedResponse: `{"success":true,"data":{"shape":"Circle","dimensions":{"radius":20},"area":"1256.637","perimeter":"125.664"},"timestamp":"` + time.Now().Format(time.RFC822) + `"}` + "\n", expectedStatusCode: http.StatusOK},

		{name: "Invalid Shape field - Circle", requestBody: `{"shape":"Circle","radius":-20}`, method: "POST", expectedResponse: `{"success":false,"message":"invalid circle radius","statusCode":417}` + "\n", expectedStatusCode: http.StatusExpectationFailed},

		//Square
		{name: "Valid Request body - Square", requestBody: `{"shape":"Square","side":10}`, method: "POST", expectedResponse: `{"success":true,"data":{"shape":"Square","dimensions":{"side":10},"area":"100.000","perimeter":"40.000"},"timestamp":"` + time.Now().Format(time.RFC822) + `"}` + "\n", expectedStatusCode: http.StatusOK},

		{name: "Invalid Shape field - Square", requestBody: `{"shape":"Square","side":-10}`, method: "POST", expectedResponse: `{"success":false,"message":"invalid square side","statusCode":417}` + "\n", expectedStatusCode: http.StatusExpectationFailed},

		//Rectangle
		{name: "Valid Request body - Rectangle", requestBody: `{"shape":"Rectangle","length":10,"breadth":20}`, method: "POST", expectedResponse: `{"success":true,"data":{"shape":"Rectangle","dimensions":{"length":10,"breadth":20},"area":"200.000","perimeter":"60.000"},"timestamp":"` + time.Now().Format(time.RFC822) + `"}` + "\n", expectedStatusCode: http.StatusOK},

		{name: "Invalid shape field - Rectangle", requestBody: `{"shape":"Rectangle","length":-10,"breadth":20}`, method: "POST", expectedResponse: `{"success":false,"message":"invalid rectangle params","statusCode":417}` + "\n", expectedStatusCode: http.StatusExpectationFailed},

		//Triangle
		{name: "Valid Request body - Triangle", requestBody: `{"shape":"Triangle","sideA":10,"sideB":15,"sideC":20}`, method: "POST", expectedResponse: `{"success":true,"data":{"shape":"Triangle","dimensions":{"sideA":10,"sideB":15,"sideC":20},"area":"72.618","perimeter":"45.000"},"timestamp":"` + time.Now().Format(time.RFC822) + `"}` + "\n", expectedStatusCode: http.StatusOK},

		{name: "Invalid shape field - Triangle", requestBody: `{"shape":"Triangle","sideA":-10,"sideB":15,"sideC":20}`, method: "POST", expectedResponse: `{"success":false,"message":"invalid triangle side","statusCode":417}` + "\n", expectedStatusCode: http.StatusExpectationFailed},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			body := strings.NewReader(tc.requestBody)

			req := httptest.NewRequest(tc.method, "/getArea", body)
			resp := httptest.NewRecorder()

			handler.GetAreaHandler(resp, req)

			result := resp.Result()
			defer req.Body.Close()

			data, err := io.ReadAll(result.Body)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			if resp.Result().StatusCode != tc.expectedStatusCode {
				t.Errorf("Expected Status Code %v, Got %v", tc.expectedStatusCode, resp.Result().StatusCode)
			}

			if string(data) != tc.expectedResponse {
				t.Errorf("Expected %v, Got %v", tc.expectedResponse, string(data))
			}

		})
	}
}
