# transact-mail

Project in charge of reading the .csv file with user transactions; send the transactions by email to the desired user and save the transactions in postgresSQL

### Setup
1. Build image

   docker pull postgres
   docker build -t transact-mail .

2. Run image

   docker run -d --name postgres-transactions -e POSTGRES_PASSWORD=pass\
   docker run -d --name transact-mail\
   -e EMAIL_HOST=host.email.com
   -e EMAIL_PORT=port
   -e EMAIL_SENDER_USER=sender@mail.com
   -e EMAIL_SENDER_PASSWORD=examplepasswordapp
   -e EMAIL_TO=receiver@mail.com
   -e EMAIL_FROM=from
   -e POSTGRES_USER=user
   -e POSTGRES_PASSWORD=pass
   -e POSTGRES_HOST=host
   -e POSTGRES_PORT=port

