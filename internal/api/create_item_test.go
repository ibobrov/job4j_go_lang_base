package api_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"

	"job4j.ru/go-lang-base/internal/api"
	"job4j.ru/go-lang-base/internal/repository"
)

func TestServer_CreateItem(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	pgContainer, err := postgres.Run(
		ctx,
		"postgres:16",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("password"),
	)
	require.NoError(t, err)
	t.Cleanup(func() { _ = pgContainer.Terminate(ctx) })

	dsn, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	db, err := sql.Open("pgx", dsn)
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })

	require.Eventually(t, func() bool {
		ctxPing, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
		return db.PingContext(ctxPing) == nil
	}, 30*time.Second, 500*time.Millisecond)

	err = goose.SetDialect("postgres")
	require.NoError(t, err)
	require.NoError(t, goose.Up(db, "../../migrations"))

	pool, err := pgxpool.New(ctx, dsn)
	require.NoError(t, err)
	t.Cleanup(pool.Close)

	repo := repository.NewRepoPg(pool)
	server := api.NewServer(repo)

	app := fiber.New()
	app.Post("/items", server.CreateItem)

	body := map[string]string{
		"name": "first item",
	}
	data, err := json.Marshal(body)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		"/items",
		bytes.NewReader(data),
	)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 5000)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	items, err := repo.List(ctx)
	require.NoError(t, err)
	require.Len(t, items, 1)
	assert.Equal(t, "first item", items[0].Name)
}
