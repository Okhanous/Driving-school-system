package utilisateurs

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Evenement qui peut être soit une leçon de conduite ou bien un examen
type Evenement struct {
	Id          int
	Date        time.Time
	Duree       int
	Description string
}

type Secretaire struct {
	User       Utilisateur
	Evenements []Evenement
	Etudiants  []Etudiant
	Moniteurs  []Moniteur
}

func (s *Secretaire) InscrireEtudiant(Nom string, Prenom string, Email string, Password string) Etudiant {
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

	etudiant := Etudiant{
		User:         user,
		PermisObtenu: false,
		Evenements:   []Evenement{},
	}

	s.Etudiants = append(s.Etudiants, etudiant)

	return etudiant

}

func (s *Secretaire) InscrireMoniteur(Nom string, Prenom string, Email string) Moniteur {
	password := "motdepasseMoniteur" // Replace with your desired password

	// Generate a hash of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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

	moniteur := Moniteur{
		User:       user,
		Evenements: []Evenement{},
		Etudiants:  []Etudiant{},
	}

	s.Moniteurs = append(s.Moniteurs, moniteur)
	return moniteur

}

func (s Secretaire) FixerEvenementEtudiant(evenement Evenement, e Etudiant) string {
	//l'opérateur := est utiliser pour assigner et déclarer une variable en même temps
	description := fmt.Sprintf("L'étudiant %s %s prend une leçon de %s le %s pour une durée de %d minutes",
		e.User.Prenom, e.User.Nom, evenement.Description, evenement.Date.Format("02/01/2006"), evenement.Duree)

	e.Evenements = append(e.Evenements, evenement)

	rand.Seed(time.Now().UnixNano())
	indexAleatoire := rand.Intn(len(s.Moniteurs))
	moniteurAleatoire := s.Moniteurs[indexAleatoire]

	fmt.Println(s.FixerEvenementMoniteur(evenement, moniteurAleatoire))
	return description
}

func (s Secretaire) FixerEvenementMoniteur(evenement Evenement, m Moniteur) string {
	//l'opérateur := est utiliser pour assigner et déclarer une variable en même temps
	description := fmt.Sprintf("Le Moniteur %s %s donne une leçon de %s le %s pour une durée de %d minutes",
		m.User.Prenom, m.User.Nom, evenement.Description, evenement.Date.Format("02/01/2006"), evenement.Duree)

	m.Evenements = append(m.Evenements, evenement)
	return description
}
