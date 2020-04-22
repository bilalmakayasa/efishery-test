package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	Users          []User `json:"users"`
	SecretKey      string `json:"secretKey"`
	ConversionRate interface{}
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type LoginScheme struct {
	Password int    `json:"password"`
	Phone    string `json:"phone"`
}

type RegistryResponse struct {
	Name     string `json:"name"`
	Password int    `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type Claims struct {
	jwt.StandardClaims
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

type Token struct {
	AccessToken string    `json:"access_token"`
	Expire      time.Time `json:"expire_at"`
}

type ResultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    interface{}
}

func RetrieveData() []byte {
	jsonFile, err := os.Open("./../data.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func passwordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

func matchPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Response(code int, message string, data interface{}) interface{} {
	result := ResultResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}

	return result
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		result := Response(405, "Your HTTP method is invalid", nil)
		json.NewEncoder(w).Encode(result)
		return
	}
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	password := rand.Intn(9999)
	retrieveData := RetrieveData()
	var data Data
	json.Unmarshal(retrieveData, &data)

	role := strings.ToLower(user.Role)

	if role != "admin" && role != "client" {
		w.WriteHeader(http.StatusBadRequest)
		result := Response(400, "your role should be either admin or client", nil)
		json.NewEncoder(w).Encode(result)
		return
	}

	for i := 0; i < len(data.Users); i++ {
		dataUser := data.Users[i]
		if user.Phone == dataUser.Phone {
			w.WriteHeader(http.StatusBadRequest)
			result := Response(400, "you already registered", nil)
			json.NewEncoder(w).Encode(result)
			return
		}
	}
	passwordString := strconv.Itoa(password)
	encryptPassword, _ := passwordHash(passwordString)
	newValue := &User{
		Name:     user.Name,
		Phone:    user.Phone,
		Password: encryptPassword,
		Role:     role,
	}

	appendedData := append(data.Users, *newValue)
	data.Users = appendedData
	newData, _ := json.MarshalIndent(data, "", "")
	ioutil.WriteFile("./../data.json", newData, 0644)
	responseData := &RegistryResponse{
		Name:     user.Name,
		Phone:    user.Phone,
		Password: password,
		Role:     role,
	}
	result := Response(200, "Registration Success, please remember your password", responseData)
	json.NewEncoder(w).Encode(result)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		result := Response(405, "Your HTTP method is invalid", nil)
		json.NewEncoder(w).Encode(result)
		return
	}

	var user LoginScheme
	json.NewDecoder(r.Body).Decode(&user)

	retrieveData := RetrieveData()
	var data Data
	json.Unmarshal(retrieveData, &data)

	for i := 0; i < len(data.Users); i++ {
		dataUser := data.Users[i]
		if user.Phone == dataUser.Phone {
			passwordString := strconv.Itoa(user.Password)
			passwordChecking := matchPassword(passwordString, dataUser.Password)
			fmt.Println(passwordChecking)
			if passwordChecking == false {
				w.WriteHeader(http.StatusBadRequest)
				result := Response(400, "Invalid Password", nil)
				json.NewEncoder(w).Encode(result)
				return
			}

			fmt.Println(time.Now().Add(8 * time.Hour).Unix())
			claims := &Claims{
				Name:  dataUser.Name,
				Phone: dataUser.Phone,
				Role:  dataUser.Role,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(20 * time.Hour).Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			rsa := []byte("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQClSEK7AtfQnRmDFCFvdbM6D6F+tP8eKOWqnvs0jnInOvrZ4c0Nzu/9ilzEhYjYMh+P79hEy0jw1+09/u6lrYo6/BAjxIpSOZDiOZlFKylUnK60fxc6auJv4NjyOBKoVXKqQUuCju6QjdYRTl+6XrSms1/fpEXFMRJP4IMthjO2DT3d6oNuzQhScAleDl9lsOYvcitn3+uZttHdVXjYsU/luw1/OPjMqf2gh7cR+GX7lX1nCa9RYAYijpYNT6LAA/XQUSYGl9qVfFR8ZaWYyTQgITUKQQHac3GpOhBBFyrkJYdNG+L34072f/Jj3qe7eKDD1SR5szyU+/S3aFd4YB+/Z2YAD+0YLud09NoIlHbisrgoYsPlyz2r+Hri0tW7IyDKhjyymbGVEBC3WKHTCW1LpOR5B/+u9nNm5AVqZ76GMOjbXyWDaKz40FnwToIpDe5oxhq/fMINbXrPViivtjfGT3ifFrzl+9j4U0qSSUAxmUE88yRubzekmEbqX++3KUJtlfclpiPAP4HB2ayVBDSUGMhSrqmmrLxggB8nTnX61jFxkBOCiN+LyOeRbwf+3RkTQ5yhu+/OBP2L6msWqTyhO77FlZGzktHqfexqspbLTjWcmMJFvt6tJgFdk3IIXAnLeCEbuoDcl2C60Y/jrrsDFeCbtvzuA4GSMtNjUs76rQ== bilal.makayasa@gmail.com")
			fmt.Println(rsa)
			tokenString, err := token.SignedString(rsa)
			if err != nil {
				panic(err)
			}

			data := &Token{
				AccessToken: tokenString,
				Expire:      time.Unix(claims.ExpiresAt, 0),
			}

			result := Response(200, "Login Success", data)
			json.NewEncoder(w).Encode(result)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	result := Response(404, "You have not registered yet", nil)
	json.NewEncoder(w).Encode(result)

	return
}

func Credential(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		result := Response(405, "Your HTTP method is invalid", nil)
		json.NewEncoder(w).Encode(result)
		return
	}
	retrieveData := RetrieveData()
	var data Data
	json.Unmarshal(retrieveData, &data)
	params := r.URL.Query().Get("token")
	rsa := []byte("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQClSEK7AtfQnRmDFCFvdbM6D6F+tP8eKOWqnvs0jnInOvrZ4c0Nzu/9ilzEhYjYMh+P79hEy0jw1+09/u6lrYo6/BAjxIpSOZDiOZlFKylUnK60fxc6auJv4NjyOBKoVXKqQUuCju6QjdYRTl+6XrSms1/fpEXFMRJP4IMthjO2DT3d6oNuzQhScAleDl9lsOYvcitn3+uZttHdVXjYsU/luw1/OPjMqf2gh7cR+GX7lX1nCa9RYAYijpYNT6LAA/XQUSYGl9qVfFR8ZaWYyTQgITUKQQHac3GpOhBBFyrkJYdNG+L34072f/Jj3qe7eKDD1SR5szyU+/S3aFd4YB+/Z2YAD+0YLud09NoIlHbisrgoYsPlyz2r+Hri0tW7IyDKhjyymbGVEBC3WKHTCW1LpOR5B/+u9nNm5AVqZ76GMOjbXyWDaKz40FnwToIpDe5oxhq/fMINbXrPViivtjfGT3ifFrzl+9j4U0qSSUAxmUE88yRubzekmEbqX++3KUJtlfclpiPAP4HB2ayVBDSUGMhSrqmmrLxggB8nTnX61jFxkBOCiN+LyOeRbwf+3RkTQ5yhu+/OBP2L6msWqTyhO77FlZGzktHqfexqspbLTjWcmMJFvt6tJgFdk3IIXAnLeCEbuoDcl2C60Y/jrrsDFeCbtvzuA4GSMtNjUs76rQ== bilal.makayasa@gmail.com")
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(params, claims, func(token *jwt.Token) (interface{}, error) {
		return rsa, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	credential := &Claims{
		Name:  claims.Name,
		Phone: claims.Phone,
		Role:  claims.Role,
	}
	result := Response(200, "Credential Acces", credential)
	json.NewEncoder(w).Encode(result)

	return
}
