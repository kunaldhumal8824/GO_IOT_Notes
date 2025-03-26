/*

Block Diagram / Pin Diagram of Arduino Uno Board Interfacing with IR Sensor/Temperature Sensor
Block Diagram:
Below is a simplified block diagram that represents how an IR sensor (or Temperature Sensor) is interfaced with an Arduino Uno Board.


  +-------------------------+
  |                         |
  |   IR SENSOR             |      +-----------------+
  |   (Signal Pin)   ----->| Pin 2 |                 |
  |   (VCC)          ----->| 5V   |                 |
  |   (GND)          ----->| GND  |    ARDUINO UNO  |
  |                         |      |                 |
  +-------------------------+      |    (Pin 13) --->| LED 1 (Output)   |
                                   |    (Pin 12) --->| LED 2 (Output)   |
                                   |    (Pin 2) ---->| IR sensor input |
                                   |                 |
                                   +-----------------+
            |
            |
            v
     +----------------------+
     |                      |
     |     BUZZER/LED       |
     |    (Output Device)   |
     +----------------------+
Explanation of the Block Diagram:
IR Sensor:

Signal Pin: Connected to the digital pin 2 of the Arduino to receive data (input).

VCC: Connected to 5V of Arduino for powering the sensor.

GND: Connected to GND of Arduino.

Arduino Uno:

Pin 2: Receives the input data from the IR sensor.

Pin 13 and Pin 12: Control the output devices like LED1 and LED2. When the sensor detects an object or a specific condition, the LEDs are toggled.

The BUZZER/LED connected as an output device, responds to the input data from the sensor.

Program to Toggle Two LEDs in Python/C++ Language (Arduino Uno)
Below is an Arduino program written in C++ that toggles two LEDs connected to pins 13 and 12.

C++ Code to Toggle Two LEDs:
**/
// Define LED Pins
int ledPin1 = 13;  // LED1 connected to digital pin 13
int ledPin2 = 12;  // LED2 connected to digital pin 12

void setup() {
  // Initialize the LED pins as outputs
  pinMode(ledPin1, OUTPUT);
  pinMode(ledPin2, OUTPUT);
}

void loop() {
  // Toggle the first LED on
  digitalWrite(ledPin1, HIGH);  
  delay(500);  // Wait for 500 milliseconds
  
  // Toggle the first LED off
  digitalWrite(ledPin1, LOW);  
  delay(500);  // Wait for 500 milliseconds
  
  // Toggle the second LED on
  digitalWrite(ledPin2, HIGH);
  delay(500);  // Wait for 500 milliseconds
  
  // Toggle the second LED off
  digitalWrite(ledPin2, LOW);
  delay(500);  // Wait for 500 milliseconds
}
/** 
Explanation of the Program:
Pin Configuration (pinMode):

The pinMode() function is used to set Pin 13 and Pin 12 as output pins because the LEDs are connected to them.

LED Toggling Logic (digitalWrite and delay):

The digitalWrite() function is used to turn the LEDs on (HIGH) and off (LOW).

delay(500) introduces a 500-millisecond delay between toggling the LEDs, making them blink every second.

Loop:

The loop() function continuously toggles the two LEDs, one after the other, creating a blinking effect.

Sample Output (Visual):
LED1 on Pin 13 and LED2 on Pin 12 will blink alternatively. The program will turn LED1 on for 500 ms, then turn it off and LED2 on for 500 ms, and so on.

Observations on Input and Output:
Input:

The program does not take any user input in this case (it's a basic toggle LED program).

The behavior of the LEDs is controlled by the program, which runs in a continuous loop.

Output:

The two LEDs connected to Pins 13 and 12 will alternate between on and off every 500 milliseconds.

As a result, you will observe LED1 and LED2 blinking alternately.

Result and Conclusion:
Result:

The program successfully toggles two LEDs connected to an Arduino Uno board. The LEDs blink alternately with a 500-millisecond delay between each toggle.

Conclusion:

This simple program demonstrates the basics of digital I/O and how to control output devices (LEDs) using an Arduino Uno. It introduces fundamental concepts of programming such as loops, delays, and digital outputs.

The use of digitalWrite() and delay() functions makes it easy to control timing and state transitions, which is essential for creating interactive systems with Arduino.

**/