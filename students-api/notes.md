# Project Initialization

### 1. Init Go mod:

- go mod is used to  manage dependency and other external libraries which are going to be used in our project.
- convention for init go module is go mod init github.com/username/repository-name . inside which your project will be sit

### 2. Folder Structure
- most of the go project has a typical folder structure


- create a folder cmd inside root folder, then inside cmd we will have another folder. which name would be same as project name (here our will be students-api).
- Inside students-api we will have our main.go file which will be the main  entry file.

### 3. config:
- this is one of important thing to manage configuration of project, especially if it's a production level project.
    1. Either we can use env to inject configuration of the project.
    2. otherwise most of the time file based config prefered in production level. because you can keep this file based configuration under version control, so we can track the version of its.

#### How to setup file based config?

- create a config folder inside root folder, then inside root folder create a local.yaml file, inside that yaml file we will keep all our configurations.
- in future we can create a production.yaml, inside which we will keep all our production level config.

SQLite is a file-based SQL database, which means the entire database—including tables, data, indexes, and other database information—is stored in a single file on your computer. Unlike MySQL or PostgreSQL, SQLite does not require a separate database server to run.

- its performant and don't required network call cause it stay inside your project, that's why its so fast

- sqllite is filebased db, so we have to declare its path inside config file. so it will be configurable in prod 


- http server setup, so that's a yaml file so we have ton take care of intended. we can do nesting using tab

- for now this much config is enough, later on we can add more configs like timeout and all the stuffs.