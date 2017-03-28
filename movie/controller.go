package movie

import (
	"fmt"
	"net/http"
)

func init()  {
	fmt.Println("movie controller was called")
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllMovies was called")
}
