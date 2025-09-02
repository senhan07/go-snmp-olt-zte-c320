package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/model"
)

func TestSendJSONResponse(t *testing.T) {
	// Initiate ResponseWriter dan Request
	rr := httptest.NewRecorder()

	response := model.OnuID{
		Board: 2,
		PON:   8,
		ID:    1,
	}

	// Call the SendJSONResponse function
	SendJSONResponse(rr, http.StatusOK, response)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code tidak sesuai: got %v want %v", status, http.StatusOK)
	}

	// Check the content type
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Content-Type tidak sesuai: got %v want %v", contentType, expectedContentType)
	}

	// Periksa Body Response
	var decodedResponse model.OnuID
	err := json.NewDecoder(rr.Body).Decode(&decodedResponse)
	if err != nil {
		t.Errorf("Gagal mendekode respons JSON: %v", err)
	}

	// Uji kasus di mana encoding JSON gagal
	// Inisialisasi ResponseWriter yang akan selalu gagal saat encoding JSON
	rrError := httptest.NewRecorder()
	// Sebagai contoh, gunakan objek yang tidak dapat di-encode sebagai respons
	errorResponse := make(chan int) // Ini akan gagal saat encoding JSON
	SendJSONResponse(rrError, http.StatusOK, errorResponse)

	// Periksa kode status respons
	if status := rrError.Code; status != http.StatusOK {
		t.Errorf("Status code tidak sesuai: got %v want %v", status, http.StatusOK)
	}

	// Periksa tipe konten
	expectedContentTypeError := "application/json"
	if contentType := rrError.Header().Get("Content-Type"); contentType != expectedContentTypeError {
		t.Errorf("Content-Type tidak sesuai: got %v want %v", contentType, expectedContentTypeError)
	}

	// Pastikan bahwa response body kosong karena encoding JSON gagal
	if body := rrError.Body.String(); body != "" {
		t.Errorf("Response body harus kosong jika encoding JSON gagal: got %v", body)
	}

}

func TestErrorBadRequest(t *testing.T) {
	rr := httptest.NewRecorder()
	err := errors.New("Bad Request Error")
	ErrorBadRequest(rr, err)

	// Periksa kode status respons
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code tidak sesuai: got %v want %v", status, http.StatusBadRequest)
	}

	// Periksa tipe konten
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Content-Type tidak sesuai: got %v want %v", contentType, expectedContentType)
	}

	// Periksa pesan kesalahan dalam respons JSON
	var response ErrorResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Gagal mendecode respons JSON: %v", err)
	}

	if response.Code != http.StatusBadRequest || response.Status != "Bad Request" || response.Message != err.Error() {
		t.Errorf("Respons JSON tidak sesuai")
	}
}

func TestErrorInternalServerError(t *testing.T) {
	rr := httptest.NewRecorder()
	err := errors.New("Internal Server Error")
	ErrorInternalServerError(rr, err)

	// Periksa kode status respons
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Status code tidak sesuai: got %v want %v", status, http.StatusInternalServerError)
	}

	// Periksa tipe konten
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Content-Type tidak sesuai: got %v want %v", contentType, expectedContentType)
	}

	// Periksa pesan kesalahan dalam respons JSON
	var response ErrorResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Gagal mendecode respons JSON: %v", err)
	}

	if response.Code != http.StatusInternalServerError || response.Status != "Internal Server Error" || response.Message != err.Error() {
		t.Errorf("Respons JSON tidak sesuai")
	}
}

func TestErrorNotFound(t *testing.T) {
	rr := httptest.NewRecorder()
	err := errors.New("Not Found Error")
	ErrorNotFound(rr, err)

	// Periksa kode status respons
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Status code tidak sesuai: got %v want %v", status, http.StatusNotFound)
	}

	// Periksa tipe konten
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Content-Type tidak sesuai: got %v want %v", contentType, expectedContentType)
	}

	// Periksa pesan kesalahan dalam respons JSON
	var response ErrorResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Gagal mendecode respons JSON: %v", err)
	}

	if response.Code != http.StatusNotFound || response.Status != "Not Found" || response.Message != err.Error() {
		t.Errorf("Respons JSON tidak sesuai")
	}
}
