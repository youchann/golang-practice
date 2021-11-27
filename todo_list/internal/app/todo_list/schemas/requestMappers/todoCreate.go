package requestMappers

type TodoCreate struct {
	Name   string `json:"name" validate:"required"`
	IsDone bool   `json:"is_done"`
}
