package app

import (
    "fmt"
    "net/http"
    netMail "net/mail"
    "appengine"
    "appengine/mail"
)

func init() {
  http.HandleFunc("/submit.go", submit)
}

func submit(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    name := r.FormValue("name")
    email := r.FormValue("email")
    subject := r.FormValue("subject")
    message := r.FormValue("message")
    msg := &mail.Message{
            Sender:  name + " <submit@sea-is-snow.appspotmail.com>",
            To:      []string{"gert.cuykens.2@gmail.com"},
            ReplyTo: email,
            Subject: subject,
            Body:    message,
            Headers: netMail.Header{
                "On-Behalf-Of": []string{email},
            },
    }
    if code := r.FormValue("code"); code != "1234" {
        c.Errorf("Wrong code: %v", code)
        fmt.Fprintf(w, fmt.Sprintf("{\"status\":\"Code invalid! %s Try again\"}", code))
    }else if err := mail.Send(c, msg); err != nil {
        c.Errorf("Couldn't send email: %v", err)
        fmt.Fprint(w, "{\"status\":\"Mail NOT send! Error\"}")
    }else{
        fmt.Fprint(w, "{\"status\":\"Mail send.\"}")
    }
}