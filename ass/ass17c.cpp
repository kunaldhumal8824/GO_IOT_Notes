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
// LEDs: Connected to digital pins on the Arduino.
// Positive Pin: Connect to digital pins (e.g., pin 8 and pin 9).
// Negative Pin: Connect to the GND pin on Arduino.


### D) C++ Program to Toggle Two LEDs

**toggle_leds.ino:**

```cpp
const int led1Pin = 8; // Pin for LED 1
const int led2Pin = 9; // Pin for LED 2

void setup() {
  pinMode(led1Pin, OUTPUT); // Set LED1 pin as an output
  pinMode(led2Pin, OUTPUT); // Set LED2 pin as an output
}

void loop() {
  // Turn on LED1 and off LED2
  digitalWrite(led1Pin, HIGH);
  digitalWrite(led2Pin, LOW);
  delay(1000); // Wait for 1 second

  // Turn off LED1 and on LED2
  digitalWrite(led1Pin, LOW);
  digitalWrite(led2Pin, HIGH);
  delay(1000); // Wait for 1 second
}


// Observations on Input and Output
// Input: The Go programs do not require any external input for the file appending program, while the first program takes two numbers as input.
// Output:
// The first Go program prints the sum, difference, product, and quotient of the two numbers.
// The second Go program appends a line of text to a file and confirms the action.
// The Arduino program toggles two LEDs on and off every second.
// Result and Conclusion
// Result:
// The Go program successfully calculates and displays the arithmetic operations.
// The file appending program successfully adds content to the specified file.
// The Arduino program successfully toggles two LEDs.
// Conclusion: These examples demonstrate basic