
# Library Backend API

Backend untuk sistem manajemen perpustakaan.
Dibuat menggunakan:

* Golang (Gin)
* PostgreSQL
* Redis



# Tech Stack

* Go
* Gin Web Framework
* PostgreSQL 18+
* Redis
* JWT Authentication


# Installation Guide (Windows)

## 1️. Install Go

Download:
[https://go.dev/dl/](https://go.dev/dl/)

Check installation:

```powershell
go version
```

---

## 2️. Install PostgreSQL (Without pgAdmin)

Download:
[https://www.postgresql.org/download/windows/](https://www.postgresql.org/download/windows/)

During installation:

* Port: `5432`

After install, open PowerShell:

```powershell
psql -U postgres
```

Create database & user:

```sql
CREATE USER library WITH PASSWORD 'library123';
CREATE DATABASE library_db OWNER library;
\q
```

---

## 3️. Install Redis (via WSL)

Open PowerShell as Administrator:

```powershell
wsl --install
```

Restart Windows.

Then open WSL:

```powershell
wsl
```

Install Redis:

```bash
sudo apt update
sudo apt install redis -y
sudo service redis-server start
```

Test:

```bash
redis-cli ping
```

Expected:

```
PONG
```

Redis will be accessible at:

```
localhost:6379
```

# Environment Configuration

```
APP_ENV=development
PORT=8080

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=library
POSTGRES_PASSWORD=library123
POSTGRES_DB=library_db

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

SMTP_HOST=localhost
SMTP_PORT=1025
SMTP_PASSWORD=
SMTP_EMAIL=test@example.com

JWT_SECRET=supersecretkey
```

# Running the Application

Go to project root folder:

```powershell
cd ordent-perpustakaan
```

Install dependencies:

```powershell
go mod tidy
```

Run server:

```powershell
go run .
```

If successful:

```
Server running on :8080
```

API available at:

```
http://localhost:8080
```


# API Documentation

1. Open [https://editor.swagger.io/](https://editor.swagger.io/) in browser
2. Copy docs/swagger.yaml
3. Paste in swagger.io


# ERD
- Open file erd.png


# Created by

Davin Bennett
Backend Developer Candidate

# Noted
    Untuk docker sudah saya sertakan, namun ketika saya build dan run, tidak bisa.