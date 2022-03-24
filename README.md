# go-rest-api

[you can see video tutorial](https://user-images.githubusercontent.com/72017753/159831538-ae5022d7-5bc4-418a-a948-2191001752e4.mp4)


1. Restful API dengan Golang
2. Gorilla Mux sebagai router
3. MongoDB sebagai database


Step untuk penginstalan
1. install dan setup mongodb [download disini jika belum ada](https://www.mongodb.com/try/download/community)
2. buat database dengan nama "linkaja"
3. pada folder cmd silahkan ```bash go install service.go ```
4. setelahnya silahkan ```bash go run service.go ```
5. setelahnya POST data pada ":3001/account" 
```bash 
{
  "account_number" : "55501", "custumer_name" : "Bob Martin", "balance" : 10000
},
{
  "account_number" : "55502", "custumer_name" : "Linus Torvalds", "balance" : 15000
}
```
6. Lakukan GET account pada ":3001/account/{account_number}"
7. Lakukan transfer saldo dengan metode POST pada ":3001/account/{from_account_number}/transfer" dengan isi
```bash
{
  "to_account_number" : "nomor akun tujuan",
  "amount": 1000
}
```
