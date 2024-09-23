# Changelog

## 1.6.0 (23/09/24)
- remove second check of authorisation
- update Gufo to v1.17.0

## 1.5.0 (21/06/24)
- Add ssl certs in Dockerfile
- Change Entrypoint logig. Now Entrypoint funxtion is running everytime when container starts. We add Update function. It is running one time if version of microservice is different

## 1.4.0 (21/06/24)
- Fixed bug with GET requests

## 1.3.0 (24/05/24)
- Bug fix with entrypoint version
- Bug fix with cron status
- Now it is possible to set microservice listen port
- Update Dockerfile
- Add version in generated files

## 1.2.0 (06/05/24)
- Bug fix with Authorisation check

## 1.1.0 (22/04/24)
- Add Cron folder
- Add admin Folder
- Add admin API endpoints for run and stop cron
- Update GUFO version in microservice go.mod
