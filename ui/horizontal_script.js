const finalLogo = document.getElementById("finalLogo");
const genLogo = document.getElementById("genLogo");
let logoExists = false;

// checkInput returns if the input field is empty or not
function checkInput(uniName) {
    if (uniName === "") {
        alert("Hmm.. a nameless university!");
        return false;
    }
    return true;
}

// showImage sets the image's source
function showImage(src) {
    finalLogo.src = src
}

// genLogo retrieves a generated logo from the given attributes
genLogo.onclick = function () {
    const uniName = document.getElementById("uniName").value;
    const imgColor = document.getElementById("colors").value;
    let opacity = 1.0;
    if (document.getElementById("opacity").checked) {
        opacity = 0.0
    }

    if (!checkInput(uniName)) {
        return
    }

    const url = "/logo-gen/api/gen?uni_name=" + uniName + "&img_color=" + imgColor + "&opacity=" + opacity + "&logo_type=2";

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

// downImg downloads the generated logo and sets file's name to university's name
document.getElementById("downImg").onclick = function () {
    if (logoExists) {
        let a = document.createElement("a");
        a.href = finalLogo.src;
        a.download = "DSC " + document.getElementById("uniName").value
            + " " + document.getElementById("colors").value + " horizontal.png";
        a.click();
    } else {
        alert("whoa hold on!\nare you trying to create a black hole in here??!");
    }
}
