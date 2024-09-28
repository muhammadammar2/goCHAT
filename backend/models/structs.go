package models

type SignupRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateRoomRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	RoomType    string `json:"room_type" validate:"required,oneof=public private"`
	RoomCode    string `json:"room_code,omitempty"`
}