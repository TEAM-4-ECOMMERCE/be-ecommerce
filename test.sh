source ./.env
go test ./domains/auth/... -coverprofile=cover.out && go tool cover -html=cover.out
go test ./domains/product/... -coverprofile=cover.out && go tool cover -html=cover.out