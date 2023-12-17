package infrastructure

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

var INSERT_QUERY = `
INSERT INTO products (name,price,discount,store) VALUES
('KALEM',10,1,'www.trendyol.com'),
('DEFTER',20,1,'www.trendyol.com'),
('SILGI',5,1,'www.trendyol.com'),
('UHU',3,1,'www.trendyol.com'),
('YORGAN',100,1,'www.hepsiburada.com'),
('YASTIK',50,1,'www.hepsiburada.com'),
('FORD CONNECT 1.8 TDCİ 110HP GLX',800000,0,'www.sahibinden.com'),
('HP VICTUS 16',21021,1,'www.hepsiburada.com');
`

func TestDataInitialize(ctx context.Context, pool *pgxpool.Pool) {
	insertProductResult, err := pool.Exec(ctx, INSERT_QUERY)
	if err != nil {
		log.Error(insertProductResult)
	} else {
		log.Infof("Products Data Created With %d Rows", insertProductResult.RowsAffected())
	}
}
