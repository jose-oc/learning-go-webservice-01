You can test the webserver with `curl`


```bash
curl --dump-header - -X GET http://localhost:3000/users
curl --dump-header - -X GET http://localhost:3000/users/
curl --dump-header - -X GET http://localhost:3000/users/1
curl --dump-header - -X POST http://localhost:3000/users/
curl --dump-header - -X PUT http://localhost:3000/users/2
curl --dump-header - -X DELETE http://localhost:3000/users/2
curl --dump-header - -X MADEUP http://localhost:3000/users/
```