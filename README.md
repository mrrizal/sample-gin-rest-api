
# sample-gin-rest-api

repository ini dibuat sebagai contoh bagaimana men deploy rest api yang dibuat menggunakan bahasa pemrograman `go` ke heroku.

Berikut adalah step by step bagaimana mendeploy ke heroku:

 - Clone repository: `git clone git@github.com:mrrizal/sample-gin-rest-api.git`
 - Change directory: `cd sample-gin-rest-api`
 - Download heroku cli: `sudo snap install --classic heroku`
 - Pastikan anda memiliki akun heroku, login heroku cli: `heroku login`
 - Create app on heroku: `heroku ceate {nama_aplikasi}`, contoh `heroku create gin-rest-api`
 - Hal yang perlu diperhatikan adalah file `Procfile`, isinya adalah perintah yang akan digunakan untuk menjalankan aplikasi. Contoh pada aplikasi ini adalah `web: bin/main`, file `bin/main`ini adalah executeable file untuk menjalankan aplikasi, Contoh lainnya dapat dilihat [disini](https://github.com/mrrizal/Rastakhans-Rumble-API/blob/master/RastakhansRumble/Procfile)
 - 
