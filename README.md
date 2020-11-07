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
.
