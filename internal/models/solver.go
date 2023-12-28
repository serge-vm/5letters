package models

type SolverRequest struct {
	Ordered   map[int]string `json:"ordered" example:"3:о,5:е"`
	Unordered []string       `json:"unordered" example:"а,б"`
	Absent    []string       `json:"absent" example:"в,г"`
}

type SolverResponse []string
