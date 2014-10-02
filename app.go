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
            Sender:  name + " <submit@sea-is-snow.appspotmail.com>",
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
        fmt.Fprintf(w,"{\"status\":\"Code invalid! Try again\", \"code\":501}")
    }else if err := mail.Send(c, msg); err != nil {
        c.Errorf("Couldn't send email: %v", err)
        fmt.Fprint(w, "{\"status\":\"Mail NOT send! Error\", \"code\":500}")
    }else{
        c.Infof("Mail send:\n %v", message)
        fmt.Fprint(w, "{\"status\":\"Mail send\", \"code\":200}")
    }
}