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

- `GET \api\partners` responds a list of `mall` objects
- `GET \api\partners\:id` responds a `mall` object
- `POST \api\partners` responds an ID of inserted `mall` object
- `PUT \api\partners\:id` responds a `mall` object right before it updated
- `DELETE \api\partners\:id` responds a `mall` object that is deleted
