[![MIT License][license-shield]][license-url]
[![Go][go-shield]][go-url]
[![Build][build-shield]][build-url]
[![Commit][commit-shield]][commit-url]
![Code][code-sheild]
[![LinkedIn][linkedin-shield]][linkedin-url]

<br>

<h1 align="center"> nitd-results </h1>
<p align="center">
  REST API for accessing the results of NIT-Delhi students.
  <br>
  <br>
  <a href="http://nitdresults.herokuapp.com/docs"><strong>Explore the docs</strong></a>
  <br>
  <a href="https://github.com/killtheverse/nitd-results/issues">Report a bug</a>
   | 
  <a href="https://github.com/killtheverse/nitd-results/issues">Request a feature</a>
</p>

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
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

## About The Project
NIT Delhi publishes the semester wise results of students on this [website][results-website-url]. 
However, if one wishes to view the results of multiple students, then they will have to enter the student's details and captcha multiple times. 
Hence, I created this easy to use API that will fetch the results of multiple students at once. 
Students can be filtered on several factors - Branch, Batch and Program. 
Data is fetched from a remote Database which is populated once every semester using the [populator][populator-url]. 

### Built With
- [Go](https://go.dev/)
- [Gorilla](https://www.gorillatoolkit.org/)

## License
Distributed under the MIT License. See [LICENSE][license-url] for more information.

## Contact
Rahul Dev Kureel - r.dev2000@gmail.com

Project Link: [https://github.com/killtheverse/nitd-results](https://github.com/killtheverse/nitd-results)

## Acknowledgments

- [Go web services](https://github.com/nicholasjackson/building-microservices-youtube)
- [MongoDB driver](https://docs.mongodb.com/drivers/go/current/)
- [JWT Authentication](https://learn.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr/)
- [Readme template](https://github.com/othneildrew/Best-README-Template)


[license-shield]: https://img.shields.io/github/license/killtheverse/nitd-results?style=for-the-badge
[license-url]: https://github.com/killtheverse/nitd-results/blob/main/LICENSE
[go-shield]: https://img.shields.io/github/go-mod/go-version/killtheverse/nitd-results?style=for-the-badge
[go-url]: https://github.com/killtheverse/nitd-results/blob/main/go.mod
[build-shield]: https://img.shields.io/github/workflow/status/killtheverse/nitd-results/go-docker-heroku-cd?style=for-the-badge
[build-url]: https://github.com/killtheverse/nitd-results/actions
[commit-shield]: https://img.shields.io/github/last-commit/killtheverse/nitd-results?style=for-the-badge
[commit-url]: https://github.com/killtheverse/nitd-results/commits/main
[code-sheild]: https://img.shields.io/tokei/lines/github/killtheverse/nitd-results?style=for-the-badge
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/rahul-dev-386454136/
[results-website-url]: https://erp.nitdelhi.ac.in/CampusLynxNITD/studentonindex.jsp
[populator-url]: https://github.com/killtheverse/nitd-results-populator
