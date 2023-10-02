# 🚀 Caddy Bandwidth Limiter Plugin

Rev up your [Caddy v2](https://caddyserver.com/v2) server with the ability to finely control the bandwidth of your HTTP responses. Perfectly designed for CDN use-cases, this is not just another plugin — it's a **key component** of our soon-to-launch _Media Edge_ software, meticulously crafted from scratch! 🛠

🔗 Dive into Caddy's magic on their [GitHub](https://github.com/caddyserver/caddy).

## 📖 Table of Contents

- [📦 Installation](#installation)
- [🖋 Usage](#usage)
  - [💡 Real-World CDN Example](#real-world-cdn-example)
- [🛠 Development](#development)
- [📜 License](#license)
- [📢 Join Our Community](#join-our-community)

## 📦 Installation

Plug into the full power of Caddy by integrating our plugin. Let's get started:

1. Grab [xcaddy](https://github.com/caddyserver/xcaddy):

```bash
go get -u github.com/caddyserver/xcaddy/cmd/xcaddy
```

2. Build Caddy with our bandwidth plugin:

```bash
xcaddy build --with github.com/mediafoundation/bandwidth
```

🎉 Voilà! You've got a `caddy` binary, now supercharged with our plugin.

## 🖋 Usage

Eager to throttle bandwidth? Use the `bandwidth` directive in your Caddyfile to set the max bytes-per-second:

```caddy
localhost
route /myroute {
    bandwidth {
        limit 100000
    }
}
```

### 💡 Real-World CDN Example

Designed with CDN use-cases in mind, you can add bandwidth limits dynamically based on headers or other conditions:

```caddy
header Server "MediaEdge vX.Y.Z"
reverse_proxy http://localhost:8080 {
    @hasBandwidthLimit header X-Bandwidth-Limit Yes
    handle_response @hasBandwidthLimit {
        bandwidth {
            limit 50000
        }
    }
}
```

- `order bandwidth before header`: Place the bandwidth module before the header module in the processing order.
- `auto_https disable_redirects`: Disables automatic HTTPS redirects.
- `on_demand_tls`: Enables on-demand TLS certificate provisioning.
- `log`: Customizes logging, for instance, using the Elastic format and specifying a file output path.
  
This allows you to have fine-grained control over bandwidth limits on a per-request basis!

## 🛠 Development

Our plugin adheres to standard Go conventions, featuring a `Middleware` struct that uses the `caddyhttp.MiddlewareHandler` interface. The `limitedResponseWriter` is meticulously designed to limit bandwidth.

💡 Ideas? Contributions are welcome! Feel free to submit issues and pull requests.

## 📜 License

Under the MIT License. Use responsibly.

## 📢 Join Our Community

- 🎮 [Discord](https://discord.gg/nyCS7ePWzf)
- 📫 [Telegram](https://t.me/Media_FDN)
