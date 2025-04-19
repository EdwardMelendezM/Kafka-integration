# 🎥 LiveStream MVP

A scalable live streaming MVP built with:

- 🧠 **Go (Gin)** for the backend and Kafka integration
- ⚡ **Kafka** as an event broker for stream lifecycle events
- 📺 **OBS Studio** as the video source
- 🌐 **Next.js (App Router)** frontend with HLS video playback
- 📡 **Nginx with RTMP module** for ingesting RTMP and serving HLS

---

## 📦 Features

- Start live streams from OBS using a unique stream key
- Publish stream metadata via Kafka
- Frontend auto-discovers available streams
- View real-time HLS video using `video.js`
- Production-ready Dockerized setup

---

## 🛠️ Tech Stack

| Layer       | Tech                                |
|------------|-------------------------------------|
| Frontend   | [Next.js 14](https://nextjs.org/), App Router, Server Actions |
| Backend    | [Go](https://golang.org/), [Gin](https://gin-gonic.com/), [Kafka](https://kafka.apache.org/) |
| Streaming  | [OBS Studio](https://obsproject.com/), [Nginx RTMP](https://github.com/arut/nginx-rtmp-module) |
| Messaging  | Kafka Topics for stream events       |
| Video Play | HLS (.m3u8) with `video.js`          |
| DevOps     | Docker, Docker Compose               |

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/your-username/livestream-mvp.git
cd livestream-mvp