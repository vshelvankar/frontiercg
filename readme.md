### Frontier Car Management Service 
#### This project is to provide a REST based API for car management supporting below use cases
- Add Car
- Delete Car
- Get all Car's
- Get a particular car

#### Internals of project
- Created using go modules
- Developed using Go 1.12
- REST router and handler are based out of go-chi + render
- Basic funtioanlities are working as expected
- WIP for unit test cases for handlers
- WIP for make file
- WIP for docker file

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
    >   http://localhost:8080/api/v1/cars/ \
    >   -H 'Accept: application/json' \
    >   -H 'Content-Type: application/json' \
    >   -H 'Postman-Token: 74a24154-f415-436d-8948-cbdac98e15a4' \
    >   -H 'cache-control: no-cache' \
    >   -d '{
    > "make" : "volvo",
    > "model": "v-1",
    > "year": "2001"
    > }'

   - RESPONSE : {"id":"53676f34-291f-4283-bf42-59d860b0cd72"}


- ###### **Get CAR by ID**
    curl -X GET \
>   http://localhost:8080/api/v1/cars/53676f34-291f-4283-bf42-59d860b0cd72 \
>   -H 'Accept: application/json' \
>   -H 'Content-Type: application/json' \
>   -H 'Postman-Token: 3629d2ed-f867-4515-a318-ee3f7630eabc' \
>   -H 'cache-control: no-cache'

 - 	Response : {"make":"volvo","model":"v-1","year":"2001"}

- ###### **Get CAR by non existing ID**
   curl -X GET \
>   http://localhost:8080/api/v1/cars/123 \
>   -H 'Accept: application/json' \
>   -H 'Content-Type: application/json' \
>   -H 'Postman-Token: 3e19261b-351a-4739-a076-14aef14de4e7' \
>   -H 'cache-control: no-cache'

 - Response : {"status_code":400,"error_message":"Car by id : 123 does not exist"}

- ###### **List all CAR's
 curl -X GET \
>   http://localhost:8080/api/v1/cars \
>   -H 'Accept: application/json' \
>   -H 'Content-Type: application/json' \
>   -H 'Postman-Token: 9abb0f3d-901c-4576-a74c-9bcf66697339' \
>   -H 'cache-control: no-cache'

 - Response : [{"id":"53676f34-291f-4283-bf42-59d860b0cd72","make":"volvo","model":"v-1","year":"2001"},{"id":"69908105-3d39-4deb-91ef-8adccd6ebb14","make":"volvo","model":"v-2","year":"2002"},{"id":"7fa743bd-334c-4011-b798-f524fb3061b2","make":"volvo","model":"v-3","year":"2003"},{"id":"e75e7400-4337-4273-9b16-abe600f71c48","make":"hyundai","model":"h-1","year":"2004"}]

- ###### **DELETE CAR by ID**
curl -X DELETE \
>   http://localhost:8080/api/v1/cars/e75e7400-4337-4273-9b16-abe600f71c48 \
>   -H 'Accept: application/json' \
>   -H 'Content-Type: application/json' \
>   -H 'Postman-Token: c3ff0b5e-f6cc-4e6e-884a-4eab4757889a' \
>   -H 'cache-control: no-cache'

	- Response : StatusCode 204.

- ###### **DELETE CAR by non existing ID**
curl -X DELETE \
>   http://localhost:8080/api/v1/cars/123 \
>   -H 'Accept: application/json' \
>   -H 'Content-Type: application/json' \
>   -H 'Postman-Token: d70c1be4-e3b8-472a-82c1-4dffe328acf3' \
>   -H 'cache-control: no-cache'

	- Response : {"status_code":400,"error_message":"Cannot delete car. Car by id : 123 does not exist"}
