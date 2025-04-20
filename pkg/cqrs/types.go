package pkgCqrs

type Middleware func(next CommandHandler) CommandHandler
