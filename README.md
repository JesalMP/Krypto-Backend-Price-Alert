[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-c66648af7eb3fe8bc4f294546bfd86ef473780cde1dea487d3c4ff354943c9ae.svg)](https://classroom.github.com/online_ide?assignment_repo_id=7942927&assignment_repo_type=AssignmentRepo)
<div id="top"></div>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://www.letskrypto.com">
    <img src="https://www.letskrypto.com/img/krypto-logo-nas.png" alt="Logo">
  </a>

<h3 align="center"GoLang Crypto Price Monitoring in Realtime With EMail Notification</h3>

  <p align="center">
    This tool is used to create and manage BTC Prices (w.r.t. USDT exchange) and Alert User if the price has falen or hiked as per Alert settings via Email Notification.
    <br />
    <a href="https://github.com/JesalMP/Krypto-Backend-Price-Alert"><strong>Explore the docs »</strong></a>
    <br />
    <br />
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

<img src="images/ss1.png" alt="About">

This Project is a GoLang Project Backend REST API software to let a user create and manage BTC Price Alerts and send/Trigger these alerts as soon as the price hit he markline. Realtime prices are taken from Binance API and There are 2 Go routines running in the Docker Container, One Serving the HTTP requests to set, view and manage the API reponses, and other iteration over Prcies of BTC/USDT and triggering the alerts and sending mails accordingly.

<p align="right">(<a href="#top">back to top</a>)</p>

## Features
- Redis Caching
- JWT Tokens
- Docker Containerization
- Go Parallel SubRoutines
- Realtime Price Tracking
- Email ALerts

<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [GoLang](https://go.dev/)
* [mongoDB](https://www.mongodb.com/)
* [DockerHub](https://hub.docker.com/)
<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

First things first, to use this software, you'll have to install Docker, create an account on Dockerhub.io and have GoLang Installed.




### Prerequisites

WSL2 updated version need to be installed in order for Docker to work. Check Online guides on Installing Docker for more.
Post Installing Docker, Install GoLang and Run the Docker Daemon.
![image](https://user-images.githubusercontent.com/84318539/181877796-db739efc-33e7-4c8b-af6c-d046f67e2a98.png)



### Installation

1. Download the code as zip or clone the repo.
2. In the project root directory, go to .env file and set your MONGODB Url to access the data base, Set your Personal email Id, set IMAP Handler Email id and password (Zoho mail is recomended, since GMAIL disabled support for IMAP), set IMAP Handler Email host address and port.
3. We are Good to Go.
![image](https://user-images.githubusercontent.com/84318539/181877883-21114630-9fd3-444c-b028-0bb23822f8df.png)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage
1. Go to Project Root Dir.
2. In the Root Directory, Run the Following Commands in windows CMD, Powershell or terminal in Linux,
- ```sh
  docker compose build
  ```

-  ![image](https://user-images.githubusercontent.com/84318539/181878056-17b30442-7dfb-435c-bd66-b832ef16ef02.png)


3. This Will Build the latest Docker Image to be used in a container of Our Go Project.
4. To run the Project, in the cmd/powershell/terminal , run the following command
- ```sh
  docker run -p 8080:8080 -it krypto-backend-price-alert_web
  ```
-![image](https://user-images.githubusercontent.com/84318539/181878329-d734d511-891a-4dc4-b9da-c8c3a0c2f399.png)

5. Notice that we are Using port 8080 in the LocalHost for the APIs

This Project Has Following APIs
## /alerts/create/{alertPrice}
- Creates a Alert with price inputted in alertPrice
- The Default status of the set Alert State as "Primed".
- ![image](https://user-images.githubusercontent.com/84318539/181878374-fe56def3-afa1-4b35-93fe-c0ba1a7f1208.png)
- ![image](https://user-images.githubusercontent.com/84318539/181878382-d7182a3d-3e5a-4912-a9cd-adda00f52503.png)

## /alerts/delete/{alertPrice}
- Deletes the Alert with price inputted in alertPrice.
- ![image](https://user-images.githubusercontent.com/84318539/181878444-72134b52-7a41-4a92-8b4b-4c35fe47544a.png)


## /alerts
- Shows all alerts present in MongoDB Database.
- ![image](https://user-images.githubusercontent.com/84318539/181878396-3f459a0c-a184-4f15-b38e-b67b45353178.png)
- ![image](https://user-images.githubusercontent.com/84318539/181878412-ee9bbcd1-29a7-435d-b8fa-f25e2128a3f0.png)
## /alerts/show/{state}
- Shows all Alerts in the Database with its State as given in state.
- e.g. /alerts/show/Primed will show all alerts that are primed, /alerts/show/Triggered will show all alerts that are Triggered.
- ![image](https://user-images.githubusercontent.com/84318539/181878465-b92a9d6f-5a67-4ddc-837e-58a8f020d242.png)


# Email Alerts
- A seperate GORoutine is running continously monitoring the Realtime Prices of BTC/USDT with Binance API.
- When user sets an Alert, There are 2 possibilty of triggering the alert
- - Let price of BTC  be X during Creation of Alert, and let ALert Price be A which will be Greater then X. When BTC Price Exceed both X and A, the Alert will happen as a "Price Hike" alert.
- - Let price of BTC  be X during Creation of Alert, and let ALert Price be A which will be Leser then X. When BTC Price Drops below both X and A, the Alert will happen as a "Price Dip" alert.
- When the price happens to be in range of any "Primed" alerts, its status will be changed to Triggered in Database, and a email will be sent to user's email id of Current price, time of triggering and Which Alert has been Triggered.
- ![image](https://user-images.githubusercontent.com/84318539/181878650-08f558e4-074a-40f3-9ebc-6564083c1889.png)





<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Patel Jesal Manoj - jesalpatel290@gmail.com
https://www.linkedin.com/in/jesal-patel-130a5b217/ 

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* https://tutorialedge.net/golang/go-docker-tutorial/ Reference for code taken.
* https://levelup.gitconnected.com/dockerized-crud-restful-api-with-go-gorm-jwt-postgresql-mysql-and-testing-61d731430bd8 For Reference and Insights.


<p align="right">(<a href="#top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/dyte-submissions/dyte-vit-2022-JesalMP/blob/main/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/jesal-patel-130a5b217/
[product-screenshot]: images/screenshot.png
