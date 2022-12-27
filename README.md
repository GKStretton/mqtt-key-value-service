# MQTT key-value Service

mqtt-key-value-service is a simple, lightweight, key-value store with an MQTT interface. This combines the simplicity and usefulness of a key-value store with the interoperability of MQTT.

---

## Usage

| Topic                   | Payload   | Description                                                                                                 |
|-------------------------|-----------|-------------------------------------------------------------------------------------------------------------|
| **root**/set/**key**    | **value** | Publish to this to set **key** equal to **value**                                                           |
| **root**/setack/**key** | -         | Empty message sent to this topic when **key** is set                                                        |
| **root**/sub/**key**    | **value** | Subscribe to this to receive **value** for **key** whenever the value is set or req/**key** is published to |
| **root**/req/**key**    | -         | Publish here to request the value for key to be published to the `sub/` topic                               |

---

## Configuration

| Environment Variable | Usage                                              |
|----------------------|----------------------------------------------------|
| KV_ROOT              | **root**, as seen in topics above                  |
| KV_STORAGE_DIR       | Absolute path where the key-value pairs are stored |