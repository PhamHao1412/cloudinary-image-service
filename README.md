# 🖼️ Image Service with Cloudinary

A modern image processing REST API built with **Go (Gin)**, **Cloudinary**, and **PostgreSQL**.  
It supports dynamic image transformations such as resize, crop, rotate, watermark, flip, compress, format conversion, and filters — all powered by Cloudinary’s CDN & transformation API.

---

## 🚀 Features

| Feature | Description |
|----------|--------------|
| 🧩 **Upload** | Upload original image to Cloudinary |
| 📏 **Resize** | Resize to specific width and height |
| ✂️ **Crop** | Crop image with width, height, x, y |
| 🔄 **Rotate** | Rotate image by any angle |
| ↔️ **Flip / Mirror** | Flip horizontally or vertically |
| 💧 **Watermark** | Overlay watermark image (Cloudinary public ID) |
| 🗜️ **Compress** | Optimize image size with `q_auto` |
| 🔁 **Convert** | Change format (JPEG, PNG, WEBP, etc.) |
| 🎨 **Filter** | Apply grayscale, sepia, blur, sharpen filters |
| 🧠 **Metadata** | Store and query image info in PostgreSQL |

---

## 🧱 Architecture

```
┌──────────────────────────────────┐
│            Client / API           │
│   (curl / REST Client / Postman)  │
└──────────────────────────────────┘
                │
                ▼
┌──────────────────────────────────┐
│            Gin Server             │
│  - REST routes                    │
│  - JSON handlers                  │
└──────────────────────────────────┘
                │
                ▼
┌──────────────────────────────────┐
│        Service Layer (Go)         │
│  - Image metadata                 │
│  - Transformation logic           │
└──────────────────────────────────┘
                │
                ▼
┌────────────────────────────┐      ┌──────────────────────┐
│      PostgreSQL (DB)       │◄───►│  Cloudinary Storage   │
│  - Store metadata (ID, URL)│      │  - Upload / Transform │
└────────────────────────────┘      └──────────────────────┘
```

---

## ⚙️ Tech Stack

- **Language:** Go 1.22
- **Framework:** Gin
- **Storage:** Cloudinary
- **Database:** PostgreSQL (GORM ORM)
- **Containerization:** Docker + Makefile (no docker-compose required)

---

## 🧰 Prerequisites

- [Go 1.22+](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Cloudinary account](https://cloudinary.com/)
- [PostgreSQL running locally or via Docker]

---

## 🔧 Configuration

Create `.env` file in the root:

```env
PORT=8085
DATABASE_URL=postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable&search_path=<schema>
CLOUDINARY_URL=cloudinary://<api_key>:<api_secret>@<cloud_name>
```

---

## 🏗️ Build & Run with Makefile

### 🔨 Build Docker image
```bash
make build
```

### 🚀 Run container
```bash
make run
```

### 📜 View logs
```bash
make logs
```

### 🛑 Stop and remove container
```bash
make stop
```

### 🧽 Clean image
```bash
make clean
```

---

## 🔥 API Endpoints

| Method | Endpoint | Description |
|--------|-----------|-------------|
| POST | `/upload` | Upload new image |
| POST | `/resize` | Resize image |
| POST | `/crop` | Crop section |
| POST | `/rotate` | Rotate |
| POST | `/flip` | Mirror / Flip |
| POST | `/watermark` | Add watermark |
| POST | `/compress` | Compress image |
| POST | `/convert` | Change format |
| POST | `/filter` | Apply filter |
| GET  | `/images/:id` | Get metadata |

---

## 🧪 Quick Demo with REST Client

Use the included `docs/api.http` file (VS Code REST Client):

1. Install [REST Client extension](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)
2. Open `docs/api.http`
3. Click **“Send Request”** above each API block
4. See JSON response inline

Or use curl manually:
```bash
curl -F "file=@black_goku.jpg" http://localhost:8085/api/v1/upload
```

---

## 🧾 Example Response

```json
{
  "id": "a8e04a73-3e93-4cb3-9a54-29cfd4dd293b",
  "url": "https://res.cloudinary.com/dlqwa0yhj/image/upload/v1730000000/image-service/originals/a8e04a73.jpg",
  "format": "jpg",
  "uploaded_at": "2025-10-22T09:00:00Z"
}
```

---

## 🧠 Notes

- Gin logs use **colored console output**
- GORM logs **show SQL queries**
- Designed with **Clean Architecture** (api → service → storage → db)

---

## 🧾 License

MIT © 2025 — Created by [Hao Pham](mailto:hao.pham@kyanon.digital)

---

## 🌐 References

- [Cloudinary Go SDK](https://github.com/cloudinary/cloudinary-go)
- [Gin Framework](https://gin-gonic.com/)
- [GORM ORM](https://gorm.io/)
