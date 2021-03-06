# BACKEND RESTFULL API GOODJOBS OFFICE BOOKING SYSTEM WRITTEN IN GO
![goodJob Logo](goodJobs_logo.jpeg?raw=true "goodJob Logo")

### What GoodJobs Office Booking System is
goodJobs is an web-based platform to rent office spaces across Jakarta. goodjobs aims for the convenience of user when searching for a suitable working space in Jakarta’s business district areas.

## API Documentation
##### ERD
<a href="https://whimsical.com/erd-obs-A7j4gQLQWCwanKBf4kq19o">Entity Relationship Diagram - GoodJobs Office Booking System</a>

#### Swagger
<a href="https://app.swaggerhub.com/apis-docs/iskandev/goodJobs/1.0.0">GoodJobs Office Booking System API</a>


### Technology and Frameworks yang digunakan

* [GO](https://go.dev/doc/) as Programming Language  
* [echo](https://labstack.com/echo) as high performance framework
* [GORM](https://gorm.io/docs/) for the initial migration and creation of the database schema
* [MySQL](https://dev.mysql.com/doc/)
* [MongoDB](https://docs.mongodb.com/)
* Implement custom Error Handling
* Using JWT as Token via [jwt-go package](https://github.com/dgrijalva/jwt-go)
* Implement Role base authorization
* Write unit test for API endpoint and middlewares
* Using [Docker](https://docs.docker.com/) for packaging applications into containers
* Using [AWS](https://aws.amazon.com/) for Deployment

### Project Dependencies
1. Install Golang (tested with Go 1.6)
2. Install and run MongoDB service on your localhost for storing data


### How to use from this API
##### Clone the repository
```
git clone https://github.com/alvyfz/backend-goodjobs.git
cd backend-goodjobs
```


##### Build, serve or test using make

```
make build
make serve
make test
```
