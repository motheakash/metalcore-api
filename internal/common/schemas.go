package common

// PaginationRequest represents the query parameters for pagination
type PaginationRequest struct {
	Page     int `form:"page" binding:"omitempty,min=1"`      // Current page number (default: 1)
	PageSize int `form:"page_size" binding:"omitempty,min=1"` // Items per page (default: 10)
}

// GetPage returns the page number with a default value of 1
func (p *PaginationRequest) GetPage() int {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}

// GetPageSize returns the page size with a default value of 10
// Also enforces a maximum of 100 items per page
func (p *PaginationRequest) GetPageSize() int {
	if p.PageSize <= 0 {
		return 10
	}
	if p.PageSize > 100 {
		return 100
	}
	return p.PageSize
}

// GetOffset returns the offset for database queries
func (p *PaginationRequest) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPageSize()
}

// GetLimit returns the limit for database queries
func (p *PaginationRequest) GetLimit() int {
	return p.GetPageSize()
}

// PaginationMetadata represents pagination information in the response
type PaginationMetadata struct {
	Page       int   `json:"page"`        // Current page number
	PageSize   int   `json:"page_size"`   // Items per page
	TotalItems int64 `json:"total_items"` // Total number of items across all pages
	TotalPages int   `json:"total_pages"` // Total number of pages
	HasNext    bool  `json:"has_next"`    // Whether there is a next page
	HasPrev    bool  `json:"has_prev"`    // Whether there is a previous page
}

// NewPaginationMetadata creates pagination metadata from request and total count
func NewPaginationMetadata(req *PaginationRequest, totalItems int64) *PaginationMetadata {
	page := req.GetPage()
	pageSize := req.GetPageSize()
	totalPages := int((totalItems + int64(pageSize) - 1) / int64(pageSize))

	return &PaginationMetadata{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}
}

// PaginatedResponse represents a generic paginated response wrapper
type PaginatedResponse struct {
	Data       interface{}         `json:"data"`
	Pagination *PaginationMetadata `json:"pagination"`
}

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Status  int               `json:"status"`
	Error   string            `json:"error"`
	Message string            `json:"message,omitempty"`
	Details map[string]string `json:"details,omitempty"`
}

// SuccessResponse represents a standard success response
type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}
