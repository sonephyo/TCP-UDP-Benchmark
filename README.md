# Network Latency and Throughput Measurement

## Assignment Overview
This project measures the **latency** and **throughput** of TCP and UDP protocols across multiple machine pairs and networks. The results are visualized on a web page using graphs.

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
### **1. TCP Latency Measurement**
- Measure **RTT** for message sizes: **8, 64, 256, and 512 bytes**.
- Messages are echoed and validated.

### **2. TCP Throughput Measurement**
- Send **1 MB of data** using different message sizes:
  - **1024 messages of 1024 bytes**
  - **2048 messages of 512 bytes**
  - **4096 messages of 256 bytes**
- 8-byte acknowledgment is sent in the reverse direction.

### **3. UDP Latency Measurement**
- Same as TCP latency measurement but using UDP.

### **4. UDP Throughput Measurement**
- Same as TCP throughput measurement but using UDP.

## Tools & Libraries
- **Go (Golang)** (for networking and timing with `time.Now().UnixNano()`)
- **Go net package** (`net.Listen`, `net.Dial`, `net.UDPConn`)
- **React.js** (for visualization using libraries like Chart.js or D3.js)

## Instructions
1. **Set Up the Environment**
   - Ensure Go is installed (`go version`).
   - Compile the Go programs from the shells provided (Note: You can configure the shells to suit according to your environment)

2. **Run the Tests**
   - Start the **server**
   - Start the **client**

## Considerations
- **Use known message contents** (e.g., numbered sequences) for validation.
- **Minimize unnecessary traffic** when testing on non-CS networks.
- **Beware of firewalls** that may block UDP packets.

## References
- [Go Networking Package](https://pkg.go.dev/net)
- [React Charting Libraries (D3.js)](https://d3js.org/)

## Author
**Phone Pyae Sone Phyo**

---
Feel free to update the README with additional details as needed!
