package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	mux := registerRoutes()
	log.Print("Listening on :4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func registerRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	// productFilterHandler := http.HandlerFunc(controllers.NewProductFilterController)
	// mux.Handle("/product-filters", middleware(productFilterHandler))
	// restaurantHandler := http.HandlerFunc(controllers.NewRestaurantController)
	// mux.Handle("/restaurants", middleware(restaurantHandler))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString("hmacSampleSecret")

	fmt.Println(tokenString, err)
	return mux
}

// func middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				log.Println("Recovered from panic:", err)
// 				pc, file, line, _ := runtime.Caller(1)
// 				functionName := runtime.FuncForPC(pc).Name()
// 				log.Printf("Function:%s in %s:%d\n", functionName, file, line)
// 				debug.PrintStack()
// 				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			}
// 		}()
// 		fmt.Println("Middleware is executing...", r.URL)

// 		// Allow requests from all origins
// 		w.Header().Set("Access-Control-Allow-Origin", "*")

// 		// Allow specific HTTP methods
// 		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")

// 		// Allow specific HTTP headers
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, X-User-Type, Authorization")

// 		dbConn := getDBConn()
// 		ctx := context.WithValue(r.Context(), "dbConn", dbConn)
// 		nR := r.Clone(ctx)
// 		next.ServeHTTP(w, nR)
// 	})
// }
