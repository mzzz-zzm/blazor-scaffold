# blazor-scaffold

## how to run the app:

## prerequisites
#compile protobuffer for server/client
```sh
docker-compose build --no-cache protobuf-gen-svr
docker-compose up protobuf-gen-svr
```


```sh
#in development mode
pwd
/workspaces/blazor-scaffold
docker-compose build --no-cache
docker-compose up #this will generate grpc stubs and spin up web
# then access http://localhost:5001 from a browser
```

```sh
#shutdonw
pwd
/workspaces/blazor-scaffold
ctrl-c to stop the docker
docker-compose down #clean up
```

## TODO
client gRPC is done (razor page); now impl service on server-side
check udemy as a reference

add .gitignore for set web/BlazorApp/obj/* and protobuf-gen/generated/* to be ignored from commit
