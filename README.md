#### This is no longer supported.

# Dynamornr
Dynamornr is a CLI-based task runner for DynamoDB (written in Go).

It leverages the Golang aws-sdk-go library in order to execute it's tasks and was born out of my frustration with interfacing with DynamoDB directly.

## Commands:
* `dynamornr tables:list` List the current tables in the running DynamoDB database.
* `dynamornr tables:create` Create tables (according to a `schema.yml` file).
* `dynamornr items:list [TABLENAME]` List items belonging in the `TABLENAME` table.
* `dynamornr items:seed` Seeds items to tables according to the [config/seeds.json file](https://github.com/nycdavid/dynamornr/blob/master/test/config/seeds.json)

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
* For supported features in DynamoDB, take a look at our [example schema.yml file](https://github.com/nycdavid/dynamornr/blob/master/test/config/schema.yml)
