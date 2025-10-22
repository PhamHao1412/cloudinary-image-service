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
- **Container:** Docker & Docker Compose  

---

## 🧰 Prerequisites

- [Go 1.22+](https://go.dev/)
- [Docker & Docker Compose](https://www.docker.com/)
- [Cloudinary account](https://cloudinary.com/)
- (Optional) [VS Code REST Client extension](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)

---

## 🏗️ Setup & Run

### 1️⃣ Clone repository
```bash
git clone https://github.com/<your-username>/image-service-cloudinary.git
cd image-service-cloudinary
```

### 2️⃣ Configure Cloudinary credentials
Edit in `internal/storage/cloudinary_storage.go`:
```go
cloudinary.NewFromParams("<cloud_name>", "<api_key>", "<api_secret>")
```
or export environment variable:
```bash
export CLOUDINARY_URL=cloudinary://<api_key>:<api_secret>@<cloud_name>
```

### 3️⃣ Start services (API + Postgres)
```bash
docker compose up --build
```

App runs at 👉 **http://localhost:8085**

---

## 🔥 API Examples

See `api.http` file (VS Code REST client)  
or run manually with curl:

```bash
curl -F "file=@black_goku.jpg" http://localhost:8085/api/v1/upload
```

**Response:**
```json
{
  "id": "uuid",
  "url": "https://res.cloudinary.com/xxx/image/upload/...jpg",
  "format": "jpg",
  "uploaded_at": "2025-10-22T04:35:00Z"
}
```

Then you can:
```bash
curl -X POST http://localhost:8085/api/v1/resize   -H "Content-Type: application/json"   -d '{"id":"<uuid>","width":400,"height":300}'
```

Full list of endpoints:
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

Use the included `api.http` file to quickly test:
- Open in VS Code
- Click **Send Request**
- See JSON response inline

---

## 🧑‍💻 Development Notes

- Uses GORM for ORM and migrations.
- Logs:
  - Gin colored HTTP logs.
  - GORM query logs enabled.
- Clean Architecture pattern (api → service → storage/db).

---

## 🧾 License

MIT © 2025 — Created by Hao Pham  
This project is for educational and open-source purposes.

---

## 🌐 References

- [Cloudinary Go SDK](https://github.com/cloudinary/cloudinary-go)
- [Gin Framework](https://gin-gonic.com/)
- [GORM ORM](https://gorm.io/)
