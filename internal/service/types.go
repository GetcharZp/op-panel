package service

type TaskDetailResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Spec string `json:"spec"`
	Data string `json:"data"`
}
