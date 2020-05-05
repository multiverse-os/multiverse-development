#!/bin/bash

sudo systemctl stop NetworkManager
sudo systemctl disable NetworkManager

sudo systemctl stop networking
sudo systemctl disable networking
