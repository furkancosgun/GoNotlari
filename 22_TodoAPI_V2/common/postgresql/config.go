package postgresql

// Postgresql Baglantısı İçin Yapı
type Config struct {
	Host                  string
	Port                  string
	UserName              string
	Password              string
	DbName                string
	MaxConnection         string
	MaxConnectionIdleTime string
}
