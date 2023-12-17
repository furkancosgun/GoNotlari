package infrastructure

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

func TruncateTestData(ctx context.Context, pool *pgxpool.Pool) {
	_, err := pool.Exec(ctx, "TRUNCATE products RESTART IDENTITY")
	if err != nil {
		log.Errorf("TruncateTestData: An Error Occurred: %v", err)
	} else {
		log.Info("TruncateTestData:Product Table Truncated")
	}
}
