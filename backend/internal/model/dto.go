package model

// --- Auth ---

type CustomerLoginRequest struct {
	StoreID     string `json:"storeId" binding:"required"`
	TableNumber int    `json:"tableNumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type AdminLoginRequest struct {
	StoreID  string `json:"storeId" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Role         string `json:"role"`
	StoreID      string `json:"storeId"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

// --- Menu ---

type MenuCreateRequest struct {
	CategoryID  uint   `json:"categoryId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Price       uint   `json:"price" binding:"required"`
	ImageURL    string `json:"imageUrl"`
	IsAvailable *bool  `json:"isAvailable"`
}

type MenuUpdateRequest struct {
	CategoryID  *uint   `json:"categoryId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Price       *uint   `json:"price"`
	ImageURL    *string `json:"imageUrl"`
	IsAvailable *bool   `json:"isAvailable"`
}

type MenuOrderRequest struct {
	ID        uint `json:"id" binding:"required"`
	SortOrder int  `json:"sortOrder" binding:"required"`
}

// --- Order ---

type OrderCreateRequest struct {
	Items []OrderItemRequest `json:"items" binding:"required,min=1"`
}

type OrderItemRequest struct {
	MenuID   uint `json:"menuId" binding:"required"`
	Quantity uint `json:"quantity" binding:"required,min=1"`
}

type StatusUpdateRequest struct {
	Status string `json:"status" binding:"required"`
}

// --- Table ---

type TableSetupRequest struct {
	TableNumber int    `json:"tableNumber" binding:"required"`
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
