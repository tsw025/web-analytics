# Web Analyzer API Documentation

## Overview

The Web Analyzer API provides endpoints for user registration, authentication, and website analysis. This document includes details on each endpoint, parameters, sample requests, and possible responses.

## Table of Contents
- [Base URL](#base-url)
- [Endpoints](#endpoints)
    - [POST /auth/register](#post-authregister)
    - [POST /auth/login](#post-authlogin)
    - [POST /analyse](#post-analyse)
    - [GET /websitesbyID](#get-websitesbyid)
    - [GET /websites](#get-websites)
- [Authorization](#authorization)
- [Error Handling](#error-handling)

## Base URL
```
http://localhost:8000/api
```

## Endpoints

### POST /auth/register
- **Description**: Registers a new user with a unique username and password.
- **URL**: `/auth/register`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "tharinda22",
    "password": "Tharinda@123"
  }
  ```
- **Success Response**:
    - **Status Code**: `200 OK`
    - **Body**:
      ```json
      {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ3ZWIi..."
      }
      ```
- **Error Responses**:
    - **Status Code**: `400 Bad Request`
        - **Body**:
          ```json
          {
            "errors": [
              {
                "field": "username",
                "message": "Username already exists"
              }
            ],
            "message": "validation failed"
          }
          ```
    - **Status Code**: `401 Unauthorized`
        - **Body**:
          ```json
          {
            "error": "Invalid username or password.",
            "message": "Invalid username or password."
          }
          ```

### POST /auth/login
- **Description**: Authenticates an existing user and returns a bearer token.
- **URL**: `/auth/login`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "tharinda22",
    "password": "Tharinda@123"
  }
  ```
- **Success Response**:
    - **Status Code**: `200 OK`
    - **Body**:
      ```json
      {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ3ZWIi..."
      }
      ```

### POST /analyse
- **Description**: Submits a website URL for analysis.
- **URL**: `/analyse`
- **Method**: `POST`
- **Headers**:
    - `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
    "url": "https://golangci-lint.run/"
  }
  ```
- **Success Response**:
    - **Status Code**: `200 OK`
    - **Body**:
      ```json
      {
        "ID": 6,
        "WebsiteID": 6,
        "Data": {
          "links": {
            "external_links": 16,
            "internal_links": 44,
            "inaccessible_links": 2
          },
          "page_title": "Introduction | golangci-lint",
          "html_version": "HTML5",
          "headings_count": {
            "h1": 1,
            "h2": 8,
            "h3": 1,
            "h4": 0,
            "h5": 0,
            "h6": 0
          },
          "contains_login_form": false,
          "analysis_completed_at": "2024-10-07T03:29:34Z"
        },
        "Status": "completed",
        "CreatedAt": "2024-10-07T03:29:30.573784Z",
        "UpdatedAt": "2024-10-07T03:29:34.803617Z"
      }
      ```
- **Pending Analysis Response**:
    - **Status Code**: `200 OK`
    - **Body**:
      ```json
      {
        "ID": 8,
        "WebsiteID": 8,
        "Data": null,
        "Status": "pending",
        "CreatedAt": "2024-10-07T03:30:42.958160273Z",
        "UpdatedAt": "2024-10-07T03:30:42.958160273Z"
      }
      ```
- **Error Response**:
    - **Status Code**: `400 Bad Request`
        - **Body**:
          ```json
          {
            "errors": [
              {
                "field": "URL",
                "message": "failed on the url tag"
              }
            ],
            "message": "validation failed"
          }
          ```

### GET /websitesbyID/{id}
- **Description**: Retrieves detailed information for a specific website by its ID.
- **URL**: `/websitesbyID/{id}`
- **Method**: `GET`
- **Headers**:
    - `Authorization: Bearer <token>`
- **Path Parameter**:
    - `id`: The ID of the website to retrieve.
- **Success Response**:
    - **Status Code**: `200 OK`
    - **Body**:
      ```json
      {
        "id": 2,
        "url": "https://web.facebook.com/",
        "created_at": "2024-10-07T01:15:23.783626Z",
        "updated_at": "2024-10-07T01:15:23.795117Z",
        "analytics": {
          "ID": 2,
          "WebsiteID": 2,
          "Data": {
            "links": {
              "external_links": 20,
              "internal_links": 28,
              "inaccessible_links": 0
            },
            "page_title": "Facebook – log in or sign up",
            "html_version": "HTML5",
            "headings_count": {
              "h1": 0,
              "h2": 1,
              "h3": 0,
              "h4": 0,
              "h5": 0,
              "h6": 0
            },
            "contains_login_form": true,
            "analysis_completed_at": "2024-10-07T06:45:25+05:30"
          },
          "status": "completed",
          "CreatedAt": "2024-10-07T01:15:23.799547Z",
          "UpdatedAt": "2024-10-07T01:15:25.811175Z"
        }
      }
      ```

### GET /websites
- **Description**: Retrieves a list of all websites with their respective information.
- **URL**: `/websites`
- **Method**: `GET`
- **Headers**:
    - `Authorization: Bearer <token>`
- **Success Response**:
    - **Status Code**: `200 OK`
    - **Body**:
      ```json
      [
        {
          "id": 1,
          "url": "https://stackoverflow.com/",
          "created_at": "2024-10-07T01:14:05.071162Z",
          "updated_at": "2024-10-07T01:14:05.083857Z"
        },
        {
          "id": 2,
          "url": "https://web.facebook.com/",
          "created_at": "2024-10-07T01:15:23.783626Z",
          "updated_at": "2024-10-07T01:15:23.795117Z"
        }
        // More entries...
      ]
      ```

## Authorization

Some endpoints require an authorization token. Include this token in the header for such requests:
```
Authorization: Bearer <token>
```

## Error Handling

- **400 Bad Request**: Occurs when the request body contains invalid data.
- **401 Unauthorized**: Authentication failure due to invalid credentials.
- **403 Forbidden**: Access denied due to missing or invalid bearer token.
- **404 Not Found**: Resource not found (e.g., website ID does not exist).

--- 