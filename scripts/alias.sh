project_name='entry-notificator-development'

alias docker-compose="docker-compose -f deployments/development/docker-compose.yml -p $project_name"

alias build="docker-compose build"
alias up="docker-compose up app"
alias stop="docker-compose stop"
alias down="docker-compose down"

alias migrate="docker-compose run --rm migrate"

alias run="docker-compose run --rm"
alias exec="docker-compose exec"
