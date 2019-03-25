#!/bin/bash

sudo systemctl disable NetworkManager
sudo systemctl disable networking
sudo systemctl stop NetworkManager
sudo systemctl stop networking
