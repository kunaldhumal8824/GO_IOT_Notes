/* 
1. Block Diagram of Arduino Uno Board Interfacing with IR Sensor:
The block diagram below illustrates how an IR sensor can be connected to an Arduino Uno board for a basic application.

lua
Copy
       +-------------------------+
       |                         |
       |       IR SENSOR          |
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
       |      ARDUINO UNO         |
       |                         |
       |    - Pin 2  -----> Input |
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

Signal Pin: Connected to Pin 2 (Digital Pin) on the Arduino to read the signal from the IR sensor.

VCC Pin: Connected to the 5V pin on the Arduino to power the IR sensor.

GND Pin: Connected to GND on the Arduino to complete the circuit.

Arduino Uno:

The Arduino board will process the input signal from the IR sensor.

The Signal Pin (Pin 2) reads the input from the IR sensor.

The Output Pin (Pin 13) controls an external output device (LED or Buzzer) based on the signal received from the IR sensor.

Output Device (LED or Buzzer):

The output device (LED or Buzzer) is connected to Pin 13 on the Arduino and is controlled based on the signal received from the IR sensor. For instance, the Arduino might turn on the LED or buzzer when the IR sensor detects an object.

2. C++ Program to Blink an LED on Arduino Uno:
Below is a C++ program for Arduino to blink an LED connected to Pin 13. */


// Define the LED Pin
int ledPin = 13;  // LED connected to digital pin 13

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
/* Explanation of the Program:
Pin Configuration (pinMode):

In the setup() function, the LED pin (Pin 13) is initialized as an OUTPUT pin using the pinMode() function.

LED Blinking Logic (digitalWrite):

In the loop() function, digitalWrite() is used to turn the LED on by setting the pin to HIGH (on).

After that, the delay(1000) function waits for 1000 milliseconds (1 second) before turning the LED off by setting the pin to LOW (off).

Another delay(1000) is used to wait for 1 second before repeating the cycle.

Continuous Loop:

The loop() function runs continuously, making the LED blink on and off every second.

3. Observations on Input and Output:
Input:

The input here is simply the code running on the Arduino, and there are no external input devices like buttons or sensors in this program. The program is simply controlling the output device (LED) using Arduino.

Output:

The output is the LED blinking on Pin 13. The LED turns on for 1 second, then turns off for 1 second, and this process repeats indefinitely.

4. Sample Output (Visual):
The LED will blink every second. That is:

LED turns ON for 1 second and then turns OFF for 1 second.

This process repeats continuously as long as the Arduino is powered.

5. Result and Conclusion:
Result:

The program successfully blinks an LED connected to Pin 13 of the Arduino Uno. The LED turns on for 1 second, then turns off for 1 second, and this cycle continues indefinitely.

Conclusion:

The program demonstrates how to control an LED using an Arduino Uno board.

This example introduces the concept of digital output control on the Arduino platform.

The digitalWrite() function is used to control the output devices like LEDs and buzzers.

This can be extended to include more sensors and output devices based on the requirements of the project, such as interfacing with IR sensors, ultrasonic sensors, etc.
 */
