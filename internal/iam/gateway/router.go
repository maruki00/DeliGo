package api

import (
	"github.com/gin-gonic/gin"
)

var userRouter = func(router *gin.Engine) {
	var api = router.Group("api/v1/")
	var user = api.Group("user")

	user.POST("/create-policy", nil)
	user.POST("/create-policy", nil)
	user.POST("/create-policy", nil)
	user.POST("/create-policy", nil)
}

var identityRouter = func(router *gin.Engine) {
	var api = router.Group("api/v1/")
	var identity = api.Group("identity")

	identity.POST("add-policy", nil)
	identity.POST("add-group-policy", nil)
	identity.POST("remove-policy", nil)
	identity.POST("remove-group-policy", nil)
	identity.POST("affect-permission", nil)

}

var authRouter = func(router *gin.Engine) {
	var api = router.Group("api/v1/")
	var auth = api.Group("auth")

	auth.POST("login", nil)
}

/*


-- Core Users table
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY, -- UUID v4
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    phone VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    role ENUM('admin', 'customer', 'restaurant_owner', 'courier') NOT NULL,
    status ENUM('created','active', 'suspended', 'banned') DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Casbin Rules table (Automatically managed by Casbin GORM adapter)
CREATE TABLE casbin_rule (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ptype VARCHAR(100) NOT NULL, -- "p" for policy, "g" for role inheritance group
    v0 VARCHAR(100),             -- sub (subject / role)
    v1 VARCHAR(100),             -- obj (resource / API path)
    v2 VARCHAR(100),             -- act (GET, POST, PUT, DELETE)
    v3 VARCHAR(100),
    v4 VARCHAR(100),
    v5 VARCHAR(100),
    INDEX idx_casbin_rule (ptype, v0, v1, v2)
);
```

user/aprove
user/ban
user/get
user/create



identity/add-policy
identity/add-group-policy
identity/remove-policy
identity/remove-group-policy
identity/affect-permission

auth/login
auth/grante-access //grpc

*/
