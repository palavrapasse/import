# import

Imports leak files into SQLite

## Hooks

This repository is configured with client-side Git hooks which you need to install by running the following command:

```bash
./hooks/INSTALL
```

## Docker

To build the tool image:

```bash
docker_tag=import:latest
docker build \
    -f ./deployments/Dockerfile \
    . -t $docker_tag
```

To run the tool container:

```bash
docker run import --help
```

---

To build the **import web api** service image:

```bash
docker_tag=import-web-api:latest
docker build \
    -f ./deployments/import-web-api.dockerfile \
    third_party/import-web-api -t $docker_tag
```

To run the service container:

```bash
export $(grep -v '^#' third_party/import-web-api/.env | xargs)

docker run \
    -p $server_port:$server_port \
    -v "type=bind,src=$leaksdb_fp,dst=$leaksdb_fp" \
    --env-file third_party/import-web-api/.env \
    -t $docker_tag
```

---

To build the **import web** tool image:

```bash
docker_tag=import-web:latest
docker build \
    -f ./deployments/import-web.dockerfile \
    third_party/import-web -t $docker_tag
```

To run the service container:

```bash
docker run -p 3000:3000 -t $docker_tag -
```