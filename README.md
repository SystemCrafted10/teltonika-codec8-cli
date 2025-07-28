# teltonika-codec8-cli

**CLI tool to parse Teltonika Codec 8 and Codec 8 Extended packets**  
[GitHub: chrisfarrugia.dev/teltonika-codec8-cli](https://chrisfarrugia.dev/teltonika-codec8-cli)

---

## Overview

`teltonika-codec8-cli` is a lightweight Go CLI tool for parsing Teltonika device packets using Codec 8 and Codec 8 Extended (8EX).

- Use it directly from your terminal/bash
    
- Or integrate it with your backend stack (Node.js, Python, PHP, etc.)
    

### What does it do?

- Accepts a single hex string (representing a raw Teltonika Codec 8/8EX packet) as input
    
- Outputs a parsed JSON representation to stdout, ready for further processing
    

---

## Quick Start

### 1. **Convert Device Data to Hex String**

Your Teltonika device sends data as a byte array/slice.  
**You must convert this byte array to a hex string** before passing to this tool.

### 2. **Parse from CLI**

#### In development:

```bash
go run ./cmd/teltonika-codec8/. 00000056...
```

#### After building (recommended):

```bash
go build -o teltonika-codec8-cli ./cmd/teltonika-codec8/main.go
./teltonika-codec8-cli 00000056...
```

Or, if installed globally:

```bash
teltonika-codec8-cli 00000056...
```

#### Example Output


``` json
{
  "Packet": "00000056...",
  "Preamble": 0,
  "Data_Length": 88,
  "CodecID": 8,
  "CodecType": "data sending",
  "Quantity1": 1,
  "Content": {
    "AVL_Datas": [
      {
        "Timestamp": "2025-06-03T08:44:33Z",
        "Priority": 0,
        "GPSelement": {
          "Longitude": 14.5250283,
          "Latitude": 35.8488933,
          "Altitude": 49,
          "Angle": 92,
          "Satellites": 16,
          "Speed": 0
        },
        "IOelement": {
          "EventID": 0,
          "ElementCount": 17,
          "Elements": {
            "1": 1,
            "12": 31906,
            "...": "..."
          }
        }
      }
    ]
  },
  "Quantity2": 1,
  "CRC": 31438
}

```

---
### Protocol Note: IMEI Handshake

> **Important:**  
> This tool is designed to parse Codec 8/8EX data packets **after** the IMEI handshake is completed.
> 
> - The Teltonika protocol starts with an IMEI handshake when a device connects.
>     
> - This CLI tool **expects a hex string representing the data packet only** (not the IMEI handshake packet).
>     
> - Make sure your application handles the IMEI handshake separately before passing subsequent packets to this parser.
>

---
## Integration in Other Languages

You can call the CLI from any language that can launch shell commands:

**Node.js**

```js
const { execFile } = require("child_process"); execFile("teltonika-codec8-cli", [hexString], (error, stdout) => {   // parse stdout as JSON });
```

**PHP**

```php
$json = shell_exec("teltonika-codec8-cli $hexstring");
```

**Python**

```python
import subprocess result = subprocess.run(['teltonika-codec8-cli', hexstring], capture_output=True) print(result.stdout)
```

---

## TypeScript Usage Example

You can use this tool as a child process from Node.js/TypeScript for tight integration:

```ts
import { execFile } from "node:child_process";
export function teltonikaCodec8Parse(hexString: string): Promise<TeltonikaAvlPacket> {
    return new Promise((resolve, reject) => {
        execFile("teltonika-codec8-cli", [hexString], (error, stdout, stderr) => {
            if (error) return reject(error);
            try {
                resolve(JSON.parse(stdout));
            } catch (err) {
                reject(new Error("Failed to parse CLI output as JSON: " + err));
            }
        });
    });
}

```

```ts

export type TeltonikaAvlPacket = {
    Packet: string;
    Preamble: number;
    Data_Length: number;
    CodecID: number;
    Quantity1: number;
    CRC: number;
    Quantity2: number;
    CodecType: string;
    Content: {  AVL_Datas: AvlData[] };
};
 

export type AvlData = {
    Timestamp: string | number | Date;
    Priority: number;
    GPSelement: {
        Longitude: number;
        Latitude: number;
        Altitude: number;
        Angle: number;
        Satellites: number;
        Speed: number;
    };
    IOelement: {
        EventID: number;
        ElementCount: string | number;
        Elements: Record<string, string | number >;
    };
};
```
---

## File Structure

```go
.
├── cmd/
│   └── teltonika-codec8/
│       └── main.go
├── parser/
│   ├── codec8.go
│   ├── codec8Extended.go
│   ├── helpers.go
│   └── types.go
├── go.mod
├── teltonika-codec8-cli (binary, after build)
└── README.md

```


## Building & Installation

```bash
go build -o teltonika-codec8-cli ./cmd/teltonika-codec8/main.go
```

**Optional:**  

Move the binary to a system path (`/usr/local/bin` is recommended):

```bash
sudo mv teltonika-codec8-cli /usr/local/bin/
sudo chmod 755 /usr/local/bin/teltonika-codec8-cli
sudo chown root:root /usr/local/bin/teltonika-codec8-cli
```


> ⚠️ **Note:** For production, avoid `chmod 777`. Use `chmod 755` to give execute permissions, and restrict ownership as above.

---

## License

MIT License

---

## Author

Chris Farrugia — [chrisfarrugia.dev](https://chrisfarrugia.dev)