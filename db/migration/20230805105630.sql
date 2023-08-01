-- Create "clerk_users" table
CREATE TABLE "clerk_users" ("id" character varying(128) NOT NULL, "linked_identity" text NULL, PRIMARY KEY ("id"));
-- Create "files" table
CREATE TABLE "files" ("id" character(26) NOT NULL, "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "deleted_at" timestamp NULL, "parent_id" text NULL, "file_name" text NULL, "mime_type" text NOT NULL, "data" bytea NOT NULL, PRIMARY KEY ("id"));
-- Create "users" table
CREATE TABLE "users" ("id" character varying(128) NOT NULL, "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "first_name" text NULL, "last_name" text NULL, "email" text NULL, "status" text NULL, "profile_image_url" text NULL, PRIMARY KEY ("id"));
