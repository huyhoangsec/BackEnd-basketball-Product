# 🏀 Ocean Basketball Center - Backend RESTful API Service

Hệ thống Server Backend RESTful API phục vụ quản lý đào tạo, điểm danh học viên, lịch học và hóa đơn học phí cho trung tâm bóng rổ **Ocean Basketball Center**. Xây dựng bằng ngôn ngữ **Go (Golang)**, **Gin Framework** và **GORM (SQLite/PostgreSQL)**.

---

## 🚀 Tính năng chính

### 1. Xác thực & Phân quyền (Auth & RBAC)
- **JWT Authentication**: Đăng nhập, đăng ký và xác thực Token an toàn.
- **Role-based Access Control**: Phân quyền chi tiết cho Admin, Huấn luyện viên và Học viên.

### 2. Quản lý Đào tạo & Điểm danh (Attendance & Coaching)
- **Batch Attendance Processing**: API điểm danh hàng loạt học viên theo buổi học.
- **Coach & Class Stats**: Thống kê số giờ dạy, tỉ lệ chuyên cần học viên của từng HLV.

### 3. Quản lý Học viên & Hóa đơn (Student & Tuition Management)
- **Student CRUD**: Lưu trữ profile học viên, thông tin phụ huynh và lớp đăng ký.
- **Invoice & Revenue Stats**: Quản lý trạng thái hóa đơn học phí, báo cáo doanh thu.

### 4. Hệ thống CMS & Public APIs
- **Public Endpoints**: Cung cấp thông tin sân tập, tin tức, lịch giải đấu cho Frontend.
- **Trial Class Handler**: Tiếp nhận và xử lý đơn đăng ký học thử của học viên mới.

---

## 🛠️ Công nghệ sử dụng

- **Language**: [Go 1.21+](https://go.dev/)
- **Web Framework**: [Gin Gonic](https://gin-gonic.com/)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: SQLite (Development) / PostgreSQL / MySQL (Production)
- **Security & Auth**: JWT (golang-jwt), bcrypt password hashing
- **Deployment**: Docker, Railway / Render

---

## 📂 Thư mục dự án (Project Structure)

```
backend-ocean-basketball/
├── cmd/
│   └── api/
│       └── main.go          # Entrypoint của ứng dụng Go
├── config/
│   └── config.go        # Đọc cấu hình biến môi trường
├── internal/
│   ├── handlers/        # HTTP Handlers (Admin, Auth, Attendance, Student, CMS,...)
│   ├── middleware/      # JWT Auth & CORS Middleware
│   ├── models/          # GORM Models, DB Connection & Seed Data
│   └── routes/          # Khai báo các RESTful API Routes
├── pkg/
│   └── utils/           # Helper functions (JWT Token, Hashing)
├── Dockerfile           # Multi-stage Docker build config
├── go.mod               # Dependencies module
└── railway.json         # Railway deployment configuration
```

---

## 📦 Hướng dẫn Khởi chạy Server Backend

### 1. Yêu cầu hệ thống
- **Go**: `v1.21` trở lên
- **SQLite3** hoặc **PostgreSQL**

### 2. Cài đặt các gói phụ thuộc
```bash
go mod download
```

### 3. Cấu hình Biến môi trường
Tạo file `.env` tại thư mục gốc:
```env
PORT=8080
ENV=development
DB_DRIVER=sqlite
DB_SOURCE=ocean_basketball.db
JWT_SECRET=super-secret-key-ocean-basketball
ALLOWED_ORIGIN=http://localhost:3000
```

### 4. Khởi chạy Server Backend
```bash
go run cmd/api/main.go
```
API Server sẽ lắng nghe tại `http://localhost:8080`.

### 5. Chạy bằng Docker
```bash
docker build -t ocean-basketball-api .
docker run -p 8080:8080 --env-file .env ocean-basketball-api
```
