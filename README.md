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

### Run script on MacOS

```shell
./build/gmailbutler
```

### OR use Bash Script on MacOS

```shell
chmod +x run.sh
./run.sh
```

### Run script on Windows

```shell
./build/gmailbutler.exe
```

### Testing

```shell
go test .
```

### TODO

- [] add cli ability to specify how many unreads emails it should delete
- [] add cli ability to specify timeframe from-to with dates
- [] delete labeled emails
- [] create notification system to push notification if email contains specific characters in email
