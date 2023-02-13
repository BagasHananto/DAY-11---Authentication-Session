function getData() {
    let name = document.getElementById('name').value
    let email = document.getElementById('email').value
    let phone = document.getElementById('phone').value
    let subject = document.getElementById('subject').value
    let description = document.getElementById('description').value

    document.getElementById("alert-contact").style.backgroundColor = "rgb(255, 0, 0)";
    document.getElementById("alert-contact").style.color = "rgb(255, 255, 255)";

    if (name == "") {
        document.getElementById("alert-contact").innerHTML = "Name cannot be empty";
    } else if (email == "") {
        document.getElementById("alert-contact").innerHTML = "Email cannot be empty";
    } else if (phone == "") {
        document.getElementById("alert-contact").innerHTML = "Phone cannot be empty";
    } else if (subject == "") {
        document.getElementById("alert-contact").innerHTML = "Subject cannot be empty";
    } else if (description == "") {
        document.getElementById("alert-contact").innerHTML = "Description cannot be empty";
    } else {



        document.getElementById("alert-contact").style.backgroundColor = "rgb(3, 123, 3)";
        document.getElementById("alert-contact").innerHTML = "Success add message"

        const defEmail = "danandjaya123@gmail.com"
        let ebody = "Hallo Bagas, nama saya" + " " + name + "." + " " + "Saya seorang" + subject + "." + " " + "Saya ingin" + " " + description + "." + " " + "Kamu bisa hubungi nomor:" + " " + phone + "." + " " + "Terimakasih."

        window.location.href = `mailto:${defEmail}?subject=${subject}&body=${ebody}`
        let audience = {
            name,
            email,
            phone,
            subject,
            description,
        }
    }
}