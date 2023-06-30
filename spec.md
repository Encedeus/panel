# Encedeus specification

## Basic architecture

- ### Skyhook
    - written in rust
    - uses an actix web server
    - gRPC connection to Backend
    - runs in a non volatile docker container
    - SFTP server
      #### Role
        - controls the node machine
        - starts the servers inside docker containers
        - enables file read/write/transfer with a SFTP server
        - receives data for starting and managing servers
- ### Backend
    - written in go
    - uses a echo web server
    - REST api to frontend
    - gRPC connection to Skyhook nodes
    - gRPC connection to node plugin environments
    - Postgres database
    - config in hcl
    - ent orm
    - v1 / v2
      #### Role
        - interprets plugins
        - sends instructions to Skyhook
        - Provides a REST API service for the frontend
- ### Frontend
    - written in ts with svelte
    - displays data provided by the REST API
    - tailwind css
    - post css
      #### Role
        - visual representation of server data
        - visual representation of resource usage
        - interface for interaction with servers
        - UI
- ### Plugins
    - written in js
    - node

## Functionality

- ### Skyhook
    - gRPC server (running subprocesses)
    - WebSocket server (resource usage data)
    - bash commands
    - container spin up
    - container spin down
    - container restart
    - console input / output
    - usage data reporting
    - server spec reporting

## Endpoints documentation

- ### Auth
    - `POST /auth/login`
        - logging in using the email or username and the password
            - request body
              ```
              {
                 "email": <email, not required if username is provided>,
                 "username": <username, not required if email is provided>,
                 "password": <password>
              }
              ```
            - response body
              ```
                 {
                     "accessToken": <access token>,
                     "refreshToken": <refresh token>
                 }
              ```
    - `GET /auth/refresh`
        - refreshing the access token using the refresh token
            - **note: request has no body**
            - request header
              ```
                 Authorization: Bearer <refresh token>
              ```
            - response body
              ```
                 {
                     "accessToken": <access token>
                 }
              ```
- ### User
    - `POST /user/create`
        - creating a user
            - request header
              ```
                 Authorization: Bearer <access token>
              ```
            - request body
              ```
                 {
                     "name": <username>,
                     "password": <password>,
                     "email": <email>,
                     "role_id": <role id, not required if role name is provided>
                     "role_name": <role name, not required if role id is provided>
                 }
              ```
    - `POST /user/setpfp`
        - setting a user pfp, used in user creation and updating
            - request header
              ```
                 Authorization: Bearer <access token>
              ```
            - request body (multipart form)
              ```
                 file=<file path>
              ```
    - `GET /user/pfp/:uuid`
        - getting a user pfp
            - parameter set in url (uuid after `/user/pfp/`)
    - `POST /user/update`
        - updating a user
            - request header
              ```
                 Authorization: Bearer <access token>
              ```
            - request body
              ```
                 {
                     "id": <user uuid>
                     "name": <username>,
                     "password": <password>,
                     "email": <email>,
                     "role_id": <role id, not required if role name is provided>
                     "role_name": <role name, not required if role id is provided>
                 }
              ```  
    - `DELETE /user/delete`
        - updating a user
            - request header
              ```
                 Authorization: Bearer <access token>
              ```
            - request body
              ```
                 {
                     "id": <user uuid>
                 }
              ```
- ### Role
    - `POST /role/create`
        - creating a role
            - request header
              ```
                 Authorisation: Bearer <access token>  
              ```
            - request body
              ```
                 {
                     "name": <role name>,
                     "permissions": <permission, array of strings>
                 }                
              ```
    - `POST /role/update`
        - updating a role
            - request header
              ```
                 Authorisation: Bearer <access token>  
              ```
            - request body
              ```
                 {
                     "id": <role id>,
                     "name": <role name>,
                     "permissions": <permission, array of strings>
                 }                
              ```
    - `DELETE /role/delete`
        - creating a role
            - request header
              ```
                 Authorisation: Bearer <access token>  
              ```
            - request body
              ```
                 {
                     "id": <role id>
                 }                
              ```