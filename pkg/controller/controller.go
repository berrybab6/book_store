package controller

// Controller MovieGo
type Controller struct {
}

// NewController MovieGo
func NewController() *Controller {
	return &Controller{}
}

// Message MovieGo
type Message struct {
	Message string `json:"message" example:"message"`
}
