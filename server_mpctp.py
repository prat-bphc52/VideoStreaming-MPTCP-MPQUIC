import argparse
import socket
import cv2
import utils

def startServer(host,port):
    s = socket.socket(socket.AF_INET6, socket.SOCK_STREAM)             # Create a socket object
    s.bind((host, port))            # Bind to the port
    s.listen(5)                     # Now wait for client connection.
    
    print('Server listening....')
    
    conn, addr = s.accept()     # Establish connection with client.
    print('Got connection from', addr)
    
    cap  = cv2.VideoCapture(0)
    if not cap.isOpened():
        conn.close()
        raise IOError("Cannot open webcam")

    count = 0
    while (True):
        ret, frame = cap.read()
        frame = cv2.resize(frame, None, fx=  0.5, fy = 0.5, interpolation = cv2.INTER_AREA)
        count = count + 1
        if count>500:
            break
        print(type(frame[0][0][0]))
        conn.sendall(utils.encodeNumPyArray(frame))
        print('Sent ', count)
    cap.release()
    cv2.destroyAllWindows()
    print('Ending video streaming')
    conn.close()

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Send and receive over MP-TCP')
    parser.add_argument('host', help='interface the server listens at;'
                        ' host the client sends to')
    parser.add_argument('-p', metavar='PORT', type=int, default=6000,
                        help='TCP port (default 6000)')
    args = parser.parse_args()
    startServer(args.host, args.p)