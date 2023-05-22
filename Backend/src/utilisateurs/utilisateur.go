package utilisateurs

import (
	"fmt"
)

type Utilisateur struct {
	Id         string
	Nom        string
	Prenom     string
	Email      string
	MotDePasse string
}

//Définition des méthodes de l'étudiant sous la forme func (receiver Type) MethodName(arguments)

func (u Utilisateur) SeConnecter() string {
	return fmt.Sprintf("l'utilisateur : %s %s s'est connecté", u.Prenom, u.Nom)
}

func (u Utilisateur) SeDeconnecter() string {
	return fmt.Sprintf("l'utilisateur : %s %s s'est déconnecté", u.Prenom, u.Nom)
}

func (u *Utilisateur) ChangerMotDePasse(AncienMotDePasse string, NouveauMotDePasse string) string {
	if AncienMotDePasse == u.MotDePasse {
		u.MotDePasse = NouveauMotDePasse
	} else {
		return "Ancien mot de passe incorrecte"
	}
	return "Mot de passe changé avec succés"
}
