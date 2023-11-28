# Traefik Block Paths

A [Traefik](https://github.com/containous/traefik) plugin to block access to certain paths using a list of regex values and return a defined status code.

## Configuration

Sample configuration in Traefik.

### Configuration as local plugin

traefik.yml

```yaml
log:
  level: INFO
experimental:
  localPlugins:
    block-paths:
      moduleName: github.com/JonasSchubert/traefik-block-paths
```

dynamic-configuration.yml

```yaml
http:
  middlewares:
    block-wp:
      plugin:
        block-paths:
          regex:
            - "^/wp(.*)"
          statusCode: 404
```

docker-compose.yml

```yaml
version: "3.7"
services:
  traefik:
    image: traefik
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - /docker/config/traefik/traefik.yml:/etc/traefik/traefik.yml
      - /docker/config/traefik/dynamic-configuration.yml:/etc/traefik/dynamic-configuration.yml
      - /plugin/block-paths:/plugins-local/src/github.com/JonasSchubert/traefik-block-paths/
    ports:
      - "80:80"
  hello:
    image: containous/whoami
    labels:
      - traefik.enable=true
      - traefik.http.routers.hello.entrypoints=http
      - traefik.http.routers.hello.rule=Host(`hello.localhost`)
      - traefik.http.services.hello.loadbalancer.server.port=80
      - traefik.http.routers.hello.middlewares=my-plugin@file
```

## Sample configuration

- `regex`:  List of regex values to use for path blocking.
- `statusCode`: Return value of the status code.

```yaml
my-block-paths:
  plugin:
    block-paths:
      regex:
        - "^/wp(.*)"
      statusCode: 404
```

## Contributors

| [<img alt="JonasSchubert" src="https://secure.gravatar.com/avatar/835215bfb654d58acb595c64f107d052?s=180&d=identicon" width="117"/>](https://code.schubert.zone/jonas-schubert) |
| :---------------------------------------------------------------------------------------------------------------------------------------: |
| [Jonas Schubert](https://code.schubert.zone/jonas-schubert) |

## License

traefik-block-paths is distributed under the MIT license. [See LICENSE](LICENSE) for details.

```
MIT License

Copyright (c) 2023 Jonas Schubert

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
