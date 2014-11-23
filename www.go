package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
)

func main() {
	index := http.FileServer(http.Dir("www"))
	http.Handle("/", index)
	http.HandleFunc("/submit.go", submit)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func submit(w http.ResponseWriter, r *http.Request) {
	//from := mail.Address{r.FormValue("name"), r.FormValue("email")}
	from := mail.Address{Name: "gert", Address: "gert.cuykens@gmail.com"}
	to := mail.Address{Name: "gert", Address: "gert.cuykens.2@gmail.com"}
	body := "<html><body><h3>Contact:</h3><p>Phone: " + r.FormValue("phone") + "<br/>Message: " + r.FormValue("message") + "</p><h3>Order:</h3><p>" + r.FormValue("order") + "</p></body></html>"
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = "seaissnow.nl"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	//"On-Behalf-Of": []string{email},

	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	/*if code := r.FormValue("code"); code != "12345" {
	        log.Printf("Wrong code: %v", code)
	        http.Error(w, "{\"status\":\"Code invalid! Try again\", \"code\":401}", http.StatusUnauthorized)
			return
	}*/

	auth := smtp.PlainAuth()

	err := smtp.SendMail(
		"smtp.gmail.com:587", //"ssl0.ovh.net:587",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(msg),
	)

	if err != nil {
		log.Printf("Couldn't send email: %v", err)
		http.Error(w, "{\"status\":\"Mail NOT send! Error\", \"code\":500}", http.StatusInternalServerError)
	} else {
		log.Printf("Mail send:\n %v", msg)
		fmt.Fprint(w, "{\"status\":\"Mail send\", \"code\":200}")
	}
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
		fn(w, r)
	}
}
