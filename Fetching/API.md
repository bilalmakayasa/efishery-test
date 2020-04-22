openapi: 3.0.0
servers:
# Added by API Auto Mocking Plugin
  - description: Fetching app
    url: localhost:8082

info:
  description: This is a simple API
  version: "1.0.0"
  title: Fetching app
paths:
  /fetching:
    get:
      tags:
        - developers
      summary: fetching
      operationId: fetching
      description: Fetch data and add USD, the conversion rate will be updated every 6 hours based on hit time
      parameters:
        - in: header
          name: x-access-token
          description: access token
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Reveal Data

  /aggregate:
    get:
      tags:
        - developers
      summary: aggregate
      operationId: aggregate
      description: Showing max, min, average and Median grouped by weekly and province
      parameters:
        - in: header
          name: x-access-token
          description: access token
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Reveal Data
  
  /retrieve:
    get:
      tags:
        - developers
      summary: retrieve
      operationId: retrieve
      description: This endpoint is used for showing JWT claims
      parameters:
        - in: header
          name: x-access-token
          description: access token
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Reveal Data