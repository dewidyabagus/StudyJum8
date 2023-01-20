# Download Instalasi
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Register Environment Variable
# Linux / Mac
export PATH="$PATH:/usr/local/go/bin"
# Windows
set PATH="$PATH:/usr/local/go/bin"

# REGISTER ENV FILES
export DB_DSN="mysql://root:root123@tcp(localhost:3307)/db?charset=utf8&parseTime=True&loc=Local&x-no-lock=true"
# driver://user:password@tcp(host:password)/db_name?charset=utf8&parseTime=True&loc=Local&x-no-lock=true

# PENAMAAN FILE MIGRATION
# 0001_create_user_table.down
# <sequencial>_<costom name>.down.sql
# <sequencial>_<costom name>.up.sql

# MENJALANKAN MIGRATION
# Cara panjang up schema tables
./migrate -path ./migrations -database mysql://root:root123@tcp(localhost:3307)/db?charset=utf8&parseTime=True&loc=Local&x-no-lock=true up

# Cara pendek up schema tables
./migrate -path ./migrations -database ${DB_DSN} up

# Cara pendek down schema tables
./migrate -path ./migrations -database ${DB_DSN} up

# Untuk menghapus semuanya schema tables pada database
./migrate -path ./migrations -database ${DB_DSN} drop -f

# Jika file binary ini dijalankan ketika table schema_migrations belum ada akan auto create dan
# versioning migration akan disimpan disana

# Berpindah versi migration
./migrate -path ./migrations -database ${DB_DSN} goto 1
