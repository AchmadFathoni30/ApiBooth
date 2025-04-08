# Api Store Booth

## Technology
- **go1.24.1**
- **Gin-gonic**
- **GORM**
- **PostgreSQL**

## Fitur
- **Basic Auth** : Menggunakan basic authentication untuk akses endpoint API.
- **Database**: Koneksi ke PostgreSQL menggunakan GORM dengan konfigurasi dari file `.env`.
- **API Endpoint**: Mengambil data dari tabel : 
                    `api_pajak`

## Persyaratan
Sebelum menjalankan aplikasi, pastikan memiliki perangkat lunak yang terinstall :
- **go1.24.1**

### 2. **Instalasi dependency dengan Go mod**
```bash
go mod tidy