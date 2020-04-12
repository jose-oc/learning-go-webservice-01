# GoLang REST service

This project is just a limited playground to try and learn things in GoLang. 

## Run and test it

You can run the project with the command `go run main.go` in the directory where the project lives.

Then run these `curl` commands to use it. 
If you don't have `jq` installed, remove it from these commands:


```bash
curl --dump-header - -X GET http://localhost:3000/users

curl --show-error --silent --location --request POST 'http://localhost:3000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
        "FirstName": "Pepe",
        "LastName": "Blanco"
}' | jq

curl --show-error --silent --location --dump-header - --request POST 'http://localhost:3000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
        "FirstName": "Pepa",
        "LastName": "Zaragoza"
}'

curl --show-error --silent --location --dump-header - --request POST 'http://localhost:3000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
        "FirstName": "Pepa",
        "LastName": "Zaragoza"
}'

curl -X GET http://localhost:3000/users | jq

curl --dump-header - -X DELETE http://localhost:3000/users/2

curl --show-error --silent --location --dump-header - --request POST 'http://localhost:3000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
        "FirstName": "Pepa",
        "LastName": "Zaragoza"
}'

curl -X GET http://localhost:3000/users | jq

curl --show-error --silent --location --dump-header - --request PUT 'http://localhost:3000/users/3' \
--header 'Content-Type: application/json' \
--data-raw '{
        "ID": 3,
        "FirstName": "Pepa Juana",
        "LastName": "Zaragoza Ramírez"
}'

curl --dump-header - -X GET http://localhost:3000/users/3

curl --dump-header - -X MADEUP http://localhost:3000/users/
```