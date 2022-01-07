# implementasi google Oauth

## Apa itu oauth
Oauth merupakan sebuah standard protocol untuk melakukan authorisasi dan authentikasi.
akan tetapi alih-alih kita impelemntasi authorisasi dan authntikasi sendiri, disini kita justru bertindak seperti
client dan yang mengatur authorisasi nya adalah aplikasi pihak ke 3.
lalu aplikasi kita menggunakan secret key agar bisa melakukan komunikasi dengan pihak ke 3 tersebut.
nah aplikasi yang menyediakan service ouath ada banyak, salah satunya adalah google.


## Cara kerja oauth
<p align="center"><img src="./img/mekanisme_oauth.png" /></p>

- user melakukan login dengan google oauth melalui aplikasi kita
- lalu aplikasi kita akan melakukan redirect ke halaman login google
- pada halaman url redirect terdapat parameter akses yang akan diminta
- setelah user berhasil login, google akan mengirim jsonwebtoken kepada aplikasi kita melalui protocol http
- jsonwebtoken yang dikirim ada 2, yaitu access_token dan refresh_token
- access token digunakan untuk untuk mengakses service yang dilindungi
- sedangkan refresh token untuk meminta access_token baru jika access_tokenya sudah expired
- selanjutnya kita dapat melakukan kita dapat melakukan api call ke google menggunakan token tadi, contohnya
  seperti mendapatkan informasi user dari token tersebut

## Run project di local
- clone repository nya 
- isi client id dan secret id di file `env.example`
- setelah itu rename file `env.example` menjadi `.env`
- setting redirect url di file `config.go`
- sesuaikan variabel `devRedirectUri` dan `prodRedirectUri`
- install dependencies : `go mod download`
- run poject : `go run .` 
