/*
1. Block Diagram of Arduino Uno Board Interfacing with IR Sensor:
Below is a block diagram illustrating how an IR sensor is interfaced with an Arduino Uno board. This setup can be used to detect obstacles or objects using the IR sensor and respond by turning on/off an LED or a buzzer based on the sensor's detection.

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

Signal Pin: Connected to Pin 2 (Digital Pin) on the Arduino, used to receive input data from the IR sensor (whether an object is detected or not).

VCC Pin: Connected to the 5V pin on the Arduino to power the IR sensor.

GND Pin: Connected to GND on the Arduino to complete the circuit.

Arduino Uno:

The Arduino Uno processes the input from the IR sensor and can control an output device (e.g., LED or Buzzer) based on the data from the sensor.

Pin 2 reads the input from the IR sensor.

Pin 13 controls the output device (LED or Buzzer).

Output Device (LED/Buzzer):

An LED or buzzer is connected to Pin 13 on the Arduino and is controlled based on the sensor's signal. For example, the LED might turn on when the IR sensor detects an object.

2. C++ Program to Blink an LED on Arduino Uno:
Here is a C++ program that blinks an LED connected to Pin 13 on the Arduino Uno board.

*/
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
/*
Explanation of the Program:
Pin Configuration (pinMode):

In the setup() function, the LED Pin (Pin 13) is initialized as an OUTPUT pin using pinMode(). This tells Arduino that Pin 13 will control an output device (LED in this case).

LED Blinking Logic (digitalWrite):

In the loop() function:

The digitalWrite(ledPin, HIGH) command turns the LED ON by applying a HIGH voltage (5V).

The delay(1000) function pauses the program for 1000 milliseconds (1 second), leaving the LED on for 1 second.

The digitalWrite(ledPin, LOW) command turns the LED OFF by applying a LOW voltage (0V).

Another delay(1000) pauses the program for another 1 second, leaving the LED off for 1 second.

Continuous Loop:

The loop() function runs indefinitely, causing the LED to blink on and off every second.

3. Observations on Input and Output:
Input:

There is no external input like buttons or sensors in this simple example. The input here is the onboard program running on the Arduino, which is responsible for controlling the LED.

Output:

The output is the LED blinking on Pin 13. The LED will turn ON for 1 second and turn OFF for 1 second repeatedly in a continuous loop.

4. Sample Output (Visual):
The LED will blink as follows:

ON for 1 second, then OFF for 1 second, and this cycle continues indefinitely.

5. Result and Conclusion:
Result:

The C++ program successfully blinks an LED connected to Pin 13 of the Arduino Uno. The LED turns ON and OFF in a repeating cycle every 1 second.

Conclusion:

This simple program demonstrates basic digital output control on the Arduino platform.

The digitalWrite() function is used to control output devices such as LEDs and buzzers.

The delay() function is used to create pauses between actions, such as turning the LED on and off.

This program can be extended to control more complex output devices or interface with sensors to create interactive applications.

*/