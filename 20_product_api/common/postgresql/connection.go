package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

func GetConnectionPool(context context.Context, config Config) *pgxpool.Pool {

	//Config Dosyasını Kullanarak Bir Tane Connection Strign Oluştrulur
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable statement_cache_mode=describe pool_max_conns=%s pool_max_conn_idle_time=%s",
		config.Host,
		config.Port,
		config.UserName,
		config.Password,
		config.DbName,
		config.MaxConnection,
		config.MaxConnectionIdleTime)

	//Connection string parse edilerek bir connenctionConfig referansi alınır
	connConfig, parseConfigErr := pgxpool.ParseConfig(connString)
	if parseConfigErr != nil {
		panic(parseConfigErr)
	}

	//daha sonra connectionConfig referansı ile baglantı saglanır
	conn, err := pgxpool.ConnectConfig(context, connConfig)
	if err != nil {
		log.Errorf("Unable to connect to database: %v \n", err)
		panic(err)
	}

	return conn
}
