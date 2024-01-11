#See https://aka.ms/customizecontainer to learn how to customize your debug container and how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
WORKDIR /app
COPY Transcription/names_table.jsonl /app
COPY Transcription/surnames_table.jsonl /app
COPY Transcription/midnames_table.jsonl /app
EXPOSE 8002

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY ["Transcription/Transcription.csproj", "Transcription/"]
RUN dotnet restore "Transcription/Transcription.csproj"
COPY . .
WORKDIR "/src/Transcription"
RUN dotnet build "Transcription.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "Transcription.csproj" -c Release -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "Transcription.dll"]