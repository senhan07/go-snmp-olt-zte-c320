package pagination

import (
	"net/http"
	"strconv"
)

// Constants for default and maximum page sizes, and query parameter names
var (
	DefaultPageSize = 10
	MaxPageSize     = 100
	PageVar         = "page"
	PageSizeVar     = "limit"
)

// Pages struct defines the structure for paginated responses
type Pages struct {
	Code      int32       `json:"code"`
	Status    string      `json:"status"`
	Page      int         `json:"page"`
	PageSize  int         `json:"limit"`
	PageCount int         `json:"page_count"`
	TotalRows int         `json:"total_rows"`
	Data      interface{} `json:"data"`
}

// New creates a new Pages instance with the provided parameters
func New(page, pageSize, total int) *Pages {
	if page <= 0 {
		page = 0
	}
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	pageCount := -1
	if total >= 0 {
		pageCount = (total + pageSize - 1) / pageSize
	}
	return &Pages{
		Code:      200,
		Status:    "OK",
		Page:      page,
		PageSize:  pageSize,
		TotalRows: total,
		PageCount: pageCount,
	}
}

// GetPaginationParametersFromRequest extracts pagination parameters from the HTTP request
func GetPaginationParametersFromRequest(r *http.Request) (pageIndex, pageSize int) {
	pageIndex = parseInt(r.URL.Query().Get(PageVar), 1)
	pageSize = parseInt(r.URL.Query().Get(PageSizeVar), DefaultPageSize)
	return pageIndex, pageSize
}

// parseInt is a helper function to parse string to int with a default value
func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}
