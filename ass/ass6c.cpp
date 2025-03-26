/*
Block Diagram of Arduino Uno Board Interfacing with IR Sensor
Below is the Block Diagram that shows the interface between an Arduino Uno Board and an IR Sensor:

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

Signal Pin: The signal from the IR sensor is connected to Pin 2 on the Arduino Uno. This pin will receive data or signal from the IR sensor (detecting an object, for example).

VCC Pin: Provides power to the IR sensor, connected to the 5V pin of the Arduino Uno.

GND Pin: The ground pin of the IR sensor is connected to the GND pin of the Arduino.

Arduino Uno:

Pin 2 (Input): The Arduino reads the signal from the IR sensor here.

Pin 13 (Output): This pin can be used to control an output device, like an LED or Buzzer, based on the signal received from the IR sensor.

Output Device (LED or Buzzer):

The Arduino can control this device based on the input from the IR sensor, such as turning on/off an LED or Buzzer.

2. C++ Program to Toggle Two LEDs
Below is the C++ program for Arduino Uno to toggle two LEDs connected to Pin 13 and Pin 12.
*/

// Define LED pins
int ledPin1 = 13; // LED connected to pin 13
int ledPin2 = 12; // LED connected to pin 12

void setup() {
  // Initialize the LED pins as outputs
  pinMode(ledPin1, OUTPUT);
  pinMode(ledPin2, OUTPUT);
}

void loop() {
  // Toggle LED 1 (Turn on)
  digitalWrite(ledPin1, HIGH);  
  // Toggle LED 2 (Turn off)
  digitalWrite(ledPin2, LOW);   
  delay(1000);  // Wait for 1 second

  // Toggle LED 1 (Turn off)
  digitalWrite(ledPin1, LOW);   
  // Toggle LED 2 (Turn on)
  digitalWrite(ledPin2, HIGH);  
  delay(1000);  // Wait for 1 second
}

/*
Explanation of the Program:
Pin Configuration (pinMode):

pinMode(ledPin1, OUTPUT) and pinMode(ledPin2, OUTPUT) are used to set Pin 13 and Pin 12 as output pins. These pins are connected to the LEDs.

LED Toggling Logic (digitalWrite and delay):

digitalWrite(ledPin1, HIGH) turns on the LED connected to Pin 13, and digitalWrite(ledPin2, LOW) turns off the LED connected to Pin 12.

digitalWrite(ledPin1, LOW) turns off the LED connected to Pin 13, and digitalWrite(ledPin2, HIGH) turns on the LED connected to Pin 12.

delay(1000) introduces a 1-second delay between toggling the LEDs on and off.

Loop (loop function):

The loop continuously toggles the LEDs every second, turning one LED on and the other off.

3. Observations on Input and Output:
Input:

There is no external input in this case. The program simply toggles the LEDs based on the code logic.

However, if an IR sensor or other input devices were added, the LEDs could toggle based on the sensor input (like detecting an object).

Output:

The output is the toggling of two LEDs connected to Pin 13 and Pin 12 of the Arduino Uno.

LED 1 turns on while LED 2 turns off for 1 second, and then they toggle every second.

4. Sample Output (Visual):
LED 1 (Pin 13): Turns on for 1 second, then off for 1 second.

LED 2 (Pin 12): Turns off for 1 second, then on for 1 second.

This alternating pattern continues indefinitely.

5. Result and Conclusion:
Result:

The program successfully toggles two LEDs connected to Pin 13 and Pin 12 on the Arduino Uno.

Each LED turns on for 1 second and off for 1 second, creating a toggle effect.

Conclusion:

This program demonstrates digital I/O control on the Arduino platform.

It introduces how to control multiple outputs (LEDs) using digital pins.

The program can be expanded to include more complex interactions, like controlling LEDs based on sensor input (e.g., using an IR sensor).
*/