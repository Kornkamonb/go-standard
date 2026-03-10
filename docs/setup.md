# Version Go 1.25
https://hub.docker.com/_/golang/

# Install
mkdir -p ~/download
cd ~/download
wget https://go.dev/dl/go1.25.7.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.25.7.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version

# Compile and install the application
go build -o build/ ./cmd/cron
go build -o build/ ./cmd/main
go build -o build/ ./cmd/template-postgres
go build -o build/ ./cmd/template-oracle

# Run Compile
./build/cron
./build/main
./build/template-postgres
./build/template-oracle

# Test cron
go build -o build/ ./cmd/main
./build/cron "--run-now" "*/5 * * * * *" "./build/main"

# Proxy
export http_proxy="http://10.17.77.184:34567/"
export https_proxy="http://10.17.77.184:34567/"
export no_proxy="localhost,127.0.0.1,::1,192.168.0.0/16,10.0.0.0/8,172.16.0.0/12"
export HTTP_PROXY="http://10.17.77.184:34567/"
export HTTPS_PROXY="http://10.17.77.184:34567/"
export NO_PROXY="localhost,127.0.0.1,::1,192.168.0.0/16,10.0.0.0/8,172.16.0.0/12"

# Install direnv
apt-get update
apt-get install direnv
echo 'eval "$(direnv hook bash)"' >> ~/.bashrc
