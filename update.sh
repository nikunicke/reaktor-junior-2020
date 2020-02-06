service reaktor-app stop
go build -o reaktor-app
cp reaktor-app /var/www/reaktor-app/bin
service reaktor-app start
service reaktor-app status