package playercharacter

var createOneSQL = `
INSERT INTO player_character (
	id,
	player_id,
	character_id,
	created_at
) VALUES (
	:id,
	:player_id,
	:character_id,
	:created_at
)
RETURNING *
`

var updateOneSQL = `
UPDATE player_character SET
    id                = :id,
    player_id        = :player_id,
    character_id      = :character_id,
    updated_at        = :updated_at
WHERE id 		      = :id
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
