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
-- Create "games" table
CREATE TABLE "games" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "description" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "games_name_key" to table: "games"
CREATE UNIQUE INDEX "games_name_key" ON "games" ("name");
-- Create "nodes" table
CREATE TABLE "nodes" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "ipv4_address" character varying NOT NULL, "fqdn" character varying NOT NULL, "skyhook_version" character varying NULL, "skyhook_api_key" character varying NOT NULL, "os" character varying NULL, "cpu" character varying NULL, "cpu_base_clock" bigint NULL, "cores" bigint NULL, "logical_cores" bigint NULL, "ram" bigint NULL, "storage" bigint NULL, PRIMARY KEY ("id"));
-- Create index "nodes_fqdn_key" to table: "nodes"
CREATE UNIQUE INDEX "nodes_fqdn_key" ON "nodes" ("fqdn");
-- Create index "nodes_ipv4_address_key" to table: "nodes"
CREATE UNIQUE INDEX "nodes_ipv4_address_key" ON "nodes" ("ipv4_address");
-- Create index "nodes_skyhook_api_key_key" to table: "nodes"
CREATE UNIQUE INDEX "nodes_skyhook_api_key_key" ON "nodes" ("skyhook_api_key");
-- Create "servers" table
CREATE TABLE "servers" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "ram" bigint NOT NULL, "storage" bigint NOT NULL, "logical_cores" bigint NOT NULL, "port" smallint NOT NULL, "crater_provider" character varying NOT NULL, "crater" character varying NOT NULL, "crater_variant" character varying NOT NULL, "crater_options" jsonb NULL, "game_games" uuid NULL, "node_nodes" uuid NULL, "user_owners" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "servers_games_games" FOREIGN KEY ("game_games") REFERENCES "games" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "servers_nodes_nodes" FOREIGN KEY ("node_nodes") REFERENCES "nodes" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "servers_users_owners" FOREIGN KEY ("user_owners") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
