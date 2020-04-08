# LINE Adapter
Here, we are implementing LINEBot to interact with customers/users through automated LINE Official Account. This serves as a prototyping and validating channel for the team which is much important before our engineers have to develop a whole mobile application. Noting that there is no stable release yet, one can go to `dev` branch to view the current stage of the application. If someone is interested in contributing code, he/she can contact our email.(`general@eaten.team`)

## LINE Action Usage

**NOTE:** Deprecated

| Number | Action | Type | Data |
| ------ | ------ | ---- | ---- |
| 2.00 | Main Menu | POSTBACK | action: hungry |
| 2.10 | List Restaurants | POSTBACK | action: restaurants, restaurant: code |
| 2.11 | List Product | POSTBACK | action: products, restaurant: code, food: id |
| 2.12 | Add to Basket | POSTBACK | action: addToBasket, restaurant: code, food: id |
| 2.13 | View Product | POSTBACK | action: getProduct | Product, restaurant: code, food: id |
| 2.14 | View Basket | POSTBACK | action: getBasket, basket: id |
| 2.15 | Remove One from Basket | POSTBACK | action: removeOneFromBasket, basket: id, number: id |
| 2.16 | Remove Order from Basket | POSTBACK | action: removeFromBasker, basket: id, number: id |
| 2.20 | List Promotion | POSTBACK | action: promotions |
| 2.13 | List Product | POSTBACK | action: products, promotions: id |
| 2.30 | Coupon | POSTBACK | action: coupons |
| 3.00 | Basket | POSTBACK | action: getBasket, id |
| 3.10 | ------ | ---- | ---- |
| 4.00 | Payment | POSTBACK | action: createPayment |
| 4.20 | Apply Coupon | POSTBACK | action: applyCoupon |
| 4.11 | View Receipt | POSTBACK | action: getPayment |
| 5.00 | Tracking | POSTBACK | action: getBasketLocation |
