using Microsoft.AspNetCore.Components.Web;
using Microsoft.AspNetCore.Components.WebAssembly.Hosting;
using Grpc.Net.Client;
using Microsoft.AspNetCore.Cors.Infrastructure;
using Grpc.Net.Client.Web;
using BlazorApp;
using BlazorGrpcWebApp.Shared;
using System.Net.Http;

var builder = WebAssemblyHostBuilder.CreateDefault(args);
builder.RootComponents.Add<App>("#app");
builder.RootComponents.Add<HeadOutlet>("head::after");

// Set the gRPC server URL - make sure this is the URL as accessed from the browser
// NOTE: since the url is called from a browser, which is outside the container, it won't resolve docker-container's service name
// unless the browser is backed by a reverse proxy like Nginx running on a container
var grpcServerUrl = "http://localhost:8080";

builder.Services.AddGrpcClient<Greeter.GreeterClient>(options =>
{
    options.Address = new Uri(grpcServerUrl);
}).ConfigurePrimaryHttpMessageHandler(() =>
    new GrpcWebHandler(GrpcWebMode.GrpcWeb, new HttpClientHandler()));

Console.WriteLine($"Connecting to gRPC server at: {grpcServerUrl}");

await builder.Build().RunAsync();