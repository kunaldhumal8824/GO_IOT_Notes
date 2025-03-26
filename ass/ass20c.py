# C. Block Diagram for Raspberry Pi/Arduino with Sensors
# While I can't draw diagrams directly here, I can describe how you might set up a block diagram for interfacing a temperature sensor (like the DHT11) with a Raspberry Pi or Arduino.

# Block Diagram Description
# Raspberry Pi/Arduino: The main microcontroller or microprocessor.
# Temperature Sensor (DHT11): Connected to a GPIO pin on the Raspberry Pi/Arduino.
# Power Supply: Ensure the sensor is powered correctly (usually 5V).
# Wiring:
# Connect the VCC pin of the sensor to the 5V pin on the Raspberry Pi/Arduino.
# Connect the GND pin of the sensor to the ground.
# Connect the data pin of the sensor to a GPIO pin (e.g., GPIO 4 on Raspberry Pi or pin 2 on Arduino).
# D. Python Program to Toggle Two LEDs

# toggle_led.py:


import RPi.GPIO as GPIO
import time

# Set up GPIO pins
LED1 = 17  # GPIO pin for LED 1
LED2 = 27  # GPIO pin for LED 2

GPIO.setmode(GPIO.BCM)
GPIO.setup(LED1, GPIO.OUT)
GPIO.setup(LED2, GPIO.OUT)

try:
    while True:
        # Turn on LED1 and off LED2
        GPIO.output(LED1, GPIO.HIGH)
        GPIO.output(LED2, GPIO.LOW)
        time.sleep(1)  # Wait for 1 second

        # Turn off LED1 and on LED2
        GPIO.output(LED1, GPIO.LOW)
        GPIO.output(LED2, GPIO.HIGH)
        time.sleep(1)  # Wait for 1 second

except KeyboardInterrupt:
    pass  # Exit on Ctrl+C

finally:
    GPIO.cleanup()  # Clean up GPIO settings

# 	Observations on Input and Output
# Input: The program receives no external input; it runs indefinitely until interrupted.
# Output: The two LEDs toggle on and off every second, creating a blinking effect.
# Result and Conclusion
# Result: The program successfully toggles two LEDs connected to the Raspberry Pi/Arduino.
# Conclusion: This demonstrates basic GPIO control using Python, allowing for simple output devices like LEDs to be controlled programmatically.