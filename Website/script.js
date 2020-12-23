document.getElementById("genLogo").onclick = function () {
    const url = "http://127.0.0.1:6969/api/image/" + document.getElementById("uniName").value;

    let xhttp = new XMLHttpRequest();
    xhttp.open("GET", url, false);
    xhttp.setRequestHeader("Content-type", "text/html");

    try {
        xhttp.send();
        var response = JSON.parse(xhttp.responseText);
    } catch (error) {
        console.log(error.message);
    }
    showImage("data:image/png;base64," + response["image"])
}

function showImage(src) {
    document.getElementById("finalLogo").src = src
}