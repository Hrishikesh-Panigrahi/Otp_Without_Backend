# Otp_Without_Backend
This project implements a One-Time Password (OTP) system without a backend server, utilizing hashing and local storage to securely generate and verify OTPs. The project is built using Golang and HTML/Tailwind CSS.

## Features
- Generate OTPs and send them to users via email
- Verify OTPs without server communication using hashing
- Securely store OTPs in local storage
- Limit the number of OTP attempts a user can make to prevent brute-force attacks.
- Clean up the local storage for memory management

## Technologies used
1. Golang and GIN framework
2. Html/Tailwind CSS

## Usage
1. Enter your email address.
2. Click on the "Generate OTP" button.
3. An OTP will be sent on the entered Email adress.
4. Enter the OTP in the verification field and click "Verify".

## How to Run the Project
1. open the terminal and run the following command
```bash
go run main.go
```