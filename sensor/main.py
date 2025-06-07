import time

import machine
import network
import requests
import ubinascii
import urandom

ssid = ""
password = ""

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


def readTemperature(sensor) -> float:
    adc_value = sensor.read_u16()
    volt = (3.3 / 65535) * adc_value
    temperature = 27 - (volt - 0.706) / 0.001721
    return round(temperature, 1)


def connect(ssid, password):
    wlan = network.WLAN(network.STA_IF)
    wlan.active(True)
    wlan.connect(ssid, password)
    while wlan.isconnected() == False:
        print("Waiting for connection...")
        time.sleep(1)


sensor = machine.ADC(4)
led = machine.Pin("LED", machine.Pin.OUT)

led.high()
connect(ssid, password)
led.low()

session = generate_uuid4()

while True:
    led.high()

    now = time.localtime()
    sensor_time = "%04d-%02d-%02dT%02d:%02d:%02d.000" % (timestamp[0:6])

    temperature = readTemperature(sensor)

    request_body = {
        "session": session,
        "sensorTime": sensor_time,
        "temperature": temperature,
        "pressure": 0,
    }

    print("Request body: ", request_body)

    try:
        response = requests.put(
            "http://192.168.100.7:8080/sensors/ingest",
            json=request_body,
        )

        response_code = response.status_code
        response_content = response.content

        print("Response code: ", response_code)
        print("Response content:", response_content)

        print()
    except:
        print("Failed sending")

    led.low()

    time.sleep(sensing_period_seconds)
