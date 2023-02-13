let projects = []

function getData(event) {
    event.preventDefault()

    let title = document.getElementById('name').value
    let startDate = document.getElementById('start').value
    let endDate = document.getElementById('end').value
    let description = document.getElementById('description').value
    let check1 = document.getElementById('nodejs').checked
    let check2 = document.getElementById('php').checked
    let check3 = document.getElementById('java').checked
    let check4 = document.getElementById('laravel').checked
    let image = document.getElementById('formFile').files

    let icon1 = (check1 == true) ? '<i class="fa-brands fa-node fa-2x"></i>' : '';
    let icon2 = (check2 == true) ? '<i class="fa-brands fa-php fa-2x"></i>' : '';
    let icon3 = (check3 == true) ? '<i class="fa-brands fa-java fa-2x"></i>' : '';
    let icon4 = (check4 == true) ? '<i class="fa-brands fa-laravel fa-2x"></i>' : '';

    //let start = new Date(startDate)
    //let end = new Date(endDate)


    document.getElementById("alert").style.backgroundColor = "rgb(255, 0, 0)";
    document.getElementById("alert").style.color = "rgb(255, 255, 255)";

    if (title == "") {
        document.getElementById("alert").innerHTML = "Project Name cannot be empty";
    } else if (startDate == "") {
        document.getElementById("alert").innerHTML = "Start Date cannot be empty";
    } else if (endDate == "") {
        document.getElementById("alert").innerHTML = "End Date cannot be empty";
    } else if (description == "") {
        document.getElementById("alert").innerHTML = "Description cannot be empty";
    } else if (icon1 == false && icon2 == false && icon3 == false && icon4 == false) {
        document.getElementById("alert").innerHTML = "Technology cannot be empty";
    } else if (image.length == 0) {
        document.getElementById("alert").innerHTML = "Image cannot be empty";
    } else {

        image = URL.createObjectURL(image[0]);

        document.getElementById("alert").style.backgroundColor = "rgb(3, 123, 3)";
        document.getElementById("alert").innerHTML = "Success add project"


        let icon = [
            icon1,
            icon2,
            icon3,
            icon4
        ]

        let startDate = new Date(document.getElementById("start").value)
        let endDate = new Date(document.getElementById("end").value)
        let diff = endDate.getTime() - startDate.getTime() //in milisec

        if (diff < 0) {
            //to call alert
            document.getElementById("alert").innerHTML = "Duration is invalid";
            duration = ""
        } else {
            let durYear = Math.floor(diff / (12 * 30 * 24 * 3600 * 1000)) //in year
            let durMonth = Math.floor(diff / (30 * 24 * 3600 * 1000)) //in month
            let durWeek = Math.floor(diff / (7 * 24 * 3600 * 1000)) //in week
            let durDay = Math.floor(diff / (24 * 3600 * 1000)) //in day

            if (durYear > 0) {
                duration = durYear + " Year(s)"
            } else if (durMonth > 0) {
                duration = durMonth + " Month(s)"
            } else if (durWeek > 0) {
                duration = durWeek + " Week(s)"
            } else if (durDay > 0) {
                duration = durDay + " Day(s)"
            }

            let iconFil = icon.filter(elem => elem[1]).reverse()
            console.log(iconFil)

            let dataBlog = {
                image,
                title,
                description,
                iconFil,
                startDate,
                endDate,
                duration

            }
            projects.push(dataBlog)
            showProject()

        }
    }

    function showProject() {
        document.getElementById("content-project").innerHTML = ''
        for (let i = 0; i <= projects.length; i++) {

            if (projects[i].iconFil[0] == undefined) {
                projects[i].iconFil[0] = ""
            }
            if (projects[i].iconFil[1] == undefined) {
                projects[i].iconFil[1] = ""
            }
            if (projects[i].iconFil[2] == undefined) {
                projects[i].iconFil[2] = ""
            }
            if (projects[i].iconFil[3] == undefined) {
                projects[i].iconFil[3] = ""
            }
            if (projects[i].iconFil[4] == undefined) {
                projects[i].iconFil[4] = ""
            }


            document.getElementById("content-project").innerHTML += `
        <div>
            <div class="card m-3 shadow p-3" style="width: 18rem;">
                <img src="${projects[i].image}" class="card-img-top" alt="">
                <div class="card-body">
                    <div>
                    <h5 class="card-title fw-bold">${projects[i].title}</h5>
                    <p style=" margin-bottom: 10px; color: gray;">Duration : ${projects[i].duration}</p>
                    </div>
                    <p class="card-text">${projects[i].description}</p>
                    <div class="d-flex mb-3 gap-3 justify-content-center">
                        <div>${projects[i].iconFil[0]}</div>
                        <div>${projects[i].iconFil[1]}</div>
                        <div>${projects[i].iconFil[2]}</div>
                        <div>${projects[i].iconFil[3]}</div>
                    </div>
                    <div class="d-flex gap-3 justify-content-between">
                        <a href="#" class="btn btn-primary">Edit</a>
                        <a href="#" class="btn btn-danger">Delete</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
    `

        }
    }
}


//statement duration timeout
// function duration(startDate, endDate) {
//    let deadline = new Date(endDate) - new Date(startDate)

//Pengkondisian

//    if (deadline < 0) {
//        document.getElementById("alert").innerHTML = "Duration is invalid";
//        duration = ""
//   } else {
//        let day = Math.floor(deadline / (1000 * 60 * 60 * 24))
//        let month = Math.floor(deadline / (1000 * 60 * 60 * 24 * 30))
//        let week = Math.floor(deadline / (1000 * 60 * 60 * 24 * 7))
//        if (month > 0) {
//            return `${month} Month`
//        } else if (week > 0) {
//            return `${week} Week`
//        } else if (day > 0) {
//            return `${day} Day`
//        }

//  }
//}