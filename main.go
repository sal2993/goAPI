package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"

  jwt "github.com/darwin_amd64/jwt-go"
)


func ValidatePerson( w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  log.Println(params["token"])
  token := params["token"]
  log.Println(token)
}

func theIndex( w http.ResponseWriter, r *http.Request ) {
  paramz := mux.Vars(r);
  log.Println(paramz)
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("everythin okay bro\n"))
}

func main() {
  tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    // Don't forget to validate the alg is what you expect:
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }

    // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
    return []byte("my_secret_key"), nil
  })

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    fmt.Println(claims["foo"], claims["nbf"])
  } else {
    fmt.Println(err)
  }


  router := mux.NewRouter()
  router.HandleFunc("/people/{token}", ValidatePerson).Methods("POST")
  router.HandleFunc("/", theIndex).Methods("GET")

  log.Println("Listening...")
  log.Fatal(http.ListenAndServe(":8000", router))
}
