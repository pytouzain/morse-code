# morse-code

Small project that translate messages into morse code thanks to a LED linked to a raspberry pi

setup:
GOARM=6 GOARCH=arm GOOS=linux go build -o morse morse.go
scp morse pi@ip.of.your.raspberry.pi:path/where/to/copy/on/your/raspberry