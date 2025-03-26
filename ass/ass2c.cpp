/* Arduino Uno Board Interfacing with IR Sensor (Temperature Sensor or Camera):
For this task, we will assume Arduino Uno interfacing with an IR sensor and how to control a buzzer based on the sensor input.

Hardware Components:
Arduino Uno board

IR Sensor (KY-022 or similar)

Buzzer

Jumper wires

Breadboard

Connections:
IR Sensor:

VCC → 5V (on Arduino)

GND → GND (on Arduino)

Signal Pin → Digital Pin 2 (on Arduino)

Buzzer:

One terminal of Buzzer → Digital Pin 13 (on Arduino)

Other terminal → GND (on Arduino)

Arduino Code to Turn ON/OFF Buzzer Based on IR Sensor: */


int irSensorPin = 2;    // IR Sensor input pin (Digital Pin 2)
int buzzerPin = 13;     // Buzzer output pin (Digital Pin 13)
int irStatus = 0;       // Variable to store IR sensor status

void setup() {
  pinMode(irSensorPin, INPUT);  // Set IR sensor pin as input
  pinMode(buzzerPin, OUTPUT);   // Set buzzer pin as output
  Serial.begin(9600);           // Initialize serial communication
}

void loop() {
  irStatus = digitalRead(irSensorPin);  // Read the value from the IR sensor

  if (irStatus == HIGH) {   // If the sensor detects an object (HIGH)
    digitalWrite(buzzerPin, HIGH);  // Turn ON the buzzer
    Serial.println("Object detected! Buzzer ON.");
  } else {
    digitalWrite(buzzerPin, LOW);   // Turn OFF the buzzer
    Serial.println("No object detected! Buzzer OFF.");
  }

  delay(500);   // Delay for half a second before checking again
}
/* Explanation of the Code:
Pin Configuration:

irSensorPin: This is the pin connected to the signal output of the IR sensor.

buzzerPin: This is the pin controlling the buzzer.

Reading IR Sensor Input:

The digitalRead() function reads the signal from the IR sensor.

If the sensor detects an object, it sends a HIGH signal; otherwise, it sends a LOW signal.

Buzzer Control:

If the IR sensor detects an object, the buzzer is turned on (digitalWrite(buzzerPin, HIGH)).

If no object is detected, the buzzer is turned off (digitalWrite(buzzerPin, LOW)).

Serial Output: The program also prints whether the object was detected or not on the serial monitor for debugging.

Observations on Input and Output:
Input: The input is received from the IR sensor. It will send a HIGH signal when it detects an object and a LOW signal when no object is present.

Output: The output is sent to the buzzer, which will be turned ON when an object is detected (IR sensor gives HIGH) and OFF when no object is detected (IR sensor gives LOW).

Sample Output on Serial Monitor:
pgsql
Copy
Object detected! Buzzer ON.
No object detected! Buzzer OFF.
Object detected! Buzzer ON.
No object detected! Buzzer OFF.
Result:
The IR sensor successfully detects objects and triggers the buzzer accordingly.

The program works as expected by turning the buzzer on when the sensor detects an object and turning it off when there is no detection.

Conclusion:
This task demonstrates the basic concept of Arduino interfacing with sensors and controlling output devices such as a buzzer.

The use of digitalRead() and digitalWrite() in Arduino allows easy control of input and output devices.

This simple setup can be expanded for other applications such as security systems, where a buzzer could be triggered based on the presence of an object. */