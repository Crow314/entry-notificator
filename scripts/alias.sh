project_name='entry-notificator-development'

alias docker-compose="docker-compose -f deployments/development/docker-compose.yml -p $project_name"

alias build="docker-compose build"
alias up="docker-compose up app"
alias stop="docker-compose stop"
alias down="docker-compose down"

alias migrate="docker-compose run --rm migrate"
alias migrate-run="migrate -database postgres://postgres:postgres@postgres:5432/development?sslmode=disable -path /migrations"
alias migrate-create="migrate create -ext sql -dir /migrations -seq"

alias run="docker-compose run --rm"
alias exec="docker-compose exec"
