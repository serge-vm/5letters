package models

type Unordered struct {
	P int    `example:"1"`
	L string `example:"а"`
}

type SolverRequest struct {
	Ordered   map[int]string `json:"ordered" example:"3:о,5:е"`
	Unordered []Unordered    `json:"unordered"`
	Absent    []string       `json:"absent" example:"в,г"`
}

type SolverResponse []string
