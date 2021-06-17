# Books Api
# Documentation

<br />
**Categories**
----
    Create a new categoria.
<br />

* **URL**
    /category_
<br />

* **Method:**
    `POST`
<br />

*  **URL Params** <br />
    None
<br />

* **Data Params** <br />
    **Required:**

    ```json
    { 
        "name": "<name>" 
    }
    ```
    <br />
    
* **Success Response:**

* **Code:** 201 <br />
    **Content:** 
    ```json
    {   
        "id" : 12, 
        "name" : "Michael Bloom" 
    }
    ```
    <br />

* **Error Response:** <br />

* **Code:** 400 <br />
    **Description: if you don't send the body.** <br />
    **Content:** 
   ```json
    {
        "message": "no data on body found"
    }
    ```
    <br />

* **Code:** 400 <br />
    **Description: if name of category already exists.** <br />
    **Content:** 
   ```json
    {
        "message": "category.name already exists"
    }
    ```
    <br />

* **Code:** 400 <br />
    **Description: if name of category is small.** <br />
    **Content:** 
   ```json
    {
        "message": "category.name is small"
    }
    ```
    <br />

* **Code:** 400 <br />
    **Description: if name of category is large.** <br />
    **Content:** 
   ```json
    {
        "message": "category.name is large"
    }
    ```
    <br />

* **Code:** 500 <br />
    **Description: if the service has a problem.** <br />
    **Content:** 
   ```json
    {
        "message": "service has a problem"
    }
    ```
    <br />