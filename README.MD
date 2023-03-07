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
```shell
go mod vendor
go test -cover ./...
```
The current results are:
```shell
?       github.com/danielbcnicode/timeslot/cmd  [no test files]
ok      github.com/danielbcnicode/timeslot/internal     0.002s  coverage: 89.3% of statements
ok      github.com/danielbcnicode/timeslot/pkg/booking  0.002s  coverage: 80.8% of statements

```
### Build
Execute :
```shell
docker-compose -f Docker-compose.yaml build
```
### Up and Run
```shell
docker-compose -f Docker-compose.yaml up -d
```

## End-point details
The microservice has 2 endpoints and the http server is listening on the port 8088 by default.
We can specify the port in the .env file

### In the folder _postman_ there is a Postman json that can be used to test the end-points

### /stats end-point
You can use `curl`:
```shell
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

### /maximize end-point
```shell
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
## Loading test
To make a loading test I prepared in the load_test folder a file to be used with K6 https://k6.io/

The duration test is 2 minutes and it uses until 1000 concurrent calls.

Requirement: Have the service up and running throughout docker-compose.

To launch the test you need execute the next command:
```shell
 docker run --rm -i --network="test"  grafana/k6 run -  <./load_test/test.js
```

In my machine the result is:
```
running (2m00.0s), 0000/1000 VUs, 2340648 complete and 0 interrupted iterations
default ✓ [ 100% ] 0000/1000 VUs  2m0s

     ✓ maximize is successfully

     checks.........................: 100.00% ✓ 4681296      ✗ 0      
     data_received..................: 880 MB  7.3 MB/s
     data_sent......................: 2.2 GB  19 MB/s
     http_req_blocked...............: avg=5.37µs  min=648ns   med=1.78µs  max=118.98ms p(90)=2.44µs  p(95)=2.91µs 
     http_req_connecting............: avg=645ns   min=0s      med=0s      max=75.96ms  p(90)=0s      p(95)=0s     
   ✓ http_req_duration..............: avg=16.5ms  min=61.46µs med=10.22ms max=252.66ms p(90)=39.71ms p(95)=53.86ms
       { expected_response:true }...: avg=16.5ms  min=61.46µs med=10.22ms max=252.66ms p(90)=39.71ms p(95)=53.86ms
     http_req_failed................: 0.00%   ✓ 0            ✗ 4681296
     http_req_receiving.............: avg=68.89µs min=6.12µs  med=18.98µs max=138.12ms p(90)=28.39µs p(95)=43.95µs
     http_req_sending...............: avg=51.88µs min=3.96µs  med=10.54µs max=145.82ms p(90)=14.94µs p(95)=21.08µs
     http_req_tls_handshaking.......: avg=0s      min=0s      med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=16.38ms min=45.95µs med=10.14ms max=252.6ms  p(90)=39.48ms p(95)=53.56ms
     http_reqs......................: 4681296 39010.534719/s
     iteration_duration.............: avg=38.36ms min=220.2µs med=29.94ms max=401.5ms  p(90)=83.53ms p(95)=105.1ms
     iterations.....................: 2340648 19505.267359/s
     vus............................: 3       min=3          max=1000 
     vus_max........................: 1000    min=1000       max=1000 


running (2m00.0s), 0000/1000 VUs, 2340648 complete and 0 interrupted iterations
default ✓ [ 100% ] 0000/1000 VUs  2m0s

```