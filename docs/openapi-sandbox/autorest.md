# Creating an OpenAPI Swagger doc

1. Go to https://editor.swagger.io/
2. Make a call to every endpoint by filling in your URI, methods (POST, GET, PUT, etc) and any other parameters and authentication details for the API (See editor window screenshot below).
You will notice the Swagger doc UI simultaneously update based on what you input.
3. Once completed, export the definition file as a .yaml and .json file.
#Generating Client SDKs with Autorest
4. Install Autorest: 
  ```
  npm install autorest
  ```
