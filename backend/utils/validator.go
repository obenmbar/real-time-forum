package forum

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func ValidUserdata(data Users) error {

	isvalidnickname := true
	datahavespace := true
	for _, i := range data.Nickname {
		if !unicode.IsDigit(i) && !unicode.IsLetter(i) {
			isvalidnickname = false
			break
		} else if unicode.IsSpace(i) {
			datahavespace = false
			break
		}
	}
	if !isvalidnickname {
		return fmt.Errorf("valid nickname contient seulement de caracter et number ")
	}
	if !datahavespace {
		return fmt.Errorf("nickname ne contien pas space")
	}
	if len(data.Nickname) < 6 || len(data.Nickname) > 15 {
		return fmt.Errorf("valid nickname  conteient entre 6 et 15 caracter")
	}

	isvalidfirstname := true

	for _, v := range data.FirstName {
		if !unicode.IsLetter(v) {
			isvalidfirstname = false
			break
		} else if unicode.IsSpace(v) {
			datahavespace = false
			break
		}
	}
	if !isvalidfirstname {
		return fmt.Errorf("valid first name seulement lettre ")
	}
	if !datahavespace {
		return fmt.Errorf("firstname ne contien pas space")
	}
	if len(data.FirstName) < 6 || len(data.FirstName) > 15 {
		return fmt.Errorf("valid firstname  conteient entre 6 et 15 caracter")
	}

	isvalidlastname := true
	for _, r := range data.LastName {
		if !unicode.IsLetter(r) {
			isvalidlastname = false
			break
		} else if unicode.IsSpace(r) {
			datahavespace = false
		}
	}
	if !datahavespace {
		return fmt.Errorf("lastname ne contien pas space")
	}
	if !isvalidlastname {
		return fmt.Errorf("valid last name seulement lettre ")
	}

	if len(data.LastName) < 6 || len(data.LastName) > 15 {
		return fmt.Errorf("valid firstname  conteient entre 6 et 15 caracter")
	}

	age, err := strconv.Atoi(data.Age)
	if err != nil {
		return fmt.Errorf("age n'est pas number")
	}

	if age < 10 {
		return fmt.Errorf("age dois rendre plus de 10 ")
	} else if age > 130 {
		return fmt.Errorf("age doit rendre moint de 130")
	}

	if data.Gender != "Male" && data.Gender != "Female" {
		return fmt.Errorf("gender not valid")
	}

	regexemail := `^[^\s@]+@[^\s@]+\.[^\s@]{2,}$`
	reg, err := regexp.Compile(regexemail)
	if err != nil {
		return fmt.Errorf("compile regex error")
	}
	if !reg.MatchString(data.Email) {
		return fmt.Errorf("invalid email, email dois contien @ and . ")
	}
	if len(strings.TrimSpace(data.Password)) == 0 || len(data.Password) < 8 || len(data.Password) > 30 {
		return fmt.Errorf("valid password minimum 8")
	}

	var upercase, lowercase, number bool
	for _, i := range data.Password {
		switch {
		case unicode.IsUpper(i):
			upercase = true
		case unicode.IsLower(i):
			lowercase = true
		case unicode.IsDigit(i):
			number = true
		}
	}
	if !upercase || !lowercase || !number {
		return fmt.Errorf("le mot de pass dpit contient Uper and LOWer character et au mois number")
	}
	if data.Password != data.Confirm_password {
		return fmt.Errorf("password not match")
	}

	return nil
}

func ValidLoginData(db *sql.DB, data Login) (string, error) {
	query := "SELECT id,password FROM users WHERE nickname = ? OR email = ?"

	var userid string
	var hashpassword string

	err := db.QueryRow(query, data.Nicknameoremail, data.Nicknameoremail).Scan(&userid, &hashpassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "walo", fmt.Errorf("internl server error, try later")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(data.Password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	return userid, nil
}
