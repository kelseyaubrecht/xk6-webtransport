![ ](https://github.com/kelseyaubrecht/xk6-webtransport/actions/workflows/test.yaml/badge.svg)

# xk6-webtransport

[k6](https://github.com/grafana/k6) extension to k6 extension to use the WebTransport protocol.
Implemented using the [xk6](https://github.com/grafana/xk6) system and [webtransport-go](https://github.com/quic-go/webtransport-go).

Supports:

- bidirectional streams
- datagrams

## Work in progress

This project is a work in progress. Feedback and contributions are welcome!

## Contents

- [Build](#build)
- [Supported features](#supported-features)
- [Usage example](#usage-example)
- [Metrics](#metrics)

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- Go
- Git

Then:

- Install `xk6`:

```shell
go install go.k6.io/xk6/cmd/xk6@latest
```

- Build the binary:

```shell
xk6 build --with github.com/kelseyaubrecht/xk6-webtransport@latest
```

## Supported features

- Bidirectional stream
  - Write
  - ReadAll
  - ReadFull
  - ReadAtLeast

## Usage example

An example of using the extension to create a bidirection stream.

```javascript
import wt from "k6/x/webtransport";

export default function () {
  const url = "https://localhost:443/webtransport";
  wt.connect(url);

  let data = [0, 1, 2, 3, 4];
  wt.write(data);

  const response = wt.readAll();
  // handle response

  wt.close();
}
```

## Metrics

| Metric                   | Type    | Description                        |
| ------------------------ | ------- | ---------------------------------- |
| webtransport_read_bytes  | counter | Total bytes read                   |
| webtransport_read_count  | counter | Total read count                   |
| webtransport_read_size   | trend   | Trends of read size per operation  |
| webtransport_write_bytes | counter | Total bytes written                |
| webtransport_write_count | counter | Total write count                  |
| webtransport_write_size  | trend   | Trends of write size per operation |
