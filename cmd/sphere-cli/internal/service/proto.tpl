syntax = "proto3";

package {{.PackageName}};

import "buf/validate/validate.proto";
import "entpb/entpb.proto";
import "google/api/annotations.proto";
import "sphere/binding/binding.proto";

service {{.ServiceName}}Service {
  rpc List{{plural .ServiceName}}(List{{plural .ServiceName}}Request) returns (List{{plural .ServiceName}}Response) {
    option (google.api.http) = {get: "/api/{{.RouteName}}/list"};
  }
  rpc Create{{.ServiceName}}(Create{{.ServiceName}}Request) returns (Create{{.ServiceName}}Response) {
    option (google.api.http) = {
      post: "/api/{{.RouteName}}/create"
      body: "*"
    };
  }
  rpc Update{{.ServiceName}}(Update{{.ServiceName}}Request) returns (Update{{.ServiceName}}Response) {
    option (google.api.http) = {
      post: "/api/{{.RouteName}}/update"
      body: "*"
    };
  }
  rpc Get{{.ServiceName}}(Get{{.ServiceName}}Request) returns (Get{{.ServiceName}}Response) {
    option (google.api.http) = {get: "/api/{{.RouteName}}/detail/{id}"};
  }
  rpc Delete{{.ServiceName}}(Delete{{.ServiceName}}Request) returns (Delete{{.ServiceName}}Response) {
    option (google.api.http) = {delete: "/api/{{.RouteName}}/delete/{id}"};
  }
}

message List{{plural .ServiceName}}Request {
  option (sphere.binding.default_location) = BINDING_LOCATION_QUERY;

  int64 page = 1 [
    (buf.validate.field).required = false,
    (buf.validate.field).int64.gte = 0
  ];
  int64 page_size = 2 [(buf.validate.field).int64.gte = 0];
}

message List{{plural .ServiceName}}Response {
  repeated entpb.{{.ServiceName}} {{plural .EntityName}} = 1;
  int64 total_size = 2;
  int64 total_page = 3;
}

message Create{{.ServiceName}}Request {
  entpb.{{.ServiceName}} {{.EntityName}} = 1;
}

message Create{{.ServiceName}}Response {
  entpb.{{.ServiceName}} {{.EntityName}} = 1;
}

message Update{{.ServiceName}}Request {
  entpb.{{.ServiceName}} {{.EntityName}} = 1;
}

message Update{{.ServiceName}}Response {
  entpb.{{.ServiceName}} {{.EntityName}} = 1;
}

message Get{{.ServiceName}}Request {
  int64 id = 1 [(sphere.binding.location) = BINDING_LOCATION_URI];
}

message Get{{.ServiceName}}Response {
  entpb.{{.ServiceName}} {{.EntityName}} = 1;
}

message Delete{{.ServiceName}}Request {
  int64 id = 1 [(sphere.binding.location) = BINDING_LOCATION_URI];
}

message Delete{{.ServiceName}}Response {}