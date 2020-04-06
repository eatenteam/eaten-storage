# Eaten Storage

An official API for manipulating Eaten Storage. We GoLang as a language we used to develop this API and the design is based on CRUD principles as far as we are in an early stage. Through the development in the future, we are looking forward to extend our API to serve customized and essential operation other than that of CRUD operation. Below this it will be an API specification for each of our business objects which are `mall`, `shop`, and `product`.

## Mall

An object that contains details of shopping malls the we partner with.

### Model

| Key | Value |
| --- | ----- |
| id | string |
| brand | string |
| active | boolean |
| shops | array of `shop` objects |

### API Endpoint

- `GET \api\malls` responds a list of `mall` objects
- `GET \api\malls\:id` responds a `mall` object
- `POST \api\malls` responds an ID of inserted `mall` object
- `PUT \api\malls\:id` responds a `mall` object right before it updated
- `DELETE \api\malls\:id` responds a `mall` object that is deleted

## Shop

An object that outlines the details of restaurants(shops) in the database.

### Model

| Key | Value |
| --- | ----- |
| id | string |
| brand | string |
| tel | boolean |
| location | array of `location` objects |
| stock | array of `stock` objects |
| open | string (`hhmm`) |
| close | string (`hhmm`) |

#### Stock

| Key | Value |
| --- | ----- |
| product | string |
| quantity | int |

#### Location

| Key | Value |
| --- | ----- |
| mall | string |
| addr | string |

### API Endpoint

- `GET \api\shops` responds a list of `shop` objects
- `GET \api\shops\:id` responds a `shop` object
- `POST \api\shops` responds an ID of inserted `shop` object
- `PUT \api\shops\:id` responds a `shop` object right before it updated
- `DELETE \api\shops\:id` responds a `shop` object that is deleted

## Product

An object which define a product of a particular shops.

### Model

| Key | Value |
| --- | ----- |
| id | string |
| name | string |
| price | int |

### API Endpoint

- `GET \api\products` responds a list of `product` objects
- `GET \api\products\:id` responds a `product` object
- `POST \api\products` responds an ID of inserted `product` object
- `PUT \api\products\:id` responds a `product` object right before it updated
- `DELETE \api\products\:id` responds a `product` object that is deleted
