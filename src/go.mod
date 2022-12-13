module gorest

go 1.19

require (
	app/dbstorage v0.0.0-00010101000000-000000000000 // indirect
	app/user/actions v0.0.0-00010101000000-000000000000 // indirect
	app/user/controllers v0.0.0-00010101000000-000000000000 // indirect
	app/user/exceptions v0.0.0-00010101000000-000000000000 // indirect
	app/user/handlers v0.0.0-00010101000000-000000000000 // indirect
	app/user/models v0.0.0-00010101000000-000000000000 // indirect
	app/user/repositories v0.0.0-00010101000000-000000000000 // indirect
	app/user/transformers v0.0.0-00010101000000-000000000000 // indirect
	github.com/cosmtrek/air v1.40.4 // indirect
	github.com/creack/pty v1.1.18 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gin-contrib/cors v1.4.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.8.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/sqlite v1.4.3 // indirect
	gorm.io/gorm v1.24.2 // indirect
)

replace app/dbstorage => ./app/dbstorage

replace app/user/actions => ./app/user/actions

replace app/user/models => ./app/user/models

replace app/user/exceptions => ./app/user/exceptions

replace app/user/handlers => ./app/user/handlers

replace app/user/repositories => ./app/user/repositories

replace app/user/transformers => ./app/user/transformers

replace app/user/controllers => ./app/user/controllers
