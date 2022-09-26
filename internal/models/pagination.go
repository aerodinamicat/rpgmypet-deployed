package models

type Pagination struct {
	OrderBy        string `json:"orderBy"`
	FilterBySpecie string `json:"specie,omitempty"`

	PageSize   string `json:"pageSize"`
	PageToken  int    `json:"pageToken"`
	TotalPages string `json:"totalPages"`
	TotalItems string `json:"totalItems"`
}
