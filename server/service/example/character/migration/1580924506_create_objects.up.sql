
-- Character Type
CREATE TYPE "entity_type" AS ENUM (
  'mage',
  'familliar',
  'player-mage',
  'player-familliar',
  'starter-mage',
  'starter-familliar'
);

-- Character
CREATE TABLE "entity" (
  "id"                uuid CONSTRAINT entity_pk PRIMARY KEY DEFAULT gen_random_uuid(),
  "entity_type"       entity_type NOT NULL,
  "name"              text NOT NULL,
  "avatar"            text NOT NULL,
  "strength"          int NOT NULL DEFAULT 0,
  "dexterity"         int NOT NULL DEFAULT 0,
  "intelligence"      int NOT NULL DEFAULT 0,
  "attribute_points"  int NOT NULL DEFAULT 0,
  "experience_points" bigint NOT NULL DEFAULT 0,
  "coins"             bigint NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

-- Character Item
CREATE TABLE "entity_item" (
  "id"         uuid CONSTRAINT entity_item_pk PRIMARY KEY DEFAULT gen_random_uuid(),
  "entity_id"  uuid NOT NULL,
  "item_id"    uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

ALTER TABLE "entity_item" ADD CONSTRAINT "entity_item_entity_id_fk" FOREIGN KEY ("entity_id") REFERENCES "entity" ("id");

COMMENT ON COLUMN "entity_item"."item_id" IS 'Remote "item" service reference';

-- Character Spell
CREATE TABLE "entity_spell" (
  "id"         uuid CONSTRAINT entity_spell_pk PRIMARY KEY DEFAULT gen_random_uuid(),
  "entity_id"  uuid NOT NULL,
  "spell_id"   uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

ALTER TABLE "entity_spell" ADD CONSTRAINT "entity_spell_entity_id_fk" FOREIGN KEY ("entity_id") REFERENCES "entity" ("id");

COMMENT ON COLUMN "entity_spell"."spell_id" IS 'Remote "spell" service reference';

-- Player Character
CREATE TABLE "account_entity" (
  "id"         uuid CONSTRAINT account_entity_pk PRIMARY KEY DEFAULT gen_random_uuid(),
  "account_id" uuid NOT NULL, 
  "entity_id"  uuid NOT NULL, 
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

ALTER TABLE "account_entity" ADD CONSTRAINT "account_entity_entity_id_fk" FOREIGN KEY ("entity_id") REFERENCES "entity" ("id");

COMMENT ON COLUMN "account_entity"."account_id" IS 'Remote "account" service reference';

-- Character Instance
CREATE TABLE "entity_instance" (
  "id" uuid CONSTRAINT entity_instance_pk PRIMARY KEY DEFAULT gen_random_uuid(),
  "entity_id" uuid NOT NULL,
  "health_points" int NOT NULL DEFAULT 0,
  "action_points" int NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

ALTER TABLE "entity_instance" ADD CONSTRAINT "entity_instance_entity_id_fk" FOREIGN KEY ("entity_id") REFERENCES "entity" ("id");

COMMENT ON COLUMN "entity_instance"."health_points" IS 'Computed from attribute points';

COMMENT ON COLUMN "entity_instance"."action_points" IS 'Computed from attribute points';


