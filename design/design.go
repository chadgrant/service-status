package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("servicestatus", func() {
	Title("Service Status")
	Description("Service for managing and monitoring services")
	Server("servicestatus", func() {
		Host("localhost", func() {
			URI("http://localhost:6066/{version}")
			URI("https://localhost:6443/{version}")
			URI("grpc://localhost:6666/{version}")
			Variable("version", String, "API version", func() {
				Default("v1")
			})
		})
	})
})

var _ = Service("environment", func() {
	Description("The environment service performs operations on environments")

	HTTP(func() {
		Path("/v1/environment")
	})

	Method("list", func() {
		Description("List all environments")
		Result(CollectionOf(Environment))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("add", func() {
		Description("adds a new environment and returns the url of where to retrieve it")
		Payload(EnvironmentBase)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("update", func() {
		Description("update an existing environment")
		Payload(EnvironmentBase)
		Error("not_found", NotFound, "Environment not found")
		HTTP(func() {
			PUT("/{key}")
			Params(func() {
				Param("key", String)
			})
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
})

var EnvironmentBase = Type("EnvironmentBase", func() {
	Description("Environment describes environments such as Development, Production, Staging")
	Attribute("name", String, "Name of environment", func() {
		MaxLength(150)
		Example("Production")
		Meta("rpc:tag", "1")
	})
	Attribute("friendly", String, "url friendly name of environment used for REST based operations", func() {
		MaxLength(150)
		Example("development")
		Meta("rpc:tag", "2")
	})
	Attribute("active", Boolean, "is environment active", func() {
		Default(true)
		Example(true)
		Meta("rpc:tag", "3")
	})
	Attribute("sort", UInt32, "useful for sorting in UI", func() {
		Default(0)
		Example(0)
		Minimum(0)
		Maximum(5000)
		Meta("rpc:tag", "4")
	})
	Attribute("key", String, "used for update case of changing key, PUT /environment/key", func() {
		Meta("rpc:tag", "5")
	})
	Required("name", "friendly")
})

var Environment = ResultType("application/vnd.service-status.environment", func() {
	Reference(EnvironmentBase)
	TypeName("Environment")

	Field(1, "name")
	Field(2, "friendly")
	Field(3, "active")
	Field(4, "sort")

	Attribute("created", String, "timestamp of when environment was created", func() {
		Format(FormatDateTime)
		Example("1996-12-19T16:39:57-08:00")
		Meta("rpc:tag", "5")
	})
	Attribute("updated", String, "timestamp of when environment was updated", func() {
		Format(FormatDateTime)
		Example("1996-12-19T16:39:57-08:00")
		Meta("rpc:tag", "7")
	})
	Required("name", "friendly")
	View("default", func() {
		Attribute("friendly")
		Attribute("name")
		Attribute("active")
		Attribute("sort")
		Attribute("created")
		Attribute("updated")
	})
})

var NotFound = Type("NotFound", func() {
	Description("NotFound when an environment does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("environment test not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "name", String, "friendly name of environment")
	Required("message", "name")
})

var _ = Service("swagger", func() {
	Description("The swagger service serves the API swagger definition.")
	HTTP(func() {
		Path("/swagger")
	})
	Files("/swagger.json", "../../gen/http/openapi.json", func() {
		Description("JSON document containing the API swagger definition")
	})
})
