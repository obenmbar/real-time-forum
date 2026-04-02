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
    // Nickname
    isvalidnickname := true
    for _, i := range data.Nickname {
        if !unicode.IsDigit(i) && !unicode.IsLetter(i) {
            isvalidnickname = false
            break
        }
    }
    if !isvalidnickname || strings.Contains(data.Nickname, " ") {
        return fmt.Errorf("nickname must contain only letters and numbers without spaces")
    }
    if len(data.Nickname) < 6 || len(data.Nickname) > 15 {
        return fmt.Errorf("nickname must be between 6 and 15 characters long")
    }

    // First Name & Last Name
    if len(data.FirstName) < 2 || len(data.FirstName) > 15 || strings.Contains(data.FirstName, " ") {
        return fmt.Errorf("please enter a valid first name (2-15 characters, no spaces)")
    }
    if len(data.LastName) < 2 || len(data.LastName) > 15 || strings.Contains(data.LastName, " ") {
        return fmt.Errorf("please enter a valid last name (2-15 characters, no spaces)")
    }

    // Age
    age, err := strconv.Atoi(data.Age)
    if err != nil {
        return fmt.Errorf("invalid age format")
    }
    if age < 10 || age > 120 {
        return fmt.Errorf("age must be between 10 and 120 years")
    }

    // Gender
    if data.Gender != "Male" && data.Gender != "Female" {
        return fmt.Errorf("please select a valid gender")
    }

    // Email
    regexemail := `^[^\s@]+@[^\s@]+\.[^\s@]{2,}$`
    reg, _ := regexp.Compile(regexemail)
    if !reg.MatchString(data.Email) {
        return fmt.Errorf("please provide a valid email address (e.g., name@example.com)")
    }

    // Password complexity
    if len(data.Password) < 8 || len(data.Password) > 30 {
        return fmt.Errorf("password must be between 8 and 30 characters")
    }

    var upercase, lowercase, number bool
    for _, i := range data.Password {
        switch {
        case unicode.IsUpper(i): upercase = true
        case unicode.IsLower(i): lowercase = true
        case unicode.IsDigit(i): number = true
        }
    }
    if !upercase || !lowercase || !number {
        return fmt.Errorf("password must include uppercase, lowercase letters and at least one number")
    }

    if data.Password != data.Confirm_password {
        return fmt.Errorf("passwords do not match, please check again")
    }

    return nil
}

func ValidLoginData(db *sql.DB, data Login) (string, error) {
    // كنقلبو بـ nickname أو email
    query := "SELECT id, password FROM users WHERE nickname = ? OR email = ?"

    var userid string
    var hashpassword string

    err := db.QueryRow(query, data.Nicknameoremail, data.Nicknameoremail).Scan(&userid, &hashpassword)
    if err != nil {
        if err == sql.ErrNoRows {
            // ميساج موحد للحماية
            return "", fmt.Errorf("invalid username/email or password")
        }
        return "walo", fmt.Errorf("technical issue: please try again in a few moments")
    }

    err = bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(data.Password))
    if err != nil {
        return "", fmt.Errorf("invalid username/email or password")
    }

    return userid, nil
}