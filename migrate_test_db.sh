#!/bin/bash

rm -rf ./ent/migrate/migrations
atlas migrate diff test_2 \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "docker://postgres/15/test?search_path=public"

atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url "postgres://postgres:root@localhost:5432/PanelDB?search_path=public&sslmode=disable"