# go-payments
Make payments to users via bank deposits (VISA, Mastercard) or Mobile Money

## Process
- User with the role of `PaymentInitiator` logs into the system via `OAauth` with possible `MFA` integrated.
- Initiate the payment process by adding a small narrative, selecting from a list of users to send payments to, provide the amount.
- Submit the request
- A request entry is created for each person chosen.
- A user with a role `PaymentApprover` is notified of the request, logs into the system, views and approves.
- After approval, the payment is made to the users based on their preferred mode of payment.  

## Technologies
- OAuth0
- Gin
- Redis
- Postgres
- etc . . .




