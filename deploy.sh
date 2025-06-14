# /bin/sh

cd ..
zip -r TemperatureMonitor.zip TemperatureMonitor

cd TemperatureMonitor
scp ../TemperatureMonitor.zip pi@rpi.nidus:/home/pi/Projects/TemperatureMonitor.zip
scp build-and-start.sh pi@rpi.nidus:/home/pi/Projects/build-and-start-TemperatureMonitor.sh
ssh pi@rpi.nidus 'cd Projects && chmod +x build-and-start-TemperatureMonitor.sh && ./build-and-start-TemperatureMonitor.sh'
