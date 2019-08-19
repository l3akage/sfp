# SFP Shield for Raspberry Pi

With this shield you can read SFP informations and firmware through i2c.

It has LEDs for TX_FAULT and RX_LOS which are also readable via GPIO

# Pins
GPIO     | Description
---------|-------------
4 | RATE_SELECT
17 | TX_DISABLE (set to 0 to turn on laser)
22 | RX_LOS
27 | TX_FAULT
