document.getElementById("genLogo").onclick = function () {
    const uni_name = document.getElementById("uniName").value;
    const img_color = document.getElementById("colors").value;
    let bg_color = document.getElementById("bgColor").value;

    if (bg_color.length < 6) {
        bg_color = "-16"
    }
    if (uni_name === "" || bg_color === "") {
        alert("Enter fields blyat!");
        return;
    }

    const url = "http://127.0.0.1:6969/api/uni_name/" + uni_name + "/img_color/" + img_color +"/bg_color/" + bg_color;

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