// Code generated by builtingen ../language/builtin.hlb ../builtin/builtin.go; DO NOT EDIT.

package builtin

import "github.com/openllb/hlb/parser"

type BuiltinLookup struct {
	ByType map[parser.ObjType]LookupByType
}

type LookupByType struct {
	Func map[string]FuncLookup
}

type FuncLookup struct {
	Params []*parser.Field
}

var (
	Lookup = BuiltinLookup{
		ByType: map[parser.ObjType]LookupByType{
			parser.Filesystem: LookupByType{
				Func: map[string]FuncLookup{
					"scratch": FuncLookup{
						Params: []*parser.Field{},
					},
					"image": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "ref", false),
						},
					},
					"http": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "url", false),
						},
					},
					"git": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "remote", false),
							parser.NewField(parser.Str, "ref", false),
						},
					},
					"local": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "path", false),
						},
					},
					"frontend": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "source", false),
						},
					},
					"shell": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "arg", true),
						},
					},
					"run": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "arg", true),
						},
					},
					"env": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "key", false),
							parser.NewField(parser.Str, "value", false),
						},
					},
					"dir": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "path", false),
						},
					},
					"user": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "name", false),
						},
					},
					"mkdir": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "path", false),
							parser.NewField(parser.Int, "filemode", false),
						},
					},
					"mkfile": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "path", false),
							parser.NewField(parser.Int, "filemode", false),
							parser.NewField(parser.Str, "content", false),
						},
					},
					"rm": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "path", false),
						},
					},
					"copy": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Filesystem, "input", false),
							parser.NewField(parser.Str, "src", false),
							parser.NewField(parser.Str, "dst", false),
						},
					},
					"dockerPush": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "ref", false),
						},
					},
					"dockerLoad": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "ref", false),
						},
					},
					"download": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "localPath", false),
						},
					},
					"downloadTarball": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "localPath", false),
						},
					},
					"downloadOCITarball": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "localPath", false),
						},
					},
					"downloadDockerTarball": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "localPath", false),
							parser.NewField(parser.Str, "ref", false),
						},
					},
				},
			},
			"option::copy": LookupByType{
				Func: map[string]FuncLookup{
					"followSymlinks": FuncLookup{
						Params: []*parser.Field{},
					},
					"contentsOnly": FuncLookup{
						Params: []*parser.Field{},
					},
					"unpack": FuncLookup{
						Params: []*parser.Field{},
					},
					"createDestPath": FuncLookup{
						Params: []*parser.Field{},
					},
				},
			},
			"option::frontend": LookupByType{
				Func: map[string]FuncLookup{
					"input": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "key", false),
							parser.NewField(parser.Filesystem, "value", false),
						},
					},
					"opt": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "key", false),
							parser.NewField(parser.Str, "value", false),
						},
					},
				},
			},
			"option::git": LookupByType{
				Func: map[string]FuncLookup{
					"keepGitDir": FuncLookup{
						Params: []*parser.Field{},
					},
				},
			},
			"option::http": LookupByType{
				Func: map[string]FuncLookup{
					"checksum": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "digest", false),
						},
					},
					"chmod": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Int, "filemode", false),
						},
					},
					"filename": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "name", false),
						},
					},
				},
			},
			"option::image": LookupByType{
				Func: map[string]FuncLookup{
					"resolve": FuncLookup{
						Params: []*parser.Field{},
					},
				},
			},
			"option::local": LookupByType{
				Func: map[string]FuncLookup{
					"includePatterns": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "pattern", true),
						},
					},
					"excludePatterns": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "pattern", true),
						},
					},
					"followPaths": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "path", true),
						},
					},
				},
			},
			"option::mkdir": LookupByType{
				Func: map[string]FuncLookup{
					"createParents": FuncLookup{
						Params: []*parser.Field{},
					},
					"chown": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "owner", false),
						},
					},
					"createdTime": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "created", false),
						},
					},
				},
			},
			"option::mkfile": LookupByType{
				Func: map[string]FuncLookup{
					"chown": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "owner", false),
						},
					},
					"createdTime": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "created", false),
						},
					},
				},
			},
			"option::mount": LookupByType{
				Func: map[string]FuncLookup{
					"readonly": FuncLookup{
						Params: []*parser.Field{},
					},
					"tmpfs": FuncLookup{
						Params: []*parser.Field{},
					},
					"sourcePath": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "path", false),
						},
					},
					"cache": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "cacheid", false),
							parser.NewField(parser.Str, "sharingmode", false),
						},
					},
				},
			},
			"option::rm": LookupByType{
				Func: map[string]FuncLookup{
					"allowNotFound": FuncLookup{
						Params: []*parser.Field{},
					},
					"allowWildcards": FuncLookup{
						Params: []*parser.Field{},
					},
				},
			},
			"option::run": LookupByType{
				Func: map[string]FuncLookup{
					"readonlyRootfs": FuncLookup{
						Params: []*parser.Field{},
					},
					"env": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "key", false),
							parser.NewField(parser.Str, "value", false),
						},
					},
					"dir": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "path", false),
						},
					},
					"user": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "name", false),
						},
					},
					"network": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "networkmode", false),
						},
					},
					"security": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "securitymode", false),
						},
					},
					"host": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "hostname", false),
							parser.NewField(parser.Str, "address", false),
						},
					},
					"ssh": FuncLookup{
						Params: []*parser.Field{},
					},
					"secret": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "localPath", false),
							parser.NewField(parser.Str, "mountPoint", false),
						},
					},
					"mount": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Filesystem, "input", false),
							parser.NewField(parser.Str, "mountPoint", false),
						},
					},
				},
			},
			"option::secret": LookupByType{
				Func: map[string]FuncLookup{
					"uid": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Int, "id", false),
						},
					},
					"gid": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Int, "id", false),
						},
					},
					"mode": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Int, "filemode", false),
						},
					},
				},
			},
			"option::ssh": LookupByType{
				Func: map[string]FuncLookup{
					"target": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "mountPoint", false),
						},
					},
					"localPaths": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "path", true),
						},
					},
					"uid": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Int, "id", false),
						},
					},
					"gid": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Int, "id", false),
						},
					},
					"mode": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Int, "filemode", false),
						},
					},
				},
			},
			parser.Str: LookupByType{
				Func: map[string]FuncLookup{
					"format": FuncLookup{
						Params: []*parser.Field{
							parser.NewField(parser.Str, "formatString", false),
							parser.NewField(parser.Str, "values", true),
						},
					},
				},
			},
		},
	}
)
