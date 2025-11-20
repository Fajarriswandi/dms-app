# Database Connection Guide - DBeaver

## Koneksi ke SQLite Database

### Informasi Database
- **Type**: SQLite
- **Database File**: `backend/dms.db`
- **Full Path**: `/Users/f/Documents/Projects/dms-app/backend/dms.db`

## Konfigurasi DBeaver

### 1. Pilih Driver
- **Driver name**: SQLite
- Pastikan SQLite driver sudah terinstall di DBeaver

### 2. Connection Settings

#### Option 1: Menggunakan Path (Recommended)

**Connect by:** Pilih **"Host"**

**JDBC URL:**
```
jdbc:sqlite:/Users/f/Documents/Projects/dms-app/backend/dms.db
```

**Atau biarkan default dan isi Path:**
- **Path:** `/Users/f/Documents/Projects/dms-app/backend/dms.db`
- Klik button **"Open ..."** untuk browse ke file

#### Option 2: Menggunakan URL

**Connect by:** Pilih **"URL"**

**JDBC URL:**
```
jdbc:sqlite:/Users/f/Documents/Projects/dms-app/backend/dms.db
```

### 3. Test Connection
1. Klik **"Test Connection ..."**
2. Jika driver belum terinstall, DBeaver akan menawarkan untuk download
3. Setelah berhasil, klik **"Finish"**

### 4. Connection Name (Optional)
- **Name:** DMS App Database
- **Type:** SQLite

## Informasi Tambahan

### Lokasi File Database
```
Absolute Path: /Users/f/Documents/Projects/dms-app/backend/dms.db
Relative Path: backend/dms.db (dari project root)
```

### Tables
Setelah connect, Anda akan melihat:
- **users** - Table untuk user accounts

### Query Contoh

```sql
-- Lihat semua users
SELECT * FROM users;

-- Lihat schema table
.schema users

-- Lihat semua tables
SELECT name FROM sqlite_master WHERE type='table';
```

## Troubleshooting

### File tidak ditemukan
- Pastikan backend sudah running minimal sekali
- Database file dibuat saat pertama kali `InitDB()` dipanggil
- Jika belum ada, restart backend: `make restart`

### Driver SQLite tidak tersedia
1. Klik **"Driver Settings"** di DBeaver
2. Download SQLite driver jika belum ada
3. Atau install via DBeaver: **Help > Install New Software > Add SQLite Driver**

### Path dengan space
Jika path memiliki space, gunakan absolute path dengan quotes atau escape spaces.

## Catatan Penting

⚠️ **Database Lock:**
- Pastikan backend **tidak running** saat mengakses dari DBeaver
- SQLite tidak support concurrent writes dari multiple applications
- Atau gunakan **read-only mode** di DBeaver

### Workflow Recommended:
1. **Development:** Stop backend (`make down` atau `docker-compose down`)
2. **View Database:** Connect dari DBeaver
3. **Continue Development:** Start backend lagi (`make dev`)

