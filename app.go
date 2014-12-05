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
    message := "<html><body><h3>Contact:</h3><p>Phone: " + r.FormValue("phone") +"<br/>Message: "+ r.FormValue("message") +"</p><h3>Order:</h3><p>"+ r.FormValue("order") + "</p></body></html>"
    msg := &mail.Message{
            Sender:  name + " <seaissnow.nl@sea-is-snow.appspotmail.com>",
            To:      []string{"gert.cuykens.2@gmail.com"},
            ReplyTo: email,
            Subject: "seaissnow.nl",
            //Body:    message,
            HTMLBody: message,
            Headers: netMail.Header{
                //"Content-Type": []string{"text/html; charset=UTF-8"},
                "On-Behalf-Of": []string{email},
            },
    }
    if code := r.FormValue("code"); code != "12345" {
        c.Errorf("Wrong code: %v", code)
        http.Error(w, "{\"status\":\"Code invalid! Try again\", \"code\":401}", http.StatusUnauthorized)
    }else if err := mail.Send(c, msg); err != nil {
        c.Errorf("Couldn't send email: %v", err)
        http.Error(w, "{\"status\":\"Mail NOT send! Error\", \"code\":500}", http.StatusInternalServerError)
    }else{
        c.Infof("Mail send:\n %v", message)
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