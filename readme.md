# Install and setup

go get -u github.com/spf13/cobra@latest

go install github.com/spf13/cobra-cli@latest

cobra-cli init

create command
cobra-cli add << cmd name >>

create sub command to an existing one
cobra-cli add -p < < parent cmd name > > < < child cmd name > > 

# run

go run main.go compare env x
