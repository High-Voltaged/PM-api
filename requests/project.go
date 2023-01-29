package requests

type CreateProjectBody struct {
	Name        string `json:"name" binding:"required,min=6,max=16"`
	Description string `json:"description" binding:"min=20,max=100"`
	Start_at    string `json:"start_at" binding:"required"`
	End_at      string `json:"end_at" binding:"required"`
}
