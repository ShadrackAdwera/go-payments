# go-payments
Make payments to users via bank deposits (VISA, Mastercard) or Mobile Money

## Process
- A user with the permission of `payment:intiator` logs into the system via `Auth0` with possible `MFA` integrated.
- Initiate the payment process by adding a small narrative, selecting from a list of users to send payments to, provide the amount.
- The list of users are managed from the admin module.
- Users have a preferred payment type. If MPESA, funds are sent to their MPESA, if bank, payouts are made ito their bank accouts.
- Submit the request.
- A request entry is created for each person chosen.
- A user with the permission of `payment:approver` is notified of the request, logs into the system, views and approves.
- After approval, the payment is made to the users based on their preferred mode of payment (Bank Payout / MPESA).  

## Modules

1. Admin Module

- Manage users - add / modify user roles / remove users.
- Manage permissions - view list of permissions.
- View all requests - option to export payments.

2. Requests

- Make / View new requests from this view. Requires `payment:initiator` permission.

3. Approvals

- View / Approve payment requests from this module. Requires `payment:approver` permission.

4. Dashboard

- Dashboard view of all requests made, payments made, money in the account, mpesa balance etc


## Technologies
- Auth0
- Gin
- Redis
- Postgres
- etc . . .




