## To Run:
CD into the project folder and write the following 
```
go get -u github.com/gorilla/max
go get -u gorm.io/gorm gorm.io/driver/sqlite
```
After that, you will be able to run the program
```
go run main.go
```
The program will check if :8080 is open. If it is, it will occupy it and if not the program will throw a fatal error
