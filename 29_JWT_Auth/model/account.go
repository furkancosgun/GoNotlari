package model

import (
	"errors"
	"jwt_auth/common"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// JWT struct
type Token struct {
	UserId   uint
	Username string
	jwt.StandardClaims
}

// Kullanıcı tablosu struct
type Account struct {
	gorm.Model        // Migrasyon işlemi yapılırken, veritabanı üzerinde accounts tablosu yaratılması için belirtilir
	Email      string `json:"email"`
	Password   string `json:"password"`
	Token      string `json:"token";sql:"-"`
}

// Gelen bilgileri doğrulama fonksiyonu
func (account *Account) Validate() error {

	if !strings.Contains(account.Email, "@") {
		return errors.New(common.INVALID_EMAIL_ADDR)
	}

	if len(account.Password) < 8 {
		return errors.New(common.PASSSWORD_MUST_BE_GREATHER_THAN)
	}

	temp := &Account{}

	// Email adresinin kayıtlı olup olmadığı kontrol edilir
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil {
		return err
	}

	if temp.Email != "" {
		return errors.New(common.EMAIL_ALREADY_USING)
	}

	return nil
}

// Kullanıcı hesabı yaratma fonksiyonu
func (account *Account) Register() error {

	err := account.Validate()
	if err != nil {
		return err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return errors.New(common.CONNECTION_ERROR)
	}

	// Yaratılan hesap için JWT oluşturulur
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = "" // Yanıt içerisinden parola silinir
	return nil
}

// Giriş yapma fonksiyonu
func (account *Account) Login() error {

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(account).Error
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(account.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { // Parola eşleşmedi
		return errors.New(common.EMAIL_OR_PASSWORD_INCORRECT)
	}

	// Giriş başarılı
	account.Password = ""

	// JWT yaratılır
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString // JWT yanıta eklenir
	return nil
}

func GetAccountById(u uint) (*Account, error) {
	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" {
		return nil, errors.New(common.ACCOUNT_NOT_FOUND)
	}
	acc.Password = ""
	return acc, nil
}

func GetAllAccounts() []Account {
	accounts := []Account{}

	GetDB().Table("accounts").Find(&accounts)

	for _, acc := range accounts {
		acc.Password = ""
	}

	return accounts
}
