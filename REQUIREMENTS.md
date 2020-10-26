## Requirements

- Implement a Rest API with features like "AddFriend", "Unfriend", "Block", "Receive Updates" etc.
- Database: PostgreSQL.
- Unit test

## Detail of endpoints

#### 1. Add Friends

- URL path: `/api/users/addfriend`
- HTTP method: `POST`
- Body: Two email address in friends array

        Required|Data type|Description|
        | --- | --- | --- |
        | yes | string | email of first user    
        | yes | string | email of second user

    Validate is email address and could not be empty any field
- Response:
    - Content type: `application/json` 
    - HTTP status: `200 OK`
    - Body: Result status success

        |Name|Data type|Description|
        | --- | --- | --- |
        | `success` |Boolean| Status Add Friends Both Users |

- Example:  GET `/api/users/addfriend`
  - Request body:
    ```json
    {
        "friends":
            [
            "andy@example.com",
            "john@example.com"
            ]
    }
     ```  
  - Response:
    ```json
    {
        "success":true
    }
    ```

