// # A. Block Diagram for Arduino Uno with IR Sensor/Temperature Sensor
// Block Diagram Description
// Arduino Uno: The main microcontroller board.
// IR Sensor / Temperature Sensor: The sensor connected to the Arduino.
// For an IR Sensor (e.g., HC-SR501):
// VCC: Connect to the 5V pin on Arduino.
// GND: Connect to the GND pin on Arduino.
// OUT: Connect to a digital pin (e.g., pin 2).
// For a Temperature Sensor (e.g., DHT11):
// VCC: Connect to the 5V pin on Arduino.
// GND: Connect to the GND pin on Arduino.
// DATA: Connect to a digital pin (e.g., pin 2).
// Buzzer: Connected to another digital pin on the Arduino.
// Positive Pin: Connect to a digital pin (e.g., pin 8).
// Negative Pin: Connect to the GND pin on Arduino.


### B. C++ Program to Turn ON/OFF a Buzzer

**buzzer.ino:**


const int buzzerPin = 8; // Pin for the buzzer
const int sensorPin = 2; // Pin for the IR/Temperature sensor

void setup() {
  pinMode(buzzerPin, OUTPUT); // Set the buzzer pin as an output
  pinMode(sensorPin, INPUT);   // Set the sensor pin as an input
  Serial.begin(9600);          // Initialize serial communication
}

void loop() {
  int sensorValue = digitalRead(sensorPin); // Read the sensor value

  if (sensorValue == HIGH) { // If the sensor detects something
    digitalWrite(buzzerPin, HIGH); // Turn ON the buzzer
    Serial.println("Buzzer ON");
  } else {
    digitalWrite(buzzerPin, LOW); // Turn OFF the buzzer
    Serial.println("Buzzer OFF");
  }

  delay(1000); // Wait for 1 second before the next reading
}

// Observations on Input and Output
// Input: The program reads the state of the IR sensor or temperature sensor connected to the specified pin.
// Output:
// When the sensor detects an object (or a certain temperature), the buzzer turns ON, and the message "Buzzer ON" is printed to the Serial Monitor.
// When the sensor does not detect anything, the buzzer turns OFF, and the message "Buzzer OFF" is printed.
// Result and Conclusion
// Result: The program successfully turns the buzzer ON when the sensor detects an object (or a certain temperature) and
//  turns it OFF when it does not.
// Conclusion: This demonstrates basic input/output control using Arduino. 
// The integration of sensors and actuators (like buzzers) is fundamental in embedded systems and IoT applications. 
// This setup can be used in various applications, such as security systems or temperature monitoring systems.