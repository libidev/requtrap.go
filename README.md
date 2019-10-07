![REQUTRAP](./assets/images/banner.png)

## Configuration
You can configure RequTrap by creating a YAML file which 
contain configuration schema. It will look like example below.

```yml
gateway:
  name: book-store
  host: 127.0.0.1
  port: 8080
  services:
    - path: /books
      upstream: http://127.0.0.1:8001
    - path: /authors
      upstream: http://127.0.0.1:8002
```

## CLI Commands
### Starting API Gateway
`requtrap start [path/to/configuration.yml]`

### Stoping API Gateway
`requtrap stop [gateway-name]`

### More
`requtrap help [command]`
