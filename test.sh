source ./.env
go test ./domains/users/... -coverprofile=cover.out && go tool cover -html=cover.out
go test ./domains/auth/... -coverprofile=cover.out && go tool cover -html=cover.out