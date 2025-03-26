/*
# Arduino Uno Board Interfacing with IR Sensor and LED Blink Program

## 1. Block Diagram of Arduino Uno Board Interfacing with IR Sensor

### Block Diagram Description:
In this project, we are interfacing an **IR sensor** with an **Arduino Uno board**. The block diagram below illustrates how the IR sensor is connected to the Arduino Uno board.

### Block Diagram:

+-------------------------+ | | | IR SENSOR | | (Signal Pin) ----->| Pin 2 (Digital Pin) | | (VCC) ----->| 5V (Power) | | (GND) ----->| GND | | | +-------------------------+ | | v +-------------------------+ | | | ARDUINO UNO | | | | - Pin 2 -----> Input| | - Pin 13 ---> Output | | - 5V ------------> VCC| | - GND -----------> GND| | | +-------------------------+ | v +----------------------+ | | | BUZZER/LED | | (Output Device) | +----------------------+

markdown
Copy

### Explanation:

- **IR Sensor:**
  - The **IR sensor** is connected to the Arduino. The sensor has three pins:
    - **Signal Pin**: This pin is connected to **digital pin 2** of the Arduino to receive the sensor's data.
    - **VCC**: This pin is connected to the **5V pin** on the Arduino to provide power to the sensor.
    - **GND**: This pin is connected to the **GND pin** on the Arduino to complete the circuit.

- **Arduino Uno:**
  - The **Arduino Uno** is used to process the input from the IR sensor and can control an output device like a **buzzer** or **LED** based on the IR sensor data.
  - The signal from the IR sensor is read by **Pin 2 (Digital Pin)** of Arduino.
  - The output device (such as an **LED** or **Buzzer**) is connected to **Pin 13 (Digital Pin)**.

---

## 2. WAP in C/C++ Language to Blink LED (Arduino Uno)

### Below is an Arduino program written in **C++** to blink an LED connected to **Pin 13** of the Arduino Uno board:

```cpp
*/
// Define LED Pin
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

The pinMode() function is used in the setup() function to configure Pin 13 as an output pin because the LED is connected to it.

LED Blinking Logic:

In the loop() function, the digitalWrite() function is used to set the LED pin to HIGH (on) and LOW (off).

delay(1000) makes the program wait for 1000 milliseconds (1 second) between turning the LED on and off, resulting in a blink effect.

3. Observations on Input and Output:
Input:
The input is not explicitly used in this case (no sensors or external inputs). The program simply uses the onboard LED connected to Pin 13 of the Arduino.

Output:
The output is the LED blinking on Pin 13. It will blink on for 1 second and off for 1 second in a continuous loop.

*/