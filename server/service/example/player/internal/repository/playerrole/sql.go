package accountrole

var createOneSQL = `
INSERT INTO account_role (
    id,
    account_id,
    role,
    created_at
) VALUES (
    :id,
    :account_id,
    :role,
    :created_at
)
RETURNING *
`

var updateOneSQL = `
UPDATE account_role SET
    account_id   = :account_id,
    role         = :role,
    updated_at   = :updated_at
WHERE id = :id
AND   deleted_at IS NULL
RETURNING *
`

// CreateOneSQL -
func (r *Repository) CreateOneSQL() string {
	return createOneSQL
}

// UpdateOneSQL -
func (r *Repository) UpdateOneSQL() string {
	return updateOneSQL
}

// OrderBy -
func (r *Repository) OrderBy() string {
	return "created_at desc"
}
