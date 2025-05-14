# Tesodev Product API

## 📌 Proje Hakkında

Bu proje, Tesodev için bir staj başvurusu kapsamında geliştirilmiş bir **CRUD (Create, Read, Update, Delete)** API uygulamasıdır. Proje, kullanıcıların ürün ekleyebileceği, listeleyebileceği, güncelleyebileceği ve silebileceği bir ürün yönetim sistemi sağlar. Proje Go programlama dili ve Echo web framework'ü kullanılarak geliştirilmiştir ve veritabanı olarak MongoDB kullanılmıştır.

---

## 🚀 Kullanılan Teknolojiler

* **Programlama Dili:** Go
* **Web Framework:** Echo
* **Veritabanı:** MongoDB
* **Veritabanı İstemcisi:** MongoDB Shell (mongosh)
* **API Test Aracı:** Postman

---

## 📌 Proje Kurulumu

### 1️⃣ MongoDB Kurulumu ve Çalıştırılması

* MongoDB, yerel olarak yüklendi ve 27018 portunda çalışacak şekilde yapılandırıldı.
* MongoDB Shell (mongosh) ile veritabanı yönetimi sağlandı.
* MongoDB'nin manuel kurulumu sırasında yaşanan zorluklar ve çözüm:

  * MongoDB istemcisi (mongo) yerine mongosh yüklendi.
  * PATH ayarları dikkatlice yapıldı.

### 2️⃣ Proje Kurulumu

```bash
# Proje dizinine gidin
cd ~/desktop/tesodev-product-api

# Gerekli paketleri yükleyin
go get github.com/labstack/echo/v4
```

### 3️⃣ Projeyi Başlatma

```bash
# API Sunucusunu başlatın
go run main.go
```

### 4️⃣ MongoDB ile Bağlantı

* MongoDB 27018 portunda çalışıyor ve API otomatik olarak bu porta bağlanıyor.
* Bağlantı sorunsuz bir şekilde sağlandığında şu mesajı görmelisiniz:

```bash
Connected to MongoDB
⇨ http server started on [::]:8080
```

---

## 📌 API Kullanımı

### 1️⃣ Ürün Ekleme (POST)

* URL: `http://localhost:8080/products`
* Method: `POST`
* JSON Body:

```json
{
  "name": "Product1",
  "price": 100.0
}
```

### 2️⃣ Ürünleri Listeleme (GET)

* URL: `http://localhost:8080/products`

### 3️⃣ Ürün Güncelleme (PUT)

* URL: `http://localhost:8080/products/{id}`
* JSON Body:

```json
{
  "name": "Updated Product",
  "price": 120.0
}
```

### 4️⃣ Ürün Silme (DELETE)

* URL: `http://localhost:8080/products/{id}`

---

## 📌 Proje Yapısı

* **`main.go`**: API'nin giriş noktası.
* **`config/database.go`**: MongoDB bağlantısı yapılandırması.
* **`controllers/product_controller.go`**: API işlemlerinin kontrolü.
* **`models/product.go`**: Ürün model tanımı.
* **`routes/routes.go`**: API rotalarının tanımlandığı yer.

---

## 📌 Öğrendiklerimiz

* MongoDB Shell (mongosh) kullanarak veritabanı yönetimi.
* Go'da Echo framework ile API geliştirme.
* API testi için Postman kullanımı.
* Hata yönetimi ve debugging.

---

## 📌 Sorunlar ve Çözümler

* MongoDB istemcisi (mongo) yerine mongosh kullanıldı.
* PATH ayarları manuel yapıldı.
* API testi için curl ve Postman alternatifleri denendi.
