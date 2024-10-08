// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: rate_limits.sql

package dbsqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const bulkUpdateRateLimits = `-- name: BulkUpdateRateLimits :many
UPDATE
    "RateLimit" rl
SET
    "value" = get_refill_value(rl) - input."units",
    "lastRefill" = CASE
        WHEN NOW() - rl."lastRefill" >= rl."window"::INTERVAL THEN
            CURRENT_TIMESTAMP
        ELSE
            rl."lastRefill"
    END
FROM
    (
        SELECT
            unnest($2::text[]) AS "key",
            unnest($3::int[]) AS "units"
    ) AS input
WHERE
    rl."key" = input."key"
    AND rl."tenantId" = $1::uuid
RETURNING rl."tenantId", rl.key, rl."limitValue", rl.value, rl."window", rl."lastRefill"
`

type BulkUpdateRateLimitsParams struct {
	Tenantid pgtype.UUID `json:"tenantid"`
	Keys     []string    `json:"keys"`
	Units    []int32     `json:"units"`
}

func (q *Queries) BulkUpdateRateLimits(ctx context.Context, db DBTX, arg BulkUpdateRateLimitsParams) ([]*RateLimit, error) {
	rows, err := db.Query(ctx, bulkUpdateRateLimits, arg.Tenantid, arg.Keys, arg.Units)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*RateLimit
	for rows.Next() {
		var i RateLimit
		if err := rows.Scan(
			&i.TenantId,
			&i.Key,
			&i.LimitValue,
			&i.Value,
			&i.Window,
			&i.LastRefill,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRateLimitsForSteps = `-- name: ListRateLimitsForSteps :many
SELECT
    units, "stepId", "rateLimitKey", "tenantId"
FROM
    "StepRateLimit" srl
WHERE
    srl."stepId" = ANY($1::uuid[])
    AND srl."tenantId" = $2::uuid
`

type ListRateLimitsForStepsParams struct {
	Stepids  []pgtype.UUID `json:"stepids"`
	Tenantid pgtype.UUID   `json:"tenantid"`
}

func (q *Queries) ListRateLimitsForSteps(ctx context.Context, db DBTX, arg ListRateLimitsForStepsParams) ([]*StepRateLimit, error) {
	rows, err := db.Query(ctx, listRateLimitsForSteps, arg.Stepids, arg.Tenantid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*StepRateLimit
	for rows.Next() {
		var i StepRateLimit
		if err := rows.Scan(
			&i.Units,
			&i.StepId,
			&i.RateLimitKey,
			&i.TenantId,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRateLimitsForTenant = `-- name: ListRateLimitsForTenant :many
WITH refill AS (
    UPDATE
        "RateLimit" rl
    SET
        "value" = CASE
            WHEN NOW() - rl."lastRefill" >= rl."window"::INTERVAL THEN
                get_refill_value(rl)
            ELSE
                rl."value"
        END,
        "lastRefill" = CASE
            WHEN NOW() - rl."lastRefill" >= rl."window"::INTERVAL THEN
                CURRENT_TIMESTAMP
            ELSE
                rl."lastRefill"
        END
    WHERE
        rl."tenantId" = $1::uuid
    RETURNING "tenantId", key, "limitValue", value, "window", "lastRefill"
)
SELECT
    refill."tenantId", refill.key, refill."limitValue", refill.value, refill."window", refill."lastRefill",
    -- return the next refill time
    (refill."lastRefill" + refill."window"::INTERVAL)::timestamp AS "nextRefillAt"
FROM
    refill
`

type ListRateLimitsForTenantRow struct {
	TenantId     pgtype.UUID      `json:"tenantId"`
	Key          string           `json:"key"`
	LimitValue   int32            `json:"limitValue"`
	Value        int32            `json:"value"`
	Window       string           `json:"window"`
	LastRefill   pgtype.Timestamp `json:"lastRefill"`
	NextRefillAt pgtype.Timestamp `json:"nextRefillAt"`
}

func (q *Queries) ListRateLimitsForTenant(ctx context.Context, db DBTX, tenantid pgtype.UUID) ([]*ListRateLimitsForTenantRow, error) {
	rows, err := db.Query(ctx, listRateLimitsForTenant, tenantid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*ListRateLimitsForTenantRow
	for rows.Next() {
		var i ListRateLimitsForTenantRow
		if err := rows.Scan(
			&i.TenantId,
			&i.Key,
			&i.LimitValue,
			&i.Value,
			&i.Window,
			&i.LastRefill,
			&i.NextRefillAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const upsertRateLimit = `-- name: UpsertRateLimit :one
INSERT INTO "RateLimit" (
    "tenantId",
    "key",
    "limitValue",
    "value",
    "window"
) VALUES (
    $1::uuid,
    $2::text,
    $3::int,
    $3::int,
    COALESCE($4::text, '1 minute')
) ON CONFLICT ("tenantId", "key") DO UPDATE SET
    "limitValue" = $3::int,
    "window" = COALESCE($4::text, '1 minute')
RETURNING "tenantId", key, "limitValue", value, "window", "lastRefill"
`

type UpsertRateLimitParams struct {
	Tenantid pgtype.UUID `json:"tenantid"`
	Key      string      `json:"key"`
	Limit    int32       `json:"limit"`
	Window   pgtype.Text `json:"window"`
}

func (q *Queries) UpsertRateLimit(ctx context.Context, db DBTX, arg UpsertRateLimitParams) (*RateLimit, error) {
	row := db.QueryRow(ctx, upsertRateLimit,
		arg.Tenantid,
		arg.Key,
		arg.Limit,
		arg.Window,
	)
	var i RateLimit
	err := row.Scan(
		&i.TenantId,
		&i.Key,
		&i.LimitValue,
		&i.Value,
		&i.Window,
		&i.LastRefill,
	)
	return &i, err
}
