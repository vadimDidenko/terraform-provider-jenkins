NAME=terraform-provider-jenkins


build:
	go build -o $(NAME) main.go

check:
	make build
	terraform init
	terraform plan