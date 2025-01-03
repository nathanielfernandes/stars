package github

import "strings"

type LanguageIdentifier struct {
	Aliases    []string
	Extensions []string
	Filenames  []string
	Color      string
}

func (id LanguageIdentifier) Matches(fn string) bool {
	for _, filename := range id.Filenames {
		if fn == filename {
			return true
		}
	}

	for _, ext := range id.Extensions {
		if ext != "" && strings.HasSuffix(fn, ext) {
			return true
		}
	}

	for _, alias := range id.Aliases {
		if alias == fn {
			return true
		}
	}

	return false
}

func GetByAlias(alias string) (LanguageIdentifier, bool) {
	for _, id := range LANGUAGE_IDENTIFIERS {
		for _, a := range id.Aliases {
			if a == alias {
				return id, true
			}
		}
	}

	return LanguageIdentifier{}, false
}

func AliasMatches(alias string, fn string) bool {
	id, ok := GetByAlias(alias)
	if !ok {
		return alias == fn
	}

	return id.Matches(fn)
}

func GetAlias(lang string) string {
	return LANGUAGE_IDENTIFIERS[lang].Aliases[0]
}

func GetColor(lang string) string {
	return LANGUAGE_IDENTIFIERS[lang].Color

}

func GetColorFromAlias(alias string) string {
	for _, id := range LANGUAGE_IDENTIFIERS {
		for _, a := range id.Aliases {
			if a == alias {
				return id.Color
			}
		}
	}

	return "#ffffff"
}

var /* const */ LANGUAGE_IDENTIFIERS = map[string]LanguageIdentifier{
	"bat": {
		[]string{
			"Batch",
			"bat",
		},
		[]string{
			".bat",
			".cmd",
		},
		[]string{
			"",
		},
		"#C1F12E",
	},
	"clojure": {
		[]string{
			"Clojure",
			"clojure",
		},
		[]string{
			".clj",
			".cljs",
			".cljc",
			".cljx",
			".clojure",
			".edn",
		},
		[]string{
			"",
		},
		"#db5855",
	},
	"coffeescript": {
		[]string{
			"CoffeeScript",
			"coffeescript",
			"coffee",
		},
		[]string{
			".coffee",
			".cson",
			".iced",
		},
		[]string{
			"",
		},
		"#244776",
	},
	"jsonc": {
		[]string{
			"JSON with Comments",
		},
		[]string{
			".code-workspace",
			"language-configuration.json",
			"icon-theme.json",
			"color-theme.json",
			".code-snippets",
			".jsonc",
			".eslintrc",
			".eslintrc.json",
			".jsfmtrc",
			".jshintrc",
			".swcrc",
			".hintrc",
			".babelrc",
		},
		[]string{
			"settings.json",
			"launch.json",
			"tasks.json",
			"keybindings.json",
			"extensions.json",
			"argv.json",
			"profiles.json",
			".devcontainer.json",
			"babel.config.json",
			".babelrc.json",
			".ember-cli",
			"tsconfig.json",
			"jsconfig.json",
		},
		"#292929",
	},
	"c": {
		[]string{
			"C",
			"c",
		},
		[]string{
			".c",
			".i",
		},
		[]string{
			"",
		},
		"#555555",
	},
	"cpp": {
		[]string{
			"C++",
			"Cpp",
			"cpp",
		},
		[]string{
			".cpp",
			".cc",
			".cxx",
			".c++",
			".hpp",
			".hh",
			".hxx",
			".h++",
			".h",
			".ii",
			".ino",
			".inl",
			".ipp",
			".ixx",
			".tpp",
			".txx",
			".hpp.in",
			".h.in",
		},
		[]string{
			"",
		},
		"#f34b7d",
	},
	"cuda-cpp": {
		[]string{
			"CUDA C++",
			"Cuda",
		},
		[]string{
			".cu",
			".cuh",
		},
		[]string{
			"",
		},
		"#3A4E3A",
	},
	"csharp": {
		[]string{
			"C#",
			"csharp",
		},
		[]string{
			".cs",
			".csx",
			".cake",
		},
		[]string{
			"",
		},
		"#178600",
	},
	"css": {
		[]string{
			"CSS",
			"css",
		},
		[]string{
			".css",
		},
		[]string{
			"",
		},
		"#563d7c",
	},
	"dart": {
		[]string{
			"Dart",
		},
		[]string{
			".dart",
		},
		[]string{
			"",
		},
		"#00B4AB",
	},
	"diff": {
		[]string{
			"Diff",
			"diff",
		},
		[]string{
			".diff",
			".patch",
			".rej",
		},
		[]string{
			"",
		},
		"#ba595e",
	},
	"dockerfile": {
		[]string{
			"Docker",
			"Dockerfile",
			"Containerfile",
		},
		[]string{
			".dockerfile",
			".containerfile",
		},
		[]string{
			"Dockerfile",
			"Containerfile",
		},
		"#384d54",
	},
	"ignore": {
		[]string{
			"Ignore",
			"ignore",
		},
		[]string{
			".gitignore_global",
			".gitignore",
			".npmignore",
		},
		[]string{
			".vscodeignore",
			".prettierignore",
			".dockerignore",
		},
		"#000000",
	},
	"fsharp": {
		[]string{
			"F#",
			"FSharp",
			"fsharp",
		},
		[]string{
			".fs",
			".fsi",
			".fsx",
			".fsscript",
		},
		[]string{
			"",
		},
		"#b845fc",
	},
	"git-commit": {
		[]string{
			"Git Commit Message",
			"git-commit",
		},
		[]string{
			"",
		},
		[]string{
			"COMMIT_EDITMSG",
			"MERGE_MSG",
		},
		"#F44D27",
	},
	"git-rebase": {
		[]string{
			"Git Rebase Message",
			"git-rebase",
		},
		[]string{
			"",
		},
		[]string{
			"git-rebase-todo",
		},
		"#F44D27",
	},
	"go": {
		[]string{
			"Go",
		},
		[]string{
			".go",
		},
		[]string{
			"",
		},
		"#00ADD8",
	},
	"groovy": {
		[]string{
			"Groovy",
			"groovy",
		},
		[]string{
			".groovy",
			".gvy",
			".gradle",
			".jenkinsfile",
			".nf",
		},
		[]string{
			"Jenkinsfile",
		},
		"#4298b8",
	},
	"handlebars": {
		[]string{
			"Handlebars",
			"handlebars",
		},
		[]string{
			".handlebars",
			".hbs",
			".hjs",
		},
		[]string{
			"",
		},
		"#f7931e",
	},
	"hlsl": {
		[]string{
			"HLSL",
			"hlsl",
		},
		[]string{
			".hlsl",
			".hlsli",
			".fx",
			".fxh",
			".vsh",
			".psh",
			".cginc",
			".compute",
			".sf",
			".hs",
			".ds",
			".ps",
		},
		[]string{
			"",
		},
		"#aace60",
	},
	"html": {
		[]string{
			"HTML",
			"htm",
			"html",
			"xhtml",
		},
		[]string{
			".html",
			".htm",
			".shtml",
			".xhtml",
			".xht",
			".mdoc",
			".jsp",
			".asp",
			".aspx",
			".jshtm",
			".volt",
			".ejs",
			".rhtml",
		},
		[]string{
			"",
		},
		"#e34c26",
	},
	"ini": {
		[]string{
			"Ini",
			"ini",
		},
		[]string{
			".ini",
		},
		[]string{
			".flake8",
			".pep8",
			".pylintrc",
			".pypirc",
		},
		"#d1dbe0",
	},
	"properties": {
		[]string{
			"Properties",
			"properties",
		},
		[]string{
			".properties",
			".cfg",
			".conf",
			".directory",
			".gitattributes",
			".gitconfig",
			".gitmodules",
			".editorconfig",
			".npmrc",
		},
		[]string{
			"gitconfig",
		},
		"#F44D27",
	},
	"jupyter": {
		[]string{
			"Jupyter (JSON)",
		},
		[]string{
			".ipynb",
		},
		[]string{
			"",
		},
		"#DA5B0B",
	},
	"java": {
		[]string{
			"Java",
			"java",
		},
		[]string{
			".java",
			".jav",
		},
		[]string{
			"",
		},
		"#b07219",
	},
	"javascriptreact": {
		[]string{
			"JavaScript React",
			"jsx",
		},
		[]string{
			".jsx",
		},
		[]string{
			"",
		},
		"#f1e05a",
	},
	"javascript": {
		[]string{
			"JavaScript",
			"javascript",
			"js",
		},
		[]string{
			".js",
			".es6",
			".mjs",
			".cjs",
			".pac",
		},
		[]string{
			"jakefile",
		},
		"#f1e05a",
	},
	"json": {
		[]string{
			"JSON",
			"json",
		},
		[]string{
			".json",
			".bowerrc",
			".jscsrc",
			".webmanifest",
			".js.map",
			".css.map",
			".ts.map",
			".har",
			".jslintrc",
			".jsonld",
			".geojson",
		},
		[]string{
			"composer.lock",
			".watchmanconfig",
			".prettierrc",
			"Pipfile.lock",
		},
		"#292929",
	},
	"julia": {
		[]string{
			"Julia",
			"julia",
		},
		[]string{
			".jl",
		},
		[]string{
			"",
		},
		"#a270ba",
	},
	"juliamarkdown": {
		[]string{
			"Julia Markdown",
			"juliamarkdown",
		},
		[]string{
			".jmd",
		},
		[]string{
			"",
		},
		"#083fa1",
	},
	"tex": {
		[]string{
			"TeX",
			"tex",
		},
		[]string{
			".sty",
			".cls",
			".bbx",
			".cbx",
		},
		[]string{
			"",
		},
		"#3D6117",
	},
	"latex": {
		[]string{
			"LaTeX",
			"latex",
		},
		[]string{
			".tex",
			".ltx",
			".ctx",
		},
		[]string{
			"",
		},
		"#3D6117",
	},
	"bibtex": {
		[]string{
			"BibTeX",
			"bibtex",
		},
		[]string{
			".bib",
		},
		[]string{
			"",
		},
		"#778899",
	},
	"cpp_embedded_latex": {
		[]string{
			"",
		},
		[]string{
			"",
		},
		[]string{
			"",
		},
		"#3D6117",
	},
	"markdown_latex_combined": {
		[]string{
			"",
		},
		[]string{
			"",
		},
		[]string{
			"",
		},
		"#3D6117",
	},
	"less": {
		[]string{
			"Less",
			"less",
		},
		[]string{
			".less",
		},
		[]string{
			"",
		},
		"#1d365d",
	},
	"log": {
		[]string{
			"Log",
		},
		[]string{
			".log",
			"*.log.?",
		},
		[]string{
			"",
		},
		"#295b9a",
	},
	"lua": {
		[]string{
			"Lua",
			"lua",
		},
		[]string{
			".lua",
		},
		[]string{
			"",
		},
		"#000080",
	},
	"makefile": {
		[]string{
			"Makefile",
			"makefile",
		},
		[]string{
			".mak",
			".mk",
		},
		[]string{
			"Makefile",
			"makefile",
			"GNUmakefile",
			"OCamlMakefile",
		},
		"#427819",
	},
	"markdown": {
		[]string{
			"Markdown",
			"markdown",
		},
		[]string{
			".md",
			".mkd",
			".mdwn",
			".mdown",
			".markdown",
			".markdn",
			".mdtxt",
			".mdtext",
			".workbook",
		},
		[]string{
			"",
		},
		"#083fa1",
	},
	"markdown-math": {
		[]string{
			"",
		},
		[]string{
			"",
		},
		[]string{
			"",
		},
		"#083fa1",
	},
	"objective-c": {
		[]string{
			"Objective-C",
		},
		[]string{
			".m",
		},
		[]string{
			"",
		},
		"#438eff",
	},
	"objective-cpp": {
		[]string{
			"Objective-C++",
		},
		[]string{
			".mm",
		},
		[]string{
			"",
		},
		"#6866fb",
	},
	"perl": {
		[]string{
			"Perl",
			"perl",
		},
		[]string{
			".pl",
			".pm",
			".pod",
			".t",
			".PL",
			".psgi",
		},
		[]string{
			"",
		},
		"#0298c3",
	},
	"perl6": {
		[]string{
			"Perl 6",
			"perl6",
		},
		[]string{
			".p6",
			".pl6",
			".pm6",
			".nqp",
		},
		[]string{
			"",
		},
		"#0298c3",
	},
	"php": {
		[]string{
			"PHP",
			"php",
		},
		[]string{
			".php",
			".php4",
			".php5",
			".phtml",
			".ctp",
		},
		[]string{
			"",
		},
		"#4F5D95",
	},
	"powershell": {
		[]string{
			"PowerShell",
			"powershell",
			"ps",
			"ps1",
		},
		[]string{
			".ps1",
			".psm1",
			".psd1",
			".pssc",
			".psrc",
		},
		[]string{
			"",
		},
		"#012456",
	},
	"jade": {
		[]string{
			"Pug",
			"Jade",
			"jade",
		},
		[]string{
			".pug",
			".jade",
		},
		[]string{
			"",
		},
		"#a86454",
	},
	"python": {
		[]string{
			"Python",
			"py",
		},
		[]string{
			".py",
			".rpy",
			".pyw",
			".cpy",
			".gyp",
			".gypi",
			".pyi",
			".ipy",
			".pyt",
		},
		[]string{
			"Snakefile",
			"SConstruct",
			"SConscript",
		},
		"#3572A5",
	},
	"r": {
		[]string{
			"R",
			"r",
		},
		[]string{
			".r",
			".rhistory",
			".rprofile",
			".rt",
		},
		[]string{
			"",
		},
		"#198CE7",
	},
	"razor": {
		[]string{
			"Razor",
			"razor",
		},
		[]string{
			".cshtml",
		},
		[]string{
			"",
		},
		"#512be4",
	},
	"ruby": {
		[]string{
			"Ruby",
			"rb",
		},
		[]string{
			".rb",
			".rbx",
			".rjs",
			".gemspec",
			".rake",
			".ru",
			".erb",
			".podspec",
			".rbi",
		},
		[]string{
			"rakefile",
			"gemfile",
			"guardfile",
			"podfile",
			"capfile",
			"cheffile",
			"hobofile",
			"vagrantfile",
			"appraisals",
			"rantfile",
			"berksfile",
			"berksfile.lock",
			"thorfile",
			"puppetfile",
			"dangerfile",
			"brewfile",
			"fastfile",
			"appfile",
			"deliverfile",
			"matchfile",
			"scanfile",
			"snapfile",
			"gymfile",
		},
		"#701516",
	},
	"rust": {
		[]string{
			"Rust",
			"rust",
			"rs",
		},
		[]string{
			".rs",
		},
		[]string{
			"",
		},
		"#dea584",
	},
	"scss": {
		[]string{
			"SCSS",
			"scss",
		},
		[]string{
			".scss",
		},
		[]string{
			"",
		},
		"#c6538c",
	},
	"search-result": {
		[]string{
			"Search Result",
		},
		[]string{
			".code-search",
		},
		[]string{
			"",
		},
		"#814CCC",
	},
	"shaderlab": {
		[]string{
			"ShaderLab",
			"shaderlab",
		},
		[]string{
			".shader",
		},
		[]string{
			"",
		},
		"#222c37",
	},
	"shellscript": {
		[]string{
			"Shell Script",
			"shellscript",
			"bash",
			"sh",
			"zsh",
			"ksh",
			"csh",
		},
		[]string{
			".sh",
			".bash",
			".bashrc",
			".bash_aliases",
			".bash_profile",
			".bash_login",
			".ebuild",
			".profile",
			".bash_logout",
			".xprofile",
			".xsession",
			".xsessionrc",
			".Xsession",
			".zsh",
			".zshrc",
			".zprofile",
			".zlogin",
			".zlogout",
			".zshenv",
			".zsh-theme",
			".ksh",
			".csh",
			".cshrc",
			".tcshrc",
			".yashrc",
			".yash_profile",
		},
		[]string{
			"APKBUILD",
			"PKGBUILD",
			".envrc",
			".hushlogin",
			"zshrc",
			"zshenv",
			"zlogin",
			"zprofile",
			"zlogout",
			"bashrc_Apple_Terminal",
			"zshrc_Apple_Terminal",
		},
		"#89e051",
	},
	"sql": {
		[]string{
			"SQL",
			"PLpgSQL",
		},
		[]string{
			".sql",
			".dsql",
		},
		[]string{
			"",
		},
		"#e38c00",
	},
	"swift": {
		[]string{
			"Swift",
			"swift",
		},
		[]string{
			".swift",
		},
		[]string{
			"",
		},
		"#F05138",
	},
	"typescript": {
		[]string{
			"TypeScript",
			"ts",
			"typescript",
		},
		[]string{
			".ts",
			".cts",
			".mts",
		},
		[]string{
			"",
		},
		"#2b7489",
	},
	"typescriptreact": {
		[]string{
			"TypeScript React",
			"tsx",
		},
		[]string{
			".tsx",
		},
		[]string{
			"",
		},
		"#2b7489",
	},
	"vb": {
		[]string{
			"Visual Basic",
			"vb",
		},
		[]string{
			".vb",
			".brs",
			".vbs",
			".bas",
			".vba",
		},
		[]string{
			"",
		},
		"#15dcdc",
	},
	"xml": {
		[]string{
			"XML",
			"xml",
		},
		[]string{
			".xml",
			".xsd",
			".ascx",
			".atom",
			".axml",
			".axaml",
			".bpmn",
			".cpt",
			".csl",
			".csproj",
			".csproj.user",
			".dita",
			".ditamap",
			".dtd",
			".ent",
			".mod",
			".dtml",
			".fsproj",
			".fxml",
			".iml",
			".isml",
			".jmx",
			".launch",
			".menu",
			".mxml",
			".nuspec",
			".opml",
			".owl",
			".proj",
			".props",
			".pt",
			".publishsettings",
			".pubxml",
			".pubxml.user",
			".rbxlx",
			".rbxmx",
			".rdf",
			".rng",
			".rss",
			".shproj",
			".storyboard",
			".svg",
			".targets",
			".tld",
			".tmx",
			".vbproj",
			".vbproj.user",
			".vcxproj",
			".vcxproj.filters",
			".wsdl",
			".wxi",
			".wxl",
			".wxs",
			".xaml",
			".xbl",
			".xib",
			".xlf",
			".xliff",
			".xpdl",
			".xul",
			".xoml",
		},
		[]string{
			"",
		},
		"#0060ac",
	},
	"xsl": {
		[]string{
			"XSL",
			"xsl",
		},
		[]string{
			".xsl",
			".xslt",
		},
		[]string{
			"",
		},
		"#EB8CEB",
	},
	"dockercompose": {
		[]string{
			"Compose",
			"compose",
		},
		[]string{
			"",
		},
		[]string{
			"",
		},
		"#384d54",
	},
	"yaml": {
		[]string{
			"YAML",
			"yaml",
		},
		[]string{
			".yml",
			".eyaml",
			".eyml",
			".yaml",
			".cff",
		},
		[]string{
			".condarc",
		},
		"#cb171e",
	},
	"toml": {
		[]string{
			"TOML",
			"toml",
		},
		[]string{
			".toml",
		},
		[]string{
			"Pipfile",
			"poetry.lock",
		},
		"#9c4221",
	},
	"graphql": {
		[]string{
			"",
		},
		[]string{
			".graphql",
		},
		[]string{
			"",
		},
		"#e10098",
	},
	"vue": {
		[]string{
			"",
		},
		[]string{
			".vue",
		},
		[]string{
			"",
		},
		"#41b883",
	},
	"go.mod": {
		[]string{
			"Go Module File",
		},
		[]string{
			"",
		},
		[]string{
			"go.mod",
			"gopls.mod",
		},
		"#88562A",
	},
	"go.work": {
		[]string{
			"Go Work File",
		},
		[]string{
			"",
		},
		[]string{
			"go.work",
		},
		"#00ADD8",
	},
	"go.sum": {
		[]string{
			"Go Checksum File",
		},
		[]string{
			"",
		},
		[]string{
			"go.sum",
		},
		"#82937f",
	},
	"gotmpl": {
		[]string{
			"Go Template File",
		},
		[]string{
			".tmpl",
			".gotmpl",
		},
		[]string{
			"",
		},
		"#00ADD8",
	},
	"ra_syntax_tree": {
		[]string{
			"",
		},
		[]string{
			".rast",
		},
		[]string{
			"",
		},
		"#358a5b",
	},
	"jinja": {
		[]string{
			"Jinja",
		},
		[]string{
			".j2",
			".jinja2",
		},
		[]string{
			"",
		},
		"#a52a22",
	},
	"pip-requirements": {
		[]string{
			"pip requirements",
			"requirements.txt",
		},
		[]string{
			"",
		},
		[]string{
			"constraints.txt",
			"requirements.in",
			"requirements.txt",
		},
		"#000000",
	},
	"commonlisp": {
		[]string{
			"Common Lisp",
			"common lisp",
			"commonlisp",
		},
		[]string{
			".lisp",
			".lsp",
			".l",
			".cl",
			".asd",
			".asdf",
		},
		[]string{
			"",
		},
		"#3fb68b",
	},
	"glsl": {
		[]string{
			"GLSL",
			"OpenGL Shading Language",
			"glsl",
		},
		[]string{
			".vs",
			".fs",
			".gs",
			".comp",
			".vert",
			".tesc",
			".tese",
			".frag",
			".geom",
			".glsl",
			".glslv",
			".glslf",
			".glslg",
			".mesh",
			".task",
			".rgen",
			".rint",
			".rahit",
			".rchit",
			".rmiss",
			".rcall",
		},
		[]string{
			"",
		},
		"#5686a5",
	},
	"cg": {
		[]string{
			"Cg",
			"C for Graphics",
			"cg",
		},
		[]string{
			".cg",
		},
		[]string{
			"",
		},
		"#555555",
	},
	"svelte": {
		[]string{
			"Svelte",
			"svelte",
		},
		[]string{
			".svelte",
		},
		[]string{
			"",
		},
		"#ff3e00",
	},
	"elixir": {
		[]string{
			"Elixir",
			"elixir",
		},
		[]string{
			".exs",
			".ex",
			".heex",
		},
		[]string{
			"",
		},
		"#9542f5",
	},
	"quilt": {
		[]string{
			"CodeQL",
			"Quilt",
			"quilt",
		},
		[]string{
			".ql",
		},
		[]string{
			"",
		},
		"#d66eff",
	},
}
