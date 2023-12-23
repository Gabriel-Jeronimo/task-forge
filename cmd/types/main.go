package types

type Task struct {
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	Done      bool   `json:"done"`
}