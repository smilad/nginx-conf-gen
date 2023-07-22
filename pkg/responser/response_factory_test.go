package responser

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestErrorResponse(t *testing.T) {
	// Test case: Creating an error response
	errorResp := NewErrorBuilder().
		SetStatusCode(http.StatusBadRequest).
		SetMessage("Invalid input").
		SetDetail("username", "Username cannot be empty").
		Build()

	expectedJSON := `{"status":400,"message":"Invalid input","details":[{"field":"username","description":"Username cannot be empty"}]}`
	jsonData, err := json.Marshal(errorResp)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(jsonData))

	// Test case: Responding to an Echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err = errorResp.Respond(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, "Invalid input", response.Message)
	assert.Len(t, response.Details, 1)
	assert.Equal(t, "username", response.Details[0].Field)
	assert.Equal(t, "Username cannot be empty", response.Details[0].Description)
}

func TestResponse(t *testing.T) {
	// Test case: Creating a success response
	data := map[string]string{"key": "value"}
	successResp := NewResponse(http.StatusOK, "Success", data)

	expectedJSON := `{"status":200,"message":"Success","data":{"key":"value"}}`
	jsonData, err := json.Marshal(successResp)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(jsonData))

}
