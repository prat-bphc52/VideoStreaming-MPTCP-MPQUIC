import argparse
import cv2
import numpy as np

def encodeNumPyArray(fr):
    mybytes = bytearray()
    for i in fr:
        for j in i:
            mybytes.append(int(j[0]))
            mybytes.append(int(j[1]))
            mybytes.append(int(j[2]))
    return mybytes

if __name__ == '__main__':
    print("Video Streaming !")
    cap  = cv2.VideoCapture(0)
    if not cap.isOpened():
        raise IOError("Cannot open webcam")

    count = 0
    while True:
        ret, frame = cap.read()
        frame = cv2.resize(frame, None, fx=  0.5, fy = 0.5, interpolation = cv2.INTER_AREA)
        cv2.imshow('Input', frame)
        count = count+1
        temp = encodeNumPyArray(frame)
        print('Frame count ', count)
        print(type(frame))
        print(len(temp))
        print(len(frame))
        print(len(frame[0]))
        print(frame)
        if count==1:
            break
        c = cv2.waitKey(1)
        if c==27:
            break
    cap.release()
    cv2.destroyAllWindows()