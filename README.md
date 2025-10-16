 # HNG STAGE 0 TASK
 
 ## Task Description
 This project makes us of the RESTful API endpoint GET / me  to return my profile information (email, name, backend stack), a random cat fact fetched from [CAT API URL](https://catfact.ninja/fact) and the current UTC timestamp in ISO 8601 format


# Getting Started
1.  Clone the Repository

``` bash
git clone https://github.com/Emeey-Lanr/hng_stage_0.git

cd repo hng_stage_0

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

Your API will be available at http://locahost:8080/me


## API DOCUMENTATION

### GET /me

Example Response 
```json
{
 "status":"success",  
"user" :{
    "email":"<Your Email>",
    "name":"<Your Name>",
    "stack":"Go/Gin"
}  ,
"timestamp":"2025-10-16T10:48:08Z",
"fact":"Tylenol and chocolate are both poisionous to cats."
} 
```

```bash
Content-Type: application/json
```

