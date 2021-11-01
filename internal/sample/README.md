# Sample

Sample app / service

## Architecture

This project follows SOLID principle and Clean Architecture

### Getting Started

Below are the steps if you want to run locally without docker

1. Set required environment variable

   ENV=local

2. Set configuration

   Change config/local.yaml configuration value properly
   and make sure can connect to blank MySQL database properly

3. Run migration

   > make migrate

4. Add seed data

   > make seed

5. Run app

   > make run

6. Open the browser

   Visit the url

   > http://localhost:8080/v1/users

### Run Test

go test -v ./internal/user/...

## Get Started Using Docker Compose

Below are the steps if you want to run locally with docker & docker compose

1. Build docker image

   Generate the docker imgae by this command:

   > docker build --rm -t starterkitapi -f dockerfile.api .

2. Run with docker compose

   Copy the sample.env to .env and adjust it then run the docker compose with this command:

   > docker-compose up

3. Open the browser

   Visit the url

   > http://localhost:8080/v1/users

4. Quit & Cleanup

   Click Ctrl+C to quit in console then run these commands below to clean up all things:

   > docker-compose rm -v

### Project Structure

The starter kit uses the following project layout:

```
.
├── api					api documentation i.e. protobuf or openapi
│   ├── openapi			openapi doc files in case want contract first for develop
│   └── protobuf		protobuf contract files
├── cmd					docker and main applications of the project
│   ├── admin			admin tools application
│   └── server			the API server application
├── config				configuration files for different environments
├── internal			private application and library code (implement `Screaming Architecture`)
│   ├── healthcheck		healthcheck module (just handler)
│   ├── <firstmodule>	your modules project, i.e. user
│   ├── ...				your modules project
│   ├── <nmodule>		your modules project, i.e. user
│   └── test			helpers for testing purpose
├── migrations			database migrations i.e. have sub dir of its impl db
│   ├── mysql			mysql database migrations files
│   └── mongodb			mongodb database migrations files
├── pkg					public library code
│   └── protobuf		shared project's protobuf files
└── testdata			test data scripts
```

### TODO

- [x] Configuration
- [x] Viper for better handling env and config
- [x] Logging
- [x] Health Check
- [x] Data Seed
- [x] DB Migration
- [x] Makefile (Optional)
- [x] Docker & Docker Compose
- [x] Hot Reload for development with docker
- [ ] Use google wire for Dependency Injection
- [ ] Error Handling
- [ ] Validation
- [ ] Versioning
- [ ] Pagination
- [ ] Run & Manage via CLI Command
- [ ] Use Cobra for CLI
- [ ] Linter
- [ ] Unit Test
- [ ] Integration Test sample
- [ ] Add more examples for service, repository and test
- [ ] Graceful Shutdown with timeout for active connection
- [ ] JWT base Authentication
- [ ] Observability/Metrics
- [ ] Kubernetes deployment

### References
- https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html