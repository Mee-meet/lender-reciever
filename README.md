# lender-reciever
.
Assignment
Frontend - 
1. React Native
2. 4 fields to take input from user
   2.1 Lender Name
   2.2 Receiver Name
   2.3 Date
   2.5 Total Amount
3. Submit Button that will call the POST call to backend
   3.1 POST /transactions
       Body - {
		"lender_name": "mehul"
		"receiver_name": "gaurav"
		"amount": 100
		"date": "19/06/2023"
      }  


Backend - 
1. Backend will expose the above post api (R/D topic - try to figure out a way how we can expose an api from golang, there are multiple ways)
2. Read the content of the body
3. Save the content in the mysql database table
4. Backend will connect with DB (R/D topic - try to figure out a way how a golang backend connect with mysql)


Database - 
1. mysql database
2. Table Transaction Content
   2.1 Primary Key
   2.2 LenderName
   2.3 ReceiverName
   2.4 Amount
   2.5 Date
