# sbomdbs

## mongodb

### run

```bash
docker network create some-network
```

```bash
docker run \
  --rm \
  -d \
  --network some-network \
  --name some-mongo \
  -p 9999:27017 \
  -e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
  -e MONGO_INITDB_ROOT_PASSWORD=secret \
  mongo:7.0.7
```

### query

<https://www.mongodb.com/docs/manual/reference/sql-comparison/>

```bash
docker run \
  -it \
  --rm \
  --network some-network \
  mongo:7.0.7 \
    mongosh \
    --host some-mongo \
    -u mongoadmin \
    -p secret \
    --authenticationDatabase admin \
    some-db
```

use test
show collections
db.people.find()

| SQL                                      | MongoDB find()                           |
| ---------------------------------------- | ---------------------------------------- |
| SELECT \* FROM people WHERE status = "A" | db.people.find( { name: "alpine:3.15" }) |

find the docker images that contain the vulnerable zlib package

```bash
db.people.find( {"components.name": "zlib","components.version":"1.2.11-r3"},{"metadata.component.name":1,"metadata.component.purl":1,"metadata.component.bom-ref":1} )
```

show vulnerabilities:

```bash
trivy sbom data.json
```

db.people.find( { "components.type": "library" }, { "components.name": 1 ,"components.version": 1 } )

db.people.count()

## postgres
