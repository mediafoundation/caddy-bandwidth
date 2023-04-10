Caddy Bandwidth Throttling Module

This repository contains a bandwidth throttling module for Caddy v2, developed by Media Foundation. This module allows you to limit the bandwidth per subdomain in a CDN environment, providing granular control over resource consumption.

## Features

- Bandwidth throttling per subdomain
- Seamless integration with Caddy v2
- Simple configuration using Caddyfile

## Installation

To install and use the bandwidth throttling module, follow these steps:

1. Clone this repository:

``git clone https://github.com/mediafoundation/caddy-bandwidth.git``

2. Build Caddy with the custom module:

``xcaddy build --with github.com/mediafoundation/caddy-bandwidth@latest``


3. Replace your existing Caddy binary with the newly built binary:

``sudo mv caddy /usr/bin/caddy``

## Usage

Update your Caddyfile with the `bandwidth` directive and specify the bandwidth limit for each subdomain:

```Caddyfile
client1.example.com {
 handle {
     bandwidth {
         limit 100000
     }
     reverse_proxy your_upstream_server
 }
}

client2.example.com {
 handle {
     bandwidth {
         limit 200000
     }
     reverse_proxy your_upstream_server
 }
}
```

In this example, client1.example.com has a bandwidth limit of 100 KB/s and client2.example.com has a bandwidth limit of 200 KB/s.

## Support
If you encounter any issues or require assistance, please open an issue on the GitHub repository or contact Media Foundation directly.

## Credits
This module is developed and maintained by Media Foundation. Special thanks to the Caddy community for their support and guidance.
