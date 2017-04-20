package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"io"
	"log"
	"net/http"
	"strings"
	"os"
	"io/ioutil"
	"mime/multipart"
	"github.com/google/uuid"
)

// Determine which handler needs to be used
func handleConnection(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" && r.FormValue("id") == "" {
		w.WriteHeader(200)
		io.WriteString(w, "To upload, make a POST request with the file attached in the 'image' field\n")
		io.WriteString(w, "To view, make a GET request with the file id in the 'id' parameter")

		return
	}

	if r.Method == "POST" {
		handlePost(w, r)
	}

	if r.Method == "GET" {
		handleGet(w, r)
	}
}

// Handle GET (file fetch) requests
func handleGet(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		w.WriteHeader(400)
		io.WriteString(w, "Invalid 'id' parameter")

		return
	}

	file, err := http.Dir(uploadDir).Open(id)

	if os.IsNotExist(err) {
		log.Println(err.Error())
		http.NotFound(w, r)

		return
	}

	_, fileType, imageError  := image.Decode(file)

	defer file.Close()

	if imageError != nil {
		log.Println(imageError.Error())
		http.NotFound(w, r)

		return
	}

	fileinfo, err := file.Stat()
	data, err := ioutil.ReadFile(uploadDir + "/" + fileinfo.Name())

	w.Header().Set("Content-Type", "image/"+fileType)

	io.WriteString(w, string(data))
}

// Handle POST (file upload) requests
func handlePost(w http.ResponseWriter, r *http.Request) {
	var err error

	if key != "" {
		if key != r.PostFormValue("key") {
			w.WriteHeader(400)
			io.WriteString(w, "Invalid API key")

			return
		}
	}

	if strings.Index(r.Header.Get("Content-Type"), "multipart/form-data") != 0 {
		w.WriteHeader(400)
		io.WriteString(w, "Please set the Content-Type to multipart/form-data")

		return
	}

	if verbose {
		log.Println("Receiving file...")
	}

	file, fileHeader, _ := r.FormFile("image")

	if verbose {
		log.Println("Validating file...")
	}

	if _, _, err = image.Decode(file); nil != err {
		if verbose {
			log.Println(err.Error())
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Open uploaded file
	var infile multipart.File
	if infile, err = fileHeader.Open(); nil != err {
		if verbose {
			log.Println(err.Error())
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if verbose {
		log.Println("Checking upload directory existance")
	}

	// Check if the target directory exists
	if _, err = os.Stat(uploadDir); nil != err {
		if verbose {
			log.Println(err.Error())
		}

		if os.IsNotExist(err) {
			os.MkdirAll(uploadDir, 0755)
		} else {
			log.Fatalln(err.Error())
		}
	}

	if verbose {
		log.Println("Opening target file")
	}

	// Open the target file
	var outfile *os.File
	filename, _ := uuid.NewUUID()
	if outfile, err = os.Create(uploadDir + "/" + filename.String()); nil != err {
		if verbose {
			log.Println(err.Error())
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if verbose {
		log.Println("Writing target file...")
	}

	// Write the destination file
	if _, err = io.Copy(outfile, infile); nil != err {
		if verbose {
			log.Println(err.Error())
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the UUID filename to the client
	w.Write([]byte(filename.String()))

	return
}
