# BGP-Golang-TRPL3A 

Dokumenatasi / *source* Tugas Mata Kuliah **Project 2** (*Web Programming*)\
Bayu Gusti Paraya [link to touch!](http://itsbabay.xyz/)\
1811082018\
D4 - TRPL 3A
TI - PNP
<br><br>

## ~ Materi Pertemuan 5  (Struct di dalam struct pada JSON dan XML)
1. Contoh penarapan pada pertemuan 5 yaitu menggunakan tabel order (database northwind) pada OrderDetail , OrderDetailXML dan ReadJson [(Pertemuan 5)](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/4e733374b8703b9cc9ae030853c24b4a4021588c)
2. Untuk tugas membuat database baru yang berkaitan dengan nilai mahasiswa, yaitu database [bgp_akademik.sql](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/88190654e62a27e54de25fbb40cc562cf9a5c851)
3. [Tugas Pertemuan 5](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/b38588cf24472da5e691016260bdd24458ed8a82) yaitu pada database bgp_akademik. Dimana ada 4 tabel (dosen, mahasiswa, matkul, nilai) yang saling berelasi. Dan menerapkan pada MahasiswaNilai, MahasiswaNilaiXML dan MahasiswaReadJson
<br><br>

## ~ Materi Pertemuan 6  (Integrasi go dengan html)
1. Menggunakan database [Northwind.sql](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/88190654e62a27e54de25fbb40cc562cf9a5c851)
2. [Tugas Pertemuan 6](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/07c33a79d821f728f08fff4e8c8666093ff6077b) yaitu implementasi / mengubah HtmlPage dan HtmlPostData untuk tabel northwind.employees 

## ~ Materi Pertemuan 7  (Implementasi Go dengan Framework)
### + Tugas 4
1. Menggunakan database [Northwind.sql](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/88190654e62a27e54de25fbb40cc562cf9a5c851)
2. Membuat struct untuk product dan customer sesuai tabel di database. [add struct customer, product](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/7a2eab3f256de187879f8251c88904b750958ece) . <br>_loc file:_ /Framework/git/order/common/global.go
3. [add rootUlr yml](Framework/git/order/common/global.go) Menambahkan rootUrl untuk customer dan product. Contohnya localhost:9000/__getCustomer__/... . <br>_loc file:_ /Framework/git/order/conf-dev.yml
4. [add rootUrl on Configuration Struct](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/e766359d4969e878d463f6e3b8593f4c5555dedf) <br>_loc file:_ /Framework/git/order/common/config.go
5. Menambahkan link / url api request yg bisa diakses user. [add url api request](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/da7899897d834fb2a7d1fc8af17d2586ba9a01f7) <br>_loc file:_ /Framework/git/order/main.go
6. Menambahkan endpoint untuk customer dan product. [add endpoints](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/5ea0d4dc40b8c9e58225de59cb4090c4b54ae7b6) <br>_loc file:_ /Framework/git/order/transport/endpoint.go
7. Menambahkan decode request dari json ke customer atau product struct. [add decodeRequest](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/d5d3d3bac0c007398b1c5fc6a51afd8a811e4f19) <br>_loc file:_ /Framework/git/order/transport/transport.go
8. Menambahkan item interface untuk customer dan product. [add interface items](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/6137c9e3f6b712a08d5c287f3c6b9d0ec1bd2f2c) <br>_loc file:_ /Framework/git/order/services/services.go
9. Membuat customer handler, dimana salah satunya untuk menghandel atau eksekusi pengambilan data customer dg query sql . [add customer handler](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/c8247cbb4485bf1024ac67be78fe32464317adcb) <br>_loc file:_ /Framework/git/order/services/handler_customer.go
10. Membuat product handler, dimana salah satunya untuk menghandel atau eksekusi pengambilan data product dg query sql . [add product handler](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/a681c8243d2698c2f46250cbb5345512997553a0) <br>_loc file:_ /Framework/git/order/services/handler_product.go
11. Menambahkan middleware , agar controller fokus dalam menyelesaikan logika alur bisnis dari suatu flow aplikasi tanpa harus untuk melakukan hal — hal di luar itu seperti validasi input untuk setiap flow. [add middleware ](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/7eef7f7fc35ad5040aa5300363fd1b8d77a9aa76) <br>_loc file:_ /Framework/git/order/middleware/basicmw.go
<br>

__POSTMAN Api Dokumentasi silahkan lihat disini [link](https://documenter.getpostman.com/view/7157328/TVYF8Jsh/)__
<br>
## ~ Materi Pertemuan 8  (Implementasi FastPay dimana JSON request dan responsenya berbeda)
### + Tugas 5
1. Menggunakan database [Northwind.sql](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/88190654e62a27e54de25fbb40cc562cf9a5c851) <br>_loc file:_ /Northwind.sql
2. Menambahkan tabel [northwind.trans](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/98ecf6351801e2f216ff91f5bd651d1c32d31605) <br>_loc file:_ /Northwind.trans.sql
3. [add rootUrlData on yaml](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/69cc1708379fc58d5d1f90ac873c548c936bf63e) <br>_loc file:_ /Framework/git/order/conf-dev.yml
4. [add rootUrlData on struct](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/e21c3d4496db039020956d2d2c68975688d7f8fe) . <br>_loc file:_ /Framework/git/order/common/config.go
5. Menambahkan struct untuk FastPay (request, response dan detail). [add FastPay req res struct](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/607b3d8071f8e2aab1d2a01f9e5c0d09385fd412) <br>_loc file:_ /Framework/git/order/common/global.go
6. Menambahkan url api untuk FastPay. [add FastPay url](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/0d5ecf73ab4cffdc809695a7f8486bfca7fe1fba) <br>_loc file:_ /Framework/git/order/main.go
7. Menambahkan endpoint untuk FastPay. [add FastPay ENdpoint](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/e8b49bda83d7dee893e9610835f385e7e4bb4149) <br>_loc file:_ /Framework/git/order/transport/endpoint.go
8. Menambahkan decode request dari json ke FastPay. [add DecodeFastPayRequest](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/4835f16010e4d8937d8579cb2ba2ee4ef382eb62) <br>_loc file:_ /Framework/git/order/transport/transport.go
9. Menambahkan item interface untuk FastPay. [add fastpay interface item](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/0f03623934a1630c6c25bbfdbc91dcb3c17756a6) <br>_loc file:_ /Framework/git/order/services/services.go
10. Membuat FastPay handler (handler_fastpay.go), dimana salah satunya untuk menghandel atau eksekusi pengambilan data customer dg query sql . [add fastpay handler](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/ee49cd58a3dccd377dc80b148ba002fa421bffe4) <br>_loc file:_ /Framework/git/order/services/handler_fastpay.go
11. Menambahkan middleware , agar controller fokus dalam menyelesaikan logika alur bisnis dari suatu flow aplikasi tanpa harus untuk melakukan hal — hal di luar itu seperti validasi input untuk setiap flow. [add fastpay middleware](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/e3be241522566ce129c8d35ce1927f5c142d426a) <br>_loc file:_ /Framework/git/order/middleware/basicmw.go
<br>

__POSTMAN Api Dokumentasi silahkan lihat disini [link](https://documenter.getpostman.com/view/7157328/TVYF8Jsh/)__
<br>
## ~ Materi Pertemuan 9  (eksekusi insert data ke dbms mysql dari response trips)
### + Tugas 6
__Api URL : http://localhost:9000/getOrder/trips__
<br>

1. Menggunakan database [Northwind.sql](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/88190654e62a27e54de25fbb40cc562cf9a5c851) <br>_loc file:_ /Northwind.sql
2. Menambahkan tabel [Northwind.trips.sql](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/69953b9cb986f0680ec6daf74226ccb9c06a05db) <br>_loc file:_ /Northwind.trips.sql
3. Menambahkan struct untuk Trips (request, response dan detail). [add Trips req res struct](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/0b5fd647860302fb68e8ac406388bcdde0a601e0) <br>_loc file:_ /Framework/git/order/common/global.go
4. Menambahkan url api untuk Trips. [add Trips url req](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/05c9567492f83fa5874492987bd8d1c4b1430a9c) <br>_loc file:_ /Framework/git/order/main.go
5. Menambahkan endpoint untuk Trips. [add Trips Endpoint](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/ae8201e015c0a1ec19de5d27bbf5d177f5aa5990) <br>_loc file:_ /Framework/git/order/transport/endpoint.go
6. Menambahkan decode request dari json ke Trips. [add DecodeTripRequest](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/11b14ba1307cf2ba32ec86a4f206d96fd45a7105) <br>_loc file:_ /Framework/git/order/transport/transport.go
7. Menambahkan item interface untuk FastPay. [add trip interface item](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/47db5f09904c5d0ca67d8d2c18c2716e3c3910d6) <br>_loc file:_ /Framework/git/order/services/services.go
8. Membuat Trips handler, yang berguna untuk menghandel atau eksekusi, pengambilan data trip dari api http://35.186.147.192/travel/GetTripsSample.php dan data disimpan pada tabel northwind.trips . [add trips handler](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/ea679782489446be6ef1775681bd34dcf0f52479) <br>_loc file:_ /Framework/git/order/services/handler_trips.go
9. Menambahkan middleware , agar controller fokus dalam menyelesaikan logika alur bisnis dari suatu flow aplikasi tanpa harus untuk melakukan hal — hal di luar itu seperti validasi input untuk setiap flow. [add trips middleware](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/2c9bed01b54d5c4c71b3bc015011ef2941618563) <br>_loc file:_ /Framework/git/order/middleware/basicmw.go
<br>

__POSTMAN Api Dokumentasi silahkan lihat disini [link](https://documenter.getpostman.com/view/7157328/TVYF8Jsh/#a4428df5-4372-4c8f-b1fe-7fc6552d5b65)__
<br>
## ~ Materi Pertemuan 10  (Echo Framework)
Gambaran Struktur File / Folder<br>
__Api URL : http://localhost:3000/groupRoute/route__
```bash
./common/common.go               | berisi struct json
./main.go                        | main (file utama)
  -> ./db/db.go                  | membuat koneksi ke dbms 
    -> ./config/config.json      | berisi json data dbms
    -> ./config/config.go        | memanggil config.json
  -> ./routes/routes.go          | kumpulan group route
    -> ./routes/group_route.go   | berisi route-route
      -> ./controllers/nama_controller.go
        -> ./models/nama_model.go  | inti / core Framework (CRUD)
          -> ./common/common.go  | berisi struct json
          -> ./db/db.go          | membuat koneksi 
-> ./helpers                     | tambahan
-> ./middleware                  |
-> ./go.mod                      | semua library yg dibutuhkan
```
1. Menambahkan tabel [northwind.users](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/55cd79eb789e18e53f828f3b12b0e5b7dccc2334) <br>_loc file:_ /Northwind.users.sql
2. Penerapan [pertemuan 10](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/52102fbd1018b3aa7cfc108a343454fea2cf5a2f)
### + Tugas 7
1. Menggunakan database [Northwind.sql](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/88190654e62a27e54de25fbb40cc562cf9a5c851) <br>_loc file:_ /Northwind.sql tabel suppliers
2. Menambahkan [Suppliers struct](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/73e6db805db8381f108a56612d93cd087ba1b35e) pada common. <br>_loc file:_ /echo-rest/common/common.go
3. Menambahkan group route Suppliers. [add suppliers route group](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/cff6564ed1ae30bdc236bbba8c227245713a8b1c) <br>_loc file:_ /echo-rest/routes/routes.go
4. Membuat route route Supplier pada 1 file. [add suppliers_route](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/84a985dc20a8a2336b33657de57b76cd74c8339d) <br>_loc file:_ /echo-rest/routes/suppliers_route.go
5. Membuat Supplier Controller. [add supplier_controller](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/37098b28e1aaaa1c7249aae386e0cf6c53549f68) <br>_loc file:_ /echo-rest/controllers/supplier_controller.go
6. Membuat model / proses CRUD Suppliers. [add suppliers_model](https://github.com/bayugustiparaya/BGP-Golang-TRPL3A/commit/be83460c8cf377e7091c441c54244e3aaf81a5c3) <br>_loc file:_ /echo-rest/models/suppliers_model.go
<br>
__POSTMAN Api Dokumentasi silahkan lihat disini [link](https://documenter.getpostman.com/view/7157328/TVYF8Jsh/#a4428df5-4372-4c8f-b1fe-7fc6552d5b65)__
<br><br>
```javascript
API URL menggunakan Method POST
Bentuk Umum : http://localhost:3000/groupRoute/route
http://localhost:3000/suppliers/add
http://localhost:3000/suppliers/list
http://localhost:3000/suppliers/update
http://localhost:3000/suppliers/delete
```
<br>