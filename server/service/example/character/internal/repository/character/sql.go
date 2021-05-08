package character

var createOneSQL = `
INSERT INTO character (
	id,
	character_type,
	name,
	avatar,
	strength,
	dexterity,
	intelligence,
	attribute_points,
	experience_points,
	coins,
	created_at
) VALUES (
	:id,
	:character_type,
	:name,
	:avatar,
	:strength,
	:dexterity,
	:intelligence,
	:attribute_points,
	:experience_points,
	:coins,
	:created_at
)
RETURNING *
`

var updateOneSQL = `
UPDATE character SET
	id                = :id,
	character_type    = :character_type,      
    name              = :name,
    avatar            = :avatar,
    strength          = :strength,
    dexterity         = :dexterity,
    intelligence      = :intelligence,
    attribute_points  = :attribute_points,
    experience_points = :experience_points,
    coins             = :coins,
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
