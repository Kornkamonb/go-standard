# Tool
# go install github.com/air-verse/air@latest

go clean -cache
go build -x

# Database postgres
go get github.com/jackc/pgx/v5@v5.8.0

# Database oracle old version
go get github.com/godror/godror@v0.50.0

# Utility
go get github.com/robfig/cron/v3@v3.0.1
go get github.com/fsnotify/fsnotify@v1.7.0
go get gopkg.in/yaml.v3@v3.0.1
go get github.com/joho/godotenv@v1.5.1
go get github.com/xuri/excelize/v2@v2.10.0

# API
# go get github.com/gin-gonic/gin@v1.11.0
# go get github.com/gin-contrib/cors@v1.7.6

go mod tidy
go mod vendor
go mod verify