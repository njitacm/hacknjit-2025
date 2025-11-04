package backend

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
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
	RegistrationTime    string // Time registration was processed; Formatted using TIMEFORMAT
}

var TIMEFORMAT = time.RFC850 // "Monday, 02-Jan-06 15:04:05 MST"

var csvWriter *csv.Writer
var CsvFile *os.File
var csvMutex sync.Mutex

var outputDir string = "./responses"
var resumeDir string = fmt.Sprintf("%s/resumes", outputDir) // Must be a subdirectory of outputDir
var csvFilePath string = fmt.Sprintf("%s/registrations-%d.csv", outputDir, time.Now().UnixNano())

func init() {
	// Create output directories (if they do not exist) and csv file
	// MkDirAll will create resume directory and its parent directory (outputDir)
	os.MkdirAll(resumeDir, os.ModePerm)
	csvFile, err := os.Create(csvFilePath)
	if err != nil {
		log.Fatalf("[FATAL] Failed to create csv file: %s\n", err)
	}

	csvWriter = csv.NewWriter(csvFile)

	err = writeRow([]string{"FirstName", "LastName", "PreferredName", "Age", "Phone", "Email", "Country", "Uni", "LvlOfStudy", "FirstHack", "Major", "ShirtSize", "DietaryRestrictions", "MLH Checkbox 0", "MLH Checkbox 1", "MLH Checkbox 2", "Minority", "Gender", "Race", "ResumePath", "LinkedIn", "RegistrationTime"})
	if err != nil {
		log.Fatalf("[FATAL] Failed to write headers for csv file %s: %s\n", csvFilePath, err)
	}
}

func HandleFormSubmission(w http.ResponseWriter, r *http.Request) {
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
		MLHCheckbox0:     r.PostFormValue("mlh_checkbox_0") != "",
		MLHCheckbox1:     r.PostFormValue("mlh_checkbox_1") != "",
		MLHCheckbox2:     r.PostFormValue("mlh_checkbox_2") != "",
		Minority:         r.PostFormValue("minority"),
		Gender:           r.PostFormValue("gender"),
		Race:             r.PostFormValue("race"),
		LinkedIn:         r.PostFormValue("linkedin"),
		RegistrationTime: time.Now().Format(TIMEFORMAT),
	}

	// Hacky fix because I can't figure out the proper way of getting checkboxes
	log.Println(r.Form["dietaryrestrictions"])
	for i := range 5 {
		v := (r.FormValue(fmt.Sprintf("dietaryrestrictions[%d]", i)))
		if v != "" {
			form.DietaryRestrictions = append(form.DietaryRestrictions, v)
		}
	}

	// Handle resume PDF upload
	if err = handleResumeUpload(&form, r); err != nil {
		log.Printf("[ERROR] Error saving resume: %s\n", err.Error())
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
	}

	err = saveRegistrationToCSV(form)
	if err != nil {
		log.Printf("[ERROR] Failed to save registration: \n%s\n", err)
		http.Error(w, "Failed to save registration", http.StatusInternalServerError)
		return
	}
	log.Printf("[INFO] Finished saving form: %+v\n", form)
	if form.PreferredName != "" {
		sendConfirmationEmail(form.PreferredName, form.Email)
	} else {
		sendConfirmationEmail(form.FirstName, form.Email)
	}

	// Automatically returns 200 success code
	_, err = w.Write([]byte("Successfully Processed Form"))
	if err != nil {
		log.Printf("[INFO] Failed to write success header: %s\n", err)
	}
}

func saveRegistrationToCSV(form RegistrationForm) error {
	csvRowUnsanitized := []string{
		form.FirstName,
		form.LastName,
		form.PreferredName,
		form.Age,
		form.Phone,
		form.Email,
		form.Country,
		form.Uni,
		form.LvlOfStudy,
		form.FirstHack,
		form.Major,
		form.ShirtSize,
		strings.Join(form.DietaryRestrictions,
			","),
		strconv.FormatBool(form.MLHCheckbox0),
		strconv.FormatBool(form.MLHCheckbox1),
		strconv.FormatBool(form.MLHCheckbox2),
		form.Minority,
		form.Gender,
		form.Race,
		form.ResumePath,
		form.LinkedIn,
		form.RegistrationTime,
	}

	csvRow := []string{}
	for _, col := range csvRowUnsanitized {
		csvRow = append(csvRow, sanitizeCSVColumn(col))
	}

	return writeRow(csvRow)
}

func writeRow(csvRow []string) error {
	csvMutex.Lock()         // Ensures only one write operation is being done at a given time
	defer csvMutex.Unlock() // Unlock on exit
	err := csvWriter.Write(csvRow)
	if err != nil {
		return fmt.Errorf("csvWriter failed to write row to csv file %s:\n\tRow: %s\n\tError: %w\n", csvFilePath, strings.Join(csvRow, ","), err)
	}

	csvWriter.Flush() // Must flush output buffer to actually write to the file
	// If an error occurs in csvWriter.flush, it is stored internally and can be retrieved with csvWriter.Error()
	if err := csvWriter.Error(); err != nil {
		return fmt.Errorf("csvWriter failed to flush output buffer when writing row to csv file %s:\n\tRow: %s\n\tError: %w\n", csvFilePath, strings.Join(csvRow, ","), err)
	}

	return nil
}

func handleResumeUpload(form *RegistrationForm, r *http.Request) error {
	file, _, err := r.FormFile("resume")
	if err == http.ErrMissingFile {
		log.Printf("[INFO] User did not submit a resume. Continuing...\n")
		return nil
	} else if err != nil {
		return err
	}

	log.Printf("[INFO] Beginning saving resume...")
	defer file.Close()
	dst := fmt.Sprintf("./responses/resumes/%s-%s-%d-resume.pdf", sanitizeString(form.FirstName), sanitizeString(form.LastName), time.Now().UnixNano())

	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}
	form.ResumePath = dst
	log.Printf("[INFO] Resume succesfully saved at %s", dst)

	return nil
}

func sanitizeString(str string) string {
	// Only keep letters, numbers, underscore, dash
	validChars := regexp.MustCompile(`[^a-zA-Z0-9_-]`)
	sanitizedStr := validChars.ReplaceAllString(str, "_")

	// Condense repeating underscores
	multiUnderscores := regexp.MustCompile(`_+`)
	sanitizedStr = multiUnderscores.ReplaceAllString(sanitizedStr, "_")

	// In case the result is only invalid characters
	if sanitizedStr == "" {
		return "ERROR"
	}
	return sanitizedStr
}

func sanitizeCSVColumn(str string) string {
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, "\n", " ")
	str = strings.ReplaceAll(str, "\r", " ")
	return str
}
