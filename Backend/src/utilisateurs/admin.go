package utilisateurs

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	User        Utilisateur
	Secretaires []Secretaire
}

func (a *Admin) InscrireSecretaire(Nom string, Prenom string, Email string, Password string) Secretaire {

	// Generate a hash of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating password hash:", err)
	}
	user := Utilisateur{
		Id:         uuid.New().String(),
		Nom:        Nom,
		Prenom:     Prenom,
		Email:      Email,
		MotDePasse: string(hashedPassword),
	}

	secretaire := Secretaire{
		User:       user,
		Evenements: []Evenement{},
		Etudiants:  []Etudiant{},
		Moniteurs:  []Moniteur{},
	}
	a.Secretaires = append(a.Secretaires, secretaire)
	return secretaire

}
