/*
Block Diagram of Arduino Uno Board Interfacing with IR Sensor
Below is a simple block diagram showing how an IR sensor is interfaced with an Arduino Uno board:

lua
Copy
       +-------------------------+
       |                         |
       |        IR SENSOR         |
       |   (Signal Pin)   -----> | Pin 2 (Digital Pin)  |
       |   (VCC)          -----> | 5V (Power)           |
       |   (GND)          -----> | GND                   |
       |                         |
       +-------------------------+
                |
                |
                v
       +-------------------------+
       |                         |
       |      ARDUINO UNO        |
       |                         |
       |    - Pin 2  -----> Input|
       |    - Pin 13 ---> Output |
       |    - 5V ------------> VCC|
       |    - GND -----------> GND|
       |                         |
       +-------------------------+
                |
                v
     +-----------------------+
     |                       |
     |     BUZZER/LED        |
     |   (Output Device)     |
     +-----------------------+
Explanation of the Block Diagram:
IR Sensor:

Signal Pin: This pin is connected to Pin 2 (Digital Pin) on the Arduino board to receive the signal data.

VCC Pin: This is connected to the 5V pin on the Arduino to power the IR sensor.

GND Pin: This is connected to the GND pin on the Arduino to complete the circuit.

Arduino Uno:

Pin 2 (Input): Reads data from the IR sensor. This pin will be triggered when the IR sensor detects a certain object or condition.

Pin 13 (Output): This pin can control an LED or a Buzzer based on the data received from the IR sensor.

The Arduino board provides 5V and GND to power the IR sensor and other connected components.

C++ Program to Blink an LED on Arduino Uno
Now, let's write a C++ program for Arduino Uno to blink an LED connected to pin 13.
*/

// Define the LED pin
int ledPin = 13; // LED connected to digital pin 13

void setup() {
  // Initialize the LED pin as an output
  pinMode(ledPin, OUTPUT);
}

void loop() {
  // Turn the LED on (HIGH is the voltage level)
  digitalWrite(ledPin, HIGH);  
  delay(1000);  // Wait for 1 second

  // Turn the LED off by making the voltage LOW
  digitalWrite(ledPin, LOW);   
  delay(1000);  // Wait for 1 second
}

/**
Explanation of the Program:
Pin Configuration (pinMode):

The pinMode(ledPin, OUTPUT) function configures pin 13 as an output pin because an LED is connected to it.

LED Blinking Logic (digitalWrite and delay):

The digitalWrite(ledPin, HIGH) function turns the LED on by supplying HIGH voltage.

The digitalWrite(ledPin, LOW) function turns the LED off by supplying LOW voltage.

The delay(1000) function causes a 1000-millisecond (1 second) delay, ensuring the LED stays on for 1 second, and then off for 1 second.

Loop:

The loop() function repeats the above steps indefinitely, making the LED blink on and off.

Observations on Input and Output:
Input:

The program does not have user input; it just runs in a loop to blink the LED.

If an IR sensor were added to the system, it would trigger the LED blink based on the sensor input.

Output:

The output is the blinking LED on pin 13 of the Arduino board.

The LED will turn on for 1 second and then turn off for 1 second. This cycle will repeat indefinitely.

Sample Output (Visual):
The LED will blink every second, i.e., it will turn on for 1 second, then off for 1 second, and the process will repeat.

Result and Conclusion:
Result:

The program successfully blinks the LED connected to pin 13 of the Arduino Uno board. The LED turns on for 1 second and off for 1 second, repeating this cycle indefinitely.

Conclusion:

This program demonstrates basic digital I/O functionality on the Arduino. The program uses the digitalWrite function to control the LED, and the delay function to control the timing. This exercise serves as an introduction to programming and controlling hardware devices like LEDs with Arduino.

**/