package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSettingGetAppearanceEndpoint(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1.Setting/GetAppearance", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
