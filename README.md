# Category RESTful API

API Spec for Category RESTful API

**Version**: 1.0.0

## Table of Contents

- [Overview](#overview)
- [Base URL](#base-url)
- [Endpoints](#endpoints)
  - [List all Categories](#list-all-categories)
  - [Create new Category](#create-new-category)
  - [Get Category by Id](#get-category-by-id)
  - [Update Category by ID](#update-category-by-id)
  - [Delete Category by Id](#delete-category-by-id)
- [Security](#security)
- [Data Models](#data-models)

## Overview

This document provides the API specification for the Category RESTful API. It includes information about the API's version, base URL, available endpoints, security, and data models.

## Base URL

The Category RESTful API is hosted at:

```
http://localhost:3000/api
```

## Endpoints

### List all Categories

- **URL:** `/categories`
- **Method:** GET
- **Description:** List all Categories
- **Summary:** List of all categories
- **Security:** CategoryAuth (Authentication for Category API)
- **Response:**
  - Status 200 OK
    - Description: Success get all categories
    - Content-Type: application/json
    - Schema:
      ```json
      {
        "code": "number",
        "status": "string",
        "data": ["Category"]
      }
      ```

### Create new Category

- **URL:** `/categories`
- **Method:** POST
- **Description:** Create new Category
- **Summary:** Create new Category
- **Security:** CategoryAuth (Authentication for Category API)
- **Request Body:**
  - Content-Type: application/json
  - Schema: [CreateOrUpdateCategory](#createorupdatecategory)
- **Response:**
  - Status 200 OK
    - Description: Success create category
    - Content-Type: application/json
    - Schema:
      ```json
      {
        "code": "number",
        "status": "string",
        "data": ["Category"]
      }
      ```

### Get Category by Id

- **URL:** `/categories/{categoryId}`
- **Method:** GET
- **Description:** Get Category by Id
- **Summary:** Get Category by Id
- **Security:** CategoryAuth (Authentication for Category API)
- **Parameters:**
  - `categoryId` (path) - Category ID
- **Response:**
  - Status 200 OK
    - Description: Success Get Category
    - Content-Type: application/json
    - Schema:
      ```json
      {
        "code": "number",
        "status": "string",
        "data": ["Category"]
      }
      ```

### Update Category by ID

- **URL:** `/categories/{categoryId}`
- **Method:** PUT
- **Description:** Update Category by ID
- **Summary:** Update Category by ID
- **Security:** CategoryAuth (Authentication for Category API)
- **Parameters:**
  - `categoryId` (path) - Category ID
- **Request Body:**
  - Content-Type: application/json
  - Schema: [CreateOrUpdateCategory](#createorupdatecategory)
- **Response:**
  - Status 200 OK
    - Description: Success get all categories
    - Content-Type: application/json
    - Schema:
      ```json
      {
        "code": "number",
        "status": "string",
        "data": ["Category"]
      }
      ```

### Delete Category by Id

- **URL:** `/categories/{categoryId}`
- **Method:** DELETE
- **Description:** Delete Category by Id
- **Summary:** Delete Category by Id
- **Security:** CategoryAuth (Authentication for Category API)
- **Parameters:**
  - `categoryId` (path) - Category ID
- **Response:**
  - Status 200 OK
    - Description: Success delete Category
    - Content-Type: application/json
    - Schema:
      ```json
      {
        "code": "number",
        "status": "string"
      }
      ```

## Security

The Category RESTful API uses API key authentication. You should include the `X-API-Key` header in your requests to authenticate.

### CategoryAuth

- **Type:** API Key
- **Location:** Header
- **Name:** X-API-Key
- **Description:** Authentication for Category API

## Data Models

### CreateOrUpdateCategory

- **Type:** Object
- **Properties:**
  - `name` (string)

### Category

- **Type:** Object
- **Properties:**
  - `id` (number)
  - `name` (string)

Please make sure to include the necessary headers and parameters when making requests to the API endpoints.
