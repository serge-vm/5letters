package models

type SolverRequest struct {
	Ordered   map[int]string `json:"ordered" example:"3:о,5:е"`
	Unordered map[int]string `json:"unordered" example:"1:а,2:а"`
	Absent    []string       `json:"absent" example:"в,г"`
}

type SolverResponse []string
