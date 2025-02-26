FROM mcr.microsoft.com/dotnet/sdk:9.0 AS wasm-build
WORKDIR /app

COPY web/BlazorApp.sln ./
COPY web/BlazorApp/BlazorApp.csproj ./BlazorApp/
RUN mkdir -p BlazorApp/proto
COPY ../../protobuf-gen/*.proto ./BlazorApp/proto/

RUN dotnet restore
COPY ./web .
RUN dotnet publish -c Release -o out

FROM nginx:1.27-alpine
WORKDIR /app
EXPOSE 8080
COPY web/nginx.conf /etc/nginx/nginx.conf
COPY --from=wasm-build /app/out/wwwroot /usr/share/nginx/html
#keep the container running
# CMD ["tail", "-f", "/dev/null"]