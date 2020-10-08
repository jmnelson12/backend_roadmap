# Internet

## How does the internet work?

### 1. parse url

### 2. is it a url or search term

- url has a `protocol` (**_http_**)
- if not url, feed the text to the browsers default web search engine

### 3. check `HSTS` list

- HTTP Strict Transport Security list
- List of websites that have request to be contact via `HTTPS` only
- If website not in list, the initial request is sent via `HTTP`
  - potentially vulnerable to a [downgrade attack](http://en.wikipedia.org/wiki/SSL_stripping)

### 4. `DNS` lookup

### 5. `ARP` process

- Address Resolution Protocol

  - Needs the target IP, MAC address of the interface it will use to send the broadcast

    ```
     // ARP Request
     Sender MAC: interface:mac:address:here
     Sender IP: interface.ip.goes.here
     Target MAC: FF:FF:FF:FF:FF:FF (Broadcast)
     Target IP: target.ip.goes.here

     // ARP Reply
     Sender MAC: target:mac:address:here
     Sender IP: target.ip.goes.here
     Target MAC: interface:mac:address:here
     Target IP: interface.ip.goes.here
    ```

### 6. Opening of a socket

- `TCP` --> `Network Layer` --> `Link Layer`
  - `TCP`
    - destination port added to header
    - source port is chosen
  - `Network Layer`
    - wraps an additional IP header
  - `Link Layer`
    - frame header is added. Includes: MAC address of machine and gateway (local router)

At this point the packet is ready to be transmitted through either:

- Ethernet
- WiFi
- Cellular data network

### 7. TLS handshake

- client computer send `ClientHello` message to the server
  - contains: **_version_**, **_list of cipher algos_**, and **_available compression methods_**
- server replies with `ServerHello` message
  - contains **_version_**, **_selected cipher_**, **_selected compression methods_**, and the servers **_public certificate_** (signed by a CA - **_Certificate Authority_**)
- client verifies servers certificate against a list of its trusted CAs
  - if trusted, client generates string of data and encrypts it with the **server's** public key
- server decrypts the data and does stuff
- client sends a `Finished` message
- server sends a `Finished` message

### 8. HTTP protocol

- covered below at [What is HTTP](#what-is-http?)

### 9. Behind the scenes of the browser

- covered below at [Browsers and how they work](#browsers-and-how-they-work?)

## What is HTTP?

- https://github.com/alex/what-happens-when#http-protocol
- https://daniel.haxx.se/http2/

## Browsers and how they work?

- https://github.com/alex/what-happens-when#behind-the-scenes-of-the-browser

## DNS and how it works?

## What is Domain Name?

## What is hosting?
