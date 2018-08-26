# The Church Web Platform
A template-less golang Content Management System (CMS) based on the Echo web framework (https://github.com/labstack/echo) and SQLBoiler (https://github.com/vattle/sqlboiler) and an innovative, think-outside-the-box golang architecture. 

## Intro
The project started from the need to build a sustainable website for a small church pastored by a friend. The site needed to allow admins to maintain events, sermons, and articles. Here's a temporary link to that site: http://www.cceastmetro.com/pages/home
- The start of commit history here represents 1.5 years of previous work - so this sits on top of a great amount of innovation and effort, yet this is just the beginning...
- Note: while I feel the project, once configured is production ready, please use at your own discretion. 
- Note also that over the next couple months there may be some change to the architecture.
- The project will adopt semantic versioning very soon

## Architecture
- The **Echo** web framework (https://github.com/labstack/echo) was chosen because of its maturity and simplicity - hey, and speed too
- **Postgres** is our main database. **SQLBoiler** provides a code-generated, thus compiled and performant database interface layer
- **Redis** provides our session and (soon) our caching store
- Template - none. This is a **code-first** platform (anyone here love Go?), as such HTML is generated by Go, _simply_, with **github.com/rohanthewiz/element**.
- Glide is used for dependency management - why? It still works great while we wait for the dust to settle from the transition from Dep to Vgo.
- Sermons are uploaded to the server then copied by a goroutine to an ftp site.
- Images uploaded via the online editor are resized then stored locally on the server
- Styling - See the Contributing section for the policy on styling - basically do your own, but don't contribute it back so projects can be styled differently.See the example app for an example (coming soon) using [stylus] (http://stylus-lang.com) - my CSS processor of choice.
- BTW Bootstrap is only used in the online editor Summer Note. I therefore limit Bootstrap CSS to just form scope.
- And don't worry, Docker will come. Let's just everything jiving smoothly first.  

## General Workflow
1. Make a fork of this project and clone your fork to local
1. Check out the `master` branch. Maintain your project in master.
1. (If there are non-styling and non-project-specific changes you'd like to contribute back to the project, cherry pick those into core and make a pull request from core to the source repo.)

## Platform Setup
The project requires
- Go 1.7 or greater
- Postgres around 9.5 or later
- Redis
- VIPS image library (optional - only for image resizing)
Note: feel free to use your own deployment strategy. The below is an **example install on Ubuntu Linux**.

### Server Prep
Basically everything is built on the server for now (KISS comes to mind)
- Note: I prefer not to include `$GOPATH/bin` in my PATH, but rather manually copy executables to `~/bin` which I include in my PATH
- Newb: 
    1. Create the directory (folder) with `mkdir ~/bin`
    1. To include `~/bin` in your path add a line like this to `~/.profile`: `PATH="$HOME/bin:$PATH"`

#### Install Go
- Follow install instructions at: https://golang.org/doc/install
- Run `go env` to see if everything is sane. Mine looks like this:

```bash
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/myuser/go"
GORACE=""
GOROOT="/home/myuser/apps/go"
GOTOOLDIR="/home/myuser/apps/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
CC="gcc"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build264891901=/tmp/go-build -gno-record-gcc-switches"
CXX="g++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
```

#### Install GCC and build tools
- Install gcc and essential build tools (needed for VIPS and glide install): `sudo apt-get install build-essential`

#### Install postgres server
- Follow instructions here: https://www.postgresql.org/download/linux/ubuntu/ (Get the server *and client tools*)
- For passwords to work without known system user, change an entry in */etc/postgresql/<postgres-version>/main/pg_hba.conf*
- In pg_hba.conf, change this line: 

```
local   all             all                                     peer
```
to:  

```
local   all             all                                     password
```
- Restart the postgres server `sudo service postgresql restart`
- Check that the server is running properly with `sudo service postgresql status` (it should say "online" or similar)
- Go into psql as user postgres `sudo -u postgres psql postgres` This is one way to login after a fresh reload of postgres.
- Change the postgres user's password with `\password postres`

#### Install Redis Server
- Install redis `sudo apt install redis-server`

#### Install VIPS
VIPS is used to automatically resize images added to the editor. _We most likely don't need all libraries installed below_. In fact, VIPS isn't critical to the app - you should lose only the image resizing capability.

```bash
mkdir ~/libs
cd libs
wget https://github.com/jcupitt/libvips/releases/download/v8.5.9/vips-8.5.9.tar.gz
tar -xf vips-8.5.9.tar.gz
cd vips-8.5.9/
./configure
gcc -v
which pkg-config
sudo apt-get install pkg-config
sudo apt-get install libgmodule2.0-dev
sudo apt-get install libgmodule2.0
sudo apt-get install libexpat-dev
sudo apt-get install libexpat2-dev
sudo apt-get install libexpat2.2.5-dev
sudo apt-get install libexpat2.2-dev
sudo apt-get install libexpat2
sudo apt-get install libexpat-2
sudo apt-get install libexpat2.0
sudo apt-get install libexpat1-dev
sudo apt-get install libpng
sudo apt-get install libpng2
sudo apt install libjpeg-dev
sudo apt install libpng16
sudo apt install libpng16-dev
sudo apt install giflib-dev
sudo apt install libgif-dev
sudo apt install giflib-dbg
sudo apt install libtiff
sudo apt install libtiff5-dev
sudo apt install libpoppler-glib-dev
sudo apt install libpoppler-glib8
sudo apt install librsvg2-dev
sudo apt install librsvg2
sudo apt install librsvg2-2
sudo apt install libexif-dev
./configure
make
sudo make install
vips -v
cd /etc/ld.so.conf.d/
more libc.conf # should include `/usr/local/lib`
sudo ldconfig # update shared libraries
vips -v
```

#### Go: Sqlboiler Getting Started - https://www.youtube.com/watch?v=fKmRemtmi0Y
- `go get -u -t github.com/vattle/sqlboiler`
- `go get -u gopkg.in/nullbio/null.v6`
- `go get -u gopkg.in/inconshreveable/log15.v2`
- Once the db is migrated, SQLboiler will generate/ update the model code (in the `models` directory)
- Copy the sqlboiler binary into place `cp $GOPATH/bin/sqlboiler ~/bin/`
- Make sure project_root/sqlboiler.toml has proper creds for your postgres database
- From the project root run `sqlboiler postgres`

#### Install Glide
```bash
mkdir -p go/src/github.com/Masterminds
cd go/src/github.com/Masterminds/
git clone https://github.com/Masterminds/glide.git
cd glide
make build
mv glide ~/bin/ # ~/bin is a good place to store exe's - make sure it is added to $PATH
```

## App Setup

### Migrations
- As the 'posgres' user, create a user with CREATEDB permissions
 `psql -h localhost -p 5432 -d postgres -U postgres -c "create user myuser with password 'secret' CREATEDB"`
- Create the database `postgres=# CREATE DATABASE "church_development" WITH OWNER "myuser";`
- Run goose migrations: first make sure to get and install the goose binary: https://github.com/pressly/goose
- `go get -u github.com/pressly/goose/cmd/goose`
- Note goose requires that the database is already created
- Change to the migrate directory `cd db/migrate`
- Migrate up: `goose postgres "user=myuser password=secret dbname=church_development sslmode=disable" up`
- You can manually migrate each of the sql files (though not recommended) `psql -h localhost -d church_development -U myuser -f db/migrate/20170419004813_CreateUsersTable.sql`

### Configuration
- `cp options.yml.sample options.yml`
- Update config params: `sudo nano config/options.yml`

#### Create a Random Seeds file
- Create a file: `config/random_seeds.txt` and populate it with random strings _one per line_.
- If you are not paranoid about security (you should be) you can just copy from the provided sample file `cp config/rand_seeds.txt.sample config/rand_seeds.txt`
- For the above I highly recommend Random.org's *Random File Generation Service* https://files.random.org/ (cheap)
    or manually build random_seeds.txt with strings generated by their free *Random String Generator Service*. https://www.random.org/strings/ 

## Building the Platform

```bash
# Clone the repo into a suitable place under your GOPATH
mkdir -p ~/go/src/github.com/rohanthewiz/
cd ~/go/src/github.com/rohanthewiz/
git clone <the-project-clone-url>
cd church
rm -rf vendor
glide install
glide status
go build
```

### Initial Run
- I do apologize that there is some seeding required here -- please be patient. (On the TODO is to accomplish the above seeding via Grift scripts or something similar.)
- Start the server `./church`
- (Missing: a file with a token with which to start administration)
- The slug of the Main menu should be hardwired to "main-menu". For now manually update that DB entry
- The slug for the Footer menu also should be hardwired to "footer-menu"


### Install RunIt process supervisor (Optional) for running the app resiliently
sudo apt-get install runit
- `sudo mkdir -p /etc/sv/church`
- `cd /etc/sv/church`
- `sudo nano run`
- Below is a sample Runit script
- Note: chpst (Change Process state) looks for envars in `env_dir` specified by the `-e` flag

```bash
#!/bin/sh -e
exec 2>&1
cd /home/myuser/go/src/github.com/rohanthewiz/church
exec chpst -e /etc/service/church/env_dir "/home/myuser/go/src/github.com/rohanthewiz/church/church"
```

- `sudo chmod +x run`
- `./run` # it should run manually
- `cd /etc/service`
- `sudo ln -s /etc/sv/church church`
- `sudo sv status church`
- `sudo sv up church`

## Backup the DB (Might be a good candidate for crontab)
- On the server run `pg_dump -U myuser church_development > ~/church_db-2018-0624.sql`

## Contributing
- The bulk of the code of the platform is kept in the `core` branch.
- Contribute non-styling changes back to the core branch.
- Individual projects can make styling changes in their master branch
- Individual projects can then merge changes from core into their master branch without affecting their individual styling.
- In summary: core is the shareable branch, master is your own
- _We will not accept styling changes in a PR_, so keep those out of core if you'd like to contribute back to the project
- A contribution flow might look like this:
    1. Clone the project
    1. `git checkout master`
    1. Create a new whiz-bang module
    1. Git add and commit styling changes
    1. Git add and commit code changes
    1. Checkout core
    1. Cherry pick the code change commit into the core branch
    1. Push the core branch
    1. On Github, make a pull request from core branch to core branch
    
Happy Coding :-)