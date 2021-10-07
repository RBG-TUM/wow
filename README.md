# Wow!

Simple 1-minute-setup docker-compose thingy that offers a part of Wowzas core functionallity (and some premium features):
- Accepts Ingested streams via rtmp
- Encodes streams into multiple bitrate versions
- Outputs streams via hls
    - Including DVR (seekable stream)


### Usage

- Step 0: get docker and docker compose.
- Step 1
```bash
docker-compose build
docker-compose up -d
```
That's it. Your server is running. You can now
- Step 2: ingest a stream:

This will create an adaptive stream (be carefull about your cpu though, this can get nasty):
```bash
# be a little more creative, especially if you are unsure about your input (you should go for flv or h.264 with some reasonable encoding settings)
ffmpeg -re -i my_video.mp4 -f -flv rtmp://localhost/live1/streamName
```

This stream is not recoded, saving you tons of cpu capacity and making your viewers sad:
```bash
ffmpeg -re -i my_video -f -flv rtmp://localhost/live2/streamName 
```
- Play your stream:

The stream will wait for you at http://localhost:8080/hls/streamName.m3u8

### todo

If this project will ever get some love, features could be:
- Web interface to create authenticated streamers, all others can be rejected.
- Reduce latency/fine-tune ffmpeg
- Rotate/remove old hls segments
- Make it Scalable™
    - Step 1: Multiple output nodes with docker/k8s orchestration in mind
    - Step 2: Multiple input/transcoding nodes