#https://docs.opencv.org/4.7.0/d2/d75/namespacecv.html

import cv2

video_camera = cv2.VideoCapture(0)

while True:
    #カメラから1フレーム読み込む
    _, video_frame = video_camera.read()

    #カメラから取得した1フレームを外部ウィンドウに表示
    cv2.imshow('video' , video_frame)

    # 16ms待機 
    cv2.waitKey(16)
