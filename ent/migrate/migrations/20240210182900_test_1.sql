-- Create "games" table
CREATE TABLE "games" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "description" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "games_name_key" to table: "games"
CREATE UNIQUE INDEX "games_name_key" ON "games" ("name");
-- Create "nodes" table
CREATE TABLE "nodes" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "ipv4_address" character varying NOT NULL, "fqdn" character varying NOT NULL, "skyhook_version" character varying NOT NULL, "os" character varying NOT NULL, "cpu" character varying NOT NULL, "cpu_base_clock" bigint NOT NULL, "cores" bigint NOT NULL, "logical_cores" bigint NOT NULL, "ram" bigint NOT NULL, "storage" bigint NOT NULL, PRIMARY KEY ("id"));
-- Create index "nodes_fqdn_key" to table: "nodes"
CREATE UNIQUE INDEX "nodes_fqdn_key" ON "nodes" ("fqdn");
-- Create index "nodes_ipv4_address_key" to table: "nodes"
CREATE UNIQUE INDEX "nodes_ipv4_address_key" ON "nodes" ("ipv4_address");
-- Create "servers" table
CREATE TABLE "servers" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "ram" bigint NOT NULL, "storage" bigint NOT NULL, "logical_cores" bigint NOT NULL, "port" smallint NOT NULL, "game_games" uuid NULL, "node_nodes" uuid NULL, "user_owners" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "servers_games_games" FOREIGN KEY ("game_games") REFERENCES "games" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "servers_nodes_nodes" FOREIGN KEY ("node_nodes") REFERENCES "nodes" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "servers_users_owners" FOREIGN KEY ("user_owners") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
