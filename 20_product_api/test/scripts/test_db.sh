#!/bin/bash
#bu sh kodu docker içinde postgresql kurulumu ve veritabanı oluşturması tablo oluşturması yapar

#Docker Ayagı Kaldırılır Bir Image Yaratılı Bu Image Adı:postgres-test olur
#iki env deger verilir kullanıcı adı ve şifre 
#port tanımlanması yapılır:6432 portundan 5432(defatul postgres) portuna baglancagız diye
#-d ile backgroundda çalışcagını soyleriz ve image olarakta postgres sql in son surumunu alcagını belirtiriz
docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 6432:5432 -d postgres:latest

echo "Postgresql Starting..."
sleep 3 #saniye bekletilir

# docker içine girilir image olarka oluşturdugumuz image verilir postgres-test
# daha sonra postgresql baglantısı yapılır postges kullanıcı adı ile postgres veritabanına baglancagımı belirtilir
# bir sorgu calistrilir bu sorgu da productdb veri tabanını olusturur
docker exec -it postgres-test psql -U postgres -d postgres -c "CREATE DATABASE productdb"

sleep 3
echo "Database Created"

docker exec -it postgres-test psql -U postgres -d productdb -c "
CREATE TABLE IF NOT EXISTS products
(
    id bigserial not null primary key,
    name varchar(255) not null,
    price double precision not null,
    discount double precision,
    store varchar(255) not null
);
"
sleep 3
echo "Table Products Created"