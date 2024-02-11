-- Create "roles" table
CREATE TABLE "roles" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "name" character varying NOT NULL, "permissions" jsonb NULL, PRIMARY KEY ("id"));
-- Create index "roles_name_key" to table: "roles"
CREATE UNIQUE INDEX "roles_name_key" ON "roles" ("name");
-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "email" character varying NOT NULL, "password" character varying NOT NULL, "name" character varying NOT NULL, "role_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "users_roles_role" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "users_name_key" to table: "users"
CREATE UNIQUE INDEX "users_name_key" ON "users" ("name");
-- Create "api_keys" table
CREATE TABLE "api_keys" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "description" character varying NULL, "ip_addresses" jsonb NULL, "key" character varying NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "api_keys_users_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
