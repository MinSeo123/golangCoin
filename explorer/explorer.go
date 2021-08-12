package explorer

import (
	"fmt"
	"github.com/MinSeo123/nomadcoin/blockchain"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template
const (
	port string = ":4000"
	templateDir string = "explorer/templates/"
)


type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func home(writer http.ResponseWriter, request *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(writer, "home", data)
}

func add(write http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		templates.ExecuteTemplate(write, "add", nil)
	case "POST":
		request.ParseForm()
		data := request.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(write, request, "/", http.StatusPermanentRedirect)
	}
}

func Start() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.html"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.html"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Println("Listening on ", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
