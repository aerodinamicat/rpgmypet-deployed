package models

import "time"

// swagger:model
type Pet struct {
	// Nombre de la mascota
	// required: true
	Name string `json:"name"`

	// Denominación de la especie de la mascota
	// required: true
	Specie string `json:"specie"`

	// Sexo de la mascota
	// required: true
	Sex string `json:"sex"`

	// Fecha de nacimiento de la mascota
	// required: true
	Birthdate time.Time `json:"birthdate"`

	// Id de la mascota
	// required: true
	Id string `json:"id"`
}

func (p *Pet) GetAgeInDays() float64 {
	return float64(time.Since(p.Birthdate).Hours() / 24)
}
