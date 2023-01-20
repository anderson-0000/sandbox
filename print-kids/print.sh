#!/bin/bash

lpr -P $(lpstat -d | cut -d ' ' -f 2) nazorigaki/nazorigaki-*
