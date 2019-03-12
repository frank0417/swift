package v2

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var updateReleaseRequestSchema *gojsonschema.Schema
var getReleaseStatusRequestSchema *gojsonschema.Schema
var getVersionRequestSchema *gojsonschema.Schema
var rollbackReleaseRequestSchema *gojsonschema.Schema
var installReleaseRequestSchema *gojsonschema.Schema
var getReleaseContentRequestSchema *gojsonschema.Schema
var uninstallReleaseRequestSchema *gojsonschema.Schema
var summarizeReleasesRequestSchema *gojsonschema.Schema
var getHistoryRequestSchema *gojsonschema.Schema

func init() {
	var err error
	updateReleaseRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "chartChart": {
      "description": "Chart is a helm package that contains metadata, a default config, zero or more\n\toptionally parameterizable templates, and zero or more charts (dependencies).",
      "properties": {
        "dependencies": {
          "description": "Charts that this chart depends on.",
          "items": {
            "$ref": "#/definitions/chartChart"
          },
          "type": "array"
        },
        "files": {
          "description": "Miscellaneous files in a chart archive,\ne.g. README, LICENSE, etc.",
          "items": {
            "$ref": "#/definitions/hapichartAny"
          },
          "type": "array"
        },
        "metadata": {
          "$ref": "#/definitions/chartMetadata",
          "description": "Contents of the Chartfile."
        },
        "templates": {
          "description": "Templates for this chart.",
          "items": {
            "$ref": "#/definitions/chartTemplate"
          },
          "type": "array"
        },
        "values": {
          "$ref": "#/definitions/chartConfig",
          "description": "Default config for this template."
        }
      },
      "type": "object"
    },
    "chartConfig": {
      "description": "Config supplies values to the parametrizable templates of a chart.",
      "properties": {
        "raw": {
          "type": "string"
        },
        "values": {
          "additionalProperties": {
            "$ref": "#/definitions/chartValue"
          },
          "type": "object"
        }
      },
      "type": "object"
    },
    "chartMaintainer": {
      "description": "Maintainer describes a Chart maintainer.",
      "properties": {
        "email": {
          "title": "Email is an optional email address to contact the named maintainer",
          "type": "string"
        },
        "name": {
          "title": "Name is a user name or organization name",
          "type": "string"
        },
        "url": {
          "title": "Url is an optional URL to an address for the named maintainer",
          "type": "string"
        }
      },
      "type": "object"
    },
    "chartMetadata": {
      "description": "Metadata for a Chart file. This models the structure of a Chart.yaml file.\n\n\tSpec: https://k8s.io/helm/blob/master/docs/design/chart_format.md#the-chart-file",
      "properties": {
        "annotations": {
          "additionalProperties": {
            "type": "string"
          },
          "description": "Annotations are additional mappings uninterpreted by Tiller,\nmade available for inspection by other applications.",
          "type": "object"
        },
        "apiVersion": {
          "description": "The API Version of this chart.",
          "type": "string"
        },
        "appVersion": {
          "description": "The version of the application enclosed inside of this chart.",
          "type": "string"
        },
        "condition": {
          "title": "The condition to check to enable chart",
          "type": "string"
        },
        "deprecated": {
          "title": "Whether or not this chart is deprecated",
          "type": "boolean"
        },
        "description": {
          "title": "A one-sentence description of the chart",
          "type": "string"
        },
        "engine": {
          "description": "The name of the template engine to use. Defaults to 'gotpl'.",
          "type": "string"
        },
        "home": {
          "title": "The URL to a relevant project page, git repo, or contact person",
          "type": "string"
        },
        "icon": {
          "description": "The URL to an icon file.",
          "type": "string"
        },
        "keywords": {
          "items": {
            "type": "string"
          },
          "title": "A list of string keywords",
          "type": "array"
        },
        "kubeVersion": {
          "description": "KubeVersion is a SemVer constraint specifying the version of Kubernetes required.",
          "type": "string"
        },
        "maintainers": {
          "items": {
            "$ref": "#/definitions/chartMaintainer"
          },
          "title": "A list of name and URL/email address combinations for the maintainer(s)",
          "type": "array"
        },
        "name": {
          "title": "The name of the chart",
          "type": "string"
        },
        "sources": {
          "items": {
            "type": "string"
          },
          "title": "Source is the URL to the source code of this chart",
          "type": "array"
        },
        "tags": {
          "title": "The tags to check to enable chart",
          "type": "string"
        },
        "tillerVersion": {
          "title": "TillerVersion is a SemVer constraints on what version of Tiller is required.\nSee SemVer ranges here: https://github.com/Masterminds/semver#basic-comparisons",
          "type": "string"
        },
        "version": {
          "title": "A SemVer 2 conformant version string of the chart",
          "type": "string"
        }
      },
      "type": "object"
    },
    "chartTemplate": {
      "description": "Template represents a template as a name/value pair.\n\nBy convention, name is a relative path within the scope of the chart's\nbase directory.",
      "properties": {
        "data": {
          "description": "Data is the template as byte data.",
          "type": "string"
        },
        "name": {
          "description": "Name is the path-like name of the template.",
          "type": "string"
        }
      },
      "type": "object"
    },
    "chartValue": {
      "description": "Value describes a configuration value as a string.",
      "properties": {
        "value": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "hapichartAny": {
      "properties": {
        "type_url": {
          "description": "A resource name whose content describes the type of the\nserialized data.",
          "type": "string"
        },
        "value": {
          "description": "Data for file.",
          "type": "string"
        }
      },
      "title": "Copied from https://github.com/golang/protobuf/blob/master/ptypes/any/any.proto",
      "type": "object"
    }
  },
  "description": "UpdateReleaseRequest updates a release.",
  "properties": {
    "ca_bundle": {
      "title": "` + "`" + `ca_bundle` + "`" + ` is a PEM encoded CA bundle used to sign server certificate of chart repository.\n+optional",
      "type": "string"
    },
    "chart": {
      "$ref": "#/definitions/chartChart",
      "description": "Chart is the protobuf representation of a chart."
    },
    "chart_url": {
      "title": "URL to download chart archive",
      "type": "string"
    },
    "client_certificate": {
      "title": "ClientCertificate contains PEM-encoded data passed as a client cert to chart repository.\n+optional",
      "type": "string"
    },
    "client_key": {
      "title": "ClientKey contains PEM-encoded data passed as a client key to chart repository.\n+optional",
      "type": "string"
    },
    "disable_hooks": {
      "description": "DisableHooks causes the server to skip running any hooks for the upgrade.",
      "type": "boolean"
    },
    "dry_run": {
      "title": "dry_run, if true, will run through the release logic, but neither create",
      "type": "boolean"
    },
    "force": {
      "description": "Force resource update through delete/recreate if needed.",
      "type": "boolean"
    },
    "insecure_skip_verify": {
      "title": "Skip certificate verification for chart repository.\n+optional",
      "type": "boolean"
    },
    "name": {
      "title": "The name of the release",
      "type": "string"
    },
    "password": {
      "title": "Password is the password for basic authentication to the chart repository.\n+optional",
      "type": "string"
    },
    "recreate": {
      "title": "Performs pods restart for resources if applicable",
      "type": "boolean"
    },
    "reset_values": {
      "description": "ResetValues will cause Tiller to ignore stored values, resetting to default values.",
      "type": "boolean"
    },
    "reuse_values": {
      "description": "ReuseValues will cause Tiller to reuse the values from the last release.\nThis is ignored if reset_values is set.",
      "type": "boolean"
    },
    "timeout": {
      "description": "timeout specifies the max amount of time any kubernetes client command can run.",
      "type": "integer"
    },
    "token": {
      "title": "Token is the bearer token for authentication to the chart repository.\n+optional",
      "type": "string"
    },
    "username": {
      "title": "Username is the username for basic authentication to the chart repository.\n+optional",
      "type": "string"
    },
    "values": {
      "$ref": "#/definitions/chartConfig",
      "description": "Values is a string containing (unparsed) YAML values."
    },
    "wait": {
      "title": "wait, if true, will wait until all Pods, PVCs, and Services are in a ready state\nbefore marking the release as successful. It will wait for as long as timeout",
      "type": "boolean"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	getReleaseStatusRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "description": "GetReleaseStatusRequest is a request to get the status of a release.",
  "properties": {
    "name": {
      "title": "Name is the name of the release",
      "type": "string"
    },
    "version": {
      "title": "Version is the version of the release",
      "type": "integer"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	getVersionRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "description": "GetVersionRequest requests for version information.",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	rollbackReleaseRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "disable_hooks": {
      "title": "DisableHooks causes the server to skip running any hooks for the rollback",
      "type": "boolean"
    },
    "dry_run": {
      "title": "dry_run, if true, will run through the release logic but no create",
      "type": "boolean"
    },
    "force": {
      "description": "Force resource update through delete/recreate if needed.",
      "type": "boolean"
    },
    "name": {
      "title": "The name of the release",
      "type": "string"
    },
    "recreate": {
      "title": "Performs pods restart for resources if applicable",
      "type": "boolean"
    },
    "timeout": {
      "description": "timeout specifies the max amount of time any kubernetes client command can run.",
      "type": "integer"
    },
    "version": {
      "description": "Version is the version of the release to deploy.",
      "type": "integer"
    },
    "wait": {
      "title": "wait, if true, will wait until all Pods, PVCs, and Services are in a ready state\nbefore marking the release as successful. It will wait for as long as timeout",
      "type": "boolean"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	installReleaseRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "chartChart": {
      "description": "Chart is a helm package that contains metadata, a default config, zero or more\n\toptionally parameterizable templates, and zero or more charts (dependencies).",
      "properties": {
        "dependencies": {
          "description": "Charts that this chart depends on.",
          "items": {
            "$ref": "#/definitions/chartChart"
          },
          "type": "array"
        },
        "files": {
          "description": "Miscellaneous files in a chart archive,\ne.g. README, LICENSE, etc.",
          "items": {
            "$ref": "#/definitions/hapichartAny"
          },
          "type": "array"
        },
        "metadata": {
          "$ref": "#/definitions/chartMetadata",
          "description": "Contents of the Chartfile."
        },
        "templates": {
          "description": "Templates for this chart.",
          "items": {
            "$ref": "#/definitions/chartTemplate"
          },
          "type": "array"
        },
        "values": {
          "$ref": "#/definitions/chartConfig",
          "description": "Default config for this template."
        }
      },
      "type": "object"
    },
    "chartConfig": {
      "description": "Config supplies values to the parametrizable templates of a chart.",
      "properties": {
        "raw": {
          "type": "string"
        },
        "values": {
          "additionalProperties": {
            "$ref": "#/definitions/chartValue"
          },
          "type": "object"
        }
      },
      "type": "object"
    },
    "chartMaintainer": {
      "description": "Maintainer describes a Chart maintainer.",
      "properties": {
        "email": {
          "title": "Email is an optional email address to contact the named maintainer",
          "type": "string"
        },
        "name": {
          "title": "Name is a user name or organization name",
          "type": "string"
        },
        "url": {
          "title": "Url is an optional URL to an address for the named maintainer",
          "type": "string"
        }
      },
      "type": "object"
    },
    "chartMetadata": {
      "description": "Metadata for a Chart file. This models the structure of a Chart.yaml file.\n\n\tSpec: https://k8s.io/helm/blob/master/docs/design/chart_format.md#the-chart-file",
      "properties": {
        "annotations": {
          "additionalProperties": {
            "type": "string"
          },
          "description": "Annotations are additional mappings uninterpreted by Tiller,\nmade available for inspection by other applications.",
          "type": "object"
        },
        "apiVersion": {
          "description": "The API Version of this chart.",
          "type": "string"
        },
        "appVersion": {
          "description": "The version of the application enclosed inside of this chart.",
          "type": "string"
        },
        "condition": {
          "title": "The condition to check to enable chart",
          "type": "string"
        },
        "deprecated": {
          "title": "Whether or not this chart is deprecated",
          "type": "boolean"
        },
        "description": {
          "title": "A one-sentence description of the chart",
          "type": "string"
        },
        "engine": {
          "description": "The name of the template engine to use. Defaults to 'gotpl'.",
          "type": "string"
        },
        "home": {
          "title": "The URL to a relevant project page, git repo, or contact person",
          "type": "string"
        },
        "icon": {
          "description": "The URL to an icon file.",
          "type": "string"
        },
        "keywords": {
          "items": {
            "type": "string"
          },
          "title": "A list of string keywords",
          "type": "array"
        },
        "kubeVersion": {
          "description": "KubeVersion is a SemVer constraint specifying the version of Kubernetes required.",
          "type": "string"
        },
        "maintainers": {
          "items": {
            "$ref": "#/definitions/chartMaintainer"
          },
          "title": "A list of name and URL/email address combinations for the maintainer(s)",
          "type": "array"
        },
        "name": {
          "title": "The name of the chart",
          "type": "string"
        },
        "sources": {
          "items": {
            "type": "string"
          },
          "title": "Source is the URL to the source code of this chart",
          "type": "array"
        },
        "tags": {
          "title": "The tags to check to enable chart",
          "type": "string"
        },
        "tillerVersion": {
          "title": "TillerVersion is a SemVer constraints on what version of Tiller is required.\nSee SemVer ranges here: https://github.com/Masterminds/semver#basic-comparisons",
          "type": "string"
        },
        "version": {
          "title": "A SemVer 2 conformant version string of the chart",
          "type": "string"
        }
      },
      "type": "object"
    },
    "chartTemplate": {
      "description": "Template represents a template as a name/value pair.\n\nBy convention, name is a relative path within the scope of the chart's\nbase directory.",
      "properties": {
        "data": {
          "description": "Data is the template as byte data.",
          "type": "string"
        },
        "name": {
          "description": "Name is the path-like name of the template.",
          "type": "string"
        }
      },
      "type": "object"
    },
    "chartValue": {
      "description": "Value describes a configuration value as a string.",
      "properties": {
        "value": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "hapichartAny": {
      "properties": {
        "type_url": {
          "description": "A resource name whose content describes the type of the\nserialized data.",
          "type": "string"
        },
        "value": {
          "description": "Data for file.",
          "type": "string"
        }
      },
      "title": "Copied from https://github.com/golang/protobuf/blob/master/ptypes/any/any.proto",
      "type": "object"
    }
  },
  "description": "InstallReleaseRequest is the request for an installation of a chart.",
  "properties": {
    "ca_bundle": {
      "title": "` + "`" + `ca_bundle` + "`" + ` is a PEM encoded CA bundle used to sign server certificate of chart repository.\n+optional",
      "type": "string"
    },
    "chart": {
      "$ref": "#/definitions/chartChart",
      "description": "Chart is the protobuf representation of a chart."
    },
    "chart_url": {
      "title": "URL to download chart archive",
      "type": "string"
    },
    "client_certificate": {
      "title": "ClientCertificate contains PEM-encoded data passed as a client cert to chart repository.\n+optional",
      "type": "string"
    },
    "client_key": {
      "title": "ClientKey contains PEM-encoded data passed as a client key to chart repository.\n+optional",
      "type": "string"
    },
    "disable_hooks": {
      "description": "DisableHooks causes the server to skip running any hooks for the install.",
      "type": "boolean"
    },
    "dry_run": {
      "description": "DryRun, if true, will run through the release logic, but neither create\na release object nor deploy to Kubernetes. The release object returned\nin the response will be fake.",
      "type": "boolean"
    },
    "insecure_skip_verify": {
      "title": "Skip certificate verification for chart repository.\n+optional",
      "type": "boolean"
    },
    "name": {
      "description": "Name is the candidate release name. This must be unique to the\nnamespace, otherwise the server will return an error. If it is not\nsupplied, the server will autogenerate one.",
      "type": "string"
    },
    "namespace": {
      "description": "Namepace is the kubernetes namespace of the release.",
      "type": "string"
    },
    "password": {
      "title": "Password is the password for basic authentication to the chart repository.\n+optional",
      "type": "string"
    },
    "reuse_name": {
      "description": "ReuseName requests that Tiller re-uses a name, instead of erroring out.",
      "type": "boolean"
    },
    "timeout": {
      "description": "timeout specifies the max amount of time any kubernetes client command can run.",
      "type": "integer"
    },
    "token": {
      "title": "Token is the bearer token for authentication to the chart repository.\n+optional",
      "type": "string"
    },
    "username": {
      "title": "Username is the username for basic authentication to the chart repository.\n+optional",
      "type": "string"
    },
    "values": {
      "$ref": "#/definitions/chartConfig",
      "description": "Values is a string containing (unparsed) YAML values."
    },
    "wait": {
      "title": "wait, if true, will wait until all Pods, PVCs, and Services are in a ready state\nbefore marking the release as successful. It will wait for as long as timeout",
      "type": "boolean"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	getReleaseContentRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "description": "GetReleaseContentRequest is a request to get the contents of a release.",
  "properties": {
    "format_values_as_json": {
      "title": "Format release config and values to JSON string",
      "type": "boolean"
    },
    "name": {
      "title": "The name of the release",
      "type": "string"
    },
    "version": {
      "title": "Version is the version of the release",
      "type": "integer"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	uninstallReleaseRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "description": "UninstallReleaseRequest represents a request to uninstall a named release.",
  "properties": {
    "disable_hooks": {
      "description": "DisableHooks causes the server to skip running any hooks for the uninstall.",
      "type": "boolean"
    },
    "name": {
      "description": "Name is the name of the release to delete.",
      "type": "string"
    },
    "purge": {
      "description": "Purge removes the release from the store and make its name free for later use.",
      "type": "boolean"
    },
    "timeout": {
      "description": "timeout specifies the max amount of time any kubernetes client command can run.",
      "type": "integer"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	summarizeReleasesRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "ListSortSortBy": {
      "default": "UNKNOWN",
      "description": "SortBy defines sort operations.",
      "enum": [
        "UNKNOWN",
        "NAME",
        "LAST_RELEASED"
      ],
      "type": "string"
    },
    "ListSortSortOrder": {
      "default": "ASC",
      "description": "SortOrder defines sort orders to augment sorting operations.",
      "enum": [
        "ASC",
        "DESC"
      ],
      "type": "string"
    }
  },
  "description": "SummarizeReleasesRequest requests a list of releases.\n\nReleases can be retrieved in chunks by setting limit and offset.\n\nReleases can be sorted according to a few pre-determined sort stategies.",
  "properties": {
    "all": {
      "title": "if true, list releases with all status codes",
      "type": "boolean"
    },
    "filter": {
      "description": "Filter is a regular expression used to filter which releases should be listed.\n\nAnything that matches the regexp will be included in the results.",
      "type": "string"
    },
    "limit": {
      "description": "Limit is the maximum number of releases to be returned.",
      "type": "integer"
    },
    "namespace": {
      "description": "Namespace is the filter to select releases only from a specific namespace.",
      "type": "string"
    },
    "offset": {
      "description": "Offset is the last release name that was seen. The next listing\noperation will start with the name after this one.\nExample: If list one returns albert, bernie, carl, and sets 'next: dennis'.\ndennis is the offset. Supplying 'dennis' for the next request should\ncause the next batch to return a set of results starting with 'dennis'.",
      "type": "string"
    },
    "sort_by": {
      "$ref": "#/definitions/ListSortSortBy",
      "description": "SortBy is the sort field that the ListReleases server should sort data before returning."
    },
    "sort_order": {
      "$ref": "#/definitions/ListSortSortOrder",
      "description": "SortOrder is the ordering directive used for sorting."
    },
    "status_codes": {
      "items": {
        "type": "string"
      },
      "title": "allowed values are UNKNOWN, DEPLOYED, DELETED, SUPERSEDED, FAILED, DELETING",
      "type": "array"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	getHistoryRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "description": "GetHistoryRequest requests a release's history.",
  "properties": {
    "max": {
      "description": "The maximum number of releases to include.",
      "type": "integer"
    },
    "name": {
      "description": "The name of the release.",
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *UpdateReleaseRequest) Valid() (*gojsonschema.Result, error) {
	return updateReleaseRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *UpdateReleaseRequest) IsRequest() {}

func (m *GetReleaseStatusRequest) Valid() (*gojsonschema.Result, error) {
	return getReleaseStatusRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *GetReleaseStatusRequest) IsRequest() {}

func (m *GetVersionRequest) Valid() (*gojsonschema.Result, error) {
	return getVersionRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *GetVersionRequest) IsRequest() {}

func (m *RollbackReleaseRequest) Valid() (*gojsonschema.Result, error) {
	return rollbackReleaseRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *RollbackReleaseRequest) IsRequest() {}

func (m *InstallReleaseRequest) Valid() (*gojsonschema.Result, error) {
	return installReleaseRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *InstallReleaseRequest) IsRequest() {}

func (m *GetReleaseContentRequest) Valid() (*gojsonschema.Result, error) {
	return getReleaseContentRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *GetReleaseContentRequest) IsRequest() {}

func (m *UninstallReleaseRequest) Valid() (*gojsonschema.Result, error) {
	return uninstallReleaseRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *UninstallReleaseRequest) IsRequest() {}

func (m *SummarizeReleasesRequest) Valid() (*gojsonschema.Result, error) {
	return summarizeReleasesRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SummarizeReleasesRequest) IsRequest() {}

func (m *GetHistoryRequest) Valid() (*gojsonschema.Result, error) {
	return getHistoryRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *GetHistoryRequest) IsRequest() {}
