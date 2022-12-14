gen_server:
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate types ./swagger/auth/swagger.yaml > ./internal/generated/spec/auth/types.gen.go
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate spec ./swagger/auth/swagger.yaml > ./internal/generated/spec/auth/spec.gen.go
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate chi-server ./swagger/auth/swagger.yaml > ./internal/generated/spec/auth/chi_server.gen.go

	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate types ./swagger/product/swagger.yaml > ./internal/generated/spec/product/types.gen.go
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate spec ./swagger/product/swagger.yaml > ./internal/generated/spec/product/spec.gen.go
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate chi-server ./swagger/product/swagger.yaml > ./internal/generated/spec/product/chi_server.gen.go

	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate types ./swagger/gateway/swagger.yaml > ./internal/generated/spec/gateway/types.gen.go
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate spec ./swagger/gateway/swagger.yaml > ./internal/generated/spec/gateway/spec.gen.go
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate chi-server ./swagger/gateway/swagger.yaml > ./internal/generated/spec/gateway/chi_server.gen.go

	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate types ./swagger/image/swagger.yaml > ./internal/generated/spec/image/types.gen.go
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate spec ./swagger/image/swagger.yaml > ./internal/generated/spec/image/spec.gen.go
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate chi-server ./swagger/image/swagger.yaml > ./internal/generated/spec/image/chi_server.gen.go

	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate types ./swagger/cart/swagger.yaml > ./internal/generated/spec/cart/types.gen.go
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate spec ./swagger/cart/swagger.yaml > ./internal/generated/spec/cart/spec.gen.go
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0 -package spec -generate chi-server ./swagger/cart/swagger.yaml > ./internal/generated/spec/cart/chi_server.gen.go

	go mod tidy

gen_orm:
	-go run entgo.io/ent/cmd/ent@v0.10.1 generate --target ./internal/repository/ent ./internal/repository/schema 

start:
	docker-compose up --build