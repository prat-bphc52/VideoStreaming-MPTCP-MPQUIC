import argparse
import cv2
import numpy as np

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
        print('Frame count ', count)
        if count==100:
            break
        c = cv2.waitKey(1)
        if c==27:
            break
    cap.release()
    cv2.destroyAllWindows()