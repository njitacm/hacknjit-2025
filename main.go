package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	MLHCheckbox0        bool
	MLHCheckbox1        bool
	MLHCheckbox2        bool
	Minority            string
	Gender              string
	Race                string
	ResumePath          string // path to saved PDF
	LinkedIn            string
}

var csvWriter *csv.Writer
var csvFilePath string = "./formdata/registrations.csv"
var csvFile *os.File

func setup() {
	var (
		isDev = flag.Bool("dev", false, "run in development mode")
	)
	flag.Parse()

	if *isDev {
		log.Println("Error: Dev mode not set")
	}

	os.MkdirAll("./formdata/resumes", os.ModePerm)

	if _, err := os.Stat(csvFilePath); errors.Is(err, os.ErrNotExist) {
		// file does not exist
		csvFile, err = os.Create(csvFilePath)

		// Setup CSV file
		if err != nil {
			log.Fatal("Failed to create csv file\n\t", err)
		}

		csvWriter = csv.NewWriter(csvFile)

		err = csvWriter.Write([]string{"FirstName", "LastName", "PreferredName", "Age", "Phone", "Email", "Country", "Uni", "LvlOfStudy", "FirstHack", "Major", "ShirtSize", "DietaryRestrictions", "MLH Checkbox 0", "MLH Checkbox 1", "MLH Checkbox 2", "Minority", "Gender", "Race", "ResumePath", "LinkedIn"})
		if err != nil {
			log.Fatal("Failed to write headers to csv\n\t", err)
		}
	} else {
		csvFile, err = os.OpenFile(csvFilePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Failed to open csv file\n\t", err)
		}
		csvWriter = csv.NewWriter(csvFile)
	}
}

func main() {
	setup()
	defer csvFile.Close()
	defer csvWriter.Flush()

	http.Handle("/", http.FileServer(http.Dir("./dist")))
	http.HandleFunc("/api/register", handleFormSubmission)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form (2 MB max memory for files)
	err := r.ParseMultipartForm(2 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	form := RegistrationForm{
		FirstName:     r.PostFormValue("first_name"),
		LastName:      r.PostFormValue("last_name"),
		PreferredName: r.PostFormValue("preferred_name"),
		Age:           r.PostFormValue("age"),
		Phone:         r.PostFormValue("phone"),
		Email:         r.PostFormValue("email"),
		Country:       r.PostFormValue("country"),
		Uni:           r.PostFormValue("uni"),
		LvlOfStudy:    r.PostFormValue("lvlofstudy"),
		FirstHack:     r.PostFormValue("firsthack"),
		Major:         r.PostFormValue("major"),
		ShirtSize:     r.PostFormValue("shirtsize"),
		MLHCheckbox0:  r.PostFormValue("mlh_checkbox_0") == "true",
		MLHCheckbox1:  r.PostFormValue("mlh_checkbox_1") == "true",
		MLHCheckbox2:  r.PostFormValue("mlh_checkbox_2") == "true",
		Minority:      r.PostFormValue("minority"),
		Gender:        r.PostFormValue("gender"),
		Race:          r.PostFormValue("race"),
		LinkedIn:      r.PostFormValue("linkedin"),
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
		dst := fmt.Sprintf("./formdata/resumes/%s-%s-%d-resume.pdf", form.FirstName, form.LastName, time.Now().UnixNano())

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
	err = saveRegistrationToCSV(form)
	if err != nil {
		log.Printf("Failed to save registration: %s\n", err)
	}

	log.Printf("Processed form: %+v\n", form)

	// Automatically return 200 success code
	_, err = w.Write([]byte("Successfully Processed Form"))
	if err != nil {
		log.Printf("Failed to write success header: %s\n", err)
	}
}

func saveRegistrationToCSV(form RegistrationForm) error {
	err := csvWriter.Write([]string{
		form.FirstName, form.LastName, form.PreferredName, form.Age, form.Phone, form.Email, form.Country, form.Uni, form.LvlOfStudy, form.FirstHack, form.Major, form.ShirtSize, strings.Join(form.DietaryRestrictions, ","), strconv.FormatBool(form.MLHCheckbox0), strconv.FormatBool(form.MLHCheckbox1), strconv.FormatBool(form.MLHCheckbox2), form.Minority, form.Gender, form.Race, form.ResumePath, form.LinkedIn,
	})
	csvWriter.Flush()
	return err
}
