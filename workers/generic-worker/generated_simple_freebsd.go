//go:build simple

// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package main

import (
	"encoding/json"

	tcclient "github.com/taskcluster/taskcluster/v58/clients/client-go"
)

type (
	Artifact struct {

		// Content-Encoding for the artifact. If not provided, `gzip` will be used, except for the
		// following file extensions, where `identity` will be used, since they are already
		// compressed:
		//
		// * 7z
		// * bz2
		// * deb
		// * dmg
		// * flv
		// * gif
		// * gz
		// * jpeg
		// * jpg
		// * png
		// * swf
		// * tbz
		// * tgz
		// * webp
		// * whl
		// * woff
		// * woff2
		// * xz
		// * zip
		// * zst
		//
		// Note, setting `contentEncoding` on a directory artifact will apply the same content
		// encoding to all the files contained in the directory.
		//
		// Since: generic-worker 16.2.0
		//
		// Possible values:
		//   * "identity"
		//   * "gzip"
		ContentEncoding string `json:"contentEncoding,omitempty"`

		// Explicitly set the value of the HTTP `Content-Type` response header when the artifact(s)
		// is/are served over HTTP(S). If not provided (this property is optional) the worker will
		// guess the content type of artifacts based on the filename extension of the file storing
		// the artifact content. It does this by looking at the system filename-to-mimetype mappings
		// defined in multiple `mime.types` files located under `/etc`. Note, setting `contentType`
		// on a directory artifact will apply the same contentType to all files contained in the
		// directory.
		//
		// See [mime.TypeByExtension](https://pkg.go.dev/mime#TypeByExtension).
		//
		// Since: generic-worker 10.4.0
		ContentType string `json:"contentType,omitempty"`

		// Date when artifact should expire must be in the future, no earlier than task deadline, but
		// no later than task expiry. If not set, defaults to task expiry.
		//
		// Since: generic-worker 1.0.0
		Expires tcclient.Time `json:"expires,omitempty"`

		// Name of the artifact, as it will be published. If not set, `path` will be used.
		// Conventionally (although not enforced) path elements are forward slash separated. Example:
		// `public/build/a/house`. Note, no scopes are required to read artifacts beginning `public/`.
		// Artifact names not beginning `public/` are scope-protected (caller requires scopes to
		// download the artifact). See the Queue documentation for more information.
		//
		// Since: generic-worker 8.1.0
		Name string `json:"name,omitempty"`

		// Relative path of the file/directory from the task directory. Note this is not an absolute
		// path as is typically used in docker-worker, since the absolute task directory name is not
		// known when the task is submitted. Example: `dist\regedit.exe`. It doesn't matter if
		// forward slashes or backslashes are used.
		//
		// Since: generic-worker 1.0.0
		Path string `json:"path"`

		// Artifacts can be either an individual `file` or a `directory` containing
		// potentially multiple files with recursively included subdirectories.
		//
		// Since: generic-worker 1.0.0
		//
		// Possible values:
		//   * "file"
		//   * "directory"
		Type string `json:"type"`
	}

	// Requires scope `queue:get-artifact:<artifact-name>`.
	//
	// Since: generic-worker 5.4.0
	ArtifactContent struct {

		// Max length: 1024
		Artifact string `json:"artifact"`

		// If provided, the required SHA256 of the content body.
		//
		// Since: generic-worker 10.8.0
		//
		// Syntax:     ^[a-f0-9]{64}$
		SHA256 string `json:"sha256,omitempty"`

		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		TaskID string `json:"taskId"`
	}

	// Base64 encoded content of file/archive, up to 64KB (encoded) in size.
	//
	// Since: generic-worker 11.1.0
	Base64Content struct {

		// Base64 encoded content of file/archive, up to 64KB (encoded) in size.
		//
		// Since: generic-worker 11.1.0
		//
		// Syntax:     ^[A-Za-z0-9/+]+[=]{0,2}$
		// Max length: 65536
		Base64 string `json:"base64"`
	}

	// By default tasks will be resolved with `state/reasonResolved`: `completed/completed`
	// if all task commands have a zero exit code, or `failed/failed` if any command has a
	// non-zero exit code. This payload property allows customsation of the task resolution
	// based on exit code of task commands.
	ExitCodeHandling struct {

		// If the task exits with a purge caches exit status, all caches
		// associated with the task will be purged.
		//
		// Since: generic-worker 49.0.0
		//
		// Array items:
		// Mininum:    0
		PurgeCaches []int64 `json:"purgeCaches,omitempty"`

		// Exit codes for any command in the task payload to cause this task to
		// be resolved as `exception/intermittent-task`. Typically the Queue
		// will then schedule a new run of the existing `taskId` (rerun) if not
		// all task runs have been exhausted.
		//
		// See [itermittent tasks](https://docs.taskcluster.net/docs/reference/platform/taskcluster-queue/docs/worker-interaction#intermittent-tasks) for more detail.
		//
		// Since: generic-worker 10.10.0
		//
		// Array items:
		// Mininum:    1
		Retry []int64 `json:"retry,omitempty"`
	}

	// Feature flags enable additional functionality.
	//
	// Since: generic-worker 5.3.0
	FeatureFlags struct {

		// The backing log feature publishes a task artifact containing the complete
		// stderr and stdout of the task.
		//
		// Since: generic-worker 48.2.0
		//
		// Default:    true
		BackingLog bool `json:"backingLog" default:"true"`

		// This allows you to interactively run commands from within the worker
		// as the task user. This may be useful for debugging purposes.
		// Can be used for SSH-like access to the running worker.
		// Note that this feature works differently from the `interactive` feature
		// in docker worker, which `docker exec`s into the running container.
		// Since tasks on generic worker are not guaranteed to be running in a
		// container, a bash shell is started on the task user's account.
		// A user can then `docker exec` into the a running container, if there
		// is one.
		//
		// Since: generic-worker 49.2.0
		Interactive bool `json:"interactive,omitempty"`

		// The live log feature streams the combined stderr and stdout to a task artifact
		// so that the output is available while the task is running.
		//
		// Since: generic-worker 48.2.0
		//
		// Default:    true
		LiveLog bool `json:"liveLog" default:"true"`

		// Audio loopback device created using snd-aloop.
		// An audio device will be available for the task. Its
		// location will be `/dev/snd`. Devices inside that directory
		// will take the form `/dev/snd/controlC<N>`,
		// `/dev/snd/pcmC<N>D0c`, `/dev/snd/pcmC<N>D0p`,
		// `/dev/snd/pcmC<N>D1c`, and `/dev/snd/pcmC<N>D1p`,
		// where <N> is an integer between 0 and 31, inclusive.
		// The Generic Worker config setting `loopbackAudioDeviceNumber`
		// may be used to change the device number in case the
		// default value (`16`) conflicts with another
		// audio device on the worker.
		//
		// This feature is only available on Linux. If a task
		// is submitted with this feature enabled on a non-Linux,
		// posix platform (FreeBSD, macOS), the task will resolve as
		// `exception/malformed-payload`.
		//
		// Since: generic-worker 54.5.0
		LoopbackAudio bool `json:"loopbackAudio,omitempty"`

		// Video loopback device created using v4l2loopback.
		// A video device will be available for the task. Its
		// location will be passed to the task via environment
		// variable `TASKCLUSTER_VIDEO_DEVICE`. The
		// location will be `/dev/video<N>` where `<N>` is
		// an integer between 0 and 255. The value of `<N>`
		// is not static, and therefore either the environment
		// variable should be used, or `/dev` should be
		// scanned in order to determine the correct location.
		// Tasks should not assume a constant value.
		//
		// This feature is only available on Linux. If a task
		// is submitted with this feature enabled on a non-Linux,
		// posix platform (FreeBSD, macOS), the task will resolve as
		// `exception/malformed-payload`.
		//
		// Since: generic-worker 53.1.0
		LoopbackVideo bool `json:"loopbackVideo,omitempty"`

		// The taskcluster proxy provides an easy and safe way to make authenticated
		// taskcluster requests within the scope(s) of a particular task. See
		// [the github project](https://github.com/taskcluster/taskcluster/tree/main/tools/taskcluster-proxy) for more information.
		//
		// Since: generic-worker 10.6.0
		TaskclusterProxy bool `json:"taskclusterProxy,omitempty"`
	}

	FileMount struct {

		// One of:
		//   * ArtifactContent
		//   * IndexedContent
		//   * URLContent
		//   * RawContent
		//   * Base64Content
		Content json.RawMessage `json:"content"`

		// The filesystem location to mount the file.
		//
		// Since: generic-worker 5.4.0
		File string `json:"file"`

		// Compression format of the preloaded content.
		//
		// Since: generic-worker 55.3.0
		//
		// Possible values:
		//   * "bz2"
		//   * "gz"
		//   * "lz4"
		//   * "xz"
		//   * "zst"
		Format string `json:"format,omitempty"`
	}

	// This schema defines the structure of the `payload` property referred to in a
	// Taskcluster Task definition.
	GenericWorkerPayload struct {

		// Artifacts to be published.
		//
		// Since: generic-worker 1.0.0
		Artifacts []Artifact `json:"artifacts,omitempty"`

		// One array per command (each command is an array of arguments). Several arrays
		// for several commands.
		//
		// Since: generic-worker 0.0.1
		//
		// Array items:
		// Array items:
		Command [][]string `json:"command"`

		// Env vars must be string to __string__ mappings (not number or boolean). For example:
		// ```
		// {
		//   "PATH": "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin",
		//   "GOOS": "darwin",
		//   "FOO_ENABLE": "true",
		//   "BAR_TOTAL": "3"
		// }
		// ```
		//
		// Note, the following environment variables will automatically be set in the task
		// commands:
		//   * `TASK_ID` - the task ID of the currently running task
		//   * `RUN_ID` - the run ID of the currently running task
		//   * `TASKCLUSTER_ROOT_URL` - the root URL of the taskcluster deployment
		//   * `TASKCLUSTER_PROXY_URL` (if taskcluster proxy feature enabled) - the
		//      taskcluster authentication proxy for making unauthenticated taskcluster
		//      API calls
		//   * `TASKCLUSTER_WORKER_LOCATION`. See
		//     [RFC #0148](https://github.com/taskcluster/taskcluster-rfcs/blob/master/rfcs/0148-taskcluster-worker-location.md)
		//     for details.
		//
		// Since: generic-worker 0.0.1
		//
		// Map entries:
		Env map[string]string `json:"env,omitempty"`

		// Feature flags enable additional functionality.
		//
		// Since: generic-worker 5.3.0
		Features FeatureFlags `json:"features,omitempty"`

		// Configuration for task logs.
		//
		// Since: generic-worker 48.2.0
		Logs Logs `json:"logs,omitempty"`

		// Maximum time the task container can run in seconds.
		// The maximum value for `maxRunTime` is set by a `maxTaskRunTime` config property specific to each worker-pool.
		//
		// Since: generic-worker 0.0.1
		//
		// Mininum:    1
		MaxRunTime int64 `json:"maxRunTime"`

		// Directories and/or files to be mounted.
		//
		// Since: generic-worker 5.4.0
		//
		// Array items:
		// One of:
		//   * FileMount
		//   * WritableDirectoryCache
		//   * ReadOnlyDirectory
		Mounts []json.RawMessage `json:"mounts,omitempty"`

		// By default tasks will be resolved with `state/reasonResolved`: `completed/completed`
		// if all task commands have a zero exit code, or `failed/failed` if any command has a
		// non-zero exit code. This payload property allows customsation of the task resolution
		// based on exit code of task commands.
		OnExitStatus ExitCodeHandling `json:"onExitStatus,omitempty"`

		// A list of OS Groups that the task user should be a member of. Not yet implemented on
		// non-Windows platforms, therefore this optional property may only be an empty array if
		// provided.
		//
		// Since: generic-worker 6.0.0
		//
		// Array items:
		OSGroups []string `json:"osGroups,omitempty"`

		// This property is allowed for backward compatibility, but is unused.
		SupersederURL string `json:"supersederUrl,omitempty"`
	}

	// Content originating from a task artifact that has been indexed by the Taskcluster Index Service.
	//
	// Since: generic-worker 51.0.0
	IndexedContent struct {

		// Max length: 1024
		Artifact string `json:"artifact"`

		// Max length: 255
		Namespace string `json:"namespace"`
	}

	// Configuration for task logs.
	//
	// Since: generic-worker 48.2.0
	Logs struct {

		// Specifies a custom name for the backing log artifact.
		// This is only used if `features.backingLog` is `true`.
		//
		// Since: generic-worker 48.2.0
		//
		// Default:    "public/logs/live_backing.log"
		Backing string `json:"backing" default:"public/logs/live_backing.log"`

		// Specifies a custom name for the live log artifact.
		// This is only used if `features.liveLog` is `true`.
		//
		// Since: generic-worker 48.2.0
		//
		// Default:    "public/logs/live.log"
		Live string `json:"live" default:"public/logs/live.log"`
	}

	// Byte-for-byte literal inline content of file/archive, up to 64KB in size.
	//
	// Since: generic-worker 11.1.0
	RawContent struct {

		// Byte-for-byte literal inline content of file/archive, up to 64KB in size.
		//
		// Since: generic-worker 11.1.0
		//
		// Max length: 65536
		Raw string `json:"raw"`
	}

	ReadOnlyDirectory struct {

		// One of:
		//   * ArtifactContent
		//   * IndexedContent
		//   * URLContent
		//   * RawContent
		//   * Base64Content
		Content json.RawMessage `json:"content"`

		// The filesystem location to mount the directory volume.
		//
		// Since: generic-worker 5.4.0
		Directory string `json:"directory"`

		// Archive format of content for read only directory.
		//
		// Since: generic-worker 5.4.0
		//
		// Possible values:
		//   * "rar"
		//   * "tar.bz2"
		//   * "tar.gz"
		//   * "tar.lz4"
		//   * "tar.xz"
		//   * "tar.zst"
		//   * "zip"
		Format string `json:"format"`
	}

	// URL to download content from.
	//
	// Since: generic-worker 5.4.0
	URLContent struct {

		// If provided, the required SHA256 of the content body.
		//
		// Since: generic-worker 10.8.0
		//
		// Syntax:     ^[a-f0-9]{64}$
		SHA256 string `json:"sha256,omitempty"`

		// URL to download content from.
		//
		// Since: generic-worker 5.4.0
		URL string `json:"url"`
	}

	WritableDirectoryCache struct {

		// Implies a read/write cache directory volume. A unique name for the
		// cache volume. Requires scope `generic-worker:cache:<cache-name>`.
		// Note if this cache is loaded from an artifact, you will also require
		// scope `queue:get-artifact:<artifact-name>` to use this cache.
		//
		// Since: generic-worker 5.4.0
		CacheName string `json:"cacheName"`

		// One of:
		//   * ArtifactContent
		//   * IndexedContent
		//   * URLContent
		//   * RawContent
		//   * Base64Content
		Content json.RawMessage `json:"content,omitempty"`

		// The filesystem location to mount the directory volume.
		//
		// Since: generic-worker 5.4.0
		Directory string `json:"directory"`

		// Archive format of the preloaded content (if `content` provided).
		//
		// Since: generic-worker 5.4.0
		//
		// Possible values:
		//   * "rar"
		//   * "tar.bz2"
		//   * "tar.gz"
		//   * "tar.lz4"
		//   * "tar.xz"
		//   * "tar.zst"
		//   * "zip"
		Format string `json:"format,omitempty"`
	}
)

// Returns json schema for the payload part of the task definition. Please
// note we use a go string and do not load an external file, since we want this
// to be *part of the compiled executable*. If this sat in another file that
// was loaded at runtime, it would not be burned into the build, which would be
// bad for the following two reasons:
//  1. we could no longer distribute a single binary file that didn't require
//     installation/extraction
//  2. the payload schema is specific to the version of the code, therefore
//     should be versioned directly with the code and *frozen on build*.
//
// Run `generic-worker show-payload-schema` to output this schema to standard
// out.
func JSONSchema() string {
	return `{
  "$id": "/schemas/generic-worker/simple_posix.json#",
  "$schema": "/schemas/common/metaschema.json#",
  "additionalProperties": false,
  "definitions": {
    "content": {
      "oneOf": [
        {
          "additionalProperties": false,
          "description": "Requires scope ` + "`" + `queue:get-artifact:\u003cartifact-name\u003e` + "`" + `.\n\nSince: generic-worker 5.4.0",
          "properties": {
            "artifact": {
              "maxLength": 1024,
              "type": "string"
            },
            "sha256": {
              "description": "If provided, the required SHA256 of the content body.\n\nSince: generic-worker 10.8.0",
              "pattern": "^[a-f0-9]{64}$",
              "title": "SHA 256",
              "type": "string"
            },
            "taskId": {
              "pattern": "^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$",
              "type": "string"
            }
          },
          "required": [
            "taskId",
            "artifact"
          ],
          "title": "Artifact Content",
          "type": "object"
        },
        {
          "additionalProperties": false,
          "description": "Content originating from a task artifact that has been indexed by the Taskcluster Index Service.\n\nSince: generic-worker 51.0.0",
          "properties": {
            "artifact": {
              "maxLength": 1024,
              "type": "string"
            },
            "namespace": {
              "maxLength": 255,
              "type": "string"
            }
          },
          "required": [
            "namespace",
            "artifact"
          ],
          "title": "Indexed Content",
          "type": "object"
        },
        {
          "additionalProperties": false,
          "description": "URL to download content from.\n\nSince: generic-worker 5.4.0",
          "properties": {
            "sha256": {
              "description": "If provided, the required SHA256 of the content body.\n\nSince: generic-worker 10.8.0",
              "pattern": "^[a-f0-9]{64}$",
              "title": "SHA 256",
              "type": "string"
            },
            "url": {
              "description": "URL to download content from.\n\nSince: generic-worker 5.4.0",
              "format": "uri",
              "title": "URL",
              "type": "string"
            }
          },
          "required": [
            "url"
          ],
          "title": "URL Content",
          "type": "object"
        },
        {
          "additionalProperties": false,
          "description": "Byte-for-byte literal inline content of file/archive, up to 64KB in size.\n\nSince: generic-worker 11.1.0",
          "properties": {
            "raw": {
              "description": "Byte-for-byte literal inline content of file/archive, up to 64KB in size.\n\nSince: generic-worker 11.1.0",
              "maxLength": 65536,
              "title": "Raw",
              "type": "string"
            }
          },
          "required": [
            "raw"
          ],
          "title": "Raw Content",
          "type": "object"
        },
        {
          "additionalProperties": false,
          "description": "Base64 encoded content of file/archive, up to 64KB (encoded) in size.\n\nSince: generic-worker 11.1.0",
          "properties": {
            "base64": {
              "description": "Base64 encoded content of file/archive, up to 64KB (encoded) in size.\n\nSince: generic-worker 11.1.0",
              "maxLength": 65536,
              "pattern": "^[A-Za-z0-9/+]+[=]{0,2}$",
              "title": "Base64",
              "type": "string"
            }
          },
          "required": [
            "base64"
          ],
          "title": "Base64 Content",
          "type": "object"
        }
      ]
    },
    "fileMount": {
      "additionalProperties": false,
      "properties": {
        "content": {
          "$ref": "#/definitions/content",
          "description": "Content of the file to be mounted.\n\nSince: generic-worker 5.4.0"
        },
        "file": {
          "description": "The filesystem location to mount the file.\n\nSince: generic-worker 5.4.0",
          "title": "File",
          "type": "string"
        },
        "format": {
          "description": "Compression format of the preloaded content.\n\nSince: generic-worker 55.3.0",
          "enum": [
            "bz2",
            "gz",
            "lz4",
            "xz",
            "zst"
          ],
          "title": "Format",
          "type": "string"
        }
      },
      "required": [
        "file",
        "content"
      ],
      "title": "File Mount",
      "type": "object"
    },
    "mount": {
      "oneOf": [
        {
          "$ref": "#/definitions/fileMount"
        },
        {
          "$ref": "#/definitions/writableDirectoryCache"
        },
        {
          "$ref": "#/definitions/readOnlyDirectory"
        }
      ],
      "title": "Mount"
    },
    "readOnlyDirectory": {
      "additionalProperties": false,
      "properties": {
        "content": {
          "$ref": "#/definitions/content",
          "description": "Contents of read only directory.\n\nSince: generic-worker 5.4.0",
          "title": "Content"
        },
        "directory": {
          "description": "The filesystem location to mount the directory volume.\n\nSince: generic-worker 5.4.0",
          "title": "Directory",
          "type": "string"
        },
        "format": {
          "description": "Archive format of content for read only directory.\n\nSince: generic-worker 5.4.0",
          "enum": [
            "rar",
            "tar.bz2",
            "tar.gz",
            "tar.lz4",
            "tar.xz",
            "tar.zst",
            "zip"
          ],
          "title": "Format",
          "type": "string"
        }
      },
      "required": [
        "directory",
        "content",
        "format"
      ],
      "title": "Read Only Directory",
      "type": "object"
    },
    "writableDirectoryCache": {
      "additionalProperties": false,
      "dependencies": {
        "content": [
          "format"
        ],
        "format": [
          "content"
        ]
      },
      "properties": {
        "cacheName": {
          "description": "Implies a read/write cache directory volume. A unique name for the\ncache volume. Requires scope ` + "`" + `generic-worker:cache:\u003ccache-name\u003e` + "`" + `.\nNote if this cache is loaded from an artifact, you will also require\nscope ` + "`" + `queue:get-artifact:\u003cartifact-name\u003e` + "`" + ` to use this cache.\n\nSince: generic-worker 5.4.0",
          "title": "Cache Name",
          "type": "string"
        },
        "content": {
          "$ref": "#/definitions/content",
          "description": "Optional content to be preloaded when initially creating the cache\n(if set, ` + "`" + `format` + "`" + ` must also be provided).\n\nSince: generic-worker 5.4.0",
          "title": "Content"
        },
        "directory": {
          "description": "The filesystem location to mount the directory volume.\n\nSince: generic-worker 5.4.0",
          "title": "Directory Volume",
          "type": "string"
        },
        "format": {
          "description": "Archive format of the preloaded content (if ` + "`" + `content` + "`" + ` provided).\n\nSince: generic-worker 5.4.0",
          "enum": [
            "rar",
            "tar.bz2",
            "tar.gz",
            "tar.lz4",
            "tar.xz",
            "tar.zst",
            "zip"
          ],
          "title": "Format",
          "type": "string"
        }
      },
      "required": [
        "directory",
        "cacheName"
      ],
      "title": "Writable Directory Cache",
      "type": "object"
    }
  },
  "description": "This schema defines the structure of the ` + "`" + `payload` + "`" + ` property referred to in a\nTaskcluster Task definition.",
  "properties": {
    "artifacts": {
      "description": "Artifacts to be published.\n\nSince: generic-worker 1.0.0",
      "items": {
        "additionalProperties": false,
        "properties": {
          "contentEncoding": {
            "description": "Content-Encoding for the artifact. If not provided, ` + "`" + `gzip` + "`" + ` will be used, except for the\nfollowing file extensions, where ` + "`" + `identity` + "`" + ` will be used, since they are already\ncompressed:\n\n* 7z\n* bz2\n* deb\n* dmg\n* flv\n* gif\n* gz\n* jpeg\n* jpg\n* png\n* swf\n* tbz\n* tgz\n* webp\n* whl\n* woff\n* woff2\n* xz\n* zip\n* zst\n\nNote, setting ` + "`" + `contentEncoding` + "`" + ` on a directory artifact will apply the same content\nencoding to all the files contained in the directory.\n\nSince: generic-worker 16.2.0",
            "enum": [
              "identity",
              "gzip"
            ],
            "title": "Content-Encoding header when serving artifact over HTTP.",
            "type": "string"
          },
          "contentType": {
            "description": "Explicitly set the value of the HTTP ` + "`" + `Content-Type` + "`" + ` response header when the artifact(s)\nis/are served over HTTP(S). If not provided (this property is optional) the worker will\nguess the content type of artifacts based on the filename extension of the file storing\nthe artifact content. It does this by looking at the system filename-to-mimetype mappings\ndefined in multiple ` + "`" + `mime.types` + "`" + ` files located under ` + "`" + `/etc` + "`" + `. Note, setting ` + "`" + `contentType` + "`" + `\non a directory artifact will apply the same contentType to all files contained in the\ndirectory.\n\nSee [mime.TypeByExtension](https://pkg.go.dev/mime#TypeByExtension).\n\nSince: generic-worker 10.4.0",
            "title": "Content-Type header when serving artifact over HTTP",
            "type": "string"
          },
          "expires": {
            "description": "Date when artifact should expire must be in the future, no earlier than task deadline, but\nno later than task expiry. If not set, defaults to task expiry.\n\nSince: generic-worker 1.0.0",
            "format": "date-time",
            "title": "Expiry date and time",
            "type": "string"
          },
          "name": {
            "description": "Name of the artifact, as it will be published. If not set, ` + "`" + `path` + "`" + ` will be used.\nConventionally (although not enforced) path elements are forward slash separated. Example:\n` + "`" + `public/build/a/house` + "`" + `. Note, no scopes are required to read artifacts beginning ` + "`" + `public/` + "`" + `.\nArtifact names not beginning ` + "`" + `public/` + "`" + ` are scope-protected (caller requires scopes to\ndownload the artifact). See the Queue documentation for more information.\n\nSince: generic-worker 8.1.0",
            "title": "Name of the artifact",
            "type": "string"
          },
          "path": {
            "description": "Relative path of the file/directory from the task directory. Note this is not an absolute\npath as is typically used in docker-worker, since the absolute task directory name is not\nknown when the task is submitted. Example: ` + "`" + `dist\\regedit.exe` + "`" + `. It doesn't matter if\nforward slashes or backslashes are used.\n\nSince: generic-worker 1.0.0",
            "title": "Artifact location",
            "type": "string"
          },
          "type": {
            "description": "Artifacts can be either an individual ` + "`" + `file` + "`" + ` or a ` + "`" + `directory` + "`" + ` containing\npotentially multiple files with recursively included subdirectories.\n\nSince: generic-worker 1.0.0",
            "enum": [
              "file",
              "directory"
            ],
            "title": "Artifact upload type.",
            "type": "string"
          }
        },
        "required": [
          "type",
          "path"
        ],
        "title": "Artifact",
        "type": "object"
      },
      "title": "Artifacts to be published",
      "type": "array",
      "uniqueItems": true
    },
    "command": {
      "description": "One array per command (each command is an array of arguments). Several arrays\nfor several commands.\n\nSince: generic-worker 0.0.1",
      "items": {
        "items": {
          "type": "string"
        },
        "minItems": 1,
        "type": "array",
        "uniqueItems": false
      },
      "minItems": 1,
      "title": "Commands to run",
      "type": "array",
      "uniqueItems": false
    },
    "env": {
      "additionalProperties": {
        "type": "string"
      },
      "description": "Env vars must be string to __string__ mappings (not number or boolean). For example:\n` + "`" + `` + "`" + `` + "`" + `\n{\n  \"PATH\": \"/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin\",\n  \"GOOS\": \"darwin\",\n  \"FOO_ENABLE\": \"true\",\n  \"BAR_TOTAL\": \"3\"\n}\n` + "`" + `` + "`" + `` + "`" + `\n\nNote, the following environment variables will automatically be set in the task\ncommands:\n  * ` + "`" + `TASK_ID` + "`" + ` - the task ID of the currently running task\n  * ` + "`" + `RUN_ID` + "`" + ` - the run ID of the currently running task\n  * ` + "`" + `TASKCLUSTER_ROOT_URL` + "`" + ` - the root URL of the taskcluster deployment\n  * ` + "`" + `TASKCLUSTER_PROXY_URL` + "`" + ` (if taskcluster proxy feature enabled) - the\n     taskcluster authentication proxy for making unauthenticated taskcluster\n     API calls\n  * ` + "`" + `TASKCLUSTER_WORKER_LOCATION` + "`" + `. See\n    [RFC #0148](https://github.com/taskcluster/taskcluster-rfcs/blob/master/rfcs/0148-taskcluster-worker-location.md)\n    for details.\n\nSince: generic-worker 0.0.1",
      "title": "Env vars",
      "type": "object"
    },
    "features": {
      "additionalProperties": false,
      "description": "Feature flags enable additional functionality.\n\nSince: generic-worker 5.3.0",
      "properties": {
        "backingLog": {
          "default": true,
          "description": "The backing log feature publishes a task artifact containing the complete\nstderr and stdout of the task.\n\nSince: generic-worker 48.2.0",
          "title": "Enable backing log",
          "type": "boolean"
        },
        "interactive": {
          "description": "This allows you to interactively run commands from within the worker\nas the task user. This may be useful for debugging purposes.\nCan be used for SSH-like access to the running worker.\nNote that this feature works differently from the ` + "`" + `interactive` + "`" + ` feature\nin docker worker, which ` + "`" + `docker exec` + "`" + `s into the running container.\nSince tasks on generic worker are not guaranteed to be running in a\ncontainer, a bash shell is started on the task user's account.\nA user can then ` + "`" + `docker exec` + "`" + ` into the a running container, if there\nis one.\n\nSince: generic-worker 49.2.0",
          "title": "Interactive shell",
          "type": "boolean"
        },
        "liveLog": {
          "default": true,
          "description": "The live log feature streams the combined stderr and stdout to a task artifact\nso that the output is available while the task is running.\n\nSince: generic-worker 48.2.0",
          "title": "Enable [livelog](https://github.com/taskcluster/taskcluster/tree/main/tools/livelog)",
          "type": "boolean"
        },
        "loopbackAudio": {
          "description": "Audio loopback device created using snd-aloop.\nAn audio device will be available for the task. Its\nlocation will be ` + "`" + `/dev/snd` + "`" + `. Devices inside that directory\nwill take the form ` + "`" + `/dev/snd/controlC\u003cN\u003e` + "`" + `,\n` + "`" + `/dev/snd/pcmC\u003cN\u003eD0c` + "`" + `, ` + "`" + `/dev/snd/pcmC\u003cN\u003eD0p` + "`" + `,\n` + "`" + `/dev/snd/pcmC\u003cN\u003eD1c` + "`" + `, and ` + "`" + `/dev/snd/pcmC\u003cN\u003eD1p` + "`" + `,\nwhere \u003cN\u003e is an integer between 0 and 31, inclusive.\nThe Generic Worker config setting ` + "`" + `loopbackAudioDeviceNumber` + "`" + `\nmay be used to change the device number in case the\ndefault value (` + "`" + `16` + "`" + `) conflicts with another\naudio device on the worker.\n\nThis feature is only available on Linux. If a task\nis submitted with this feature enabled on a non-Linux,\nposix platform (FreeBSD, macOS), the task will resolve as\n` + "`" + `exception/malformed-payload` + "`" + `.\n\nSince: generic-worker 54.5.0",
          "title": "Loopback Audio device",
          "type": "boolean"
        },
        "loopbackVideo": {
          "description": "Video loopback device created using v4l2loopback.\nA video device will be available for the task. Its\nlocation will be passed to the task via environment\nvariable ` + "`" + `TASKCLUSTER_VIDEO_DEVICE` + "`" + `. The\nlocation will be ` + "`" + `/dev/video\u003cN\u003e` + "`" + ` where ` + "`" + `\u003cN\u003e` + "`" + ` is\nan integer between 0 and 255. The value of ` + "`" + `\u003cN\u003e` + "`" + `\nis not static, and therefore either the environment\nvariable should be used, or ` + "`" + `/dev` + "`" + ` should be\nscanned in order to determine the correct location.\nTasks should not assume a constant value.\n\nThis feature is only available on Linux. If a task\nis submitted with this feature enabled on a non-Linux,\nposix platform (FreeBSD, macOS), the task will resolve as\n` + "`" + `exception/malformed-payload` + "`" + `.\n\nSince: generic-worker 53.1.0",
          "title": "Loopback Video device",
          "type": "boolean"
        },
        "taskclusterProxy": {
          "description": "The taskcluster proxy provides an easy and safe way to make authenticated\ntaskcluster requests within the scope(s) of a particular task. See\n[the github project](https://github.com/taskcluster/taskcluster/tree/main/tools/taskcluster-proxy) for more information.\n\nSince: generic-worker 10.6.0",
          "title": "Run [taskcluster-proxy](https://github.com/taskcluster/taskcluster/tree/main/tools/taskcluster-proxy) to allow tasks to dynamically proxy requests to taskcluster services",
          "type": "boolean"
        }
      },
      "required": [],
      "title": "Feature flags",
      "type": "object"
    },
    "logs": {
      "additionalProperties": false,
      "description": "Configuration for task logs.\n\nSince: generic-worker 48.2.0",
      "properties": {
        "backing": {
          "default": "public/logs/live_backing.log",
          "description": "Specifies a custom name for the backing log artifact.\nThis is only used if ` + "`" + `features.backingLog` + "`" + ` is ` + "`" + `true` + "`" + `.\n\nSince: generic-worker 48.2.0",
          "title": "Backing log artifact name",
          "type": "string"
        },
        "live": {
          "default": "public/logs/live.log",
          "description": "Specifies a custom name for the live log artifact.\nThis is only used if ` + "`" + `features.liveLog` + "`" + ` is ` + "`" + `true` + "`" + `.\n\nSince: generic-worker 48.2.0",
          "title": "Live log artifact name",
          "type": "string"
        }
      },
      "required": [],
      "title": "Logs",
      "type": "object"
    },
    "maxRunTime": {
      "description": "Maximum time the task container can run in seconds.\nThe maximum value for ` + "`" + `maxRunTime` + "`" + ` is set by a ` + "`" + `maxTaskRunTime` + "`" + ` config property specific to each worker-pool.\n\nSince: generic-worker 0.0.1",
      "minimum": 1,
      "multipleOf": 1,
      "title": "Maximum run time in seconds",
      "type": "integer"
    },
    "mounts": {
      "description": "Directories and/or files to be mounted.\n\nSince: generic-worker 5.4.0",
      "items": {
        "$ref": "#/definitions/mount",
        "title": "Mount"
      },
      "type": "array",
      "uniqueItems": false
    },
    "onExitStatus": {
      "additionalProperties": false,
      "description": "By default tasks will be resolved with ` + "`" + `state/reasonResolved` + "`" + `: ` + "`" + `completed/completed` + "`" + `\nif all task commands have a zero exit code, or ` + "`" + `failed/failed` + "`" + ` if any command has a\nnon-zero exit code. This payload property allows customsation of the task resolution\nbased on exit code of task commands.",
      "properties": {
        "purgeCaches": {
          "description": "If the task exits with a purge caches exit status, all caches\nassociated with the task will be purged.\n\nSince: generic-worker 49.0.0",
          "items": {
            "minimum": 0,
            "title": "Exit statuses",
            "type": "integer"
          },
          "title": "Purge caches exit status",
          "type": "array",
          "uniqueItems": true
        },
        "retry": {
          "description": "Exit codes for any command in the task payload to cause this task to\nbe resolved as ` + "`" + `exception/intermittent-task` + "`" + `. Typically the Queue\nwill then schedule a new run of the existing ` + "`" + `taskId` + "`" + ` (rerun) if not\nall task runs have been exhausted.\n\nSee [itermittent tasks](https://docs.taskcluster.net/docs/reference/platform/taskcluster-queue/docs/worker-interaction#intermittent-tasks) for more detail.\n\nSince: generic-worker 10.10.0",
          "items": {
            "minimum": 1,
            "title": "Exit codes",
            "type": "integer"
          },
          "title": "Intermittent task exit codes",
          "type": "array",
          "uniqueItems": true
        }
      },
      "required": [],
      "title": "Exit code handling",
      "type": "object"
    },
    "osGroups": {
      "description": "A list of OS Groups that the task user should be a member of. Not yet implemented on\nnon-Windows platforms, therefore this optional property may only be an empty array if\nprovided.\n\nSince: generic-worker 6.0.0",
      "items": {
        "type": "string"
      },
      "maxItems": 0,
      "title": "OS Groups",
      "type": "array",
      "uniqueItems": true
    },
    "supersederUrl": {
      "description": "This property is allowed for backward compatibility, but is unused.",
      "title": "unused",
      "type": "string"
    }
  },
  "required": [
    "command",
    "maxRunTime"
  ],
  "title": "Generic worker payload",
  "type": "object"
}`
}
