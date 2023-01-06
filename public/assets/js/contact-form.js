
const emailReceiver = "bogeymanyolo@gmail.com";


let nameEl = document.getElementById("name");
let emailEl = document.getElementById("email");
let phoneNumberEl = document.getElementById("telp");
let subjectEl = document.getElementById("subject");
let messageEl = document.getElementById("message");

function submitForm() {
    let name = nameEl.value;
    let email = emailEl.value;
    let phoneNumber = phoneNumberEl.value;
    let subject = subjectEl.value;
    let message = messageEl.value

    if (name == ""){
        return alert("Name input fields must not be empty")
    } else if (email == ""){
        return alert("Email input fields must not be empty")
    } else if (phoneNumber == ""){
        return alert("Phone Number input fields must not be empty")
    } else if (subject == ""){
        return alert("Subject input fields must not be empty")
    } else if (message == ""){
        return alert(message == "Message input fields must not be empty")
    }


    const a = document.createElement("a");
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Hello my name ${name}, ${subject}, ${message}`;
    a.target = "_blank";
    a.click()

    let dataObject = {
        name: name,
        email: email,
        phoneNumber: phoneNumber,
        subject: subject,
        message: message
    };
    console.log(dataObject);
}
    