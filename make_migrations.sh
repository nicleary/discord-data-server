if [ -z "$1" ]; then
  echo "Must provide migration name";
  exit 1
else
  export MIGRATION_NAME=$1 && docker compose up atlas_make_migrations --build;
fi

