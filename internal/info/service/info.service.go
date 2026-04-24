package service

import (
	"project/module/internal/info/model"
)


type JobService struct{}

func (s *JobService) GetTable() model.Response {
	return model.Response{Status: "OK", Data: "Tabledata"}
}

func (s *JobService) GetCardData() model.Response {
	return model.Response{Status: "OK", Data: "Carddata"}
}

func (s *JobService) GetExportData() model.Response {
	return model.Response{Status: "OK", Data: "Carddata"}
}

func (s *JobService) CreateJob(dto model.CreateJobDTO) model.Response {
	return model.Response{Status: "OK", Data: dto}
}

func (s *JobService) EditJob(id int, dto model.UpdateJobDTO) model.Response {
	return model.Response{Status: "OK", Data: dto}
}

func (s *JobService) DeleteJob(dto model.DeleteJobDTO) model.Response {
	return model.Response{Status: "OK"}
}