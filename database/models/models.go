package models

import (
	"gorm.io/gorm"
)

type EtudiantModel struct {
	gorm.Model
	SiajeId int    `json:"id"`
	Nom     string `json:"nom"`
	Prenom  string `json:"prenom"`
}

func CreateEtudiant(etudiant EtudiantModel) {

}
