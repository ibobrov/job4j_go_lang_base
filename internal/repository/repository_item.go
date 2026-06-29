package repository

import (
	"context"
	"fmt"

	"job4j.ru/go-lang-base/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepoPg struct {
	pool *pgxpool.Pool
}

func NewRepoPg(pool *pgxpool.Pool) *RepoPg {
	return &RepoPg{pool: pool}
}

func (r *RepoPg) Create(ctx context.Context, it domain.Item) (domain.Item, error) {
	_, err := r.pool.Exec(
		ctx,
		`insert into items(id, name) values($1, $2)`,
		it.ID, it.Name,
	)
	if err != nil {
		return domain.Item{}, fmt.Errorf("r.pool.Exec: %w", err)
	}
	return it, nil
}

func (r *RepoPg) List(ctx context.Context) ([]domain.Item, error) {
	rows, err := r.pool.Query(ctx, `select id, name from items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.Item
	for rows.Next() {
		var item domain.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *RepoPg) Get(ctx context.Context, id string) (domain.Item, error) {
	var it domain.Item
	err := r.pool.QueryRow(
		ctx,
		`select id, name from items where id = $1`,
		id,
	).Scan(&it.ID, &it.Name)

	return it, err
}
