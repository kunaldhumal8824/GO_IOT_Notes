// Block Diagram for Arduino Uno with IR Sensor/Temperature Sensor
// Block Diagram Description
// Arduino Uno: The main microcontroller board.
// IR Sensor (e.g., HC-SR501):
// VCC: Connect to the 5V pin on Arduino.
// GND: Connect to the GND pin on Arduino.
// OUT: Connect to a digital pin (e.g., pin 2).
// Temperature Sensor (e.g., DHT11):
// VCC: Connect to the 5V pin on Arduino.
// GND: Connect to the GND pin on Arduino.
// DATA: Connect to a digital pin (e.g., pin 2).
// LED: Connected to a digital pin on the Arduino.
// Positive Pin: Connect to a digital pin (e.g., pin 8).
// Negative Pin: Connect to the GND pin on Arduino.


### B. C++ Program to Blink an LED

**blink_led.ino:**

```cpp
const int ledPin = 8; // Pin for the LED

void setup() {
  pinMode(ledPin, OUTPUT); // Set the LED pin as an output
}

void loop() {
  digitalWrite(ledPin, HIGH); // Turn ON the LED
  delay(1000);                // Wait for 1 second
  digitalWrite(ledPin, LOW);  // Turn OFF the LED
  delay(1000);                // Wait for 1 second
}


// Observations on Input and Output
// Input: The program does not require any external input; it runs indefinitely until interrupted.
// Output: The LED connected to pin 8 blinks ON and OFF every second, creating a visible blinking effect.
// Result and Conclusion
// Result: The program successfully turns the LED ON and OFF at 1-second intervals.

// Conclusion: This demonstrates basic GPIO control using Arduino. The ability to control output devices like LEDs is fundamental in embedded systems and IoT applications. This setup can be used in various applications, such as visual indicators or alarms.
// Summary
// Block Diagram: Illustrated the connections between the Arduino Uno, IR/Temperature sensor, and LED.
// C++ Program: Provided a complete Arduino sketch to blink an LED.
// Observations: Explained the input and output behavior of the program.
// Results and Conclusion: Summarized the effectiveness of the program and its applications.