package input

type( 
	TaskInput struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
	TaskIdInput struct{
		Id     string `json:"id"`
	}
)