package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Registration form structure
type RegistrationForm struct {
	FirstName           string
	LastName            string
	PreferredName       string
	Age                 string
	Phone               string
	Email               string
	Country             string
	Uni                 string
	LvlOfStudy          string
	FirstHack           string
	Major               string
	ShirtSize           string
	DietaryRestrictions []string
	Terms               bool
	MarketingEmails     bool
	MarketingEmails1    bool
	Minority            string
	Gender              string
	Race                string
	Resume              string // path to saved PDF
	LinkedIn            string
}

func main() {
	setOptions()
	http.Handle("/", http.FileServer(http.Dir("./dist")))
	http.HandleFunc("/api/register", handleFormSubmission)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setOptions() {
	var (
		isDev = flag.Bool("dev", false, "run in development mode")
	)
	flag.Parse()

	if *isDev {
		log.Println("Error: Dev mode not set")
	}

}

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
	// Handle CORS when doing local dev on different ports
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form (10 MB max memory for files)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	form := RegistrationForm{
		FirstName:           r.FormValue("first_name"),
		LastName:            r.FormValue("last_name"),
		PreferredName:       r.FormValue("preferred_name"),
		Age:                 r.FormValue("age"),
		Phone:               r.FormValue("phone"),
		Email:               r.FormValue("email"),
		Country:             r.FormValue("country"),
		Uni:                 r.FormValue("uni"),
		LvlOfStudy:          r.FormValue("lvlofstudy"),
		FirstHack:           r.FormValue("firsthack"),
		Major:               r.FormValue("major"),
		ShirtSize:           r.FormValue("shirtsize"),
		Terms:               r.FormValue("terms") == "true",
		MarketingEmails:     r.FormValue("marketing_emails") == "true",
		MarketingEmails1:    r.FormValue("marketing_emails_1") == "true",
		Minority:            r.FormValue("minority"),
		Gender:              r.FormValue("gender"),
		Race:                r.FormValue("race"),
		DietaryRestrictions: r.MultipartForm.Value["dietaryrestrictions[]"],
		LinkedIn:            r.FormValue("linkedin"),
	}

	// Handle resume PDF upload
	file, header, err := r.FormFile("resume")
	if err == nil {
		defer file.Close()
		os.MkdirAll("./uploads", os.ModePerm)
		dst := fmt.Sprintf("./uploads/%d-%s", time.Now().UnixNano(), header.Filename)
		out, err := os.Create(dst)
		if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}
		form.Resume = dst
	}

	log.Printf("Received registration: %+v\n", form)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"success","message":"Registration received"}`)
}
