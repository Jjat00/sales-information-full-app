# Sales and buyers full app  using Golang and Vue

## Launch backend API rest 
```
go run server.go
```

## Launch database dgraph
```
docker run --rm -it -p 8080:8080 -p 9080:9080 -p 8000:8000 -v ~/dgraph:/dgraph dgraph/standalone:v20.07.0
```

### Dependencies
* go get -u github.com/go-chi/chi

# frontend sales informations

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).