package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"text/template"
	"time"
)

var (
	httpAddr   = flag.String("http", ":8080", "Listen address")
	pollPeriod = flag.Duration("poll", 5*time.Second, "Duration")
	version    = flag.String("version", "1.4", "Go version")
)

const baseChangeURL = "https://code.google.com/p/go/source/detail?r="

func main() {
	flag.Parse()

	changeURL := fmt.Sprintf("%sgo%s", baseChangeURL, *version)
	http.Handle("/", NewServer(*version, changeURL, *pollPeriod))
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

type Server struct {
	version string
	url     string
	period  time.Duration

	mu  sync.RWMutex // protects the yes variable
	yes bool
}

func NewServer(version, url string, period time.Duration) *Server {
	s := &Server{version: version, url: url, period: period}
	go s.poll()
	return s
}

func (s *Server) poll() {
	for !isTagged(s.url) {
		time.Sleep(s.period)
	}
	s.mu.Lock()
	s.yes = true
	s.mu.Unlock()
}

func isTagged(url string) bool {
	r, err := http.Head(url)
	if err != nil {
		log.Print(err)
		return false
	}
	return r.StatusCode == http.StatusOK
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	data := struct {
		URL     string
		Version string
		Yes     bool
	}{
		s.url,
		s.version,
		s.yes,
	}
	s.mu.RUnlock()
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
	}
}

var tmpl = template.Must(template.New("tmpl").Parse(`
	<!DOCTYPE html><html>
		<body><center>
			<h2>Is Go {{.Version}} out yet?</h2>
			<h1>
				{{ if .Yes }}
					<a href="{{.URL}}">YES!</a>
				{{ else }}
					No. :-(
				{{ end}}
			</h1>
		</center></body>
	</html>
`))
