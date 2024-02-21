# 👨‍💼 G-Mail Butler

Deletion of all unread messages in Gmail using GmailAPI written in Go.

## Quickguide

Make sure you have installed `Go` lang on your machine. You can do so here: https://go.dev/doc/install.
Run

```shell
go install
```

### Enable the API

- Enable the API In the Google Cloud console, enable the Gmail API.
- Configure the OAuth consent screen
- Authorize credentials for a desktop application
- Make sure that `modify` is in the permission Scope. Needed for `batchDeletion`

### Build script

```shell
go build -o build/gmailbutler
```

### Run script

```shell
./build/gmailbutler
```

### OR use Bash Script

This should work only on MacOS.

```shell
chmod +x run.sh
./run.sh
```
