# Translation Manager
An API endpoint that translate arabic sentences in dialogue to english using OpenAi.

## Technologies
- **GoLang**

### Structure
- **Gateway**: Contains any external gateway.
- **Gateway.factory**: Factory pattern that create instance of translator gateway.
- **Gateway.open_ai**: OpenAi gateway to translate texts.
- **Internal**: All internal logic for translation manager.
- **Internal.handler**: Request handler. e.g.HTTP, gRPC.
- **Internal.model**: Contains request and response models.
- **Internal.service**: Service that responsible for selecting the translation gateway and passing the request to it.
- **Internal.pkg**: Shared packages e.g.request validator.


## How to run
- **Manual**
    - Install Golang.
    - Run `cd transaltion_manager/`
    - Run `go run main.go`
- **Docker**
  - 
  
## Translate endpoint
- **cURL**
  ```bash
  curl --location --request POST 'localhost:8080/api/v1/translation' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "dialogue":
  [{"speaker":"Samar","time":"00:00:04","sentence":"صباح الخير"},
  {"speaker":"Maria","time":"00:00:09","sentence":"Good Morning"}]
  }'
