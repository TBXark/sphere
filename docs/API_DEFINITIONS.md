# Defining HTTP APIs with Protobuf

Sphere allows you to define HTTP interfaces for your services using standard Protobuf definitions and `google.api.http`
annotations. This document outlines the rules and conventions for mapping your gRPC methods to RESTful HTTP endpoints.

## Getting Started: A Basic Example

To expose a gRPC method as an HTTP endpoint, you need to define it in a `.proto` file and add an HTTP annotation.

Here is a basic example of a `greeter.proto` file that defines a simple `Greeter` service with one method, `SayHello`,
exposed as an HTTP `GET` request.

```protobuf
syntax = "proto3";

package your.service.v1;

import "google/api/annotations.proto";

// The Greeter service definition.
service Greeter {
  // Sends a greeting.
  rpc SayHello(SayHelloRequest) returns (SayHelloResponse) {
    option (google.api.http) = {
      get: "/v1/greeter/{name}"
    };
  }
}

// The request message for the SayHello RPC.
message SayHelloRequest {
  string name = 1;
}

// The response message for the SayHello RPC.
message SayHelloResponse {
  string message = 1;
}
```

### Key Components

1. **`import "google/api/annotations.proto";`**: This import is required to use HTTP annotations.
2. **`service Greeter { ... }`**: Defines your gRPC service.
3. **`rpc SayHello(...) returns (...)`**: Defines a method within the service.
4. **`option (google.api.http) = { ... };`**: This is the core of the HTTP mapping.
    * **`get: "/v1/greeter/{name}"`**: This specifies that the `SayHello` method should be exposed as an HTTP `GET`
      request. The path is `/v1/greeter/{name}`, where `{name}` is a path parameter mapped from the `name` field in the
      `SayHelloRequest` message.

Sphere uses these definitions to automatically generate server-side stubs and routing information.

## API Definition Rules

When defining HTTP transcoding rules, Sphere follows specific conventions to map your service methods to RESTful HTTP
endpoints.

### URL Path Mapping

Sphere converts gRPC-Gateway style URL paths from your `.proto` definitions into Gin-compatible routes. This includes support for path parameters, wildcards, and complex segments.

The following table shows how Protobuf URL paths are translated into Gin routes.

| Protobuf Path Template                           | Generated Gin Route                         |
|--------------------------------------------------|---------------------------------------------|
| `/users/{user_id}`                               | `/users/:user_id`                           |
| `/users/{user_id}/posts/{post_id}`               | `/users/:user_id/posts/:post_id`            |
| `/files/{file_path=**}`                          | `/files/*file_path`                         |
| `/files/{name=*}`                                | `/files/:name`                              |
| `/static/{path=assets/*}`                        | `/static/assets/:path`                      |
| `/static/{path=assets/**}`                       | `/static/assets/*path`                      |
| `/projects/{project_id}/locations/{location=**}` | `/projects/:project_id/locations/*location` |
| `/v1/users/{user.id}`                            | `/v1/users/:user_id`                        |
| `/api/{version=v1}/users`                        | `/api/v1/users`                             |
| `/users/{user_id}/posts/{post_id=drafts}`        | `/users/:user_id/posts/drafts`              |
| `/docs/{path=guides/**}`                         | `/docs/guides/*path`                        |
| `users`                                          | `/users`                                    |

#### Important Considerations

* **Routing Conflicts**: Avoid overly broad wildcard patterns like `/{path_test1:.*}` in path parameters, as this can
  lead to unexpected routing behavior.
* **Body Parsing**: Avoid using `body: "*"` in conjunction with path parameters, as it can cause conflicts during
  request parsing.

### HTTP Method and Field Binding

The binding of request message fields to the HTTP request (URL path, query parameters, or request body) depends on the
HTTP method.

* **GET / DELETE**: Fields in the request message that are not part of the URL path are automatically treated as URL
  query parameters.

* **POST / PUT / PATCH**:
    * By default, all fields in the request message not bound to the URL path are expected in the JSON request body.
    * To bind a field to URL query parameters instead, you can use a field binding annotation.

#### Field Binding Annotations

You can override the default binding behavior for `POST`, `PUT`, and `PATCH` methods by adding annotations in the
comments above the field definition in your `.proto` file. The `retags` command recognizes tags with the `// @sphere:`
prefix.

* `@sphere:uri` or `@sphere:uri="xxx"`: Binds the field to the URL path.
* `@sphere:form` or `@sphere:form="xxx"`: Binds the field to URL query parameters.
* `@sphere:json` or `@sphere:json="xxx"`: Binds the field to the request body.
* `@sphere:!json`: Excludes the field from JSON serialization by adding a `json:"-"` tag.

##### Automatic JSON Omission

By default, when you use `@sphere:uri` or `@sphere:form` to bind a field to the URL path or query parameters, Sphere's
`retags` command automatically adds a `json:"-"` tag to that field. This prevents fields from being accidentally exposed
in both the URL and the request body.

This behavior is controlled by the `--auto_omit_json` flag in the `sphere-cli retags` command and is enabled by default.

---

For a complete, practical example, see the [`test.proto`](../layout/proto/shared/v1/test.proto) file and its generated output in [`test_sphere.pb.go`](../layout/api/shared/v1/test_sphere.pb.go).