import time
from typing import Tuple

import dht
import machine
import network
import requests
import ubinascii
import urandom

ssid = ""
password = ""

ingestion_url = "http://192.168.100.7:8080/ingest"

sensing_period_seconds = 5


def generate_uuid4():
    random_bytes = bytearray(16)
    for i in range(16):
        random_bytes[i] = urandom.getrandbits(8)

    # Set version (UUID4)
    random_bytes[6] = (random_bytes[6] & 0x0F) | 0x40
    # Set variant
    random_bytes[8] = (random_bytes[8] & 0x3F) | 0x80

    uuid = ubinascii.hexlify(random_bytes).decode().upper()
    return f"{uuid[0:8]}-{uuid[8:12]}-{uuid[12:16]}-{uuid[16:20]}-{uuid[20:32]}"


def readTemperatureAndPressure(sensor: dht.DHT22) -> Tuple[float, float]:
    sensor.measure()
    temperature = sensor.temperature()
    humiditiy = sensor.humidity()
    return (temperature, humiditiy)


def connect(ssid, password):
    wlan = network.WLAN(network.STA_IF)
    wlan.active(True)
    wlan.connect(ssid, password)
    while wlan.isconnected() == False:
        print("Waiting for connection...")
        time.sleep(1)


sensor = dht.DHT22(machine.Pin(22))
led = machine.Pin("LED", machine.Pin.OUT)

led.high()
# connect(ssid, password)
led.low()

session = generate_uuid4()

while True:
    led.high()

    now = time.localtime()
    sensor_time = "%04d-%02d-%02dT%02d:%02d:%02d.000" % (now[0:6])

    (temperature, humidity) = readTemperatureAndPressure(sensor)

    request_body = {
        "session": session,
        "sensorTime": sensor_time,
        "temperature": temperature,
        "humidity": humidity,
    }

    print("Request body: ", request_body)

    # try:
    #     response = requests.put(
    #         ingestion_url,
    #         json=request_body,
    #     )
    #
    #     response_code = response.status_code
    #     response_content = response.content
    #
    #     print("Response code: ", response_code)
    #     print("Response content:", response_content)
    #
    #     print()
    # except:
    #     print("Failed sending")

    led.low()

    time.sleep(sensing_period_seconds)
