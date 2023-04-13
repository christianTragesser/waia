package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"gitlab.FQDN_HERE/root/waia/system"
)

var version = "dev-build"

//go:embed templates/info.html.tmpl
var htmlTemplate []byte

func renderTemplate(w http.ResponseWriter, tmpl []byte) {
	host := &system.System{}
	hostInfo, _ := host.GetInfo()

	tmplData := map[string]string{
		"Hostname": hostInfo.Hostname,
		"Address":  hostInfo.Address,
		"Version":  version,
		"BgColor":  hostInfo.BgColor,
		"TxtColor": hostInfo.TxtColor,
	}

	t, _ := template.New("template").Parse(string(tmpl))
	if err := t.Execute(w, tmplData); err != nil {
		log.Fatal(err)
	}
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, htmlTemplate)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{"version": version}

	jdata, err := json.Marshal(status)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jdata)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", appHandler)
	http.HandleFunc("/health/liveness", healthHandler)
	http.HandleFunc("/health/readiness", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
