# video_audio_streaming
# Audio Video Streaming using Go

This project demonstrates how to stream audio and video files using Go. It includes a simple web server that serves HLS (HTTP Live Streaming) video and audio content.

## Prerequisites

Before you begin, ensure you have the following installed:

- Go 
- Git

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/anjush-bhargavan/video_audio_streaming.git
    ```

2. Change into the project directory:

    ```bash
    cd video_audio_streaming
    ```

## Running the Server

To run the server, execute the following command:

```bash
go run main.go
```
3.You can access the HLS streams using the following URL format:

```bash
http://localhost:8080/<video/song>/<filename>/outputlist.m3u8
```

Available files:

Videos: got (Game of Thrones trailer), st4 (Stranger Things trailer), mm (Monkey Man trailer)
Songs: husn, sorry-justinbeiber, shape of u
Replace <video/song> with either video or song, <filename> with one of the available file names, and access the provided link in an HLS-compatible player.

For example:

```bash
http://localhost:8080/video/got/outputlist.m3u8
```

HLS Player
You can use [HLS.js](https://hls-js-latest.netlify.com/demo/) or any other HLS-compatible player to stream the content.
