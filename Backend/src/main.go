package main

import (
	"driving-school-backend/src/utilisateurs"
	"fmt"
	"net/http"
	"text/template"
)

type PageData struct {
	AdminData      utilisateurs.Admin
	SecretaireData utilisateurs.Secretaire
	EtudiantData   utilisateurs.Etudiant
}

var data PageData

func main() {
	fmt.Println("Serving at: http://localhost:8080")
	// Add this code before starting the server
	fs := http.FileServer(http.Dir("../templates/images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	http.HandleFunc("/adminAuthentification", adminAuthentification)
	http.HandleFunc("/adminDashboard", adminDashboard)
	http.HandleFunc("/inscriptionSecretaire", inscriptionSecretaire)
	http.HandleFunc("/afficherSecretaires", afficherSecretaires)
	http.HandleFunc("/", mainPage)

	http.ListenAndServe(":8080", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered main Page")
	if r.Method == "POST" {
		fmt.Println("Received POST Request")
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error parsing from:", err)
		}
		action := r.Form.Get("action")
		fmt.Println(action)
		switch action {
		case "secretaire":
			http.Redirect(w, r, "/secretaireAuthentification", http.StatusFound)
		case "admin":
			http.Redirect(w, r, "/adminAuthentification", http.StatusFound)
		case "etudiant":
			http.Redirect(w, r, "/etudiantAuthentification", http.StatusFound)
		}
	} else {
		t, err := template.ParseFiles("../templates/mainPage.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = t.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
func adminAuthentification(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered admin authentication page")
	if r.Method == "POST" {
		fmt.Println("Received POST request")
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error parsing from:", err)
		}
		nom := r.Form.Get("nom")
		prenom := r.Form.Get("prenom")
		email := r.Form.Get("email")
		motDePasse := r.Form.Get("mdp")

		admin := utilisateurs.Admin{
			User: utilisateurs.Utilisateur{
				Nom:        nom,
				Prenom:     prenom,
				Email:      email,
				MotDePasse: motDePasse,
			},
			Secretaires: []utilisateurs.Secretaire{},
		}
		data.AdminData = admin

		http.Redirect(w, r, "/adminDashboard", http.StatusFound)
	} else {
		t, err := template.ParseFiles("../templates/admin-authentification.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func adminDashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered admin Dashboard page")
	// Register a secretary
	if r.Method == "POST" {
		fmt.Println("Received POST Request")
		//Avant la lecture d'un formulaire il faut le parser avec r.ParseForm()
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error parsing from:", err)
		}
		action := r.PostForm.Get("action")
		fmt.Println(action)
		switch action {
		case "afficherSecretaires":
			http.Redirect(w, r, "/afficherSecretaires", http.StatusFound)
		case "inscriptionSecretaire":
			http.Redirect(w, r, "/inscriptionSecretaire", http.StatusFound)
		}
	} else {
		t, err := template.ParseFiles("../templates/adminDashboard.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, data.AdminData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
func afficherSecretaires(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered secretaries page page")
	t, err := template.ParseFiles("../templates/afficherSecretaires.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data.AdminData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func inscriptionSecretaire(w http.ResponseWriter, r *http.Request) {
	// Access the admin data using data.AdminData

	fmt.Println("Entered secretary inscription page")
	// Register a secretary
	if r.Method == "POST" {
		fmt.Println("Received POST Request")
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error parsing from:", err)
		}
		nom := r.Form.Get("nom")
		prenom := r.Form.Get("prenom")
		email := r.Form.Get("email")
		motDePasse := r.Form.Get("mdp")
		data.SecretaireData = data.AdminData.InscrireSecretaire(nom, prenom, email, motDePasse)
		// data.SecretaireData.InscrireMoniteur("carlo", "ancelloti", "ancelloti@gmail.com")

		// // Register an etudiant
		// etudiantA := data.SecretaireData.InscrireEtudiant("Pogba", "Paul", "paulito@gmail.com", "blabla")

		// // Fix an event for the etudiant with the secretary
		// evenement := utilisateurs.Evenement{
		// 	Id:          1,
		// 	Date:        time.Now(),
		// 	Duree:       60,
		// 	Description: "PremierCours",
		// }
		// fmt.Println(data.AdminData.Secretaires[0])c
		// fmt.Println(data.AdminData.Secretaires[0].FixerEvenementEtudiant(evenement, etudiantA))
		http.Redirect(w, r, "/adminDashboard", http.StatusFound)

	} else {
		t, err := template.ParseFiles("../templates/inscriptionSecretaire.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, data.AdminData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
