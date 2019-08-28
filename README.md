# TOP Go Frameworks

- [TOP Go Frameworks](#top-go-frameworks)
	- [Command Line](#command-line)
	- [Console UI](#console-ui)
		- [Console UI Engine/Library](#console-ui-enginelibrary)
	- [GUI Application](#gui-application)
	- [Logging](#logging)
	- [Object-Relational Mapping](#object-relational-mapping)
	- [Web Frameworks](#web-frameworks)
	- [Game](#game)
		- [Game engine](#game-engine)
	- [Go E-Books](#go-e-books)

## Command Line

|                        Repo                        |                    Stars                    |                  Forks                   |                                   Description                                    |
| -------------------------------------------------- | :-----------------------------------------: | :--------------------------------------: | -------------------------------------------------------------------------------- |
| [cobra](https://github.com/spf13/cobra)            | <span style="color:yellow">**13623**</span> | <span style="color:blue">**1163**</span> | A Commander for modern Go CLI interactions                                       |
| [cli](https://github.com/urfave/cli)               |                  **11684**                  |                 **941**                  | A simple, fast, and fun package for building command line apps in Go             |
| [kingpin](https://github.com/alecthomas/kingpin)   |                  **2596**                   |                 **193**                  | A Go (golang) command line and flag parser                                       |
| [go-flags](https://github.com/jessevdk/go-flags)   |                  **1520**                   |                 **179**                  | go command line option parser                                                    |
| [readline](https://github.com/chzyer/readline)     |                  **1381**                   |                 **145**                  | Readline is a pure go(golang) implementation for GNU-Readline kind library       |
| [docopt.go](https://github.com/docopt/docopt.go)   |                  **1176**                   |                  **98**                  | A command-line arguments parser that will make you smile.                        |
| [cli](https://github.com/mitchellh/cli)            |                  **1004**                   |                  **76**                  | A Go library for implementing command-line interfaces.                           |
| [pflag](https://github.com/spf13/pflag)            |                   **775**                   |                 **175**                  | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. |
| [go-arg](https://github.com/alexflint/go-arg)      |                   **753**                   |                  **42**                  | Struct-based argument parsing in Go                                              |
| [mow.cli](https://github.com/jawher/mow.cli)       |                   **626**                   |                  **44**                  | A versatile library for building CLI applications in Go                          |
| [commandeer](https://github.com/jaffee/commandeer) |                   **93**                    |                  **5**                   | Automatically sets up command line flags based on struct fields and tags.        |

## Console UI

|                           Repo                            |   Stars   |  Forks  |                                                                     Description                                                                     |
| --------------------------------------------------------- | :-------: | :-----: | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| [lazygit](https://github.com/jesseduffield/lazygit)       | **13255** | **436** | simple terminal UI for git commands                                                                                                                 |
| [lazydocker](https://github.com/jesseduffield/lazydocker) | **11305** | **392** | The lazier way to manage everything docker                                                                                                          |
| [vuls](https://github.com/future-architect/vuls)          | **6779**  | **720** | Agent-less vulnerability scanner for Linux, FreeBSD, Container Image, Running Container, WordPress, Programming language libraries, Network devices |
| [jid](https://github.com/simeji/jid)                      | **5297**  | **110** | json incremental digger                                                                                                                             |
| [httplab](https://github.com/gchaincl/httplab)            | **3424**  | **120** | The interactive web server                                                                                                                          |
| [color](https://github.com/fatih/color)                   | **3045**  | **355** | Color package for Go (golang)                                                                                                                       |
| [dry](https://github.com/moncho/dry)                      | **1836**  | **77**  | dry - A Docker manager for the terminal @                                                                                                           |
| [fac](https://github.com/mkchoi212/fac)                   | **1614**  | **33**  | Easy-to-use CUI for fixing git conflicts                                                                                                            |
| [lf](https://github.com/gokcehan/lf)                      | **1290**  | **56**  | Terminal file manager                                                                                                                               |
| [asciigraph](https://github.com/guptarohit/asciigraph)    | **1158**  | **38**  | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.                                                |
| [cointop](https://github.com/miguelmota/cointop)          | **1117**  | **82**  | The fastest and most interactive terminal based UI application for tracking cryptocurrencies                                                        |
| [rat](https://github.com/ericfreese/rat)                  | **1098**  | **34**  | Compose shell commands to build interactive terminal applications                                                                                   |
| [mop](https://github.com/mop-tracker/mop)                 | **1019**  | **201** | Stock market tracker for hackers                                                                                                                    |
| [mpb](https://github.com/vbauerster/mpb)                  |  **724**  | **44**  | multi progress bar for Go cli applications                                                                                                          |
| [aurora](https://github.com/logrusorgru/aurora)           |  **650**  | **23**  | Golang ultimate ANSI-colors that supports Printf/Sprintf methods                                                                                    |
| [pxl](https://github.com/ichinaski/pxl)                   |  **620**  | **27**  | Display images in the terminal                                                                                                                      |
| [progressbar](https://github.com/schollz/progressbar)     |  **589**  | **37**  | A really basic thread-safe progress bar for Golang applications                                                                                     |
| [diagram](https://github.com/esimov/diagram)              |  **519**  | **14**  | CLI app to convert ascii arts into hand drawn diagrams.                                                                                             |
| [clui](https://github.com/VladimirMarkelov/clui)          |  **394**  | **33**  | Command Line User Interface (Console UI inspired by TurboVision)                                                                                    |
| [gomainr](https://github.com/MichaelThessel/gomainr)      |  **122**  | **11**  | Terminal cli app that checks the availability of domains for different configurations of keywords.                                                  |

### Console UI Engine/Library

|                        Repo                        |  Stars   |  Forks  |                                   Description                                   |
| -------------------------------------------------- | :------: | :-----: | ------------------------------------------------------------------------------- |
| [termui](https://github.com/gizak/termui)          | **9006** | **560** | Golang terminal dashboard                                                       |
| [gocui](https://github.com/jroimartin/gocui)       | **5469** | **319** | Minimalist Go package aimed at creating Console User Interfaces.                |
| [termbox-go](https://github.com/nsf/termbox-go)    | **3511** | **291** | Pure Go termbox implementation                                                  |
| [tview](https://github.com/rivo/tview)             | **3226** | **189** | Rich interactive widgets for terminal-based UIs written in Go                   |
| [go-prompt](https://github.com/c-bata/go-prompt)   | **2438** | **122** | Building powerful interactive prompts in Go, inspired by python-prompt-toolkit. |
| [tui-go](https://github.com/marcusolsson/tui-go)   | **1717** | **89**  | A UI library for terminal applications.                                         |
| [uiprogress](https://github.com/gosuri/uiprogress) | **1548** | **89**  | A go library to render progress bars in terminal applications                   |
| [gcli](https://github.com/tcnksm/gcli)             | **872**  | **76**  | The easy way to build Golang command-line application.                          |
| [uilive](https://github.com/gosuri/uilive)         | **853**  | **48**  | uilive is a go library for updating terminal output in realtime                 |
| [termdash](https://github.com/mum4k/termdash)      | **820**  | **34**  | Terminal based dashboard.                                                       |
| [uitable](https://github.com/gosuri/uitable)       | **506**  | **20**  | A go library to improve readability in terminal apps using tabular data         |

## GUI Application

|                             Repo                             |  Stars   |  Forks  |                                                                                    Description                                                                                    |
| ------------------------------------------------------------ | :------: | :-----: | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [ui](https://github.com/andlabs/ui)                          | **7026** | **616** | Platform-native GUI library for Go.                                                                                                                                               |
| [fyne](https://github.com/fyne-io/fyne)                      | **6449** | **251** | Cross platform GUI in Go based on Material Design                                                                                                                                 |
| [qt](https://github.com/therecipe/qt)                        | **6159** | **448** | Qt binding for Go (Golang) with support for Windows / macOS / Linux / FreeBSD / Android / iOS / Sailfish OS / Raspberry Pi / AsteroidOS / Ubuntu Touch / JavaScript / WebAssembly |
| [webview](https://github.com/zserge/webview)                 | **4716** | **344** | Tiny cross-platform webview library for C/C++/Golang. Uses WebKit (Gtk/Cocoa) and MSHTML (Windows)                                                                                |
| [robotgo](https://github.com/go-vgo/robotgo)                 | **4455** | **397** | RobotGo, Go Native cross-platform GUI automation                                                                                                                                  |
| [walk](https://github.com/lxn/walk)                          | **3729** | **569** | A Windows GUI toolkit for the Go Programming Language                                                                                                                             |
| [app](https://github.com/maxence-charriere/app)              | **2971** | **126** | A WebAssembly framework to build GUI with Go, HTML and CSS.                                                                                                                       |
| [go-astilectron](https://github.com/asticode/go-astilectron) | **2710** | **180** | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron)                                                                                                       |
| [go-gtk](https://github.com/mattn/go-gtk)                    | **1470** | **206** | Go binding for GTK                                                                                                                                                                |
| [go-sciter](https://github.com/sciter-sdk/go-sciter)         | **1461** | **175** | Golang bindings of Sciter: the Embeddable HTML/CSS/script engine for modern UI development                                                                                        |
| [wails](https://github.com/wailsapp/wails)                   | **926**  | **50**  | Create desktop apps using Go and Web Technologies                                                                                                                                 |
| [systray](https://github.com/getlantern/systray)             | **798**  | **119** | a cross platfrom Go library to place an icon and menu in the notification area                                                                                                    |
| [gosx-notifier](https://github.com/deckarep/gosx-notifier)   | **496**  | **36**  | gosx-notifier is a Go framework for sending desktop notifications to OSX 10.8 or higher                                                                                           |

## Logging

|                         Repo                          |   Stars   |  Forks   |                                                    Description                                                    |
| ----------------------------------------------------- | :-------: | :------: | ----------------------------------------------------------------------------------------------------------------- |
| [logrus](https://github.com/sirupsen/logrus)          | **12172** | **1418** | Structured, pluggable logging for Go.                                                                             |
| [zap](https://github.com/uber-go/zap)                 | **7592**  | **576**  | Blazing fast, structured, leveled logging in Go.                                                                  |
| [go-spew](https://github.com/davecgh/go-spew)         | **3342**  | **205**  | Implements a deep pretty printer for Go data structures to aid in debugging                                       |
| [glog](https://github.com/golang/glog)                | **2345**  | **580**  | Leveled execution logs for Go                                                                                     |
| [zerolog](https://github.com/rs/zerolog)              | **2308**  | **156**  | Zero Allocation JSON Logger                                                                                       |
| [tail](https://github.com/hpcloud/tail)               | **1552**  | **315**  | Go package for reading from continously updated files (tail -f)                                                   |
| [lumberjack](https://github.com/natefinch/lumberjack) | **1483**  | **197**  | lumberjack is a log rolling package for Go                                                                        |
| [seelog](https://github.com/cihub/seelog)             | **1362**  | **224**  | Seelog is a native Go logging library that provides flexible asynchronous dispatching, filtering, and formatting. |
| [log15](https://github.com/inconshreveable/log15)     |  **916**  | **116**  | Structured, composable logging for Go                                                                             |
| [log](https://github.com/apex/log)                    |  **736**  |  **67**  | Structured logging package for Go.                                                                                |
| [onelog](https://github.com/francoispqt/onelog)       |  **334**  |  **11**  | Dead simple, super fast, zero allocation and modular logger for Golang                                            |
| [logxi](https://github.com/mgutz/logxi)               |  **334**  |  **35**  | A 12-factor app logger built for performance and happy development                                                |
| [log](https://github.com/go-playground/log)           |  **268**  |  **21**  | :green_book: Simple, configurable and scalable Structured Logging for Go.                                         |
| [logutils](https://github.com/hashicorp/logutils)     |  **251**  |  **26**  | Utilities for slightly better logging in Go (Golang).                                                             |

## Object-Relational Mapping

|                           Repo                           |   Stars   |  Forks   |                                                           Description                                                            |
| -------------------------------------------------------- | :-------: | :------: | -------------------------------------------------------------------------------------------------------------------------------- |
| [gorm](https://github.com/jinzhu/gorm)                   | **14942** | **1700** | The fantastic ORM library for Golang, aims to be developer friendly                                                              |
| [xorm](https://github.com/go-xorm/xorm)                  | **5281**  | **673**  | Simple and Powerful ORM for Go, support mysql,postgres,tidb,sqlite3,mssql,oracle                                                 |
| [gorp](https://github.com/go-gorp/gorp)                  | **3089**  | **353**  | Go Relational Persistence - an ORM-ish library for Go                                                                            |
| [pg](https://github.com/go-pg/pg)                        | **3062**  | **239**  | Golang ORM with focus on PostgreSQL features and performance                                                                     |
| [sqlboiler](https://github.com/volatiletech/sqlboiler)   | **2302**  | **226**  | Generate a Go ORM tailored to your database schema.                                                                              |
| [db](https://github.com/upper/db)                        | **1880**  | **141**  | Productive data access layer for Go.                                                                                             |
| [reform](https://github.com/go-reform/reform)            |  **813**  |  **44**  | A better ORM for Go, based on non-empty interfaces and code generation.                                                          |
| [pop](https://github.com/gobuffalo/pop)                  |  **693**  | **166**  | A Tasty Treat For All Your Database Needs                                                                                        |
| [qbs](https://github.com/coocood/qbs)                    |  **540**  | **102**  | QBS stands for Query By Struct. A Go ORM.                                                                                        |
| [go-queryset](https://github.com/jirfag/go-queryset)     |  **452**  |  **35**  | 100% type-safe ORM for Go (Golang) with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support. GORM under the hood. |
| [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) |  **242**  |  **32**  | A flexible and powerful SQL string builder library plus a zero-config ORM.                                                       |
| [zoom](https://github.com/albrow/zoom)                   |  **242**  |  **20**  | A blazing-fast datastore and querying engine for Go built on Redis.                                                              |
| [grimoire](https://github.com/Fs02/grimoire)             |  **115**  |  **13**  | Database access layer for golang                                                                                                 |
| [go-store](https://github.com/gosuri/go-store)           |  **94**   |  **8**   | A simple and fast Redis backed key-value store library for Go                                                                    |

## Web Frameworks

|                           Repo                            |   Stars   |  Forks   |                                                                                                                                                        Description                                                                                                                                                         
| --------------------------------------------------------- | :-------: | :------: | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- 
| [gin](https://github.com/gin-gonic/gin)                   | **30640** | **3527** | Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.                                                                                                                         
| [beego](https://github.com/astaxie/beego)                 | **21668** | **4393** | beego is an open-source, high-performance web framework for the Go programming language.                                                                                                                                                                                                                                                                                              
| [iris](https://github.com/kataras/iris)                   | **15948** | **1713** | The fastest community-driven web framework for Go. Webassembly, Automatic HTTPS with Public Domain, MVC, Routing on its bests, Sessions, Caching, Versioning API, Problem API, Websocket, Dependency Injection and more. Fully compatible with the standard library and 3rd-party middleware packages. https://bit.ly/謝謝  https://bit.ly/iriscandothat1  https://bit.ly/iriscandothat3 
| [echo](https://github.com/labstack/echo)                  | **14870** | **1359** | High performance, minimalist Go web framework
| [revel](https://github.com/revel/revel)                   | **11285** | **1340** | A high productivity, full-stack web framework for the Go language.
| [martini](https://github.com/go-martini/martini)          | **10675** | **1089** | Classy web framework for Go
| [mux](https://github.com/gorilla/mux)                     | **9838**  | **1063** | A powerful HTTP router and URL matcher for building Go web servers with 🦍
| [httprouter](https://github.com/julienschmidt/httprouter) | **9826**  | **980**  | A high performance HTTP request router that scales well
| [fasthttp](https://github.com/valyala/fasthttp)           | **9536**  | **833**  | Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http
| [chi](https://github.com/go-chi/chi)                      | **6200**  | **430**  | lightweight, idiomatic and composable router for building Go HTTP services
| [buffalo](https://github.com/gobuffalo/buffalo)           | **4827**  | **381**  | Rapid Web Development w/ Go
| [go-swagger](https://github.com/go-swagger/go-swagger)    | **4068**  | **663**  | Swagger 2.0 implementation for go
| [goa](https://github.com/goadesign/goa)                   | **3528**  | **402**  | Design-based APIs and microservices in Go
| [go-restful](https://github.com/emicklei/go-restful)      | **3394**  | **524**  | package for building REST-style Web Services using Go
| [go-json-rest](https://github.com/ant0ine/go-json-rest)   | **3341**  | **375**  | A quick and easy way to setup a RESTful JSON API
| [gizmo](https://github.com/nytimes/gizmo)                 | **2866**  | **179**  | A Microservice Toolkit from The New York Times
| [macaron](https://github.com/go-macaron/macaron)          | **2828**  | **245**  | Package macaron is a high productive and modular web framework in Go.
| [armor](https://github.com/labstack/armor)                | **1552**  |  **56**  | Uncomplicated, modern HTTP server
| [web](https://github.com/gocraft/web)                     | **1396**  | **105**  | Go Router + Middleware. Your Contexts.
| [tango](https://github.com/lunny/tango)                   |  **819**  | **108**  | This is only a mirror and Moved to https://gitea.com/lunny/tango
| [goji](https://github.com/goji/goji)                      |  **769**  |  **61**  | Goji is a minimalistic and flexible HTTP request multiplexer for Go (golang)
| [neo](https://github.com/ivpusic/neo)                     |  **393**  |  **40**  | Go Web Framework

## Game

|                          Repo                          |  Stars  | Forks  |                       Description                       |
| ------------------------------------------------------ | :-----: | :----: | ------------------------------------------------------- |
| [go-astar](https://github.com/beefsack/go-astar)       | **330** | **34** | Go implementation of the A* search algorithm            |
| [gotetris](https://github.com/jjinux/gotetris)         | **208** | **26** | This is a console-based version of Tetris written in Go |
| [snake-game](https://github.com/DyegoCosta/snake-game) | **140** | **33** | Terminal-based Snake game                               |
| [sokoban-go](https://github.com/rn2dy/sokoban-go)      | **33**  | **6**  | sokoban game in terminal written with go                |
| [go-tetris](https://github.com/MichaelS11/go-tetris)   | **30**  | **1**  | Golang Tetris for console window with optional AI       |

### Game engine

|                          Repo                           |  Stars   |  Forks  |                                        Description                                        |
| ------------------------------------------------------- | :------: | :-----: | ----------------------------------------------------------------------------------------- |
| [leaf](https://github.com/name5566/leaf)                | **3103** | **870** | A game server framework in Go (golang)                                                    |
| [pixel](https://github.com/faiface/pixel)               | **2473** | **144** | A hand-crafted 2D game library in Go                                                      |
| [ebiten](https://github.com/hajimehoshi/ebiten)         | **1902** | **128** | A dead simple 2D game library in Go                                                       |
| [goworld](https://github.com/xiaonanln/goworld)         | **1218** | **235** | Scalable Distributed Game Server Engine with Hot Swapping in Golang                       |
| [go-sdl2](https://github.com/veandco/go-sdl2)           | **1163** | **149** | SDL2 binding for Go                                                                       |
| [engo](https://github.com/EngoEngine/engo)              | **1094** | **92**  | Engo is an open-source 2D game engine written in Go.                                      |
| [gonet](https://github.com/xtaci/gonet)                 | **1055** | **292** | A Game Server Skeleton in golang.                                                         |
| [termloop](https://github.com/JoelOtter/termloop)       | **1029** | **66**  | Terminal-based game engine for Go, built on top of Termbox                                |
| [nano](https://github.com/lonng/nano)                   | **1017** | **158** | Lightweight, facility, high performance golang based game server framework                |
| [engine](https://github.com/g3n/engine)                 | **770**  | **73**  | Go 3D Game Engine                                                                         |
| [oak](https://github.com/oakmound/oak)                  | **637**  | **33**  | A pure Go game engine                                                                     |
| [engine](https://github.com/azul3d/engine)              | **427**  | **32**  | Azul3D - A 3D game engine written in Go!                                                  |
| [raylib-go](https://github.com/gen2brain/raylib-go)     | **388**  | **36**  | Go bindings for raylib, a simple and easy-to-use library to learn videogames programming. |
| [GarageEngine](https://github.com/vova616/GarageEngine) | **313**  | **26**  | Game engine written in Go (golang).                                                       |

## Go E-Books

|                                               Repo                                                |   Stars   |  Forks   |                      Description                      |
| ------------------------------------------------------------------------------------------------- | :-------: | :------: | ----------------------------------------------------- |
| [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang) | **31865** | **8817** | A golang ebook intro how to build a web with golang   |
| [GoBooks](https://github.com/dariubs/GoBooks)                                                     | **6869**  | **917**  | List of Golang books                                  |
| [web-dev-golang-anti-textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook)     | **2372**  | **211**  | Learn how to write webapps without a framework in Go. |
| [go101](https://github.com/go101/go101)                                                           | **2077**  | **133**  | An online book focusing on Go syntax/semantics.       |
| [learninggo](https://github.com/miekg/learninggo)                                                 |  **367**  |  **83**  | Learning Go Book in mmark                             |
| [book](https://github.com/GoBootcamp/book)                                                        |  **238**  |  **68**  | Source code of the companion book/website             |
| [go-advanced](https://github.com/zalopay-oss/go-advanced)                                         |  **182**  |  **44**  | A small Vietnamese Go book compiled by ZaloPay teams. |

*Last Update: 2019-08-28T09:59:11+07:00*
