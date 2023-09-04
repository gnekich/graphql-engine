package hasura

import (
	"io"

	"github.com/hasura/graphql-engine/cli/v2/internal/httpc"
)

type GenericSend func(requestBody interface{}) (httpcResponse *httpc.Response, responseBody io.Reader, error error)

type Client struct {
	V1Metadata V1Metadata
	V1Query    V1Query
	V2Query    V2Query
	PGDump     PGDump
	V1Graphql  V1Graphql
	V1Version  V1Version
}

type V1Query interface {
	CommonMetadataOperations
	PGSourceOps
	Send(requestBody interface{}) (httpcResponse *httpc.Response, body io.Reader, error error)
	Bulk([]RequestBody) (io.Reader, error)
}

type V1Metadata interface {
	CommonMetadataOperations
	V2CommonMetadataOperations
	CatalogStateOperations
	Send(requestBody interface{}) (httpcResponse *httpc.Response, body io.Reader, error error)
}

type CatalogStateOperations interface {
	Set(key string, state interface{}) (io.Reader, error)
	Get() (io.Reader, error)
}

type SourceKind string

const (
	SourceKindPG        SourceKind = "postgres"
	SourceKindMSSQL     SourceKind = "mssql"
	SourceKindCitus     SourceKind = "citus"
	SourceKindCockroach SourceKind = "cockroach"
	SourceKindBigQuery  SourceKind = "bigquery"
	SourceKindSnowflake SourceKind = "snowflake"
)

type V2Query interface {
	PGSourceOps
	MSSQLSourceOps
	CitusSourceOps
	CockroachSourceOps
	BigQuerySourceOps
	Send(requestBody interface{}) (httpcResponse *httpc.Response, body io.Reader, error error)
	Bulk([]RequestBody) (io.Reader, error)
}

type V1VersionResponse struct {
	Version    string  `json:"version,omitempty"`
	ServerType *string `json:"server_type,omitempty"`
}

type V1Version interface {
	GetVersion() (*V1VersionResponse, error)
}

type IntrospectionSchema interface{}
type V1Graphql interface {
	GetIntrospectionSchema() (IntrospectionSchema, error)
}
type RequestBody struct {
	Type    string      `json:"type"`
	Version uint        `json:"version,omitempty"`
	Args    interface{} `json:"args"`
}
