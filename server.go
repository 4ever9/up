package up

import (
	"html/template"
	"io"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"sort"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

var indexHTMLFiles = []string{
	"index.html",
	"index.htm",
}

type Server struct {
	Dir http.FileSystem
}

func NewServer() *Server {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir := http.Dir(pwd)
	return &Server{
		Dir: http.Dir(dir),
	}
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		format := "[%s %s %s %v] => %s \n"
		if err != nil {
			logger.Printf(format, r.Method, r.Proto, r.URL.Path, r.RemoteAddr, err)
			return
		}
		logger.Printf(format, r.Method, r.Proto, r.URL.Path, r.RemoteAddr, "request 200")
	}()

	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		return
	}

	file := r.URL.Path

	f, err := s.Dir.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return
	}

	if fi.IsDir() {
		err = s.listDir(rw, r, f)
		return
	}

	if mimeType := mime.TypeByExtension(path.Ext(fi.Name())); mimeType != "" {
		rw.Header().Set("Content-type", mimeType)
	} else {
		rw.Header().Set("Content-type", "application/octet-stream")
	}

	if _, err = io.Copy(rw, f); err != nil {
		return
	}
}

func isIndexHTML(file string) bool {
	for _, f := range indexHTMLFiles {
		if f == file {
			return true
		}
	}
	return false
}

func (s *Server) listDir(rw http.ResponseWriter, r *http.Request, f http.File) error {
	dirs, err := f.Readdir(-1)
	if err != nil {
		return err
	}
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	items := make([]*Item, 0, len(dirs))
	for _, d := range dirs {
		if isIndexHTML(d.Name()) {
			f, err := s.Dir.Open(r.URL.Path + d.Name())
			if err != nil {
				return err
			}
			if _, err := io.Copy(rw, f); err != nil {
				return err
			}

			f.Close()
			return err
		}

		name := d.Name()
		if d.IsDir() {
			name += "/"
		}
		u := url.URL{Path: name}
		items = append(items, &Item{
			Name: name,
			Type: GetItemType(d),
			URL:  u.String(),
		})
	}
	t := template.Must(template.New("listing").Parse(Template))
	return t.Execute(rw, map[string]interface{}{
		"Items": items,
	})
}
