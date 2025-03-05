# Network Latency and Throughput Measurement

## Repository Structure
```
networking-assignment1/
│-- react-app/                # React frontend for visualization
│-- milestone1/           # First milestone (Go implementation)
│-- milestone2/           # Second milestone (Go implementation)
│-- milestone3/           # Third milestone (Go implementation)
│-- milestone4/           # Fourth milestone (Go implementation)
|-- helper/
│-- README.md             # Project documentation
```

## Project Overview
This project measures the **latency** and **throughput** of TCP and UDP protocols across multiple machine pairs and networks. The results are visualized on a web page using graphs.
<details>
<summary>Project Descripition </summary>

Measure the latency and throughput of the following TCP and/or UDP-based protocols (as noted below) across at least three pairs of machines using at least two different networks. For example, two CS servers (like rho and pi), or a CS server to a laptop, wired or wireless, or off-campus. Create a web page with graphs summarizing your results. Use appropriate measurement sample sizes and readily interpretable units in graphs.

All messages must use a simple encryption scheme. One suggestion is to use an XOR encoding of 64-bit (8-byte, Java "long") values, based on a known shared initial key, updated using a custom RNG on each step, and then validated by the receiver. Here's a simple RNG update function: xorshift, requiring a non-zero initial key.
```java
long xorShift(long r) {
    r ^= r << 13;
    r ^= r >>> 7;
    r ^= r << 17;
    return r;
}
```
1. Measure round-trip latency (RTTs) and how it varies with message size in TCP, by sending and receiving (echoing and validating) messages of size 8, 64, 256, and 512 bytes.

2. Measure throughput (bits per second) and how it varies with message size in TCP, by sending 1MByte of data (with an 8-byte acknowledgment in the reverse direction) using different numbers of messages: 1024 1024-byte messages, vs 2048 512-byte messages, vs 4096 x 256-byte messages. Use known message contents (for example, number sequences) so they can be validated.

3. The same as (1), except using UDP.

4. The same as (2), using UDP.

For timing, use System.nanoTime() (or the closest equivalent if using other languages). Read through the Java networking tutorial. Also see SimpleService.java and EchoClient.java for some stripped-down examples of using server and client sockets. When using non-CS machines and networks, minimize unnecessary traffic while developing your programs. Beware of firewalls.
</details>

## Objectives
- Measure **round-trip latency (RTT)** for TCP and UDP at various message sizes.
- Measure **throughput (bits per second)** for TCP and UDP at different message sizes.
- Implement **encryption** using XOR encoding with a custom RNG.
- Collect data across **at least three pairs of machines** using two different networks.
- Visualize results in a **React-based graphical format**.

## Encryption Scheme
Messages use XOR encoding of **64-bit values (`uint64` in Go)**, initialized with a shared key and updated using a custom random number generator (RNG):

```go
func xorShift(r uint64) uint64 {
    r ^= r << 13
    r ^= r >> 7
    r ^= r << 17
    return r
}
```

## Measurement Details
### 1. TCP RTT Measurement
- Measure **RTT** for message sizes: **8, 64, 256, and 512 bytes**.
- Messages are echoed and validated.
- 8-byte acknowledgment is sent in the reverse direction.

### 2. TCP Throughput Measurement
- Send **1 MB of data** using different message sizes:
  - **1024 messages of 1024 bytes**
  - **2048 messages of 512 bytes**
  - **4096 messages of 256 bytes**
- 8-byte acknowledgment is sent in the reverse direction.

### 3. UDP RTT Measurement
- Same as TCP latency measurement but using UDP.

### 4. UDP Throughput Measurement
- Same as TCP throughput measurement but using UDP.

## Tools & Libraries
- **Go (Golang)** (for networking and timing with `time.Now().UnixNano()`)
- **Go net package** (`net.Listen`, `net.Dial`, `net.UDPConn`)
- **React.js** (for visualization using libraries like Chart.js)

## Instructions
### 1. Set Up the Environment
- Ensure Go is installed (`go version`). 
  - Version - 1.22.5
  - Only when you want to compile go files
```bash
cd ./react-app
npm install
```

### 2. Running client server programs
- For each milstone folder, put the clientSocket and serverSocket on respective server/machines.
#### ServerSocket
- Run the execuatable file depending on your machine (amd64 or arm64)
- If arm64,
```bash
./serverSocket <serverAddress> # serverAddress - <IP_address>:<port_number>
```
- If amd64,
```bash
./amd64serverSocket <serverAddress> # serverAddress - <IP_address>:<port_number>
```
Where:

serverAddress (optional) follows the format <IP_address>:<port_number>.
If not provided, it defaults to localhost, which means you can do the execution of the sockets internally

#### ClientSocket
- Ensure the serverSocket is running on the desired system
- If arm64,
```bash
./clientSocket <serverAddress> # serverAddress - <IP_address>:<port_number>
```
- If amd64,
```bash
./amd64clientSocket <serverAddress> # serverAddress - <IP_address>:<port_number>
```
Where:

serverAddress (optional) follows the format <IP_address>:<port_number>.
If not provided, it defaults to localhost
Use the address specified in serverSocket


### 3. Web App
```bash
cd ./react-app
npm run dev
```
- The web app uses static json files gather from running from running client server connection (milestone folders)

## Considerations
- **Use known message contents** (e.g., numbered sequences) for validation.
- **Minimize unnecessary traffic** when testing on non-CS networks.
- **Beware of firewalls** that may block UDP packets.

## Ports Used on the Server for serverSockets
```
129.3.20.26:26900 - milestone1
129.3.20.26:26901 - milestone2
129.3.20.26:26903 - milestone3
129.3.20.26:26904 - milestone4
```

## References
- [Go Networking Package](https://pkg.go.dev/net)
- [React Charting Libraries (D3.js)](https://d3js.org/)

## Author
**Phone Pyae Sone Phyo**

