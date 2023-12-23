package common

var NOT_AUTH_REQUIRED_LIST = []string{"/api/v1/auth/login", "/api/v1/auth/register"}

const TOKEN_MUST_BE_SENT = "Token gönderilmelidir!"
const INVALID_TOKEN = "Geçersiz token!"
const INVALID_REQUEST = "Geçersiz istek. Lütfen kontrol ediniz!"
const INVALID_EMAIL_ADDR = "Email adresi hatalidir!"
const PASSSWORD_MUST_BE_GREATHER_THAN = "Şifreniz en az 8 karakter olmalidir!"
const CONNECTION_ERROR = "Bağlanti hatasi oluştu. Lütfen tekrar deneyiniz!"
const EMAIL_ALREADY_USING = "Email adresi başka bir kullanici tarafindan kullaniliyor."
const EMAIL_OR_PASSWORD_INCORRECT = "Email veya şifre yanlis!"
const ACCOUNT_NOT_FOUND = "Hesap Bulunamadi!"
