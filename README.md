# blazor-scaffold

## how to run the app:
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

