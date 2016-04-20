#Go User API

[![codebeat badge](https://codebeat.co/badges/bfb98303-9fb5-4f15-808b-de9fd90c5641)](https://codebeat.co/projects/github-com-philippecarle-go-user-api)

##Introduction

This API is meant to be a personnal training on Golang.

It's made to expose a CRUD for users management and to deliver JWT so that other REST APIs can check their authenticity.

It has been coded thanks to the work of these guys :  
- [Gin](github.com/gin-gonic/gin),   
- [Gin-JWT](github.com/appleboy/gin-jwt)  
- [Gin-CORS](github.com/itsjamie/gin-cors)  

##Features

###Routing 

**Login :**  
POST   /login  

**Refresh my token :**  
GET    /token/refresh  

**Get my informations :**  
GET    /users/me  

**Change my password :**  
PATCH  /users/me/change-password  

**Admin (get any user informations, get all users, create a user)**  
GET    /admin/users/:username  
POST   /admin/users  
GET    /admin/users  