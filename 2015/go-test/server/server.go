package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atomaths/talks/go-test/stringutil"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir("")))
	http.HandleFunc("/reverse", reverseHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func reverseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, stringutil.Reverse(r.FormValue("in")))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, page)
}

const page = `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <script src="/static/jquery.min.js"></script>
  <script>
    $(document).ready(function() {
      $('#btn').click(function() {
	$.get('/reverse?in='+$('#in').val(), function(result) {
	  $('#result').html(result);
	});
      });
    });
  </script>
</head>
<body>
  <input id="in" name="in" type="text">
  <button id="btn">Reverse</button>
  <div id="result"></div>
</body>
</html>
`
