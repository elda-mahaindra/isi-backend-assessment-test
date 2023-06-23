# ISI Backend Assessment Test

## Description

Consists only of one service: tabungan-api

## Getting Ready

### 1. Installing Docker

Visit [docker website](https://www.docker.com/products/docker-desktop/), download and install Docker Desktop.

Run the following command to list all running containers:

```bash
  docker ps
```

Run the following command to list all available docker images to run:

```bash
  docker images
```

### 2. Pull Postgres Images

Visit [Docker Hub](https://hub.docker.com/) and search for postgres image. Use the [official image](https://hub.docker.com/_/postgres).

Run the following command to pull the latest postgres image:

```bash
  docker pull postgres
```

Notes:

- This project is developed using postgres version **15**.

### 3. Installing 'golang-migrate'

Follow the instructions from the [CLI Documentation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) of golang-migrate and install it.

## Environment Variables

To run this project, the following environment variables needed to be added.

### 1. Main Services

Create `app.env` files inside each main service directory (`/tabungan-api`) and add the following environment variables:

```bash
  PORT=<PORT>

  DB_DRIVER=<DB_DRIVER>
  DB_SOURCE=<DATA_SOURCE_NAME>
```

Notes:

- `PORT`is the port number where the net will listen to.
- `DB_DRIVER` is the driver name to open the sql connection.
- `DATA_SOURCE_NAME` is structured as: `<DRIVER_NAME>://<USERNAME>:<PASSWORD>@<ADDRESS>:<PORT>/<DB_NAME>?sslmode=disable`.

The environment variables will be something like the below:

#### /tabungan-api

```bash
  PORT=3000

  DB_DRIVER=postgres
  DB_SOURCE=postgresql://root:secret@localhost:5432/isi_backend_assessment_test?sslmode=disable
```

## Start Servers

### 1. Docker Compose Up

Make sure the environment variables have been set up before trying to start the servers. To start the servers, run the following command from the root directory where the `docker-compose.yaml` file is placed.

For the plugin version (space compose):

```bash
  docker compose up
```

For the standalone version (dash compose):

```bash
  docker-compose up
```

### 2. Run The Migration

The following command structure is used to do the **up migration**:

```bash
  migrate -path <PATH> -database "<DRIVER_NAME>://<USERNAME>:<PASSWORD>@<ADDRESS>:<PORT>/<DB_NAME>?sslmode=disable" -verbose up
```

Run the `migrateup` command inside Makefile or run the following command:

```bash
  migrate -path db_migration -database "postgresql://root:secret@localhost:5433/isi_backend_assessment_test?sslmode=disable" -verbose up
```

**If** there is a need to do the **down migration**, the following command structure can be used:

```bash
  migrate -path <PATH> -database "<DRIVER_NAME>://<USERNAME>:<PASSWORD>@<ADDRESS>:<PORT>/<DB_NAME>?sslmode=disable" -verbose down
```

Run the `migratedown` command inside Makefile or run the following command:

```bash
  migrate -path db_migration -database "postgresql://root:secret@localhost:5433/isi_backend_assessment_test?sslmode=disable" -verbose down
```

Notes:

- Port 5433 is used because the Postgres container is mapped there (see `docker-compose.yaml` file).
- _Makefile_ can be found in the root of the project directory.
- `<PATH>`is the path to migration files.
- `sslmode=disable`is used because postgres container doesn’t enable SSL by default.
- Make sure the **dirty** state inside `schema_migration` table is `FALSE` before running any migration.

## Stop Servers

### 1. Docker Compose Down

To **remove all** existing containers and networks related to this project, run the following command from the root directory where the `docker-compose.yaml` file is placed.

For the plugin version (space compose):

```bash
  docker compose down
```

For the standalone version (dash compose):

```bash
  docker-compose down
```

### 2. Remove Docker Images

The following command structure is used to remove a docker image:

```bash
  docker rmi <IMAGE_NAME>
```

The commands to remove all images related to this project will be something like the below:

```bash
  docker rmi tabungan-api
```
