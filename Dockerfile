# Sử dụng hình ảnh Golang chính thức
FROM golang:1.20

# Tạo thư mục làm việc
WORKDIR /app

# Sao chép go.mod và go.sum để quản lý module
COPY go.mod go.sum ./

# Tải về các dependencies
RUN go mod download

# Sao chép toàn bộ mã nguồn
COPY . .

# Biên dịch ứng dụng thành tệp thực thi "main"
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Chạy ứng dụng
CMD ["./main"]
