package endpoints

import (
	"fmt"
	"net/http"
)

func ContactPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Contact Page")
}