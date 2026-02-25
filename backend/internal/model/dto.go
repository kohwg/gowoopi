package model

// --- Auth ---

type CustomerLoginRequest struct {
	StoreID     string `json:"store_id" binding:"required"`
	TableNumber int    `json:"table_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type AdminLoginRequest struct {
	StoreID  string `json:"store_id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Role         string `json:"role"`
	StoreID      string `json:"store_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

// --- Menu ---

type MenuCreateRequest struct {
	CategoryID  uint   `json:"category_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Price       uint   `json:"price" binding:"required"`
	ImageURL    string `json:"image_url"`
	IsAvailable *bool  `json:"is_available"`
}

type MenuUpdateRequest struct {
	CategoryID  *uint   `json:"category_id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Price       *uint   `json:"price"`
	ImageURL    *string `json:"image_url"`
	IsAvailable *bool   `json:"is_available"`
}

type MenuOrderRequest struct {
	ID        uint `json:"id" binding:"required"`
	SortOrder int  `json:"sort_order" binding:"required"`
}

// --- Order ---

type OrderCreateRequest struct {
	Items []OrderItemRequest `json:"items" binding:"required,min=1"`
}

type OrderItemRequest struct {
	MenuID   uint `json:"menu_id" binding:"required"`
	Quantity uint `json:"quantity" binding:"required,min=1"`
}

type StatusUpdateRequest struct {
	Status string `json:"status" binding:"required"`
}

// --- Table ---

type TableSetupRequest struct {
	TableNumber int    `json:"table_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

// --- Error ---

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
