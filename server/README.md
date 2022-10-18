
# Go Server

Note related to Go server development 

## OSI Model

| Host Layers   | Process             |Device/ Protocol         |                      |
| --------------| --------------------|-------------------------|--------------------  |
| 7- Application | Network process to end-user application      | SMTP/ DNS/ HTTP/ HTTPS/ FTP| Human- computer interaction layer, where applications can access the network services|
|6- Presentation | Data presentation and encryption | JPEG/ GIF/ TIFF/ HTML/ DOC| Ensure that data is in a usable format and is where data encryption occurs|
|5- Session| Interhost communication | NFS/ RPC/ SQL/ RTP | Maintains connections and is responsible for controlling ports and sessions|
|4- Transport | End-to-end connections and reliabiliity| TCP/ UDP/ SPX| Transmits data using transmission protocols including TCP and UCP|
|3- Network | Path determination and logical addressing | IP/ ICMP/ PX | Decide which physical path the data will take|
|2- Data Link | Physical addressing | PPP/ SLIP| Defines the format of data on the network|
|1- Physical | Media, signal and binary transmission| Hub| Transmits raw bits over the physical medium|   


## Documentation

[Hypertext Transfer Protocol](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol)

---

[Transmission Control Protocol](https://en.wikipedia.org/wiki/Transmission_Control_Protocol)

---

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`API_KEY`

`ANOTHER_API_KEY`


## Run Locally

Clone the project

```bash
  git clone https://link-to-project
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  npm install
```

Start the server

```bash
  npm run start
```


## Usage/Examples

```javascript
import Component from 'my-project'

function App() {
  return <Component />
}
```

