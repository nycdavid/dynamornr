# Dynamornr
Dynamornr is a CLI-based task runner for DynamoDB (written in Go).

It leverages the Golang aws-sdk-go library in order to execute it's tasks and was born out of my frustration with interfacing with DynamoDB directly.

Currently it's able to:
* List the current tables in the running DynamoDB database. `dynamornr tables:list`
* Create tables (according to a `schema.yml` file). `dynamornr tables:create`

## Configuration Files
There are 2 main configuration Yaml files that Dynamornr looks for:
* `config/database.yml`: It looks in this file according to the `ENV` variable provided and fetches both the port and URL from it.
  * An example file looks like this:
    ```yaml
      test:
        port: 3000
        host: http://some.fake.host
    ```
* `config/schema.yml`: This file outlines the tables that should be created, their key schema/attribute definitions and their provisioned throughput.
  * An example file looks like this:
    ```yaml
      tables:
        - name: "users"
          attribute_definitions:
            Id: "N"
          key_schema:
            Id: "HASH"
          provisioned_throughput:
            read_capacity_units: 10
            write_capacity_units: 5
    ```
