 # HNG STAGE 0 TASK
 
 ## Task Description
 This project makes us of the RESTful API endpoint / me  to return my profile information (email, name, backend stack), a random cat fact fetched from [CAT API URL](https://catfact.ninja/fact) and the current UTC timestamp in ISO 8601 format


# Getting Started
1.  Clone the Repository

``` bash
git clone url

cd repo name

```

2.  Install Dependencies
``` bash
go mod tidy

```

3. Environment Variables
create a .env file in the project root
``` env
PORT = 8080
EMAIL = <your email>
CATURL = https://catfact.ninja/fact
NAME = <Your name>
```

4. Run Server locally
``` bash
go run main.go
```

Your API will be available at http://lochost:8080/me


Example Response 
```json
{
 "status":"",  
"user" :{
    "email":"<Your Email>",
    "name":"<Your Name>",
    "stack":"GO/GIN"
}  
"timestamp":"",
"fact":""
} 
```



 <!-- Instructions to run locally -->

<!-- Environmental variables  -->