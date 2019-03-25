#!/bin/bash

sudo systemctl enable NetworkManager
sudo systemctl enable networking
sudo systemctl start NetworkManager
sudo systemctl start networking
