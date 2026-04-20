package model

type CreateJobDTO struct {
	Name string `json:"name"`
}

type UpdateJobDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DeleteJobDTO struct {
	ID int `json:"id"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}