### Frontier Car Management Service 
#### This project is to provide a REST based API for car management supporting below use cases
- Add Car
- Delete Car
- Get all Car's
- Get a particular car

#### make commands
- make test
- make fmt
- make build
- make run
- make clean
- make help
- make docker-image
- make docker-run

#### Internals of project
- Created using go modules
- Developed using Go 1.12
- REST router and handler are based out of go-chi + render
- Basic funtioanlities are working as expected
- Unit test coverag of 80+ Percent for business logic functionality files. Used Mock framework as well. Negetive cases covered as well for httptest
- Better make file with docker image create and run as well
- Docker support
- In par with HTTP Status codes of 201, 200, 400, 404, 500 for different cases

### What next ?
- Test coverage
- Better Validations
- JWT based protected resource and Authentication/Authorization
- Config file and Environment variables. Plus error messages from errors.properties
- Docker and K8s support
- CI/CD
- Pagination
- Persistence
- carmgmt as a microservice

# Sample Request Response
- ###### Add A CAR
    curl -X POST \
  http://localhost:8090/api/v1/cars \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 21396afa-09a1-4784-9a8f-33ca67410b97' \
  -H 'cache-control: no-cache' \
  -d '{
	"make": "Volvo",
	"model": "m-1",
	"year": "2019"	
}'

   - RESPONSE : 
   - {
    "id": "55ec75c1-b22b-4e19-8858-c65788645718"
    }
   - Status Code : 201 Created


- ###### **Get CAR by ID**
    curl -X GET \
  http://localhost:8090/api/v1/cars/55ec75c1-b22b-4e19-8858-c65788645718 \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 4f62ede7-8bd6-4f82-b0b5-e3617b09abd0' \
  -H 'cache-control: no-cache'

 - 	Response : {
    "make": "Volvo",
    "model": "m-1",
    "year": "2019"
    }

 - Status Code : 200 Ok

- ###### **Get CAR by non existing ID**
   curl -X GET \
  http://localhost:8090/api/v1/cars/123 \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 4d5b6a2d-8892-4e26-86dd-f934f2057c59' \
  -H 'cache-control: no-cache'

 - Response : {
    "status_code": 404,
    "error_message": "Car by id : 123 not found",
    "error_code": "0001"
    }

- Status Code : 404 Not Found

- ###### **List all CAR's
 curl -X GET \
  http://localhost:8090/api/v1/cars \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: f6fc431a-5d7b-4c29-bda0-32edcd877bf0' \
  -H 'cache-control: no-cache'

 - Response : [
    {
        "id": "b792954a-cdaa-40bb-bf84-4432f9eb08bb",
        "make": "Volvo",
        "model": "m-1",
        "year": "2019"
    },
    {
        "id": "55ec75c1-b22b-4e19-8858-c65788645718",
        "make": "Volvo",
        "model": "m-1",
        "year": "2019"
    },
    {
        "id": "5ed91f0d-e3a0-4e73-b2dd-80d8c74fc3b1",
        "make": "Volvo",
        "model": "m-2",
        "year": "2020"
    }
]
- Status Code : 200 Ok

- ###### **DELETE CAR by ID**
ccurl -X DELETE \
  http://localhost:8090/api/v1/cars/5ed91f0d-e3a0-4e73-b2dd-80d8c74fc3b1 \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 887f5bb2-c7e1-4888-9fab-65b5cee21124' \
  -H 'cache-control: no-cache'

	- Response : StatusCode 204.

- ###### **DELETE CAR by non existing ID**
curl -X DELETE \
  http://localhost:8090/api/v1/cars/123 \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 4b072a48-2b9f-431e-95d5-730b0a448f23' \
  -H 'cache-control: no-cache'

- Response : {
    "status_code": 404,
    "error_message": "Car by id : 123 not found",
    "error_code": "0001"
}
- Status Code : 404 Not Found

- ###### **Add Car with invalid payload (No Year)**
curl -X POST \
  http://localhost:8090/api/v1/cars \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: be258c57-f22a-4545-b11e-ae42fc52f2d1' \
  -H 'cache-control: no-cache' \
  -d '{
	"make": "Volvo",
	"model": "m-2"
	
}'

- Response : {
    "status_code": 400,
    "error_message": "Missing required Year field in request payload",
    "error_code": "0004"
}
- Status Code : 400 Bad Request
