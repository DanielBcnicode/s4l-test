# S4L test

This is the implementations for the S4L technical test.

The main problem is to find the best combination of bookings
to maximize the profit when we receive more than one possible
solution in the end-point payload.

To solve this problem I decided to use one N-TREE structure to store
the data adding weights to the nodes (the profit value).

To find the best solution I created the tree using as 'Brothers' the 
incompatible bookings (the overlapped bookings).

I compute the tree, and then I get all the leafs and
get the best branch (from root to this leaf) as a result.

## Requirements
- Docker
- Docker-Compose
- An internet connection

## Composition
- Main booking microservice

## How to execute the tests, build and process
### Run tests
Execute this commands in the main folder:
```bash
go mod vendor
go test -cover ./...
```
The current results are:
```bash
?       github.com/danielbcnicode/timeslot/cmd  [no test files]
ok      github.com/danielbcnicode/timeslot/internal     0.002s  coverage: 89.3% of statements
ok      github.com/danielbcnicode/timeslot/pkg/booking  0.002s  coverage: 80.8% of statements

```
### Build
Execute :
```bash
docker-compose -f Docker-compose.yaml build
```
### Up and Run
```
docker-compose -f Docker-compose.yaml up -d
```

## End-point details
The microservice has 2 endpoints and the http server is listening on the port 8088 by default.
We can specify the port in the .env file

### In the folder _postman_ there is a Postman json that can be used to test the end-points

## /stats end-point
You can use `curl`:
```bash
curl --location --request POST 'localhost:8088/stats' \
--header 'Content-Type: application/json' \
--data-raw '[
{
"request_id":"bookata_XY123",
"check_in":"2020-01-01",
"nights":5,
"selling_rate":200,
"margin":20
},
{
"request_id":"kayete_PP234",
"check_in":"2020-01-04",
"nights":4,
"selling_rate":156,
"margin":22
}
]'
```

## /maximize end-point
```bash
curl --location --request POST 'localhost:8088/maximize' \
--header 'Content-Type: application/json' \
--data-raw '[
{
"request_id":"bookata_XY123",
"check_in":"2020-01-01",
"nights":5,
"selling_rate":200,
"margin":20
},
{
"request_id":"kayete_PP234",
"check_in":"2020-01-04",
"nights":4,
"selling_rate":156,
"margin":5
},
{
"request_id":"atropote_AA930",
"check_in":"2020-01-04",
"nights":4,
"selling_rate":150,
"margin":6
},
{
"request_id":"acme_AAAAA",
"check_in":"2020-01-10",
"nights":4,
"selling_rate":160,
"margin":30
}
]'
```
