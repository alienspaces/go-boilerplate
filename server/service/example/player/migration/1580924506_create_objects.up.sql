-- type account provider 
CREATE TYPE "provider" AS ENUM (
  'anonymous',
  'apple',
  'facebook',
  'github',
  'google',
  'twitter'
);

-- type account role
CREATE TYPE "role" AS ENUM (
  'default',
  'administrator'
);

-- table account
CREATE TABLE "account" (
  "id" uuid CONSTRAINT account_pk PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "email" text NOT NULL,
  "provider" provider NOT NULL,
  "provider_account_id" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

-- table account role
CREATE TABLE "account_role" (
  "id" uuid CONSTRAINT account_role_pk PRIMARY KEY DEFAULT gen_random_uuid(),
  "account_id" uuid NOT NULL,
  "role" role NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

ALTER TABLE "account_role" ADD CONSTRAINT "account_role_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "account" ("id");
