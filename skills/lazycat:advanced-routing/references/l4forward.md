# TCP/UDP Layer 4 Forwarding {#tcp-udp-ingress}

::: warning For standard HTTP traffic, please use the `application.routes` feature.
Ingress TCP/UDP forwarding is designed for use outside of microservice clients, such as for CLI tools or third-party applications. If you only need to forward a specific HTTP port of a container, use the [lzcapp HTTP routing feature](./advanced-route.md).
:::

To provide TCP/UDP services, add an `ingress` sub-field under the `application` field in the `lzc-manifest.yml` file:

```yml
application:
  ingress:
    - protocol: tcp
      port: 8080
    - protocol: tcp
      description: Database Service
      port: 3306
      service: mysql
    - protocol: tcp
      description: Forward ports 20,000-30,000 to the corresponding ports
      service: app
      publish_port: 20000-30000
    - protocol: tcp
      description: Forward all ports in 16,000-18,000 range to port 6666
      service: app
      port: 6666
      publish_port: 16000-18000
```

- `protocol`: The protocol for external service (`tcp` or `udp`).
- `description`: A description of the service to help administrators understand its purpose.
- `port`: The target service port. If omitted, it defaults to the actual inbound port. (Note: Versions prior to v1.3.8 do not support `port` 80 or 443).
- `service`: The service name used to locate the specific `service container`. Defaults to `app`.
- `publish_port`: The inbound port. Defaults to the value of `port`. Supports specific ports (e.g., `3306`) or ranges (e.g., `1000-50000`).

Once configured, you can access the service via a browser or client. For example, if your application subdomain is `app-subdomain` and your device name is `devicename`, you can access the TCP service at `app-subdomain.devicename.heiyu.space:3306`.

::: warning Security Notice
When using TCP/UDP features, the system only provides underlying virtual network protection and cannot provide an authentication workflow. Other processes on the microservice client can access these TCP/UDP ports without restriction. If users use port forwarding tools, security may be further compromised. Developers must handle authentication logic within their applications when providing TCP/UDP functionality.
:::

::: warning Ports 80/443
When your application directly handles port 443 (supported in v1.3.8+), traffic reaches your container directly. Consequently, the system cannot perform preprocessing, including but not limited to:
- User authentication
- Automatic application wakeup
- HTTPS certificate configuration
- `application.routes`, `application.upstreams`, and other configurations

In almost all cases, you should not configure port 443. The only reasonable scenario is using a microservice-assigned EIP to forward all traffic to another host/NAS using a non-microservice domain.

If you are certain you want to handle 80/443 traffic yourself, you must explicitly declare `yes_i_want_80_443: true` in the corresponding ingress entry.
:::
