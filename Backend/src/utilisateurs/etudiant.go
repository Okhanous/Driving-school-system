package utilisateurs

// creation de structure/classe se fait sous la forme: motclé(type) nomDutype struct{}
type Etudiant struct {
	User         Utilisateur
	PermisObtenu bool
	Evenements   []Evenement
	Encadrement  Moniteur
}
