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
	ResumePath          string // path to saved PDF
	LinkedIn            string
}

func main() {
	os.MkdirAll("./resumes", os.ModePerm)

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
		FirstName:        r.PostFormValue("first_name"),
		LastName:         r.PostFormValue("last_name"),
		PreferredName:    r.PostFormValue("preferred_name"),
		Age:              r.PostFormValue("age"),
		Phone:            r.PostFormValue("phone"),
		Email:            r.PostFormValue("email"),
		Country:          r.PostFormValue("country"),
		Uni:              r.PostFormValue("uni"),
		LvlOfStudy:       r.PostFormValue("lvlofstudy"),
		FirstHack:        r.PostFormValue("firsthack"),
		Major:            r.PostFormValue("major"),
		ShirtSize:        r.PostFormValue("shirtsize"),
		Terms:            r.PostFormValue("terms") == "true",
		MarketingEmails:  r.PostFormValue("marketing_emails") == "true",
		MarketingEmails1: r.PostFormValue("marketing_emails_1") == "true",
		Minority:         r.PostFormValue("minority"),
		Gender:           r.PostFormValue("gender"),
		Race:             r.PostFormValue("race"),
		LinkedIn:         r.PostFormValue("linkedin"),
	}

	// Hacky fix because I can't figure out the proper way of getting checkboxes
	for i := range 5 {
		v := (r.FormValue(fmt.Sprintf("dietaryrestrictions[%d]", i)))
		if v != "" {
			form.DietaryRestrictions = append(form.DietaryRestrictions, v)
		}
	}

	// Handle resume PDF upload
	file, _, err := r.FormFile("resume")
	if err == nil {
		log.Printf("Saving resume...")
		defer file.Close()
		dst := fmt.Sprintf("./resumes/%s-%s-%d-resume.pdf", form.FirstName, form.LastName, time.Now().UnixNano())

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
		form.ResumePath = dst
	} else {
		log.Printf("Error getting formfile: %s\n", err.Error())
	}

	log.Printf("Processed form: %+v\n", form)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"success","message":"Registration received"}`)
}
