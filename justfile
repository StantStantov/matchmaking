set dotenv-load := true
set dotenv-required := false

dev-compose := "./deployments/compose.dev.yaml"
prod-compose := "./deployments/compose.prod.yaml"

dev-context := "desktop-linux"
prod-context := "lesta-start"

build-dev:
  docker --context {{dev-context}} compose -f {{dev-compose}} build

build-prod:
  docker --context {{prod-context}} compose -f {{prod-compose}} build

up-dev:
  docker --context {{dev-context}} compose -f {{dev-compose}} up -d

up-prod:
  docker --context {{prod-context}} compose -f {{prod-compose}} up -d

down-dev:
  docker --context {{dev-context}} compose -f {{dev-compose}} down

down-prod:
  docker --context {{prod-context}} compose -f {{prod-compose}} down
