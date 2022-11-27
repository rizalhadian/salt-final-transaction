# salt-final-transaction

This is my final project on golang salt academy.
<br/><br/>
## #Transaction
Transaction is lorem ipsum dolor is amet

---
### Get List Customer's Transaction
---
#### Endpoint :
```http
GET /api/customer/{customer_id}/transaction
```
#### Query Parameters :
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `page` | `integer` | Page Transaction |

#### Response :
```javascript
{
    "success"                             : boolean,
    "message"                             : string,
    "page"                                : integer,
    "per_page"                            : integer,
    "page_count"                          : integer,
    "total_count"                         : integer,
    "data"            : [
        {
          "id"                          : integer,
          "total_price"                 : double,
          "total_discount"              : double,
          "total_price_after_discount"  : double,
          "note"                        : string,
          "created_at"                  : string,
          "status"                      : string,
        },
        {
          "id"                          : integer,
          "total_price"                 : double,
          "total_discount"              : double,
          "total_price_after_discount"  : double,
          "note"                        : string,
          "created_at"                  : string,
          "status"                      : string,
      },
    ]
}
```
<br/>

---
### Get Customer's Transaction Data Count
---
#### Endpoint :
```http
GET /api/customer/{customer_id}/transaction-data-count
```

#### Response :
```javascript
{
    "success"         : boolean,
    "message"         : string,
    "data"            : {
        "total_transaction_spend"           : double,
        "total_transaction_count"           : double,
        "first_transaction_date"            : date_time,
        "last_transaction_date"             : date_time,
    }
}
```
<br/>

---
### Get Customer's Transaction Detail
---
#### Endpoint :
```http
GET /api/customer/{customer_id}/transaction/{transaction_id}
```

#### Response :
```javascript
{
    "success"   : boolean,
    "message"   : string,
    "data"      : 
        {
            "id"                          : integer,
            "total_price"                 : double,
            "total_discount_amount"       : double,
            "final_total_price"           : double,
            "note"                        : string,
            "created_at"                  : string,
            "status"                      : string,
            "items"                       : [
                {
                    "id"                            : integer,
                    "items_type_id"                 : integer,
                    "item_id"                       : integer,
                    "price"                         : double,
                    "qty"                           : integer,
                    "total_price"                   : double,
                    "note"                          : string,
                },
                {
                    "id"                            : integer,
                    "items_type_id"                 : integer,
                    "item_id"                       : integer,
                    "price"                         : double,
                    "qty"                           : integer,
                    "total_price"                   : double,
                    "note"                          : string,
                },
            ],
            "vouchers-redeemed"            : [
                {
                    "id"                            : integer,
                    "transactions_item_id"          : integer,
                    "voucher_code"                  : string,
                    "discount-percentage"           : integer,
                    "discount-amount"               : double,
                },
                {
                    "id"                            : integer,
                    "transactions_item_id"          : integer,
                    "voucher_code"                  : string,
                    "discount-percentage"           : integer,
                    "discount-amount"               : double,
                },
            ] 
        },
}
```
<br/>


---
### Add Customer's Transaction
---
#### Endpoint :
```http
POST /api/customer/{customer_id}/transaction
```

#### Response :
```javascript
{
    "id"                          : integer,
    "note"                        : string,
    "items"                       : [
        {
            "items_type_id"                 : integer,
            "item_id"                       : integer,
            "qty"                           : integer,
            "note"                          : string,
            "voucher_code"                  : string,
        },
        {
            "items_type_id"                 : integer,
            "item_id"                       : integer,
            "qty"                           : integer,
            "note"                          : string,
            "voucher_code"                  : string,
        },
    ],
}
```
<br/>

---
### Edit Customer's Transaction
---
This process will create rollback transaction and create new updated transaction 

#### Endpoint :
```http
PUT /api/customer/{customer_id}/transaction/{transaction_id}
```

#### Response :
```javascript
{
    "id"                          : integer,
    "note"                        : string,
    "items"                       : [
        {
            "items_type_id"                 : integer,
            "item_id"                       : integer,
            "qty"                           : integer,
            "note"                          : string,
            "voucher_code"                  : string,
        },
        {
            "items_type_id"                 : integer,
            "item_id"                       : integer,
            "qty"                           : integer,
            "note"                          : string,
            "voucher_code"                  : string,
        },
    ],
}
```
<br/>

---
### Delete Customer's Transaction
---
This process will create rollback transaction and not actually delete the transaction

#### Endpoint :
```http
DELETE /api/customer/{customer_id}/transaction/{transaction_id}
```

#### Response :
```javascript
{
    "id"                          : integer,
    "note"                        : string,
    "items"                       : [
        {
            "items_type_id"                 : integer,
            "item_id"                       : integer,
            "qty"                           : integer,
            "note"                          : string,
            "voucher_code"                  : string,
        },
        {
            "items_type_id"                 : integer,
            "item_id"                       : integer,
            "qty"                           : integer,
            "note"                          : string,
            "voucher_code"                  : string,
        },
    ],
}
```
<br/>


---
### Responses
---

Many API endpoints return the JSON representation of the resources created or edited. However, if an invalid request is submitted, or some other error occurs, Gophish returns a JSON response in the following format:

The `message` attribute contains a message commonly used to indicate errors or, in the case of deleting a resource, success that the resource was properly deleted.

The `success` attribute describes if the transaction was successful or not.

The `data` attribute contains any other metadata associated with the response. This will be an escaped string containing JSON data.
<br/><br/>

---
### Status Codes
---

| Status Code | Description |
| :--- | :--- |
| 200 | `OK` |
| 201 | `CREATED` |
| 400 | `BAD REQUEST` |
| 404 | `NOT FOUND` |
| 500 | `INTERNAL SERVER ERROR` |

<br/><br/>