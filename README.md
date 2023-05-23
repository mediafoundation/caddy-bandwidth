# Caddy Bandwidth Limiter Plugin

This is a plugin for [Caddy v2](https://caddyserver.com/v2) that provides the ability to limit the bandwidth of HTTP responses.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)
- [License](#license)

## Installation

This plugin can be installed as a part of Caddy. Follow the below steps:

1. Install [xcaddy](https://github.com/caddyserver/xcaddy), a tool for building Caddy with plugins:

    ```bash
    go get -u github.com/caddyserver/xcaddy/cmd/xcaddy
    ```

2. Build Caddy with this plugin:

    ```bash
    xcaddy build --with github.com/mediafoundation/bandwidth
    ```

This will produce a `caddy` binary in your current directory which includes the plugin.

## Usage

This plugin provides the `bandwidth` directive for use in your Caddyfile. The `bandwidth` directive accepts one argument, `limit`, which is the maximum number of bytes per second that the server will send.

Here's a basic example of how to use it in a Caddyfile:

```caddy
localhost

route /myroute {
    bandwidth {
        limit 100000
    }
}
```

In this example, the bandwidth for /myroute is limited to 100,000 bytes per second.

## Development
This plugin uses standard Go conventions for its development. It consists of a Middleware struct which implements the caddyhttp.MiddlewareHandler interface, and a limitedResponseWriter which is used to limit the bandwidth of HTTP responses.

If you want to contribute to the development of this plugin, please feel free to submit issues and/or pull requests.

## License
This plugin is available under the MIT License.
