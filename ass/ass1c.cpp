/* 1. Arduino and IR Sensor Interfacing:
Components Needed:

Arduino Uno

IR Sensor (e.g., KY-022 IR Sensor)

Jumper wires

Breadboard

Connections:

Connect VCC of IR sensor to 5V on Arduino

Connect GND to Ground

Connect Signal Pin to Digital Pin 2 (or any other pin) on Arduino

Arduino Code (C++) to Read IR Sensor:
 */

int irSensorPin = 2; // IR sensor connected to digital pin 2
int irStatus = 0; // Variable to store IR sensor status

void setup() {
  pinMode(irSensorPin, INPUT); // Set the IR sensor pin as input
  Serial.begin(9600); // Initialize serial communication
}

void loop() {
  irStatus = digitalRead(irSensorPin); // Read the IR sensor value
  if (irStatus == HIGH) {
    Serial.println("IR sensor is triggered!");
  } else {
    Serial.println("IR sensor is not triggered.");
  }
  delay(500); // Delay for half a second
}
/* 
Observations:
The LED will blink on and off every second.

The program sends a high signal to pin 13, causing the LED to light up, and a low signal to turn it off.

The delay of 1000ms (1 second) controls the blink rate.

Result:
The LED blinks on and off with a 1-second interval.

The code successfully toggles the LED on the specified pin.

Conclusion:
The basic concept of controlling output pins on Arduino is demonstrated.

The program toggles the LED on and off using digitalWrite() and controls the delay using delay().

This is a simple but effective way to familiarize with Arduino programming and GPIO control.
 */
