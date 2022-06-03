package handlers

type AddRequest struct {
	NumOne float64 `json:"num_one"` // encode-decode JSON using as keys this name tags
	NumTwo float64 `json:"num_two"`
}

type MathopsRequest struct {
	NumOne float64 `json:"num_one" validate:"required"` // I use tags for validate this Structure with Validator Package
	NumTwo float64 `json:"num_two" validate:"required"`
}

type SqrtRequest struct {
	NumOne float64 `json:"num_one" validate:"required"` // I use tags for validate this Structure with Validator Package
}
