using System.Text.Json.Serialization;

var builder = WebApplication.CreateSlimBuilder(args);

builder.Services.ConfigureHttpJsonOptions(options =>
{
    options.SerializerOptions.TypeInfoResolverChain.Insert(0, AppJsonSerializerContext.Default);
});

var app = builder.Build();

var sensors = app.MapGroup("/sensors");
sensors.MapPut("ingest", (SensorRecord request) =>
{
    Console.WriteLine($"Received {request}");
    return Results.Ok("thx");
});

app.Run("http://192.168.100.7:8080");

public record SensorRecord(string Session, DateTimeOffset SensorTime, decimal Temperature, decimal Pressure);

[JsonSerializable(typeof(SensorRecord))]
internal partial class AppJsonSerializerContext : JsonSerializerContext;