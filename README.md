# Video Streaming based on MP-TCP and MP-QUIC
Network Programming Project
It allows you to stream your your webcam video from one system to the other. The protocols used for streaming are MP-TCP and MP-QUIC.
You can separately try out MP-TCP or MP-QUIC while using this code. Please check the requirement below for running MP-TCP or MP-QUIC.

## Requirements
### MP-TCP
- MP-TCP Kernel (Not supported on all OS):
For Linux based OS, please refer to the following [link](https://multipath-tcp.org/pmwiki.php/Users/AptRepository/ "link") for installing MP-TCP protocol on your kernel.

- Python (Python 3.5 recommended)

- Python libraries - **numpy** **opencv-python**
```
python3 -m pip install numpy, opencv-python
```

### MP-QUIC
- GO Language
- [quic](https://github.com/lucas-clemente/quic-go "quic Library")
- [mp-quic](https://github.com/qdeconinck/mp-quic "mpquic Library")
- GOCV - [link](https://gocv.io/ "link") - Refer to the installation instructions for GoCV

## Installation
- Clone this repository in your preferred directory

```
git clone https://github.com/prat-bphc52/VideoStreaming-MPTCP-MPQUIC
```
- Or you can also download the source code as a zip file

## Execution
### MP-TCP
Start the server on the Video Streaming Source Host

```
python3 server_mptcp.py localhost -p <port_number>
```

Start the client on the target machine
```
python3 client_mptcp.py <source_machine_IPv4_Addres> -p <port_number>
```

### MP-QUIC
Build and execute server on one host
```
go build server-mpquic.go
./server-mpquic
```
Specify the server's host name in client-mpquic.go
Build and execute client on the other host
```
go build client-mpquic.go
./client-mpquic
```

## Team Members
- [Prateek Agarwal](https://github.com/prat-bphc52/ "Prateek Agarwal") - 2017A7PS0075H
- [Naman Arora](https://github.com/namanarora00/ "Naman Arora") - 2017A7PS0175H
- [Utkarsh Grover](https://github.com/utkgrover/ "Utkarsh Grover") - 2017A7PS1428H
- Samar Kansal - 2016AAPS0196H
