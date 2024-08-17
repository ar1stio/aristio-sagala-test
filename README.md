# aristio-sagala-test
test recruitment

## Run Application
- go mod tidy (for download plugin)
- go run main.go (for run app)

## Run Test
go test ./tests -v

## Config Database
please change for your config db
user:password@tcp(127.0.0.1:3306)/todolist?parseTime=true

## execute query for dbname todolist
query DDL table tasks in folder database

## please copy curl in postman or insomnia
- created tasks
curl --request POST \
  --url http://localhost:3000/tasks \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/2023.5.8' \
  --data '{"title":"Test Task","description":"Test Description","due_date":"2024-08-17T00:00:00Z"}'
- update tasks
curl --request PUT \
  --url http://localhost:3000/tasks/2 \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/2023.5.8' \
  --data '{"title":"Test Task","description":"Test Description","due_date":"2024-08-17T00:00:00Z"}'
- get data tasks
curl --request GET \
  --url http://localhost:3000/tasks/2 \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/2023.5.8'
- soft delete tasks
curl --request DELETE \
  --url http://localhost:3000/tasks/2 \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/2023.5.8'