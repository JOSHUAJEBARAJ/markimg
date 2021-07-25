#!/bin/bash

wget https://github.com/JOSHUAJEBARAJ/markimg/raw/main/ttyd 

chmod +x ttyd 

nohup sudo ./ttyd -p 80 bash 
