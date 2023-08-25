-- Create "clerk_users" table
CREATE TABLE "public"."clerk_users" ("id" character varying(128) NOT NULL, "linked_identity" text NOT NULL, PRIMARY KEY ("id"));
-- Create "files" table
CREATE TABLE "public"."files" ("id" character(26) NOT NULL, "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "deleted_at" timestamp NULL, "parent_id" text NULL, "file_name" text NULL, "mime_type" text NOT NULL, "data" bytea NOT NULL, PRIMARY KEY ("id"));
-- Create enum type "user_state"
CREATE TYPE "public"."user_state" AS ENUM ('active', 'inactive');
-- Create "users" table
CREATE TABLE "public"."users" ("id" character varying(128) NOT NULL, "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "first_name" text NULL, "last_name" text NULL, "email" text NULL, "status" "public"."user_state" NULL, "profile_image_url" text NULL, PRIMARY KEY ("id"));
