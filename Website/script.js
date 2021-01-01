const finalLogo = document.getElementById("finalLogo");
const genLogo = document.getElementById("genLogo");
let logoExists = false;

function checkInput(uniName) {
    if (uniName === "") {
        alert("Enter University Name!");
        return false;
    }
    return true;
}

function showImage(src) {
    finalLogo.src = src
}

genLogo.onclick = function () {
    const uniName = document.getElementById("uniName").value;
    const imgColor = document.getElementById("colors").value;
    let opacity = 1.0;
    if (document.getElementById("bgColor").checked) {
        opacity = 0.0
    }

    if (!checkInput(uniName)) {
        return
    }

    const url = "http://127.0.0.1:6969/logo-gen/api/gen?" + "uni_name=" + uniName + "&img_color=" + imgColor + "&opacity=" + opacity;

    let xhttp = new XMLHttpRequest();
    xhttp.open("GET", url, false);
    xhttp.setRequestHeader("Content-type", "text/html");

    try {
        xhttp.send();
        var response = JSON.parse(xhttp.responseText);
    } catch (error) {
        console.log(error.message);
    }
    // to download logo iff generated
    logoExists = true;
    showImage("data:image/png;base64," + response["image"]);
}

document.getElementById("downImg").onclick = function () {
    if (logoExists) {
        let a = document.createElement("a");
        a.href = finalLogo.src;
        a.download = "logo.png";
        a.click();
    } else {
        alert("Generate Logo First!!")
    }
}
