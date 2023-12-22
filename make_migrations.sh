if [ -z "$1" ]; then
  echo "Must provide migration name";
  exit 1
else
  docker compose --profile run down
  docker volume rm discord-data-server_mariadb_data
  docker compose up db -d
  sleep 3
  atlas migrate diff "$1" --dir file://ent/migrate/migrations --to ent://ent/schema --dev-url mysql://mysql:mysql@localhost:3306/discorddata
fi