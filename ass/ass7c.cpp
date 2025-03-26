/*
1. Block Diagram of Arduino Uno Board Interfacing with IR Sensor
Here's the block diagram of the Arduino Uno board interfacing with an IR sensor.

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

Signal Pin: Connected to Pin 2 on the Arduino Uno. This pin reads the output signal of the IR sensor.

VCC Pin: Connected to 5V on the Arduino Uno to supply power to the IR sensor.

GND Pin: Connected to GND on the Arduino Uno to complete the circuit.

Arduino Uno:

Pin 2 (Input): The Arduino reads the input signal from the IR sensor.

Pin 13 (Output): The Arduino can control an output device such as a buzzer or LED based on the signal received from the IR sensor.

Output Device (LED or Buzzer):

The Arduino uses this output device to show feedback based on the input from the IR sensor (e.g., turning on a buzzer or LED when an object is detected).

2. C++ Program to Turn ON/OFF a Buzzer
Below is a simple C++ program for Arduino Uno that turns a buzzer ON and OFF with a delay.

*/
// Define the Buzzer pin
int buzzerPin = 8; // Buzzer connected to digital pin 8

void setup() {
  // Initialize the Buzzer pin as output
  pinMode(buzzerPin, OUTPUT);
}

void loop() {
  // Turn the Buzzer ON
  digitalWrite(buzzerPin, HIGH);
  delay(1000);  // Wait for 1 second

  // Turn the Buzzer OFF
  digitalWrite(buzzerPin, LOW);
  delay(1000);  // Wait for 1 second
}
/*
Explanation of the Program:
Pin Configuration (pinMode):

We initialize the buzzerPin (connected to digital pin 8) as an OUTPUT pin using pinMode(buzzerPin, OUTPUT).

Turning ON/OFF the Buzzer:

In the loop() function, we use digitalWrite(buzzerPin, HIGH) to turn ON the buzzer.

We wait for 1 second using delay(1000) (1000 milliseconds = 1 second).

Then we use digitalWrite(buzzerPin, LOW) to turn OFF the buzzer and wait for 1 second again using delay(1000).

Repeating the Loop:

The loop() function runs continuously, making the buzzer turn ON and OFF every second.

3. Observations on Input and Output:
Input:

There is no external input device in this program (no IR sensor or button). The program simply toggles the buzzer every second.

If you integrate this with an IR sensor, the buzzer can be turned on or off based on the signal received from the sensor.

Output:

The output of the program is the buzzer turning ON for 1 second and OFF for 1 second in a continuous loop.

4. Sample Output (Visual):
The buzzer will emit sound for 1 second, then stop for 1 second, and this process will repeat indefinitely.

5. Result and Conclusion:
Result:

The program successfully turns the buzzer ON and OFF with a 1-second delay.

This behavior is controlled by the Arduino Uno and can be modified to integrate more sensors or different devices.

Conclusion:

The program demonstrates basic digital output control using an Arduino board.

By using the digitalWrite() function, we can control devices like buzzers, LEDs, and motors connected to the Arduino.

This concept can be extended to include more complex logic and additional input devices (e.g., IR sensors or buttons) to control the output more interactively.

*/