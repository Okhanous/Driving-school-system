package utilisateurs

type Moniteur struct {
	User       Utilisateur
	Evenements []Evenement
	Etudiants  []Etudiant
}
