import argparse
import socket

def startServer(host,port):
    s = socket.socket()             # Create a socket object
    s.bind((host, port))            # Bind to the port
    s.listen(5)                     # Now wait for client connection.
    
    print('Server listening....')
    
    conn, addr = s.accept()     # Establish connection with client.
    print('Got connection from', addr)
    filename='sample_file.txt'
    
    f = open(filename,'rb')
    l = f.read(1024)

    while (l):
        conn.send(l)
        print('Sent ',repr(l))
        l = f.read(1024)
    f.close()

    print('Done sending')
    conn.close()

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Send and receive over MP-TCP')
    parser.add_argument('host', help='interface the server listens at;'
                        ' host the client sends to')
    parser.add_argument('-p', metavar='PORT', type=int, default=6000,
                        help='TCP port (default 6000)')
    args = parser.parse_args()
    startServer(args.host, args.p)