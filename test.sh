source ./.env
go test ./domains/users/... -coverprofile=cover.out && go tool cover -html=cover.out
go test ./domains/auth/... -coverprofile=cover.out && go tool cover -html=cover.out
go test ./domains/register/... -coverprofile=cover.out && go tool cover -html=cover.out
go test ./domains/product/... -coverprofile=cover.out && go tool cover -html=cover.out
go test ./domains/cart/... -coverprofile=cover.out && go tool cover -html=cover.out
<<<<<<< HEAD
go test ./domains/category/... -coverprofile=cover.out && go tool cover -html=cover.out
=======
>>>>>>> 9c7ae9202acf126cf3619b11f106b5f1c483be8e
