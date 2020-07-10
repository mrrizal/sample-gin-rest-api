
# sample-gin-rest-api

repository ini dibuat sebagai contoh bagaimana men deploy rest api yang dibuat menggunakan bahasa pemrograman `go` ke heroku.

Berikut adalah step by step bagaimana mendeploy ke heroku:

 - Clone repository: `git clone git@github.com:mrrizal/sample-gin-rest-api.git`
 - Change directory: `cd sample-gin-rest-api`
 - Download heroku cli: `sudo snap install --classic heroku`
 - Pastikan anda memiliki akun heroku, login heroku cli: `heroku login`
 - Buat app di heroku: `heroku ceate {app_name}`, contoh `heroku create gin-rest-api`
 - Hal yang perlu diperhatikan adalah file `Procfile`, isinya adalah perintah yang akan digunakan untuk menjalankan aplikasi. Contoh pada aplikasi ini adalah `web: bin/main`, file `bin/main`ini adalah executeable file untuk menjalankan aplikasi, Contoh lainnya dapat dilihat [disini](https://github.com/mrrizal/Rastakhans-Rumble-API/blob/master/RastakhansRumble/Procfile)
 - Setelah membuat app di heroku, ketik: `git remote -v` untuk melihat dan memastikan remote branch heroku muncul setelah anda mengetik perintah tesebut. 

![git remote -v](https://i.imgur.com/2DzN95v.png)
 - Push & deploy code ke heroku: `git push heroku master` , berikut tampilan jika subah berhasil push code & deploy ke heroku

![git push heroku master](https://i.imgur.com/mR8pqMA.png)
- Selanjutnya untuk membuka app kita di browser, ketikan: `heroku open`
 - Anda juga bisa men test aplikasi di local dengan mengetikan `heroku local`
 - Berikut contoh aplikasi yang sudah saya deploy di heroku [link](https://gin-rest-api.herokuapp.com/ping)
