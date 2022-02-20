package handlers

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

var JWTKey = []byte("secret_key")

var students = map[string]string{
	"student1": "password1"
	"student2": "password2"
	"student3": "password3"
}

type Credentials struct {
	StudentName string `json:"student_name"`
	Password string `json:"student_password"`
}

type Claims struct {
	StudentName string `json:"student_name"`
	jwt.StandardClaims
}

func Login(resp http.ResponserWriter, req *http.Request) {
	var creds Credentials
	err := json.NewDecoder(req.Doby).Decode(&creds)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		panic("An error has happened while trying to log you in")
		return
	}

	expectedPsswrd, ok := students[creds.StudentName]
	if !ok || expectedPsswrd != creds.Password {
		resp.WriteHeader(http.StatusUnauthorized)
		panic("You have not been authorized to access the system.")
		return
	}

	expiration := time.Now().Add(time.Minute *5)
	claims := &Claims {
		StudentName: creds.StudentName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		panic("An internal server error has happened!")
		return
	}

	http.SetCookie(resp, &http.Cookie {
		Name: "token", 
		Value: tokenString,
		Expires: expiration,
	})
}

func Home(resp http.ResponserWriter, req *http.Request) {
	cookie, err := req.Cookie("token")
	if err != nil {
		if err == ErrNoCookie {
			resp.WriteHeader(htpp.StatusUnauthorized)
			panic("No cookie has been found. Access not authorized.")
			return
		}
		resp.WriteHeader(http.StatusBadRequest)
		panic("A fatal error has happened while processing the cookie...")
		return
	}

	tokenStr := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, 
		func(t *jwt.Token) (interface{}, error){
			return JWTKey, nil
		})
	
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			resp.WriteHeader(http.StatusUnauthorized)
			panic("Access not authorized.")
			return
		}
		resp.WriteHeader(http.StatusBadRequest)
		panic("A fatal error has happened while processing inforamtion...")
		return
	}

	if !token.Valid {
		resp.WriteHeader(http.StatusUnauthorized)
		panic("INVALID TOKEN! Access has not been granted.")
		return
	}

	resp.Write([]byte(fmt.Sprintf("Hello, %s!", claims.StudentName)))
}

func Refresh(writer http.ResponserWriter, req *http.Request) {
	cookie, err := req.Cookie("token")
	if err != nil {
		if err == ErrNoCookie {
			resp.WriteHeader(htpp.StatusUnauthorized)
			panic("No cookie has been found. Access not authorized.")
			return
		}
		resp.WriteHeader(http.StatusBadRequest)
		panic("A fatal error has happened while processing the cookie...")
		return
	}

	tokenStr := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, 
		func(t *jwt.Token) (interface{}, error){
			return JWTKey, nil
		})
	
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			resp.WriteHeader(http.StatusUnauthorized)
			panic("Access not authorized.")
			return
		}
		resp.WriteHeader(http.StatusBadRequest)
		panic("A fatal error has happened while processing inforamtion...")
		return
	}

	if !token.Valid {
		resp.WriteHeader(http.StatusUnauthorized)
		panic("INVALID TOKEN! Access has not been granted.")
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30 * time.Second {
		resp.WriteHeader(http.StatusBadRequest)
		panic("Your token has expired! Access time has ended.")
		return
	}

	expirationTime := time.Now().Add(time.Minute *5)
	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		panic("An internal server error has happened!")
		return
	}

	http.SetCookie(resp, &http.Cookie {
		Name: "refreshed_token", 
		Value: tokenString,
		Expires: expiration,
	})
}