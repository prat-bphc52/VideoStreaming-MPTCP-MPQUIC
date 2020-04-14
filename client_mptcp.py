import argparse
import socket

def startClient(host, port):
    s = socket.socket()
    s.connect((host, port))

    with open('received_file', 'wb') as f:
        print('file opened')
        while True:
            data = s.recv(1024)
            if not data:
                break
            f.write(data) # write data to file

    f.close()
    print('Successfully stored the file')
    s.close()
    print('connection closed')

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Send and receive over MP-TCP')
    parser.add_argument('host', help='interface the server listens at;'
                        ' host the client sends to')
    parser.add_argument('-p', metavar='PORT', type=int, default=6000,
                        help='TCP port (default 6000)')
    args = parser.parse_args()
    startClient(args.host, args.p)