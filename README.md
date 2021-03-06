# Majoo Service

## Getting started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

1. Install Docker: https://docs.docker.com/install/
2. Install Docker-Compose: https://docs.docker.com/compose/install/
3. Install make (in case you're not unix user): https://stackoverflow.com/a/32127632

### Run Project

On the root directory, run this command:

    1. Generate virtual DB Server
        docker-compose up -d
    2. Installing dependencies
        go mod vendor
    3. Run DB Migration
        make migration-init
        make migration-up
    4. Run this project
        make run

### Testing API

1. Install Postman: https://learning.postman.com/docs/getting-started/installation-and-updates/
2. Import [this](https://github.com/afandi-syaikhu/majoo/blob/45c67dfb25fe6c335d29108efcd428bcd746932e/Majoo.postman_collection.json) Postman collection