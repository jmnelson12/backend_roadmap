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

## What is HTTP(1)?

HyperText Transfer Protocol is an application-layer protocol for transmitting hypermedia documents.

- follows a client-server model, where client opens a request, then waits until it receives a response
- stateless protocol - meaning the server doesn't keep any data between the two requests

### Things done to overcome latency pains

- **Spriting** - combining many small images into one big image
- **Inlining** - base64 in the css
- **Concatenation** - combining js files together to send one file instead of many
- **Sharding** - serving aspects of your service on as many different hosts as possible

## What is HTTP(2)?

- binary protocol
- sends binary frames
  - different frame types can be sent, but the all have the same setup
    - Length, Type, Flags, Stream Identifier, frame payload
- each frame is sent with a stream
  - stream is an independent bi-directional sequence of frames exchanged between the client and server
  - has a priority (or weight) which says which stream to consider most important
- header compression (`HPACK`)
- reset - change your mind with (`RST_STREAM`)
- server push (or cache push)
  - client asks for `x`, server knows that client will likely want `z` as well, and it sends to the client without being asked

## Browsers and how they work?

- Once server supplies the resources (HTML, CSS, JS, images, etc) the browser

  1. parses the HTML, CSS, and JS
  2. Renders - Constructs the DOM Tree --> Render Tree --> Layout of Render Tree --> Paints the render tree

**Browser High Level Structure**

copied from https://github.com/alex/what-happens-when/blob/master/README.rst#behind-the-scenes-of-the-browser

The components of the browsers are:

* **User interface:** The user interface includes the address bar,
  back/forward button, bookmarking menu, etc. Every part of the browser
  display except the window where you see the requested page.
* **Browser engine:** The browser engine marshals actions between the UI
  and the rendering engine.
* **Rendering engine:** The rendering engine is responsible for displaying
  requested content. For example if the requested content is HTML, the
  rendering engine parses HTML and CSS, and displays the parsed content on
  the screen.
* **Networking:** The networking handles network calls such as HTTP requests,
  using different implementations for different platforms behind a
  platform-independent interface.
* **UI backend:** The UI backend is used for drawing basic widgets like combo
  boxes and windows. This backend exposes a generic interface that is not
  platform specific.
  Underneath it uses operating system user interface methods.
* **JavaScript engine:** The JavaScript engine is used to parse and
  execute JavaScript code.
* **Data storage:** The data storage is a persistence layer. The browser may
  need to save all sorts of data locally, such as cookies. Browsers also
  support storage mechanisms such as localStorage, IndexedDB, WebSQL and
  FileSystem.


## DNS and how it works?

## What is Domain Name?

## What is hosting?
