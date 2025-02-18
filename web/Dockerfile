FROM mcr.microsoft.com/dotnet/sdk:9.0-alpine AS build
WORKDIR /app

COPY BlazorApp.sln ./
COPY BlazorApp/BlazorApp.csproj ./BlazorApp/

RUN dotnet restore
COPY . .
RUN dotnet publish -c Release -o out

FROM nginx:1.27-alpine
WORKDIR /app
EXPOSE 8080
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=build /app/out/wwwroot /usr/share/nginx/html
