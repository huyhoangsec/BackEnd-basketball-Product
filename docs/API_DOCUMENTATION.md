# Backend API Documentation - Ocean Basketball

Tài liệu hướng dẫn các RESTful API của hệ thống Backend bóng rổ Ocean.

## Base URL
`http://localhost:8080/api/v1`

## Endpoints

### 1. Authentication
- `POST /auth/login`: Đăng nhập hệ thống
- `POST /auth/register`: Đăng ký tài khoản

### 2. Quản lý Học Viên & Điểm danh
- `GET /students`: Lấy danh sách học viên
- `POST /attendance/batch`: Gửi danh sách điểm danh theo lớp

### HTTP Status Codes
- `200 OK`: Thành công
- `400 Bad Request`: Dữ liệu gửi lên không hợp lệ
- `401 Unauthorized`: Chưa xác thực hoặc Token hết hạn
- `500 Internal Server Error`: Lỗi hệ thống server
