```
swagger generate spec -o ./swaggerui/swagger.json --scan-models
// swagger serve -F=swagger swagger.json

go run . -kafka_brokers='127.0.0.1:9092'

// Navigate to http://localhost:8000/swaggerui/
```

Resources  
https://goswagger.io/  
https://www.ribice.ba/swagger-golang/  
https://www.ribice.ba/serving-swaggerui-golang/  
https://github.com/ribice/golang-swaggerui-example  
https://medium.com/@ribice/serve-swaggerui-within-your-golang-application-5486748a5ed4  
https://github.com/swagger-api/swagger-ui  
https://ops.tips/blog/a-swagger-golang-hello-world/
