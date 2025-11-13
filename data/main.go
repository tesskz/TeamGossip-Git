package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type Question struct {
	Text         string
	Choices      []string
	CorrectIndex int
}

type Quiz struct {
	ID        string
	Title     string
	Questions []Question
}

var quizzes = map[string]Quiz{
	"info": {
		ID:    "info", // Partie Info du Quiz
		Title: "Quiz Informatique Générale",
		Questions: []Question{
			{
				Text:         "Quel est le système d’exploitation libre le plus utilisé ?",
				Choices:      []string{"Windows", "macOS", "Linux"},
				CorrectIndex: 2,
			},
			{
				Text:         "Quel langage est compilé ?",
				Choices:      []string{"Python", "Go", "JavaScript"},
				CorrectIndex: 1,
			},
			{
				Text: "Que signifie HTML ?",
				Choices: []string{
					"HyperText Markup Language",
					"HighText Machine Learning",
					"Hyper Transfer Main Line",
				},
				CorrectIndex: 0,
			},
			{
				Text:         "Quelle unité est utilisée pour mesurer la fréquence d’un processeur ?",
				Choices:      []string{"Hertz (Hz)", "Watt (W)", "Octet (B)"},
				CorrectIndex: 0,
			},
			{
				Text:         "Que signifie l’acronyme CPU ?",
				Choices:      []string{"Central Processing Unit", "Computer Personal Unit", "Central Program Utility"},
				CorrectIndex: 0,
			},
			{
				Text:         "Quel périphérique est principalement utilisé pour saisir du texte ?",
				Choices:      []string{"La souris", "Le clavier", "L’écran"},
				CorrectIndex: 1,
			},
			{
				Text:         "Quel protocole est utilisé pour sécuriser la navigation sur le Web du?",
				Choices:      []string{"HTTP", "FTP", "HTTPS"},
				CorrectIndex: 2,
			},
			{
				Text:         "Quelle est l’extension classique d’un fichier source Go ?",
				Choices:      []string{".golang", ".go", ".gocode"},
				CorrectIndex: 1,
			},
			{
				Text:         "Quel système de gestion de versions est le plus utilisé en développement ?",
				Choices:      []string{"Git", "SVN", "Mercurial"},
				CorrectIndex: 0,
			},
			{
				Text:         "Que signifie RAM ?",
				Choices:      []string{"Random Access Memory", "Read-Only Memory", "Rapid Action Module"},
				CorrectIndex: 0,
			},
		},
	},
	"cyber": {
		ID:    "cyber",
		Title: "Quiz Cybersécurité",
		Questions: []Question{
			{
				Text: "Quel est le but principal d’un antivirus ?",
				Choices: []string{
					"Accélérer la connexion Internet",
					"Protéger contre les logiciels malveillants",
					"Sauvegarder automatiquement les fichiers",
				},
				CorrectIndex: 1,
			},
			{
				Text: "Qu’est-ce que le phishing ?",
				Choices: []string{
					"Une attaque consistant à voler des informations via de faux mails ou sites",
					"Un protocole de chiffrement de données",
					"Un type de pare-feu matériel",
				},
				CorrectIndex: 0,
			},
			{
				Text: "Quel moyen est le plus sûr pour accéder à un compte sensible ?",
				Choices: []string{
					"Un mot de passe simple",
					"Un mot de passe complexe + authentification à deux facteurs (2FA)",
					"Toujours rester connecté sans se déconnecter",
				},
				CorrectIndex: 1,
			},
			{
				Text: "Comment s’appelle une attaque qui surcharge un serveur avec trop de requêtes ?",
				Choices: []string{
					"Brute force",
					"DDoS",
					"Injection SQL",
				},
				CorrectIndex: 1,
			},
			{
				Text: "Quel mot de passe est le plus solide parmi ces choix ?",
				Choices: []string{
					"azerty",
					"12345678",
					"J4!m3LesTacos2025",
				},
				CorrectIndex: 2,
			},
			{
				Text: "À quoi sert principalement un pare-feu (firewall) ?",
				Choices: []string{
					"Refroidir l’ordinateur",
					"Filtrer le trafic réseau",
					"Augmenter la vitesse du processeur",
				},
				CorrectIndex: 1,
			},
			{
				Text: "Qu’est-ce qu’un ransomware (rançongiciel) ?",
				Choices: []string{
					"Un logiciel qui affiche de la publicité",
					"Un logiciel qui chiffre les données et demande une rançon",
					"Un logiciel de sauvegarde automatique",
				},
				CorrectIndex: 1,
			},
			{
				Text: "Que faut-il éviter de faire sur un Wi-Fi public non sécurisé ?",
				Choices: []string{
					"Regarder des vidéos",
					"Se connecter à son compte bancaire",
					"Lire des articles de presse",
				},
				CorrectIndex: 1,
			},
			{
				Text: "Que signifie l’acronyme MFA / 2FA ?",
				Choices: []string{
					"Multi-Factor Authentication",
					"Multiple Files Access",
					"Main Firewall Access",
				},
				CorrectIndex: 0,
			},
			{
				Text: "Que faire en cas de mail suspect demandant ton mot de passe ?",
				Choices: []string{
					"Cliquer sur le lien pour vérifier",
					"Répondre avec un faux mot de passe",
					"Ne pas cliquer, supprimer ou signaler le mail",
				},
				CorrectIndex: 2,
			},
		},
	},
	"ia": {
		ID:    "ia",
		Title: "Quiz Data & Intelligence Artificielle",
		Questions: []Question{
			{
				Text: "Qu’est-ce qu’un « dataset » ?",
				Choices: []string{
					"Un algorithme d’IA",
					"Un ensemble de données utilisé pour l’analyse ou l’apprentissage",
					"Un langage de programmation dédié aux statistiques",
				},
				CorrectIndex: 1,
			},
			{
				Text: "Que signifie « IA » ?",
				Choices: []string{
					"Intelligence Artificielle",
					"Internet Automatisé",
					"Interface Anonyme",
				},
				CorrectIndex: 0,
			},
			{
				Text: "Quel est l’objectif principal de l’apprentissage supervisé ?",
				Choices: []string{
					"Générer du code aléatoire",
					"Apprendre à partir de données étiquetées pour prédire ou classer",
					"Supprimer des données inutiles",
				},
				CorrectIndex: 1,
			},
			{
				Text: "Quel type d’algorithme peut être utilisé pour classer des e‑mails (spam / non spam) ?",
				Choices: []string{
					"Arbre de décision",
					"Tri à bulles",
					"Algorithme de tri rapide",
				},
				CorrectIndex: 0,
			},
			{
				Text: "En data, que représente une “feature” (caractéristique) ?",
				Choices: []string{
					"Une ligne du code source",
					"Une variable d’entrée utilisée par le modèle",
					"Le résultat final du modèle",
				},
				CorrectIndex: 1,
			},
			{
				Text: "À quoi sert principalement le langage SQL ?",
				Choices: []string{
					"Créer des pages web",
					"Programmer des microcontrôleurs",
					"Interroger et manipuler des bases de données",
				},
				CorrectIndex: 2,
			},
			{
				Text: "Un fichier CSV sert en général à :",
				Choices: []string{
					"Stocker des images",
					"Stocker des données tabulaires en texte brut",
					"Exécuter du code machine",
				},
				CorrectIndex: 1,
			},
			{
				Text: "Qu’est-ce qu’un jeu d’entraînement (training set) ?",
				Choices: []string{
					"Les données utilisées pour apprendre le modèle",
					"Les données jamais vues utilisées pour le tester",
					"Le code source de l’algorithme",
				},
				CorrectIndex: 0,
			},
			{
				Text: "Lequel de ces exemples n’est PAS une application typique de l’IA ?",
				Choices: []string{
					"La reconnaissance d’images",
					"La traduction automatique de textes",
					"Imprimer un document avec une imprimante classique",
				},
				CorrectIndex: 2,
			},
			{
				Text: "Quel risque peut apparaître si les données d’entraînement sont biaisées ?",
				Choices: []string{
					"Le modèle sera parfaitement neutre",
					"Le modèle pourra être injuste envers certains groupes",
					"Le temps d’exécution sera toujours plus rapide",
				},
				CorrectIndex: 1,
			},
		},
	},
}

var tmpl *template.Template

var funcMap = template.FuncMap{
	"add": func(a, b int) int { return a + b },
}

type HomeData struct {
	Quizzes []Quiz
}

type QuizPageData struct {
	Quiz           Quiz
	Question       Question
	QuestionIndex  int
	TotalQuestions int
	Score          int
	Feedback       string
	FeedbackClass  string
}

type ResultData struct {
	Quiz    Quiz
	Score   int
	Total   int
	Success bool
}

func main() {

	var err error
	pattern := filepath.Join("templates", "*.html")
	tmpl, err = template.New("").Funcs(funcMap).ParseGlob(pattern)
	if err != nil {
		log.Fatalf("Erreur de parsing des templates : %v", err)
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/quiz", quizHandler)

	log.Println("Serveur démarré sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erreur serveur : %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	var list []Quiz
	for _, q := range quizzes {
		list = append(list, q)
	}

	data := HomeData{Quizzes: list}
	if err := tmpl.ExecuteTemplate(w, "home.html", data); err != nil {
		http.Error(w, "Erreur de rendu", http.StatusInternalServerError)
	}
}

func quizHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		startQuizHandler(w, r)
	case http.MethodPost:
		answerQuizHandler(w, r)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

func startQuizHandler(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")
	if topic == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	quiz, ok := quizzes[topic]
	if !ok {
		http.NotFound(w, r)
		return
	}

	data := QuizPageData{
		Quiz:           quiz,
		Question:       quiz.Questions[0],
		QuestionIndex:  0,
		TotalQuestions: len(quiz.Questions),
		Score:          0,
		Feedback:       "",
		FeedbackClass:  "",
	}

	if err := tmpl.ExecuteTemplate(w, "quiz.html", data); err != nil {
		http.Error(w, "Erreur de rendu", http.StatusInternalServerError)
	}
}

func answerQuizHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	topic := r.FormValue("topic")
	quiz, ok := quizzes[topic]
	if !ok {
		http.NotFound(w, r)
		return
	}

	qIndexStr := r.FormValue("qIndex")
	scoreStr := r.FormValue("score")
	answerStr := r.FormValue("answer")

	qIndex, err := strconv.Atoi(qIndexStr)
	if err != nil {
		http.Error(w, "Index de question invalide", http.StatusBadRequest)
		return
	}
	score, err := strconv.Atoi(scoreStr)
	if err != nil {
		http.Error(w, "Score invalide", http.StatusBadRequest)
		return
	}
	answerIndex, err := strconv.Atoi(answerStr)
	if err != nil {
		http.Error(w, "Réponse invalide", http.StatusBadRequest)
		return
	}

	if qIndex < 0 || qIndex >= len(quiz.Questions) {
		http.Error(w, "Question inexistante", http.StatusBadRequest)
		return
	}
	question := quiz.Questions[qIndex]
	if answerIndex == question.CorrectIndex {
		score++
	}

	qIndex++
	if qIndex >= len(quiz.Questions) {

		total := len(quiz.Questions)
		success := score >= 7
		data := ResultData{
			Quiz:    quiz,
			Score:   score,
			Total:   total,
			Success: success,
		}
		if err := tmpl.ExecuteTemplate(w, "result.html", data); err != nil {
			http.Error(w, "Erreur de rendu", http.StatusInternalServerError)
		}
		return
	}

	nextQuestion := quiz.Questions[qIndex]
	data := QuizPageData{
		Quiz:           quiz,
		Question:       nextQuestion,
		QuestionIndex:  qIndex,
		TotalQuestions: len(quiz.Questions),
		Score:          score,
		Feedback:       "",
		FeedbackClass:  "",
	}
	if err := tmpl.ExecuteTemplate(w, "quiz.html", data); err != nil {
		http.Error(w, "Erreur de rendu", http.StatusInternalServerError)
	}
}
