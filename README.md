# import

Imports leak files into SQLite

## Hooks

This repository is configured with client-side Git hooks which you need to install by running the following command:

```bash
./hooks/INSTALL
```

## Docker

To run the service with Docker, you will first need to setup the `.git-local-credentials` file. This credentials file shall contain the git credentials config to access `damn` and `palavrapasse` private modules.

To build the service image:

```bash
docker_tag=import:latest
docker build \
    -f ./deployments/Dockerfile \
    --secret id=git-credentials,src=.local-git-credentials \
    . -t $docker_tag
```

To run the service container, for instance:

```bash
docker run import --help
```