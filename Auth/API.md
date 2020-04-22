openapi: 3.0.0
servers:
  - description: Auth app
    url: localhost:8081
info:
  description: This is a simple API
  version: "1.0.0"
  title: Auth app
paths:
  /welcome:
    get:
      tags:
        - developers
      summary: welcome
      operationId: welcome
      description: welcome
      parameters:
        - in: query
          name: token
          description: reveal claim
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Credential Acces
  /login:
    post:
      tags:
        - developers
      summary: login
      operationId: login
      description: login
      responses:
        '200':
          description: Login Success
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Login'
        description: Login
  
  /Register:
    post:
      tags:
        - developers
      summary: Register
      operationId: Register
      description: Register
      responses:
        '200':
          description: Registration Success, please remember your password
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/User'
        description: Register
    

components:
  schemas:
    User:
      type: object
      required:
        - phone
        - name
        - role
      properties:
        phone:
          type: string
          example: 085224527270
        name:
          type: string
          example: Bilal
        role:
            type: string
            example: 'admin'
    
    Login:
        type: object
        properties:
            phone:
                type: string
                example: 085224527270
            password: 
                type: string
                example: '$2a$10$4oycZz.IjxylAFk3M4b.buv9y2Of8k5fXKf8OQNBBWZ61Xt0tG0Ou'