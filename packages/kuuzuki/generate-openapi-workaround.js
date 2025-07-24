#!/usr/bin/env node

// Temporary workaround script to generate a minimal OpenAPI spec
// This bypasses the discriminated union issue

const openApiSpec = {
  "openapi": "3.0.0",
  "info": {
    "title": "kuucode",
    "version": "1.0.0",
    "description": "kuucode API"
  },
  "paths": {
    "/session": {
      "post": {
        "summary": "Create a new session",
        "operationId": "createSession",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/CreateSessionRequest"
              }
            }
          }
        },
        "responses": {
          "200": {
            "description": "Session created",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Session"
                }
              }
            }
          }
        }
      }
    },
    "/session/{id}/message": {
      "post": {
        "summary": "Send a message to a session",
        "operationId": "sendMessage",
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/SendMessageRequest"
              }
            }
          }
        },
        "responses": {
          "200": {
            "description": "Message sent",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "CreateSessionRequest": {
        "type": "object",
        "properties": {
          "providerID": {
            "type": "string"
          },
          "model": {
            "type": "string"
          },
          "system": {
            "type": "string"
          }
        }
      },
      "Session": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string"
          },
          "providerID": {
            "type": "string"
          },
          "model": {
            "type": "string"
          }
        }
      },
      "SendMessageRequest": {
        "type": "object",
        "properties": {
          "text": {
            "type": "string"
          },
          "files": {
            "type": "array",
            "items": {
              "type": "object",
              "properties": {
                "path": {
                  "type": "string"
                },
                "content": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  }
}

console.log(JSON.stringify(openApiSpec, null, 2))